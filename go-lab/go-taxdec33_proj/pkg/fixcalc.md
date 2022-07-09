/***************************************************************************************************
* Procedimiento: fixcalc.go - Corrección a montos de impuestos en conceptos fuera de límites del SAT
* Versión: 01                                                                             2018-01-19
****************************************************************************************************
* Definición de términos usados en este procedimiento:
* - Desviación: Es la fracción decimal excedente o faltante de impuesto de un concepto de la factura
*     respecto a los límites del SAT
* - Holgura: Es la fracción decimal de impuesto de un concepto de la factura que se puede
*     aumentar o disminuir en el importe de impuesto del concepto sin exceder los límites del SAT
* - Grupo de material: Es el conjunto de conceptos de una factura que se refieren a la misma
*     identificación del concepto. El caso más común es el de materiales en una factura de productos
* - Desviación objetivo: Es la cantidad de desviación en el proceso de corrección que aun esta
*     pendiente de remover
*
* Alcance:
* - Facturas con desviación de monto de impuesto en uno o más conceptos
* - Corrección de desviaciones usando holguras de uno o más grupos de material en los conceptos de
*   la factura
* - Solo queda fuera del alcance de este procedimiento el caso en que haya más de una desviación en
*   un solo grupo de material (esta caso aun no se ha detectado)
*
* Descripción general del procedimiento:
* Este procedimiento se basa en los siguientes principios:
* 1. Un concepto dentro de un grupo de material puede mostrar una desviación en monto de impuesto
*    con respecto los límites del SAT. Los demás conceptos en el mismo grupo de material pueden 
*    contener holgura respecto a los mismos límites. Lo mismo es válido para la factura en total,
*    es decir uno o más conceptos pueden mostrar desviación, pero el resto de los conceptos pueden
*    contener holguras.
* 2. El ajuste de desviación se hace en primer lugar dentro del grupo de material que contiene el
*    concepto con desviación. Si la desviación no puede ser corregida usando las holguras del propio
*    grupo de material, entonces se comienzan a usar holguras de otros grupos.  
* 3. La desviación de un concepto del grupo de material procurara ser corregida primero con las
*    holguras del resto de el o los conceptos del grupo de material. Si la desviación no puede ser
*    completamente corregida con las holguras de su grupo de material, entonces se usarán las
*    holguras de otros conceptos aun cuando estos no sean del mismo grupo de material
*
* Principio de funcionamiento:
* 1. Se asume que EDICOM cuenta con una tabla de conceptos de la factura, y que contiene en ella
*    datos del concepto, tales como, Identificación, Base del impuesto, Importe del impuesto,
*    límites del SAT para el impuesto del concepto.
* 2. Se solicita incluir implícita o explícitamente 2 datos más en la tabla de conceptos de EDICOM:
*    a). La desviación del impuesto respecto a los límites del SAT (dato con signo y precisión de
*        punto flotante para contener decimales)
*    b). La holgura del impuesto respecto a los límites del SAT (dato con signo y precisión de punto
*        flotante para contener decimales)
* 3. Adicionar dos estructuras de datos más:
*    a). Resumen de grupo de material, en forma de tabla conteniendo la suma de las holguras
*        negativas y positivas del grupo de material y la desviación del grupo de material
*    b). Resumen de factura, en forma de tabla conteniendo la identificación de los materiales con
*        desviación, y un apuntador o índice a la entrada en la tabla de conceptos, de los conceptos
*        con desviación
* 4. Este procedimiento requiere dos pasadas a la tabla de conceptos:
*    1a. Pasada - Para calcular desviación y holguras
*    2a. Pasada - Para corregir la desviación usando las holguras de los conceptos del grupo de
*                 material con desviación en primer lugar, y del resto de los conceptos en ultima
*                 instancia
*
* > Abajo se hará la descripción detallada de los pasos de esta regla compleja
* > La simulación de este procedimiento se realizó usando código escrito en el lenguaje de
*   programación GO de Google (conocido como Golang), y emplea tablas simples de registros o
*   estructuras (conocidos en Golang como slices de estructuras). El ejecutable de la simulación
*   (.EXE) no requiere de ligado dinámico y puede ser ejecutado por un programa externo o un script
*   capaz de llamar a archivos .EXE a ejecución
* > Los archivos de entrada y salida son archivos TXT planos con separación de campos con pipes (|)
*
***************************************************************************************************/

/***************************************************************************************************
* Function: EvaluayEjecutaCorreccion
*
* *** Lógica general para ajustar importes de impuesto de conceptos fuera de límites del SAT ***
* 01 EDICOM crea una tabla de datos de los conceptos de la factura
* 02 Se solicita: aplicar una lógica extra en una primera pasada a la tabla de conceptos. 
* 03   La lógica consiste en 1) Calcular 2 nuevos valores en cada concepto: Holgura y Desviación y
*      2) Hacer una sumarizacion de desviación y holguras por grupo de material
* 04 En caso de que la factura tenga error por limites excedidos entonces:
* 05   Se procede a corregir desviación por desviación en la factura (por lo general será solo una) 
* 06   Se descarga la tabla corregida de conceptos (sin desviaciones)
***************************************************************************************************/

/***************************************************************************************************
* Function: calculaDesviacionyHolguras
*
*** Lógica detallada para calcular desviación y holgura por concepto ***
* 01 En el caso en que el impuesto del concepto es mayor que el límite superior, entonces:
* 02   calcular la desviación como la resta: impuesto - límite superior (debe ser un valor positivo)
* 03 En el caso en que el impuesto del concepto es menor que el límite inferior, entonces:
* 04   calcular la desviación como la resta: impuesto - límite inferior (debe ser un valor negativo)
* 05 En el caso en que el impuesto es menor que el límite superior, entonces:
* 06   calcular la holgura como la resta: límite superior - impuesto    (debe ser un valor positivo)
* 07 En el caso en que el impuesto es mayor que el límite inferior, entonces:
* 08   calcular la holgura como la resta: límite inferior - impuesto    (debe ser un valor negativo)
* 09 Si después de esta determinación, se halla que el valor de la desviación en el concepto no es
*      cero, entonces:
* 10   Dejar registro de que la factura tiene una desviación y requiere ajuste (incluyendo en el
*      registro: Identificación del grupo de material con desviación y numero de entrada en la tabla
*      de conceptos en que se encuentra el concepto con desviación)
*
* Notas:
* - Una desviación positiva indica que el importe de impuesto del concepto excede el límite superior
*   del SAT, y una desviación negativa indica que el importe del impuesto del concepto esta por
*   debajo del límite inferior del SAT
* - Una holgura positiva es relevante cuando hay desviación positiva, e indica la fracción decimal
*   que se puede sumar al importe de impuesto del concepto sin exceder el límite superior, mientras
*   que una holgura negativa solo es relevante cuando la desviación es negativa, e indica la
*   fracción decimal que se puede restar del importe de impuesto del concepto sin caer por debajo
*   del limite inferior
***************************************************************************************************/
/***************************************************************************************************
* Section 1  
* 
* *** Lógica para mantener sumarizacion de desviación y holguras por grupo de material ***
* 01 Este resumen contendrá la identificación del material, la desviación (si la hubiere) y la suma
*    de las holguras
* 03 Se busca en la tabla de resumen de grupos de material el grupo de material del concepto en
*      turno . Cuando la línea para el material del concepto existe ya, entonces:
* 04   Se acumulan en la entrada de la tabla de grupo de material hayada, los valores de desviación
*      (negativa o positiva) y de las holguras (negativas y positivas)
* 05 Si la entrada para el grupo de material no existe aún en la tabla de grupos de material, 
*      entonces;
* 06   Se adiciona una entrada en la tabla de grupos de material con valores iniciales
***************************************************************************************************/

/***************************************************************************************************
* Function: distribuyeDesviacion
*
* *** Lógica detallada para compensar la desviación con las holguras del grupo de material ***
* (Por cuestión técnica del lenguaje GO en cuanto al manejo de datos tipo float, se hace un formateo
* en las tablas de conceptos y de resumen de materiales al número de decimales requeridos por la
* moneda. Esto es solo para propósitos de la simulación. EDICOM podría no tener que hacerlo.
*
* 01 Se revisa cada grupo de material de los conceptos de la factura 
* 02   La lógica para corrección de desviación usando holguras da prioridad en el proceso al uso de
*        holguras dentro del mismo grupo de material de la desviación
* 03     En caso de que la desviación en el grupo de materiales en revisión sea positiva entonces:
* 04       Se guarda la desviación positiva como objetivo del ajuste
* 05     En caso de que la desviación en el grupo de materiales en revisión sea negativa entonces:
* 06       Se guarda la desviación negativa como objetivo del ajuste
* 07     Se procede con el proceso detallado de ajuste 
* 08 Si después de las correcciones en el grupo de material de la desviación aún queda desviación
*      que ajustar entonces:
* 09   Hacer un barrido en los otros grupos de material para tratar de completar el ajuste de la
*        desviación
* 10     Se procede con el proceso detallado de ajuste
***************************************************************************************************/

/***************************************************************************************************
* Function: calculaDistribucion
*
* *** Lógica del proceso detallado de ajuste ***
* 01 Se hace una segunda pasada a la tabla de conceptos por cada material con desviación hasta que
*      la desviación se haga cero (se recorre la tabla de conceptos en orden inverso, de abajo hacia
*      arriba)
* 02   Se limita el proceso solo al grupo de material en revisión
* 03     Si la desviación del grupo de material en turno es positiva y la holgura del concepto en
*          turno es positiva, o si la desviación del grupo de material en turno es negativa y la
*          holgura del concepto en turno es negativa, entonces:
* 04       Reduce la desviación en la tabla de material con la cantidad de holgura disponible en el
*            concepto
* 05       Reduce también la desviación en la tabla de conceptos con la cantidad de holgura
*            disponible en el concepto
* 06       Reduce la desviación objetivo con la cantidad de la holgura disponible en el concepto
* 07       Reduca la holgura disponible en el concepto a cero
* 08       Formatea el remanente de la desviación objetivo al número de decimales de la moneda
* 09       En caso que la desviación objetivo se haya reducido a cero entonces:
* 10         Indica que el problema está resuelto, la desviación se ha corregido
*
* Notas:
* - El reducir una holgura positiva de una desviación positiva tiene el efecto de resta. Entonces,
*   en el concepto con desviación positiva, las holguras positivas (lo que falta para el limite
*   superior) reducen el monto de impuesto en el concepto con desviación, haciéndolo "bajar" al
*   límite superior
* - El reducir una holgura negativa de una desviación negativa tiene el efecto de una suma.
*   Entonces, en el concepto con desviación negativa, las holguras negativas (lo que sobra desde el
*   límite inferior) aumentan el monto de impuesto en el concepto con desviación, haciéndolo "subir"
*   al límite inferior
***************************************************************************************************/

/***************************************************************************************************
* Para esta simulación se usaron funciones que no aparecen en este archivo, pues son solamente
* accesorios para la implementación de la lógica en la simulación con el lenguaje GO. Estas
* se pueden encontrar en el archivo tools.go (que se liga en el momento de la compilación).
* Se asume que EDICOM tiene funciones similares más robustas y generales. Por eso no se detallan
* aquí.
***************************************************************************************************/

/***************************************************************************************************
*
*                 *** Descripción grafica de los terminos Desviación y Holgura ***
*
*  Caso: Desviación positiva y                :                      Caso: Desviación negativa y
*          Holgura positiva                   :                              Holgura negativa
*                                             :
*    Límite                 Límite            :                        Límite               Límite
*   inferior               superior           :                       inferior             superior
*      |                      |               :                          |                    |
*      |                      |               :                          |                    |
*      |            holgura 0 |               :                holgura 0 |                    |
*  ----|----------------------+               :   -----------------------+                    |
*      |      impuesto        |               :                impuesto  |                    |
*      |     x concepto       |               :               x concepto |                    |
*  ----|----------------------+               :   -----------------------+                    |
*      |                      |               :                          |                    |
*      |              holgura |               :                          | holgura            |
*      |             positiva |               :                          | negativa           |
*      |           <--------->|               :                          |<--------->         |
*  ----|----------+           |               :   -----------------------|----------+         |
*      | impuesto |           |               :                impuesto  |          |         |
*      |  x conc  |           |               :               x concepto |          |         |
*  ----|----------+           |               :   -----------------------|----------+         |
*      |                      |               :                          |                    |
*  ----|----------------------|----------+    :   ----------+            |                    |
*      |           impuesto   |          |    :    impuesto |            |                    |
*      |          x concepto  |          |    :     x conc  |            |                    |
*  ----|----------------------|----------+    :   ----------+            |                    |
*      |                      |<--------->    :              <---------->|                    |
*      |                      | desviación    :               desviación |                    |
*      |                      |  positiva     :                negativa  |                    |
*      |                      |               :                          |                    |
*                                             :
*                                             :
***************************************************************************************************/

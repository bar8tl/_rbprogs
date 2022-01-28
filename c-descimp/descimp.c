/*
descimp.c [2002-11-08 BAR8TL]
Genera la descripcion textual en español de un string numerico que representa
un importe en pesos (de maximo 18 digitos enteros y dos digitos decimales).
*/
#include <stdio.h>
#include <stdlib.h>
#include <mem.h>
#include <string.h>

#define RC00       0
#define RC01       1
#define RC02       2
#define RC03       3
#define BLOCK002   2
#define BLOCK004   4
#define BLOCK005   5
#define BLOCK016   16
#define BLOCK040   40
#define BLOCK080   80
#define BLOCK257   257
#define FRACCION   "/100 mn"
#define MONEDA     " peso "
#define MONEDAS    " pesos "
#define PREFIJO    "*"
#define SUFIJO     "*"
#define Y          " y "
#define CERO       "cero"
#define DIECI      "dieci"
#define VEINTI     "veinti"
#define CIEN       " cien "
#define CIENTO     "ciento "
#define CIENTOS    "cientos "
#define QUINIENTOS "quinientos "
#define MIL        " mil "
#define MILLON     " millon "
#define MILLONDE   " millon de"
#define MILLONES   " millones "
#define MILLONESDE " millones de"
#define BILLON     " billon "
#define BILLONDE   " billon de"
#define BILLONES   " billones "
#define BILLONESDE " billones de"

char *unidad[] = {
  "null"  ,
  "un"    ,
  "dos"   ,
  "tres"  ,
  "cuatro",
  "cinco" ,
  "seis"  ,
  "siete" ,
  "ocho"  ,
  "nueve"
};
char *decena[] = {
  "null"     ,
  " "        ,
  " "        ,
  "treinta"  ,
  "cuarenta" ,
  "cincuenta",
  "sesenta"  ,
  "setenta"  ,
  "ochenta"  ,
  "noventa"
};
char *decuni[] = {
  "null"   ,
  "once"   ,
  "doce"   ,
  "trece"  ,
  "catorce",
  "quince"
};
char *cenuni[] = {
  "null"  ,
  " "     ,
  "dos"   ,
  "tres"  ,
  "cuatro",
  " "     ,
  "seis"  ,
  "sete"  ,
  "ocho"  ,
  "nove"
};

int   nmtxt   (void);
void  bldprfx (void);
void  bldsufx (void);
void  bldtrec (void);
void  antrec  (void);
void  andigt  (void);
void  isrtxt  (char *work);
char *ipart   (void);
char *fpart   (void);
int   stof    (double *n);
void  reverse (void);
void  split   (void);

int  sp, nr;
char *s, **trec, *frac, *wtrec, *text, *wktxt, **row;

int rowsize = 72;

int main() { /* Programa principal */
  s = (char *) malloc(BLOCK080);
  printf("Importe en pesos: ");
  gets(s);
  switch (nmtxt()) {
  case RC03:
    printf("Error: String no numerico\n");
    break;
  case RC02:
    printf("Error: Valor negativo\n");
    break;
  }
  return (0);
}

int nmtxt(void) { /* Funcion principal de traduccion */
  int    i, rc;
  double n;

  if ((rc=stof(&n)) == RC00 || rc == RC01) {
    bldtrec();
    bldsufx();
    if (rc == RC01) {
      isrtxt(CERO);
    } else {
      antrec();
    }
    bldprfx();
    reverse();
    wktxt = (char *)  malloc(rowsize);
    row   = (char **) malloc(BLOCK040);
    for (i = 0; i <= 9; i++) {
      *(row+i) = (char *) malloc(rowsize);
    }
    split();
    for (i = 0; i <= nr; i++) {
      printf("%s\n", *(row+i));
    }
  }
  return(rc);
}

void antrec(void) { /* Analisis de tercias */
  int i, j;

  for (i = 1; i <= 5; i++) {
    j = 5 - i;
    if (atoi(*(trec+j)) == 0) {
       ;
    } else {
      if (i == 2 || i == 4) {
        isrtxt(MIL);
      } else if (i == 1) {
        ;
      } else if (i == 3) {
        if (atoi(*(trec+j)) == 1) {
          if (atoi(*(trec+3)) == 0 && atoi(*(trec+4)) == 0) {
            isrtxt(MILLONDE);
          } else {
            isrtxt(MILLON);
          }
        } else {
          if (atoi(*(trec+3)) == 0 && atoi(*(trec+4)) == 0) {
            isrtxt(MILLONESDE);
          } else {
            isrtxt(MILLONES);
          }
        }
      } else if (i == 5) {
        if (atoi(*(trec+j)) == 1) {
          if (atoi(*(trec+1)) == 0 && atoi(*(trec+2)) == 0 &&
              atoi(*(trec+3)) == 0 && atoi(*(trec+4)) == 0) {
            isrtxt(BILLONDE);
          } else {
            isrtxt(BILLON);
          }
        } else {
          if (atoi(*(trec+1)) == 0 && atoi(*(trec+2)) == 0 &&
              atoi(*(trec+3)) == 0 && atoi(*(trec+4)) == 0) {
            isrtxt(BILLONESDE);
          } else {
             isrtxt(BILLONES);
          }
        }
      }
    }
    if (atoi(*(trec+j)) == 0) {
      if (i == 3) {
        if (atoi(*(trec+1)) == 0) {
          ;
        } else {
          if (atoi(*(trec+3)) == 0 && atoi(*(trec+4)) == 0) {
            isrtxt(MILLONESDE);
          } else {
            isrtxt(MILLONES);
          }
        }
      } else {
        ;
      }
    } else {
      wtrec = *(trec+j);
      andigt();
    }
  }
}

void andigt(void) { /* Analisis de los digitos de las tercias */
  int  dig3, dig2, dig1;
  char *adig3, *adig2, *adig1;

  adig3 = (char *) malloc(BLOCK002);
  adig2 = (char *) malloc(BLOCK002);
  adig1 = (char *) malloc(BLOCK002);
  *(adig3+0) = *(wtrec+0);
  *(adig2+0) = *(wtrec+1);
  *(adig1+0) = *(wtrec+2);
  *(adig3+1) = '\0';
  *(adig2+1) = '\0';
  *(adig1+1) = '\0';
  dig3 = atoi(adig3);
  dig2 = atoi(adig2);
  dig1 = atoi(adig1);
  if (dig1 == 0) {
    if (dig2 == 0) {
      ;
    } else {
      isrtxt(decena[dig2]);
    }
  } else if ((dig1 >= 1 && dig1 <= 5) && dig2 == 1) {
    isrtxt(decuni[dig1]);
  } else {
    isrtxt(unidad[dig1]);
    if (dig2 == 0) {
      ;
    } else if (dig2 == 1) {
      isrtxt(DIECI);
    } else if (dig2 == 2) {
      isrtxt(VEINTI);
    } else {
      isrtxt(Y);
      isrtxt(decena[dig2]);
    }
  }
  if (dig3 == 0) {
    ;
  } else if (dig3 == 1) {
    if (dig2 == 0 && dig1 == 0) {
      isrtxt(CIEN);
    } else {
      isrtxt(CIENTO);
    }
  } else if (dig3 == 5) {
    isrtxt(QUINIENTOS);
  } else {
    isrtxt(CIENTOS);
    isrtxt(cenuni[dig3]);
  }
}

int stof(double *n) { /* Validacion-Conversion de un string numerico */
  int    i, sign, rc;
  double val, power;

  rc = RC00;
  for (i=0; *(s+i)==' ' || *(s+i)=='\r' || *(s+i)=='\n' || *(s+i)=='\t'; i++) {
    ;
  }
  sign = 1;
  if (*(s+i) == '+' || *(s+i) == '-') {
    sign = (*(s+i++) == '+') ? 1 : -1;
  }
  for (val = 0; *(s+i) >= '0' && *(s+i) <= '9'; i++) {
    val = 10 * val + *(s+i) - '0';
    if (*(s+i) == '.') {
      i++;
    } else if (*(s+i) != '\0') {
      rc = RC03;
    }
    for (power = 1; *(s+i) >= '0' && *(s+i) <= '9'; i++) {
      val = 10 * val + *(s+i) - '0';
      power *= 10;
    }
    if (*(s+i) != '\0') {
      rc = RC03;
    }
    *n = sign * val / power;
    if (rc != RC03) {
      if (*n < 0) {
        rc = RC02;
      }
    } else if (*n > 0) {
      rc = RC00;
    } else {
      rc = RC01;
    }
  }
  return(rc);
}

char *ipart(void) { /* Obtencion de la parte entera de un string numerico */
  int  i, j;
  char *ip;

  ip = (char *) malloc(BLOCK080);
  for (i = 0; *(s+i) == '\r' || *(s+i) == '\n' || *(s+i) == '\t' ||
              *(s+i) == ' '  || *(s+i) == '+'  || *(s+i) == '-'; i++) {
    ;
  }
  for (j = 0; *(s+i) >= '0' && *(s+i) <= '9'; i++, j++) {
    *(ip+j) = *(s+i);
  }
  *(ip+j) = '\0';
  return(ip);
}

char *fpart(void) { /* Obtencion de la parte decimal de un string numerico */
  int  i, j;
  char *fp;

  fp = (char *) malloc(BLOCK080);
  for (i = 0; *(s+i) != '.' && *(s+i) != '\0'; i++) {
    ;
  }
  if (*(s+i) == '.') {
    i++;
  }
  for (j = 0; *(s+i) >= '0' && *(s+i) <= '9'; i++, j++) {
    *(fp+j) = *(s+i);
  }
  *(fp+j) = '\0';
  return(fp);
}

void bldtrec(void) { /* Construccion de tercias */
  int i, j, k;
  char *pi, *pf, *ip, *fp, *dmmy;

  pi = ipart();
  pf = fpart();
  ip = (char *) malloc(BLOCK016);
  for (k = 0; *(pi+k) != '\0'; k++) {
    ;
  }
  for (i = k, j = 15; i >= 0 && j >= 0; i--, j--) {
    *(ip+j) = *(pi+i);
  }
  for (; j >= 0; j--) {
    *(ip+j) = '0';
  }
  fp = (char *) malloc(3);
  for (i = 0; i <= 1 && *(pf+i) != '\0'; i++) {
    *(fp+i) = *(pf+i);
  }
  if (*(pf+i) == '\0') {
    for (; i <= 1; i++) {
      *(fp+i) = '0';
    }
    *(fp+i) = '\0';
  } else {
    *(fp+i) = '\0';
  }
  frac = fp;
  trec = (char **) malloc(BLOCK005); /* se definen 5 tercias de 4 chars c/u */
  dmmy = (char * ) malloc(1); /* se define espacio de holgura de pointers */
 *dmmy = NULL;
  for (i = 0; i <= 4; i++) {
    *(trec+i) = (char *) malloc(4);
  }
  for (i = 0, k = 0; k <= 4; k++) { /* k es el indice de la tercia (5 tercias) */
    for (j = 0; j <= 2; i++, j++) { /* j es el indice del digito en la tercia */
      *(*(trec+k)+j) = *(ip+i);
    }
    *(*(trec+k)+j) = '\0';
  }
}

void isrtxt(char *work) { /* Insercion de cadenas en el texto de salida */
  int i, j;

  for (j = 0; *(work+j) != '\0'; j++) {
    ;
  }
  for (i = --j; i >= 0; i--, sp++) {
    *(text+sp) = *(work+i);
  }
  *(text+sp) = '\0';
}

void bldprfx(void) { /* Insercion de perfijo */
  isrtxt(PREFIJO);
}

void bldsufx(void) { /* Insercion de sufijo */
  text  = (char *) malloc(BLOCK257);
  wtrec = (char *) malloc(BLOCK004);
  sp = 0;
  isrtxt(SUFIJO);
  isrtxt(FRACCION);
  isrtxt(frac);
  isrtxt(MONEDAS);
}

void reverse(void) { /* Reversa del texto de salida */
  int i, j, k, temp;

  for (k = 0; *(text+k) != '\0'; k++) {
    ;
  }
  for (i = 0, j = --k; (j-i) >= 1; i++, j--) {
    temp = *(text+i);
    *(text+i) = *(text+j);
    *(text+j) = temp;
  }
}

void split(void) { /* Divide una linea de texto en varias */
  int i, j, k, l;

  for (i = j = k = l = 0; *(text+i) != '\0'; i++) {
    *(wktxt+k) = *(text+i);
    if (*(text+i) == ' ') {
      if ((j+k) >= rowsize) {
        *(*(row+nr++)+j) = '\0';
        j = 0;
      }
      for (l = 0; l <= k; l++) {
        *(*(row+nr)+j++) = *(wktxt+l);
      }
      k = 0;
    } else {
      k++;
    }
  }
  *(wktxt+k) = '\0';
  for (l = 0; l <= k; l++) {
    *(*(row+nr)+j++) = *(wktxt+l);
  }
}

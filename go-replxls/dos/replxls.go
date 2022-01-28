// replxls.go [2012-11-08 BAR8TL]
// Starts command line mode process (not through a dialog) to build and include
// Excel attachments base64 encoded within Mex XML e-invoices
package main

import "bar8tl/p/replxls/rb"
import "fmt"
import "os"
import "regexp"

func main() {
  if _, err := os.Stat(os.Args[1]); err != nil {
    fmt.Printf(
      "El archivo %s no existe, proceso detenido. Rectifique.\n", os.Args[1])
    return
  }
  if _, err := os.Stat(os.Args[2]); err != nil {
    fmt.Printf(
      "El archivo %s no existe, proceso detenido. Rectifique.\n", os.Args[2])
    return
  }
	re := regexp.MustCompile("[.]([^.]+)$")
	newxls := re.ReplaceAllString(os.Args[1], "_new.$1")
  if err := rb.ReplXls(os.Args[1], os.Args[2], newxls); err != nil {
    fmt.Printf("Error: Ocurrio un error durante la generacion del XLS base64.")
  } else {
    fmt.Printf("Exito: Se reemplazo correctamente la porcion XLS en la " +
      "factura XML.\nEl resultado esta en el archivo %s\n", newxls)
  }
}

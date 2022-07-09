package main

import "bytes"
import "fmt"
import "log"
import "os/exec"

func main() {
  var filen string
  fmt.Printf("1. Factura J92751 (Mazda)\n")
  fmt.Printf("2. Factura J92770 (Nissan)\n")
  fmt.Printf("3. Factura J92772 (Nissan)\n")
  fmt.Printf("4. Factura J92773 (Nissan)\n")
  fmt.Printf("5. Factura J92913 ()\n")
  fmt.Printf("6. Factura J92843 ()\n")
  fmt.Printf("7. Factura J93138 (Mazda)\n")
  fmt.Printf("Probar con archivo de test#: ")
  fmt.Scanf("%s", &filen)
  cmd := exec.Command("cmd", "/c", "taxdec33 2 test"+filen+".txt test"+filen+"_out.txt")
  var stderr bytes.Buffer
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
    return	}
}

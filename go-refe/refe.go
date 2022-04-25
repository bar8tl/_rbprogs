// refe.go [2018-03-01 BAR8TL]
// Start login credentials retrieval into the Win10 clipboard 
package main

import rb "bar8tl/p/refe"
import "os"
import "strings"

func main() {
  if len(os.Args) > 1 {
    tokn := strings.Split(strings.ToLower(os.Args[1]), ".")
    if len(tokn) == 1 {
      rb.Cpykey("rb", tokn[0], "pwd")
      return
    } else if len(tokn) == 2 {
      if tokn[0] == "mz" {
        rb.Cpykey(tokn[0], tokn[1], "pwd")
        return
      } else {
        rb.Cpykey("rb", tokn[0], tokn[1])
        return
      }
    } else if len(tokn) == 3 {
      rb.Cpykey(tokn[0], tokn[1], tokn[2])
      return
    }
  }
  rb.Console()
}

// rscgen.go [2012-04-12 BAR8TL] Generates and display data resources
package main

import ut "bar8tl/p/rblib"
import "flag"
import "fmt"

func main() {
  proc := flag.String("p", "na", "Procedure: [zipgen|zipdsp]")
  ifnm := flag.String("i", "na", "Input file name")
  ofnm := flag.String("o", "na", "Output file name")
  flag.Parse()
  switch *proc {
    case "zipgen":
      ut.Zipgen(*ifnm, *ofnm)
    case "zipdsp":
      ut.Zipdsp(*ifnm)
    case "na":
      fmt.Println("Option invalid")
  }
}

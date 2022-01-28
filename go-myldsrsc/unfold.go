package main

import "bufio"
import "fmt"
import "io"
import "os"
import "strings"
import "unicode"

func main() {
  var fbyte rune
  var fchar string
  ifile, _ := os.Open(os.Args[1] + "testi.xml")
  defer ifile.Close()
  rdr := bufio.NewReader(ifile)
  ofile, _ := os.Create(os.Args[2] + "testo.txt")
  defer ofile.Close()
  wrt := bufio.NewWriter(ofile)
  punto := false
  for iline, err := rdr.ReadString(byte('\n')); err != io.EOF; iline, err = rdr.ReadString(byte('\n')) {
    parts := strings.Split(iline, "\r")
    sline, lline := parts[0], len(parts[0])
    if lline == 0 {
      fmt.Fprintf(wrt, "\r\n")
      punto = false
      wrt.Flush()
      continue
    }
    if sline == "<v>" || sline == "<f>" || sline == "<c>" || sline == "<d>" || sline == "</v>" || sline == "</f>" ||
       sline == "</c>" || sline == "</d>" {
      fmt.Fprintf(wrt, "%s\r\n", sline)
      punto = false
      wrt.Flush()
      continue
    }
    fchar = sline[0:1]
    fbyte = rune(sline[0])
    fmt.Println(sline)
    if fchar == "•" {
      fmt.Println(fchar)
    }
    if unicode.IsUpper(fbyte) || fchar == "•" || fchar == "-" {
      if punto {
        if sline[lline-1:lline] == "." {
          fmt.Fprintf(wrt, "\r\n%s\r\n", sline)
          punto = true
          wrt.Flush()
          continue
        } else {
          fmt.Fprintf(wrt, "\r\n%s ", sline)
          punto = false
          wrt.Flush()
          continue
        }
      } else {
        if sline[lline-1:lline] == "." {
          fmt.Fprintf(wrt, "\r\n%s\r\n", sline)
          punto = true
          wrt.Flush()
          continue
        } else {
          fmt.Fprintf(wrt, "\r\n%s ", sline)
          punto = false
          wrt.Flush()
         continue
        }
      }
    } else {
      if punto {
        if sline[lline-1:lline] == "." {
          fmt.Fprintf(wrt, "%s\r\n", sline)
          punto = true
          wrt.Flush()
          continue
        } else {
          fmt.Fprintf(wrt, "%s ", sline)
          punto = false
          wrt.Flush()
          continue
        }
      } else {
        if sline[lline-1:lline] == "." {
          fmt.Fprintf(wrt, "%s\r\n", sline)
          punto = true
          wrt.Flush()
          continue
        } else {
          fmt.Fprintf(wrt, "%s ", sline)
          punto = false
          wrt.Flush()
         continue
        }
      }
    }
  }
}
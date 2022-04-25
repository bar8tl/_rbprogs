package main

import "bufio"
import "encoding/csv"
import "io"
import "log"
import "os"
import "path/filepath"
import "strings"

func main() {
  ifile, err := os.Open(os.Args[1]+"\\"+os.Args[2])
  if err != nil {
    log.Fatalf("Input file %s not found: %s\r\n", os.Args[1]+"\\"+os.Args[2],
    err)
  }
  defer ifile.Close()
  filid := ifile.Name()
  extn  := filepath.Ext(filid)
  file  := strings.TrimRight(filid, extn)
  filnm := file + "_out.csv"
  ofile, err := os.Create(filnm)
  if err != nil {
    log.Fatalf("Failure creating output file %s: %s\r\n", filnm, err)
  }
  defer ofile.Close()
  rdr := bufio.NewReader(ifile)
  wtr := csv.NewWriter(ofile)
  defer wtr.Flush()
  for iline, err := rdr.ReadString(byte('\n')); err != io.EOF;
    iline, err = rdr.ReadString(byte('\n')) {
    tokn  := strings.Split(iline[0:len(iline)-2], string('\t'))
    mpath, msize := tokn[0], tokn[1]
    dirs  := strings.Split(mpath, "\\")
    mtype, filen := "", ""
    if msize == "0" {
      mtype, filen = "d", ""
    } else {
      mtype, filen = "f", dirs[len(dirs)-1]
    }
    var direc [10]string
    for i := 0; i < 10; i++ {
      if msize == "0" {
        if i < (len(dirs)) {
          direc[i] = dirs[i]
        } else {
          direc[i] = ""
        }
      } else {
        if i < (len(dirs)-1) {
          direc[i] = dirs[i]
        } else {
          direc[i] = ""
        }
      }
    }
    row := []string{mtype, direc[0], direc[1], direc[2], direc[3], direc[4],
      direc[5], direc[6], direc[7], direc[8], direc[9], filen, msize}
    if err := wtr.Write(row); err != nil {
      log.Fatalf("Error writing record to file %s: %s\r\n", filnm, err)
    }
  }
}

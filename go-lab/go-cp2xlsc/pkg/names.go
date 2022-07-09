// names.go - Funtions to rename input/output files according status
// 2022-05-17 BAR8TL Version1.0 - Released
package cp2xlsc

import "log"
import "os"
import "strings"
import "time"

// Functions in this file need to be further developed, status is in progress
func PassFilter(s Settings_tp, f string) bool {
  if strings.Contains(f, "processed") {
    return false
  }
  return true
}

func RenameInpFile(inpdr, inpfl, inptp string) {
  t := time.Now()
  oldName := inpdr + inpfl + inptp
  newName := inpdr + t.Format("20060102") + "_" + inpfl +"-1.0-processed"+ inptp
  err := os.Rename(oldName, newName)
  if err != nil {
    log.Fatalf("Input file %s renaming error: %s\r\n", oldName, err)
  }
}

func RenameOutFile(outdr, outfl, outtp string) {
  t := time.Now()
  oldName := outdr + outfl + outtp
  newName := outdr + t.Format("20060102") + "_" + outfl + "-2.0" + outtp
  err := os.Rename(oldName, newName)
  if err != nil {
    log.Fatalf("Output file %s renaming error: %s\r\n", oldName, err)
  }
}

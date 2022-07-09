// cp2xlsc.go - Extend Pagos1.0 EDICOM-file with Pagos2.0 fields (Entry point)
// 2022-05-17 BAR8TL Version1.0 - Released
package main

import rb "bar8tl/p/cp2xlsc"
import ut "bar8tl/p/rblib"
import "io/ioutil"
import "log"
import "path/filepath"
import "strings"

func main() {
  s := rb.NewSettings("_config.json", "_deflts.json")
  for _, parm := range s.Cmdpr {
    if parm.Optn == "txc" { // Perform processes for tax calculation
      taxCalc(parm, s)
    }
  }
}

func taxCalc(parm ut.Param_tp, s rb.Settings_tp) {
  s.SetRunVars(parm)
  if s.Modep == s.Konst.BATCH { // For batch process browse inputs directory
    files, _ := ioutil.ReadDir(s.Inpdr)
    for _, f := range files {
      ffile  := f.Name()
      s.Flext = filepath.Ext(ffile)
      s.Flnam = strings.TrimRight(ffile, s.Flext)
      s.Flfil = s.Flnam + s.Flext
      if len(s.Ifilt) == 0 || (len(s.Ifilt) > 0 && rb.PassFilter(s, s.Flnam)) {
        procIndivFile(s, s.Inpdr, ffile)
      }
    }
  } else {                      // For individual process use specified file
    ffile  := s.Infil
    s.Flext = filepath.Ext(ffile)
    s.Flnam = strings.TrimRight(ffile, s.Flext)
    s.Flfil = s.Flnam + s.Flext
    if len(s.Ifilt) == 0 || (len(s.Ifilt) > 0 && rb.PassFilter(s, s.Flnam)) {
      procIndivFile(s, s.Fildr, ffile)
    }
  }
}

func procIndivFile(s rb.Settings_tp, dir, f string) {
  loadAssets(s)
  c   := NewCalctax(s)
  rdr := NewReader()
  rdr.OpenInpExcel(dir, f)
  rows, err := F.GetRows(s.Konst.TAB)
  if err != nil {
    log.Fatal(err)
  }
  wtr := NewWriter(s)
  wtr.CreateOutExcel()
  for _, row := range rows {
    rdr.GetLineFields(row)
    switch rds[m["documentType"]] {
      case "Document Type": wtr.FetchTitle()
      case "DZ", "PK"     : ProcessPaymentLine(c, wtr)
      case "RV"           : ProcessInvoiceLine(c)
    }
  }
  c.FetchPaymentLine(wtr).FetchInvoiceLines(wtr)
  wtr.ProduceExcelOutput(s, dir)
}

// writer.go [2022-04-06 BAR8TL] Excel file EDICOM writer functions
package cp2xlsc

import "github.com/xuri/excelize/v2"
import "fmt"
import "log"

var f1    *excelize.File
var f2    *excelize.File
var recn   int
var flfil  string
var flnam  string
var flext  string
var outpt  string
var ONE    string
var MANY   string
var FULL   string
var TAB    string

type Writer_tp struct {
  index1 int
  index2 int
}

func NewWriter(s Settings_tp) *Writer_tp {
  var w Writer_tp
  recn  = 0
  flfil = s.Flfil
  flnam = s.Flnam
  flext = s.Flext
  outpt = s.Outpt
  ONE   = s.Konst.ONE
  MANY  = s.Konst.MANY
  FULL  = s.Konst.FULL
  TAB   = s.Konst.TAB
  NewLists("lists.dat")
  return &w
}

func (w *Writer_tp) CreateOutExcel() {
  f1 = excelize.NewFile()
  f2 = excelize.NewFile()
  w.index1 = f1.NewSheet(TAB)
  w.index2 = f2.NewSheet(TAB)
}

func (w *Writer_tp) ProduceExcelOutput(dir string) {
  if outpt == ONE  || outpt == FULL {
    f1.SetActiveSheet(w.index1)
    if err := f1.SaveAs(dir+flnam+"-s"+flext); err != nil {
      log.Fatal(err)
    }
    RenameOutFile(dir, flnam+"-s", flext)
  }
  if outpt == MANY || outpt == FULL {
    f2.SetActiveSheet(w.index2)
    if err := f2.SaveAs(dir+flnam+"-m"+flext); err != nil {
      log.Fatal(err)
    }
    RenameOutFile(dir, flnam+"-m", flext)
  }
  //RenameInpFile(dir, flnam, flext)
}

func (w *Writer_tp) FetchTitle() {
  recn++
  for i, _ := range tt {
    fmt.Printf("|%d|%s|\r\n", i, tt[i])
    if (i >=  1 && i <= 37) && (outpt == ONE  || outpt == FULL) {
      f1.SetCellValue(TAB, fmt.Sprintf(cc[i]+"%d", recn), tt[i])
    }
    if (i >=  1 && i <= 37) && (outpt == MANY || outpt == FULL) {
      f2.SetCellValue(TAB, fmt.Sprintf(cc[i]+"%d", recn), tt[i])
    }
    if (i >= 38 && i <= 49) && (outpt == ONE  || outpt == FULL) {
      f1.SetCellValue(TAB, fmt.Sprintf(c1[i]+"%d", recn), tt[i])
    }
    if (i >= 50 && i <= 91) && (outpt == MANY || outpt == FULL) {
      f2.SetCellValue(TAB, fmt.Sprintf(c3[i]+"%d", recn), tt[i])
    }
  }
}

func (w *Writer_tp) PrintLineExcel() {
  recn++
  for i, _ := range cc {
    if (i >=  1 && i <= 37) && (outpt == ONE  || outpt == FULL) {
      f1.SetCellValue(TAB, fmt.Sprintf(cc[i]+"%d", recn), ps[i])
    }
    if (i >=  1 && i <= 37) && (outpt == MANY || outpt == FULL) {
      f2.SetCellValue(TAB, fmt.Sprintf(cc[i]+"%d", recn), ps[i])
    }
  }
  for i, _ := range c1 {
    if (i >= 38 && i <= 49) && (outpt == ONE  || outpt == FULL) {
      f1.SetCellValue(TAB, fmt.Sprintf(c1[i]+"%d", recn), pf[i])
    }
  }
  for i, _ := range c3 {
    if (i >= 50 && i <= 91) && (outpt == MANY || outpt == FULL) {
      f2.SetCellValue(TAB, fmt.Sprintf(c3[i]+"%d", recn), pf[i])
    }
  }
}

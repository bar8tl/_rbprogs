// writer.go - EDICOM excel file extended to Pagos 2.0 writer functions
// 2022-05-17 BAR8TL Version1.0 - In progress
package main

import rb "bar8tl/p/cp2xlsc"
import "github.com/xuri/excelize/v2"
import "fmt"
import "log"

var ONE  string
var MANY string
var FULL string
var TAB  string
var f1  *excelize.File
var f2  *excelize.File

type Writer_tp struct {
  index1 int
  index2 int
}

func NewWriter(s rb.Settings_tp) *Writer_tp {
  var w Writer_tp
  ONE  = s.Konst.ONE
  MANY = s.Konst.MANY
  FULL = s.Konst.FULL
  TAB  = s.Konst.TAB
  return &w
}

func (w *Writer_tp) CreateOutExcel() {
  f1 = excelize.NewFile()
  f2 = excelize.NewFile()
  w.index1 = f1.NewSheet(TAB)
  w.index2 = f2.NewSheet(TAB)
}

func (w *Writer_tp) ProduceExcelOutput(s rb.Settings_tp, dir string) {
  if outpt == ONE  || outpt == FULL {
    f1.SetActiveSheet(w.index1)
    if err := f1.SaveAs(dir+s.Flnam+"-s"+s.Flext); err != nil {
      log.Fatal(err)
    }
    rb.RenameOutFile(dir, s.Flnam+"-s", s.Flext)
  }
  if outpt == MANY || outpt == FULL {
    f2.SetActiveSheet(w.index2)
    if err := f2.SaveAs(dir+s.Flnam+"-m"+s.Flext); err != nil {
      log.Fatal(err)
    }
    rb.RenameOutFile(dir, s.Flnam+"-m", s.Flext)
  }
  //RenameInpFile(dir, s.Flnam, s.Flext)
}

func (w *Writer_tp) FetchTitle() {
  recn++
  for i, _ := range tt {
    if outpt == ONE  || outpt == FULL {
      switch {
      case i >=  0 && i <= 36:
        w.SetCellString(f1, cc[i], tt[i])
      case i >= 37 && i <= 48:
        w.SetCellString(f1, c1[i], tt[i])
      }
    }
    if outpt == MANY || outpt == FULL {
      switch {
      case i >=  0 && i <= 36:
        w.SetCellString(f2, cc[i], tt[i])
      case i >= 49 && i <= 90:
        w.SetCellString(f2, c3[i], tt[i])
      }
    }
  }
}

func (w *Writer_tp) PrintPaymentLine() {
  recn++
  for i, _ := range tt {
    if outpt == ONE  || outpt == FULL {
      switch {
      case i >=  0 && i <= 27:
        w.SetCellString(f1, cc[i], pms[i])
      case i >= 28 && i <= 36:
        if contains(alphac, i) {
          w.SetCellString(f1, cc[i], pms[i])
        } else {
          w.SetCellFloat (f1, cc[i], pmf[i])
        }
      case i >= 37 && i <= 48:
        if contains(alpha1, i) {
          w.SetCellString(f1, c1[i], pms[i])
        } else {
          w.SetCellFloat (f1, c1[i], pmf[i])
        }
      }
    }
    if outpt == MANY || outpt == FULL {
      switch {
      case i >=  0 && i <= 27:
        w.SetCellString(f2, cc[i], pms[i])
      case i >= 28 && i <= 36:
        if contains(alphac, i) {
          w.SetCellString(f2, cc[i], pms[i])
        } else {
          w.SetCellFloat (f2, cc[i], pmf[i])
        }
      case i >= 49 && i <= 90:
        if contains(alpha3, i) {
          w.SetCellString(f2, c3[i], pms[i])
        } else {
          w.SetCellFloat (f2, c3[i], pmf[i])
        }
      }
    }
  }
}

func (w *Writer_tp) PrintInvoiceLine(ds [91]string, df [91]float64) {
  recn++
  for i, _ := range tt {
    if outpt == ONE  || outpt == FULL {
      switch {
      case i >=  0 && i <= 27:
        w.SetCellString(f1, cc[i], ds[i])
      case i >= 28 && i <= 36:
        if contains(alphac, i) {
          w.SetCellString(f1, cc[i], ds[i])
        } else {
          w.SetCellFloat (f1, cc[i], df[i])
        }
      case i >= 37 && i <= 48:
        if contains(alpha1, i) {
          w.SetCellString(f1, c1[i], ds[i])
        } else {
          w.SetCellFloat (f1, c1[i], df[i])
        }
      case i >= 49 && i <= 90:
        if contains(alpha3, i) {
          w.SetCellString(f1, c3[i], ds[i])
        } else {
          w.SetCellFloat (f1, c3[i], df[i])
        }
      }
    }
    if outpt == MANY || outpt == FULL {
      switch {
      case i >=  0 && i <= 27:
        w.SetCellString(f2, cc[i], ds[i])
      case i >= 28 && i <= 36:
        if contains(alphac, i) {
          w.SetCellString(f2, cc[i], ds[i])
        } else {
          w.SetCellFloat (f2, cc[i], df[i])
        }
      case i >= 37 && i <= 48:
        if contains(alpha1, i) {
          w.SetCellString(f2, c1[i], ds[i])
        } else {
          w.SetCellFloat (f2, c1[i], df[i])
        }
      case i >= 49 && i <= 90:
        if contains(alpha3, i) {
          w.SetCellString(f2, c3[i], ds[i])
        } else {
          w.SetCellFloat (f2, c3[i], df[i])
        }
      }
    }
  }
}

func (w *Writer_tp) SetCellString(f *excelize.File, col string, val string) {
  f.SetCellValue(TAB, fmt.Sprintf(col+"%d", recn), val)
}

func (w *Writer_tp) SetCellFloat(f *excelize.File, col string, val float64) {
  if val == 0.0 {
    f.SetCellValue(TAB, fmt.Sprintf(col+"%d", recn), "")
  } else {
    f.SetCellValue(TAB, fmt.Sprintf(col+"%d", recn), val)
  }
}

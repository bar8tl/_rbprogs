// reader.go - EDICOM excel file reader functions
// 2022-05-17 BAR8TL Version1.0 - In progress
package main

import ut "bar8tl/p/rblib"
import "github.com/xuri/excelize/v2"
import "log"
import "strconv"
import "fmt"

var F   *excelize.File
var rds [28]string
var rdf [28]float64

type Reader_tp struct {
}

func NewReader() *Reader_tp {
  var r Reader_tp
  return &r
}

func (r *Reader_tp) OpenInpExcel(dir, fname string) {
  var err error
  F, err = excelize.OpenFile(dir+fname)
  if err != nil {
    log.Fatal(err)
  }
}

func (r *Reader_tp) GetLineFields(row []string) {
  for i, _ := range row {
    if contains(numers, i ) {
      rdf[i], _ = strconv.ParseFloat(row[i], 64)
      rds[i]    = fmt.Sprintf("%.2f", ut.Round(rdf[i], 2))
    } else if i == m["effExchangeRate"] {
      rdf[i], _ = strconv.ParseFloat(row[i], 64)
      rds[i]    = fmt.Sprintf("%.6f", ut.Round(rdf[i], 6))
      if rds[i] == "" || rds[i] == "MXN" {
        rdf[i] = 1.0
      }
      if rdf[i] == 0.0 {
        rds[i] = ""
      }
    } else {
      rds[i] = row[i]
    }
  }
  if rds[m["documentType"]] == "DZ" || rds[m["documentType"]] == "PK"{
    if rdf[m["importePago"]] == 0.0 {
      rds[m["importePago"]] = ""
    }
    if rdf[m["importeSaldoAnterior"]] == 0.0 {
      rds[m["importeSaldoAnterior"]] = ""
    }
    if rdf[m["importeSaldoInsoluto"]] == 0.0 {
      rds[m["importeSaldoInsoluto"]] = ""
    }
  }
}

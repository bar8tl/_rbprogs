// lists.go [2012-04-12 BAR8TL] Load required konst table as program resource
package cp2xlsc

import "archive/zip"
import "bytes"
import "io"
import "log"
import "strconv"
import "strings"

var m  map[string]int
var tt [92]string
var cc [38]string
var c1 [50]string
var c3 [92]string
var ps [92]string
var pf [92]float64

type Lists_tp struct {
}

func NewLists(fname string) *Lists_tp {
  var l Lists_tp
  m = make(map[string]int)
  rc, err := zip.OpenReader(fname)
  if err != nil {
    log.Fatalf("Open Lists Archive file: %v\n", err)
  }
  defer rc.Close()
  for _, f := range rc.File {
    var d io.ReadCloser
    d, err = f.Open()
    if err != nil {
      log.Fatalf("Open Lists archived file: %v\n", err)
    }
    defer d.Close()
    buf := new(bytes.Buffer)
    buf.ReadFrom(d)
    for iline, err := buf.ReadString(byte('\n')); err != io.EOF; iline,
      err = buf.ReadString(byte('\n')) {
      flds := strings.Split(string(iline), "|")
      sfl2 := strings.Split(flds[2], "\n")
      switch flds[0] {
        case "idx": l.bldMap(flds[1], sfl2[0])
        case "tit": l.bldTit(flds[1], sfl2[0])
        case "col": l.bldCol(flds[1], sfl2[0])
      }
    }
  }
  return &l
}

func (l *Lists_tp) bldMap(key, val string) {
  valn, _ := strconv.Atoi(val)
  m[key] = valn
}

func (l *Lists_tp) bldTit(key, val string) {
  keyn, _ := strconv.Atoi(key)
  tt[keyn] = val
}

func (l *Lists_tp) bldCol(key, val string) {
  keyn, _ := strconv.Atoi(key)
  switch {
    case keyn >= 1  && keyn <= 28 : cc[keyn] = val
    case keyn >= 29 && keyn <= 37 : cc[keyn] = val
    case keyn >= 38 && keyn <= 49 : c1[keyn] = val
    case keyn >= 50 && keyn <= 91 : c3[keyn] = val
  }
}

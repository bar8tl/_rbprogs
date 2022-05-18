package main

import ut "bar8tl/p/rblib"
import "bufio"
import "fmt"
import "io"
import "log"
import "os"
import "strconv"
import "strings"

const DECIMALS = 2

type t_invo struct {
  mtwdv string
  i     int
}

type t_matr struct {
  ident string
  items int
  negsl float64
  possl float64
  negdv float64
  posdv float64
}

type t_conc struct {
  ident string
  squan string
  spuni string
  staxb string
  staxr string
  staxa string
  taxbs float64
  taxrt float64
  taxam float64
  btaxa float64
  taxcl float64
  taxlw float64
  taxup float64
  bslck float64
  bdvtn float64
  slack float64
  devtn float64
}

var invo []t_invo
var matr []t_matr
var conc []t_conc
var targt float64
var staxa float64
var staxc float64
var firstMatrl bool = true
var resolved bool

func cargaTablaConceptos() {
  ifile, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatalf("Open file: %v\n", err)
  }
  defer ifile.Close()
  rdr := bufio.NewReader(ifile)
  for line, err := rdr.ReadString(byte('\n')); err != io.EOF; line, err = rdr.ReadString(byte('\n')) {
    flds := strings.Split(string(line), "|")
    sfld := strings.Split(flds[5], "\r")
    flds[5] = sfld[0]
    var nfld [5]float64
    ident := flds[0]
    for j := 1; j < len(flds); j++ {
      nfld[j-1], _ = strconv.ParseFloat(flds[j], 64)
    }
    taxbs := ut.Round(nfld[2], DECIMALS)
    taxrt := ut.Round(nfld[3], DECIMALS)
    taxam := ut.Round(nfld[4], DECIMALS)
    taxlw := ut.Ffloor(taxbs, taxrt, DECIMALS)
    taxup := ut.Fceil (taxbs, taxrt, DECIMALS)
    conc = append(conc, t_conc{ident, flds[1], flds[2], flds[3], flds[4],
      flds[5], taxbs, taxrt, taxam, taxam, (taxbs * taxrt), taxlw, taxup,
      0.0, 0.0, 0.0, 0.0})
    staxa += taxam
  }
}

func formatTablesDecimals() {
  for i := 0; i < len(conc); i++ {
    conc[i].slack = ut.Round(conc[i].slack, DECIMALS)
    conc[i].devtn = ut.Round(conc[i].devtn, DECIMALS)
    conc[i].bslck = conc[i].slack
    conc[i].bdvtn = conc[1].devtn
  }
  for j := 0; j < len(matr); j++ {
    matr[j].negsl = ut.Round(matr[j].negsl, DECIMALS)
    matr[j].possl = ut.Round(matr[j].possl, DECIMALS)
    matr[j].negdv = ut.Round(matr[j].negdv, DECIMALS)
    matr[j].posdv = ut.Round(matr[j].posdv, DECIMALS)
  }
}

func descargaResultado() {
  for i, _ := range conc {
    fmt.Printf("%s %.2f %.2f %.6f %.2f %.2f\n", conc[i].ident, conc[i].btaxa,
      conc[i].taxam, conc[i].taxcl, conc[i].taxlw, conc[i].taxup)
  }
  ofile, err := os.Create(os.Args[2])
  if err != nil {
    log.Fatalf("Open file: %v\n", err)
  }
  defer ofile.Close()
  w := bufio.NewWriter(ofile)
  for i, _ := range conc {
    fmt.Fprintf(w, "%s|%s|%s|%s|%s|%.2f\r\n", conc[i].ident, conc[i].squan,
      conc[i].spuni, conc[i].staxb, conc[i].staxr, conc[i].taxam)
    staxc += conc[i].taxam
  }
  if ut.Round(staxa, DECIMALS) != ut.Round(staxc, DECIMALS) {
    log.Fatalf("Falla en ajuste. Los totales de impuestos antes y despues " +
      "del ajuste no coinciden")
  }
  w.Flush()
}

package fixcalc

import ut "bar8tl/p/rblib"
import "fmt"
import "strconv"
import "strings"

// Data type for Invoice data
type t_invo struct {
  mtwdv string
  i     int
}

type It_invo struct {
  targt    float64
  staxa    float64
  staxc    float64
  resolved bool
  DECIMALS int
  Dec      string
  invo     []t_invo
}

func New_It_invo(dec string) *It_invo {
  var g It_invo
  d, _ := strconv.Atoi(dec)
  g.targt    = 0.0
  g.staxa    = 0.0
  g.staxc    = 0.0
  g.resolved = false
  g.DECIMALS = d
  g.Dec      = dec
  return &g
}

// Data type for Material group
type t_matr struct {
  ident string
  items int
  negsl float64
  possl float64
  negdv float64
  posdv float64
}

type It_matr struct {
  matr []t_matr
}

func New_It_matr() *It_matr {
  var m It_matr
  return &m
}

// Data type for Concepts detail
type t_conc struct {
  ident string
  squan string
  spuni string
  staxb string
  staxr string
  staxa string
  taxbs float64
  taxrt float64
  Taxam float64
  btaxa float64
  taxcl float64
  taxlw float64
  taxup float64
  bslck float64
  bdvtn float64
  slack float64
  devtn float64
}

type It_conc struct {
  Conc []t_conc
}

func New_It_conc() *It_conc {
  var t It_conc
  return &t
}

func (t *It_conc) AdicionaEntradaATabla(g *It_invo, line string) *It_conc {
  flds := strings.Split(string(line), "|")
  sfld := strings.Split(flds[6], "\r")
  flds[6] = sfld[0]
  var nfld [5]float64
  ident := flds[0]
  for j := 1; j < len(flds)-1; j++ {
    nfld[j-1], _ = strconv.ParseFloat(flds[j], 64)
  }
  taxbs := ut.Round(nfld[2], g.DECIMALS)
  taxrt := ut.Round(nfld[3], g.DECIMALS)
  taxam := ut.Round(nfld[4], g.DECIMALS)
  taxlw := ut.Ffloor(taxbs, taxrt, g.DECIMALS)
  taxup := ut.Fceil (taxbs, taxrt, g.DECIMALS)
  t.Conc = append(t.Conc, t_conc{ident, flds[1], flds[2], flds[3], flds[4],
    flds[5], taxbs, taxrt, taxam, taxam, (taxbs * taxrt), taxlw, taxup, 0.0,
    0.0, 0.0, 0.0})
  g.staxa += taxam
  return t
}

func (t *It_conc) formatTableDecimals(g *It_invo, m *It_matr) *It_conc {
  for i := 0; i < len(t.Conc); i++ {
    t.Conc[i].slack = ut.Round(t.Conc[i].slack, g.DECIMALS)
    t.Conc[i].devtn = ut.Round(t.Conc[i].devtn, g.DECIMALS)
    t.Conc[i].bslck = t.Conc[i].slack
    t.Conc[i].bdvtn = t.Conc[i].devtn
  }
  for j := 0; j < len(m.matr); j++ {
    m.matr[j].negsl = ut.Round(m.matr[j].negsl, g.DECIMALS)
    m.matr[j].possl = ut.Round(m.matr[j].possl, g.DECIMALS)
    m.matr[j].negdv = ut.Round(m.matr[j].negdv, g.DECIMALS)
    m.matr[j].posdv = ut.Round(m.matr[j].posdv, g.DECIMALS)
  }
  return t
}

func (t *It_conc) checkConsistency(g *It_invo) *It_conc {
  for i, _ := range t.Conc {
    g.staxc += t.Conc[i].Taxam
  }
  if ut.Round(g.staxa, g.DECIMALS) != ut.Round(g.staxc, g.DECIMALS) {
    dec := strconv.Itoa(g.DECIMALS)
    soutc := "Error: Falla en ajuste. Los totales de impuestos antes y " +
      "despues del ajuste no coinciden: %."+dec+"f %."+dec+"f\n"
    fmt.Printf(soutc, ut.Round(g.staxa, g.DECIMALS), ut.Round(g.staxc, g.DECIMALS))
  }
  return t
}

func (t *It_conc) DumpTable() *It_conc {
  for i, _ := range t.Conc {
    fmt.Printf("%v\n", t.Conc[i])
  }
  return t
}

func (t *It_conc) MuestraTabla(g *It_invo) *It_conc {
  dec := strconv.Itoa(g.DECIMALS)
  soutc := "%s %."+dec+"f %."+dec+"f %.6f %."+dec+"f %."+dec+"f\n"
  for i, _ := range t.Conc {
    fmt.Printf(soutc, t.Conc[i].ident, t.Conc[i].btaxa, t.Conc[i].Taxam,
      t.Conc[i].taxcl, t.Conc[i].taxlw, t.Conc[i].taxup)
  }
  return t
}

func (t *It_conc) ObtieneEntradaDeTabla(i int, g *It_invo) string {
  dec := strconv.Itoa(g.DECIMALS)
  soutc := "%s|%s|%s|%s|%s|%."+dec+"f\r\n"
  line := fmt.Sprintf(soutc, t.Conc[i].ident, t.Conc[i].squan, t.Conc[i].spuni,
    t.Conc[i].staxb, t.Conc[i].staxr, t.Conc[i].Taxam)
  return line
}

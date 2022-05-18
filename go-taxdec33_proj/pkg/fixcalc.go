package fixcalc

import ut "bar8tl/p/rblib"
import "log"

func (t *It_conc) EvaluayEjecutaCorreccion(g *It_invo, m *It_matr) {
  t.calculaDesviacionyHolguras(g, m)
  if len(g.invo) > 0 {
    t.distribuyeDesviacion(g, m)
  }
}

func (t *It_conc) calculaDesviacionyHolguras(g *It_invo, m *It_matr) {
  var negfd, posfd, negfs, posfs float64
  for i, _ := range t.Conc {
    taxam := t.Conc[i].Taxam
    taxlw := t.Conc[i].taxlw
    taxup := t.Conc[i].taxup
    switch {
      case taxam > taxup : 
        t.Conc[i].devtn = taxam - taxup
        negfd = 0
        posfd = 1
      case taxam < taxlw : 
        t.Conc[i].devtn = taxam - taxlw
        negfd = 1
        posfd = 0
      case taxam < taxup : 
        t.Conc[i].slack = taxup - taxam
        negfs = 0
        posfs = 1
      case taxam > taxlw : 
        t.Conc[i].slack = taxlw - taxam
        negfs = 1
        posfs = 0
    }
    if t.Conc[i].devtn != 0 {
      g.invo = append(g.invo, t_invo{t.Conc[i].ident, i})
    }

    mtrlFound := false
    for j := 0; j < len(m.matr) && !mtrlFound; j++ {
      if m.matr[j].ident == t.Conc[i].ident {
        m.matr[j].items += 1
        m.matr[j].negsl += negfs * t.Conc[i].slack
        m.matr[j].possl += posfs * t.Conc[i].slack
        m.matr[j].negdv += negfd * t.Conc[i].devtn
        m.matr[j].posdv += posfd * t.Conc[i].devtn
        mtrlFound = true
      }
    }
    if !mtrlFound {
      m.matr = append(m.matr, t_matr{t.Conc[i].ident, 1, negfs*t.Conc[i].slack,
        posfs*t.Conc[i].slack, negfd*t.Conc[i].devtn, posfd*t.Conc[i].devtn})
    }
  }
}

func (t *It_conc) distribuyeDesviacion(g *It_invo, m *It_matr) {
  t.formatTableDecimals(g, m)
  for k := 0; k < len(g.invo); k++ {
    for j := 0; j < len(m.matr) && !g.resolved; j++ {
      if m.matr[j].ident == g.invo[k].mtwdv {
        switch {
          case m.matr[j].posdv != 0 :
            g.targt = ut.Round(m.matr[j].posdv, g.DECIMALS)
          case m.matr[j].negdv != 0 :
            g.targt = ut.Round(m.matr[j].negdv, g.DECIMALS)
        }
        t.calculaDistribucion(j, k, g, m)
      }
    }
    log.Println(g.targt)
    if g.targt != 0.0 {
      log.Printf("Aviso: La desviacion no se pudo remover en solo un grupo\n")
      for j := 0; j < len(m.matr) && !g.resolved; j++ {
        if m.matr[j].ident != g.invo[k].mtwdv {
          t.calculaDistribucion(j, k, g, m)
        }
      }
    }
  }
  t.checkConsistency(g)
}

func (t *It_conc) calculaDistribucion(j, k int, g *It_invo, m *It_matr) {
  for i := len(t.Conc) - 1; i >= 0 && !g.resolved; i-- {
    if t.Conc[i].ident == m.matr[j].ident {
      if (g.targt > 0.0 && t.Conc[i].slack > 0.0) ||
        (g.targt < 0.0 && t.Conc[i].slack < 0.0) {
        t.Conc[g.invo[k].i].Taxam -= t.Conc[i].slack
        t.Conc[i].Taxam += t.Conc[i].slack
        g.targt         -= t.Conc[i].slack
        t.Conc[i].slack -= t.Conc[i].slack
        g.targt = ut.Round(g.targt, g.DECIMALS)
        log.Println(g.targt)
        if g.targt == 0.0 {
          g.resolved = true
        }
      }
    }
  }
}

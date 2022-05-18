package main

import ut "bar8tl/p/rblib"
import "log"

func main() {
  cargaTablaConceptos()
  for i, _ := range conc {
    calcDesviacionyHolguras(i)
  }
  if len(invo) > 0 {
    for k := 0; k < len(invo); k++ {
      distribDesviacion(k)
    }
    descargaResultado()
  }
}

func calcDesviacionyHolguras(i int) {
  var negfd float64
  var posfd float64
  var negfs float64
  var posfs float64
  taxam := conc[i].taxam
  taxlw := conc[i].taxlw
  taxup := conc[i].taxup
  switch {
  case taxam > taxup:
    conc[i].devtn = taxam - taxup
    negfd = 0
    posfd = 1
  case taxam < taxlw:
    conc[i].devtn = taxam - taxlw
    negfd = 1
    posfd = 0
  case taxam < taxup:
    conc[i].slack = taxup - taxam
    negfs = 0
    posfs = 1
  case taxam > taxlw:
    conc[i].slack = taxlw - taxam
    negfs = 1
    posfs = 0
  }
  if conc[i].devtn != 0 {
    invo = append(invo, t_invo{conc[i].ident, i})
  }

  mtrlFound := false
  for j := 0; j < len(matr) && !mtrlFound; j++ {
    if matr[j].ident == conc[i].ident {
      matr[j].items += 1
      matr[j].negsl += negfs * conc[i].slack
      matr[j].possl += posfs * conc[i].slack
      matr[j].negdv += negfd * conc[i].devtn
      matr[j].posdv += posfd * conc[i].devtn
      mtrlFound = true
    }
  }
  if !mtrlFound {
    matr = append(matr, t_matr{conc[i].ident, 1, negfs * conc[i].slack,
      posfs * conc[i].slack, negfd * conc[i].devtn, posfd * conc[i].devtn})
  }
}

func distribDesviacion(k int) {
  formatTablesDecimals()
  for j := 0; j < len(matr) && !resolved; j++ {
    if matr[j].ident == invo[k].mtwdv {
      switch {
      case matr[j].posdv != 0:
        targt = ut.Round(matr[j].posdv, DECIMALS)
      case matr[j].negdv != 0:
        targt = ut.Round(matr[j].negdv, DECIMALS)
      }
      calculaDistrib(j, k)
    }
  }
  log.Printf("La desviacion no se pudo remover en solo un grupo\n")
  if targt != 0 {
    for j := 0; j < len(matr) && !resolved; j++ {
      if matr[j].ident != invo[k].mtwdv {
        calculaDistrib(j, k)
      }
    }
  }
}

func calculaDistrib(j, k int) {
  for i := len(conc) - 1; i >= 0 && !resolved; i-- {
    if conc[i].ident == matr[j].ident {
      if (targt > 0.0 && conc[i].slack > 0.0) ||
         (targt < 0.0 && conc[i].slack < 0.0) {
        conc[invo[k].i].taxam -= conc[i].slack
        conc[i].taxam += conc[i].slack
        targt         -= conc[i].slack
        conc[i].slack -= conc[i].slack
        targt          = ut.Round(targt, DECIMALS)
        if targt == 0.0 {
          resolved = true
        }
      }
    }
  }
}

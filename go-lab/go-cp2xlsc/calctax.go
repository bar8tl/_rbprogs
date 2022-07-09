// calctax.go - New Pagos 2.0 tax fields calculation functions
// 2022-05-17 BAR8TL Version1.0 - In progress
package main

import rb "bar8tl/p/cp2xlsc"
import ut "bar8tl/p/rblib"
import "math"

var ivs [][91]string
var ivf [][91]float64

type Calctax_tp struct {
  fiv [7]bool
}

func NewCalctax(s rb.Settings_tp) *Calctax_tp {
  var c Calctax_tp
  for i, _ := range c.fiv {
    c.fiv[i] = true
  }
  return &c
}

func (c *Calctax_tp) ResetPaymentData() {
  for i, _ := range pms {
    pms[i] = ""
  }
  for i, _ := range pmf {
    pmf[i] = 0.0
  }
  for i, _ := range c.fiv {
    c.fiv[i] = true
  }
  ivs = nil
  ivf = nil
}

func (c *Calctax_tp) StorePayment() {
  for i, _ := range rds {
    pms[i] = rds[i]
  }
  for i, _ := range rdf {
    pmf[i] = rdf[i]
  }
  if rds[m["documentCurrency"]] == "MXN" {
    pmf[m["effExchangeRate"]] = 1.0
  } else {
    pmf[m["effExchangeRate"]] = rdf[m["effExchangeRate"]]
  }
  pmf[m["amountDocCurr"]] = rdf[m["amountDocCurr"]] * pmf[m["effExchangeRate"]]
  pmf[m["amountDocCurr"]] = round(pmf[m["amountDocCurr"]])
}

func (c *Calctax_tp) StoreDocRel() {
  for i, _ := range rds {
    drs[i] = rds[i]
  }
  for i, _ := range rdf {
    drf[i] = rdf[i]
  }
  drf[m["taxTrasladoBase"]]       = drf[m["trasladoBaseDR"]]
  drs[m["taxTrasladoImpuesto"]]   = drs[m["trasladoImpuestoDR"]]
  drs[m["taxTrasladoTipoFactor"]] = drs[m["trasladoTipoFactorDR"]]
  drf[m["taxTrasladoTasaOCuota"]] = drf[m["trasladoTasaOCuotaDR"]]
  drf[m["taxTrasladoImporte"]]    = drf[m["trasladoImporteDR"]]
  drf[m["taxRetncionBase"]]       = drf[m["retncionBaseDR"]]
  drs[m["taxRetncionImpuesto"]]   = drs[m["retncionImpuestoDR"]]
  drs[m["taxRetncionTipoFactor"]] = drs[m["retncionTipoFactorDR"]]
  drf[m["taxRetncionTasaOCuota"]] = drf[m["retncionTasaOCuotaDR"]]
  drf[m["taxRetncionImporte"]]    = drf[m["retncionImporteDR"]]
  ivs = append(ivs, drs)
  ivf = append(ivf, drf)
  // Reset cumulative amounts of One-taxcode payments
  if c.fiv[n["firstInvoice"]] {
    pms[m["taxTrasladoImpuesto"]]   = drs[m["taxTrasladoImpuesto"]]
    pms[m["taxTrasladoTipoFactor"]] = drs[m["taxTrasladoTipoFactor"]]
    pms[m["taxTrasladoTasaOCuota"]] = drs[m["taxTrasladoTasaOCuota"]]
    pms[m["taxRetncionImpuesto"]]   = drs[m["taxRetncionImpuesto"]]
    pms[m["taxRetncionTipoFactor"]] = drs[m["taxRetncionTipoFactor"]]
    pms[m["taxRetncionTasaOCuota"]] = drs[m["taxRetncionTasaOCuota"]]
    c.fiv[n["firstInvoice"]] = false
  }
  // Reset cumulative amounts of Multiple-taxcode payments
  if drf[m["taxTrasladoTasaOCuota"]] == 0.16 {
    if c.fiv[n["firstInvoTraslIVA16"]] {
      pms[m["trasladoImpuestoPIVA16"]]   = drs[m["taxTrasladoImpuesto"]]
      pms[m["trasladoTipoFactorPIVA16"]] = drs[m["taxTrasladoTipoFactor"]]
      pms[m["trasladoTasaOCuotaPIVA16"]] = drs[m["taxTrasladoTasaOCuota"]]
      c.fiv[n["firstInvoTraslIVA16"]] = false
    }
  }
  if drf[m["taxRetncionTasaOCuota"]] == 0.16 {
    if c.fiv[n["firstInvoRetenIVA16"]] {
      pms[m["retncionImpuestoPIVA16"]]   = drs[m["taxRetncionImpuesto"]]
      pms[m["retncionTipoFactorPIVA16"]] = drs[m["taxRetncionTipoFactor"]]
      pms[m["retncionTasaOCuotaPIVA16"]] = drs[m["taxRetncionTasaOCuota"]]
      c.fiv[n["firstInvoRetenIVA16"]] = false
    }
  }
  if drf[m["taxTrasladoTasaOCuota"]] == 0.08 {
    if c.fiv[n["firstInvoTraslIVA8"]] {
      pms[m["trasladoImpuestoPIVA8"]]    = drs[m["taxTrasladoImpuesto"]]
      pms[m["trasladoTipoFactorPIVA8"]]  = drs[m["taxTrasladoTipoFactor"]]
      pms[m["trasladoTasaOCuotaPIVA8"]]  = drs[m["taxTrasladoTasaOCuota"]]
      c.fiv[n["firstInvoTraslIVA8"]] = false
    }
  }
  if drf[m["taxRetncionTasaOCuota"]] == 0.08 {
    if c.fiv[n["firstInvoRetenIVA8"]] {
      pms[m["retncionImpuestoPIVA8"]]    = drs[m["taxRetncionImpuesto"]]
      pms[m["retncionTipoFactorPIVA8"]]  = drs[m["taxRetncionTipoFactor"]]
      pms[m["retncionTasaOCuotaPIVA8"]]  = drs[m["taxRetncionTasaOCuota"]]
      c.fiv[n["firstInvoRetenIVA8"]] = false
    }
  }
  if drf[m["taxTrasladoTasaOCuota"]] == 0.0 {
    if c.fiv[n["firstInvoTraslIVA0"]] {
      pms[m["trasladoImpuestoPIVA0"]]    = drs[m["taxTrasladoImpuesto"]]
      pms[m["trasladoTipoFactorPIVA0"]]  = drs[m["taxTrasladoTipoFactor"]]
      pms[m["trasladoTasaOCuotaPIVA0"]]  = drs[m["taxTrasladoTasaOCuota"]]
      c.fiv[n["firstInvoTraslIVA0"]] = false
    }
  }
  if drf[m["taxRetncionTasaOCuota"]] == 0.0 {
    if c.fiv[n["firstInvoRetenIVA0"]] {
      pms[m["retncionImpuestoPIVA0"]]    = drs[m["taxRetncionImpuesto"]]
      pms[m["retncionTipoFactorPIVA0"]]  = drs[m["taxRetncionTipoFactor"]]
      pms[m["retncionTasaOCuotaPIVA0"]]  = drs[m["taxRetncionTasaOCuota"]]
      c.fiv[n["firstInvoRetenIVA0"]] = false
    }
  }
}

func (c *Calctax_tp) FetchPaymentLine(w *Writer_tp) *Calctax_tp {
  importePagoCalc := pmf[m["taxTrasladoBase"]]    +
                     pmf[m["taxTrasladoImporte"]] -
                     pmf[m["taxRetncionImporte"]]
  importePagoCalc  = ut.Round(importePagoCalc, 2)
  importePago := ut.Round(pmf[m["amountDocCurr"]], 2)
  pmf[m["difImportePago1"]] = -1.0 * importePago - importePagoCalc
  if math.Abs(pmf[m["difImportePago1"]]) < 0.0000015 {
    pmf[m["difImportePago1"]] = 0.0
  }

  trasladoBaseP := 0.0
  if pmf[m["trasladoBasePIVA16"]] != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = pmf[m["trasladoBasePIVA16"]]
    }
  }
  if pmf[m["trasladoBasePIVA8"]]  != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = pmf[m["trasladoBasePIVA8"]]
    }
  }
  if pmf[m["trasladoBasePIVA0"]]  != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = pmf[m["trasladoBasePIVA0"]]
    }
  }
  trasladoImporteP := pmf[m["trasladoImportePIVA16"]] +
                      pmf[m["trasladoImportePIVA8" ]] +
                      pmf[m["trasladoImportePIVA0" ]]
  retncionImporteP := pmf[m["retncionImportePIVA16"]] +
                      pmf[m["retncionImportePIVA8" ]] +
                      pmf[m["retncionImportePIVA0" ]]
  importePagoCalc = trasladoBaseP + trasladoImporteP - retncionImporteP
  importePago = ut.Round(pmf[m["amountDocCurr"]], 2)
  pmf[m["difImportePago3"]] = -1.0 * importePago - importePagoCalc
  if math.Abs(pmf[m["difImportePago3"]]) < 0.0000015 {
    pmf[m["difImportePago3"]] = 0.0
  }
  amountDocCurr := ut.Round(pmf[m["amountDocCurr"]], 2)
  pmf[m["difMontoTotalPagos"]] = -1.0 * amountDocCurr - pmf[m["montoTotalPagos"]]
  if math.Abs(pmf[m["difMontoTotalPagos"]]) < 0.0000015 {
    pmf[m["difMontoTotalPagos"]] = 0.0
  }
  w.PrintPaymentLine()
  return c
}

func (c *Calctax_tp) FetchInvoiceLines(w *Writer_tp) *Calctax_tp {
  for i, _ := range ivs {
    importePagoCalc := ivf[i][m["taxTrasladoBase"   ]] +
                       ivf[i][m["taxTrasladoImporte"]] -
                       ivf[i][m["taxRetncionImporte"]]
    importePago := ut.Round(ivf[i][m["importePago"]], 2)
    pmf[m["difImportePago1"]] = importePago - importePagoCalc
    if math.Abs(pmf[m["difImportePago1"]]) < 0.0000015 {
      pmf[m["difImportePago1"]] = 0.0
    }
    w.PrintInvoiceLine(ivs[i], ivf[i])
  }
  return c
}

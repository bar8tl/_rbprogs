// cpm2xlsc.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields. Core Logic
package main

import rb "bar8tl/p/cp2xlsc"
import ut "bar8tl/p/rblib"

var firstLine bool = true

// Logic for Payments
func ProcessPaymentLine(c *rb.Calctax_tp, lineExcel rb.Reader_tp,
  wtr *rb.Writer_tp) {
  if firstLine {
    firstLine = false
  } else {
    c.FetchPaymentLine(wtr).FetchInvoiceLines(wtr).ResetPaymentData()
  }
  c.StorePayment(lineExcel)
}

// Logic for Invoices
func ProcessInvoiceLine(c *rb.Calctax_tp, lineExcel rb.Reader_tp) {
  var docrel   rb.Docrel_tp
  var tax, wht int
  var taxTasa, whtTasa float64
  var importePago      float64
  switch lineExcel.Src.TaxCode {
    case "A2", "B2", "CI", "CF" : tax = 16; wht =  0
    case "A5", "B5"             : tax = 16; wht = 16
    case "AA", "BA", "VA"       : tax =  8; wht =  0
    case "AB", "BB"             : tax =  8; wht =  8
    case "A0", "B0", "CG", "V0" : tax =  0; wht =  0
    case "AE", "BE" : tax = 16; wht =  8 // partial wht discontinued, no handled
    case "AF", "BF" : tax =  8; wht =  3 // partial wht discontinued, no handled
  }
  taxTasa = float64(tax) / float64(100)
  whtTasa = float64(wht) / float64(100)
  if lineExcel.Src.DocumentCurrency != "MXN" {
    importePago = lineExcel.ImportePago * c.EffExchangeRate
  } else {
    importePago = lineExcel.ImportePago
  }
  importePago = ut.Round(importePago, 6)
  docrel.ObjetoImpDR                  = rb.OBJETOIMPUESTO
  docrel.TrasladoDR.BaseDR            = ut.Round(importePago /
                                          (1.0 + taxTasa - whtTasa), rb.DEC)
  docrel.TrasladoDR.ImpuestoDR        = rb.IMPUESTO
  docrel.TrasladoDR.TipoFactorDR      = rb.TIPOFACTOR
  docrel.TrasladoDR.TasaOCuotaDR      = taxTasa
  docrel.TrasladoDR.ImporteDR         = ut.Round(docrel.TrasladoDR.BaseDR *
                                          taxTasa, rb.DEC)
  c.ImpuestosP.TrasladoP.BaseP       += docrel.TrasladoDR.BaseDR
  c.ImpuestosP.TrasladoP.ImporteP    += docrel.TrasladoDR.ImporteDR
  c.Totales.MontoTotalPagos          += importePago
  if tax == 16 {
    c.TaxPIVA16.TrasladoP.BaseP      += docrel.TrasladoDR.BaseDR
    c.TaxPIVA16.TrasladoP.ImporteP   += docrel.TrasladoDR.ImporteDR
    c.Totales.TrasladosBaseIVA16     += docrel.TrasladoDR.BaseDR
    c.Totales.TrasladosImpuestoIVA16 += docrel.TrasladoDR.ImporteDR
  } else if tax == 8 {
    c.TaxPIVA8.TrasladoP.BaseP       += docrel.TrasladoDR.BaseDR
    c.TaxPIVA8.TrasladoP.ImporteP    += docrel.TrasladoDR.ImporteDR
    c.Totales.TrasladosBaseIVA8      += docrel.TrasladoDR.BaseDR
    c.Totales.TrasladosImpuestoIVA8  += docrel.TrasladoDR.ImporteDR
  } else if tax == 0 {
    c.TaxPIVA0.TrasladoP.BaseP       += docrel.TrasladoDR.BaseDR
    c.TaxPIVA0.TrasladoP.ImporteP    += docrel.TrasladoDR.ImporteDR
    c.Totales.TrasladosBaseIVA0      += docrel.TrasladoDR.BaseDR
    c.Totales.TrasladosImpuestoIVA0  += docrel.TrasladoDR.ImporteDR
  }
  docrel.RetncionDR = rb.TaxesDR_tp{0.0, "", "", 0.0, 0.0}
  if wht != 0 {
    docrel.RetncionDR.BaseDR          = docrel.TrasladoDR.BaseDR
    docrel.RetncionDR.ImpuestoDR      = rb.IMPUESTO
    docrel.RetncionDR.TipoFactorDR    = rb.TIPOFACTOR
    docrel.RetncionDR.TasaOCuotaDR    = whtTasa
    docrel.RetncionDR.ImporteDR       = ut.Round(docrel.RetncionDR.BaseDR *
                                          whtTasa, rb.DEC)
    c.ImpuestosP.RetncionP.BaseP     += docrel.RetncionDR.BaseDR
    c.ImpuestosP.RetncionP.ImporteP  += docrel.RetncionDR.ImporteDR
    c.Totales.RetencionesIVA         += docrel.RetncionDR.ImporteDR
    if wht == 16 {
      c.TaxPIVA16.RetncionP.BaseP    += docrel.RetncionDR.BaseDR
      c.TaxPIVA16.RetncionP.ImporteP += docrel.RetncionDR.ImporteDR
    } else if wht == 8 {
      c.TaxPIVA8.RetncionP.BaseP     += docrel.RetncionDR.BaseDR
      c.TaxPIVA8.RetncionP.ImporteP  += docrel.RetncionDR.ImporteDR
    }
  }
  c.StoreDocRel(lineExcel, docrel)
}

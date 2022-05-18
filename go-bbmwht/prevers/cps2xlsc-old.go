// cps2xlsc.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields. Version for One-TaxCode Invoices per Payment.
package main

import rb "bar8tl/p/bbmwht"
import ut "bar8tl/p/rblib"
import "log"

// Core logic for Payments -----------------------------------------------------
func processPaymentLine(lineExcel rb.Line_tp) { // for lines type DZ payment
  if firstLine {
    firstLine = false
  } else {
    p.WritePaymentLine().WriteInvoiceLines().ClearPaymentData()
  }
  p.PaymentData = lineExcel
}
// Core logic for Invoices -----------------------------------------------------
func processInvoiceLine(lineExcel rb.Line_tp) { // for lines type RV invoice
  var docrel   rb.Docrel_tp
  var tax, wht int
  var taxTasa, whtTasa float64
  switch lineExcel.TaxCode {
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
  docrel.ObjetoImpDR                  = rb.OBJETOIMPUESTO
  docrel.TrasladoDR.BaseDR            = ut.Round(p.ImportePago /
                                          (1.0 + taxTasa - whtTasa), rb.DEC)
  docrel.TrasladoDR.ImpuestoDR        = rb.IMPUESTO
  docrel.TrasladoDR.TipoFactorDR      = rb.TIPOFACTOR
  docrel.TrasladoDR.TasaOCuotaDR      = taxTasa
  docrel.TrasladoDR.ImporteDR         = ut.Round(docrel.TrasladoDR.BaseDR *
                                          taxTasa, rb.DEC)
  p.ImpuestosP.TrasladoP.BaseP       += docrel.TrasladoDR.BaseDR
  p.ImpuestosP.TrasladoP.ImporteP    += docrel.TrasladoDR.ImporteDR
  p.Totales.MontoTotalPagos          += p.ImportePago
  if tax == 16 {
    p.Totales.TrasladosBaseIVA16     += docrel.TrasladoDR.BaseDR
    p.Totales.TrasladosImpuestoIVA16 += docrel.TrasladoDR.ImporteDR
  } else if tax == 8 {
    p.Totales.TrasladosBaseIVA8      += docrel.TrasladoDR.BaseDR
    p.Totales.TrasladosImpuestoIVA8  += docrel.TrasladoDR.ImporteDR
  } else if tax == 0 {
    p.Totales.TrasladosBaseIVA0      += docrel.TrasladoDR.BaseDR
    p.Totales.TrasladosImpuestoIVA0  += docrel.TrasladoDR.ImporteDR
  }
  docrel.RetncionDR = rb.TaxesDR_tp{0.00, "", "", 0.0, 0.0}
  if wht != 0 {
    docrel.RetncionDR.BaseDR          = docrel.TrasladoDR.BaseDR
    docrel.RetncionDR.ImpuestoDR      = rb.IMPUESTO
    docrel.RetncionDR.TipoFactorDR    = rb.TIPOFACTOR
    docrel.RetncionDR.TasaOCuotaDR    = whtTasa
    docrel.RetncionDR.ImporteDR       = ut.Round(docrel.RetncionDR.BaseDR *
                                          whtTasa, rb.DEC)
    p.ImpuestosP.RetncionP.BaseP     += docrel.RetncionDR.BaseDR
    p.ImpuestosP.RetncionP.ImporteP  += docrel.RetncionDR.ImporteDR
    p.Totales.RetencionesIVA         += docrel.RetncionDR.ImporteDR
  }
  p.StoreDocRel(lineExcel, docrel)
}
// Entry point -----------------------------------------------------------------
var firstLine bool = true
var p *rb.Stools_tp

func main() {
  p = rb.NewStools()
  p.OpenInpExcel("edicoms1.0_dc.xlsx")
  rows, err := p.F.GetRows("edicom")
  if err != nil {
    log.Fatal(err)
  }
  p.CreateOutExcel("Sheet1")
  for _, row := range rows {
    lineExcel := p.GetLineFields(row)
    if lineExcel.DocumentType == "Document Type" {
      p.WriteTitle(lineExcel)
    } else if lineExcel.DocumentType == "DZ" {
      processPaymentLine(lineExcel)
    } else if lineExcel.DocumentType == "RV" {
      processInvoiceLine(lineExcel)
    }
  }
  p.WritePaymentLine().WriteInvoiceLines().WriteOutExcel("edicoms2.0_dc.xlsx")
}
// ----------------------------- end of file -----------------------------------

// corelogic.go - Extend Pagos1.0 EDICOM-file with Pagos2.0 fields (Core Logic)
// 2022-05-17 BAR8TL Version1.0 - Released
package main

import ut "bar8tl/p/rblib"

var firstLine bool = true

// Logic for Payments
func ProcessPaymentLine(c *Calctax_tp, wtr *Writer_tp) {
  if firstLine {
    firstLine = false
  } else {
    c.FetchPaymentLine(wtr).FetchInvoiceLines(wtr).ResetPaymentData()
  }
  c.StorePayment()
}

// Logic for Invoices
func ProcessInvoiceLine(c *Calctax_tp) {
  var tax, wht int
  var taxTasa, whtTasa float64
  var importePago      float64
  switch rds[m["taxCode"]] {
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
  if pms[m["documentCurrency"]] == "MXN" {
    if rds[m["documentCurrency"]] == "MXN" {
      importePago = rdf[m["importePago"]]
    } else {
      importePago = rdf[m["importePago"]] / rdf[m["effExchangeRate"]]
    }
  } else {
    importePago = rdf[m["importePago"]] * pmf[m["effExchangeRate"]]
  }
  importePago = ut.Round(importePago, DEC)
  drs[m["objetoImpuesto"]]             = OBJETOIMP
  drf[m["trasladoBaseDR"]]             = round(importePago /
                                           (1.0 + taxTasa - whtTasa))
  drs[m["trasladoImpuestoDR"]]         = IMPUESTO
  drs[m["trasladoTipoFactorDR"]]       = TIPOFACTOR
  drf[m["trasladoTasaOCuotaDR"]]       = taxTasa
  drf[m["trasladoImporteDR"]]          = round(drf[m["trasladoBaseDR"]] *
                                           taxTasa)
  pmf[m["taxTrasladoBase"]]           += drf[m["trasladoBaseDR"]]
  pmf[m["taxTrasladoImporte"]]        += drf[m["trasladoImporteDR"]]
  pmf[m["montoTotalPagos"]]           += importePago
  if tax == 16 {
    pmf[m["trasladoBasePIVA16"]]      += drf[m["trasladoBaseDR"]]
    pmf[m["trasladoImportePIVA16"]]   += drf[m["trasladoImporteDR"]]
    pmf[m["trasladosBaseIVA16"]]      += drf[m["trasladoBaseDR"]]
    pmf[m["trasladosImpuestoIVA16"]]  += drf[m["trasladoImporteDR"]]
  } else if tax == 8 {
    pmf[m["trasladoBasePIVA8"]]       += drf[m["trasladoBaseDR"]]
    pmf[m["trasladoImportePIVA8"]]    += drf[m["trasladoImporteDR"]]
    pmf[m["trasladosBaseIVA8"]]       += drf[m["trasladoBaseDR"]]
    pmf[m["trasladosImpuestoIVA8"]]   += drf[m["trasladoImporteDR"]]
  } else if tax == 0 {
    pmf[m["trasladoBasePIVA0"]]       += drf[m["trasladoBaseDR"]]
    pmf[m["trasladoImportePIVA0"]]    += drf[m["trasladoImporteDR"]]
    pmf[m["trasladosBaseIVA0"]]       += drf[m["trasladoBaseDR"]]
    pmf[m["trasladosImpuestoIVA0"]]   += drf[m["trasladoImporteDR"]]
  }
  drf[m["retncionBaseDR"]]             = 0.0
  drs[m["retncionImpuestoDR"]]         = ""
  drs[m["retncionTipoFactorDR"]]       = ""
  drf[m["retncionTasaOCuotaDR"]]       = 0.0
  drf[m["retncionImporteDR"]]          = 0.0
  if wht != 0 {
    drf[m["retncionBaseDR"]]           = drf[m["trasladoBaseDR"]]
    drs[m["retncionImpuestoDR"]]       = IMPUESTO
    drs[m["retncionTipoFactorDR"]]     = TIPOFACTOR
    drf[m["retncionTasaOCuotaDR"]]     = whtTasa
    drf[m["retncionImporteDR"]]        = round(drf[m["retncionBaseDR"]] *
                                           whtTasa)
    pmf[m["taxRetncionBase"]]         += drf[m["retncionBaseDR"]]
    pmf[m["taxRetncionImporte"]]      += drf[m["retncionImporteDR"]]
    pmf[m["retencionesIVA"]]          += drf[m["retncionImporteDR"]]
    if wht == 16 {
      pmf[m["retncionBasePIVA16"]]    += drf[m["retncionBaseDR"]]
      pmf[m["retncionImportePIVA16"]] += drf[m["retncionImporteDR"]]
    } else if wht == 8 {
      pmf[m["retncionBasePIVA8"]]     += drf[m["retncionBaseDR"]]
      pmf[m["retncionImportePIVA8"]]  += drf[m["retncionImporteDR"]]
    }
  }
  c.StoreDocRel()
}

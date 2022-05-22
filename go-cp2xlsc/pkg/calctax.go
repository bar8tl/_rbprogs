// cp2xlsp.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields.
package cp2xlsc

import ut "bar8tl/p/rblib"
import "math"
import "strconv"

var IMPUESTO             string
var TIPOFACTOR           string
var OBJETOIMPUESTO       string
var DEC                  int

type linv_tp struct {
  src                    Line_tp
  docrel                 Docrel_tp
}

type Totales_tp struct {
  RetencionesIVA         float64
  TrasladosBaseIVA16     float64
  TrasladosImpuestoIVA16 float64
  TrasladosBaseIVA8      float64
  TrasladosImpuestoIVA8  float64
  TrasladosBaseIVA0      float64
  TrasladosImpuestoIVA0  float64
  MontoTotalPagos        float64
}

type Docrel_tp struct {
  ObjetoImpDR            string
  TrasladoDR             TaxesDR_tp
  RetncionDR             TaxesDR_tp
}

type TaxesDR_tp struct {
  BaseDR                 float64
  ImpuestoDR             string
  TipoFactorDR           string
  TasaOCuotaDR           float64
  ImporteDR              float64
}

type Payment_tp struct {
  TrasladoP              TaxesP_tp
  RetncionP              TaxesP_tp
}

type TaxesP_tp struct {
  BaseP                  float64
  ImpuestoP              string
  TipoFactorP            string
  TasaOCuotaP            float64
  ImporteP               float64
}

type Calctax_tp struct {
  AmountDocCurr          float64
  EffExchangeRate        float64
  PaymentData            Reader_tp
  Totales                Totales_tp
  ImpuestosP             Payment_tp
  OneTaxPaym             Payment_tp
  TaxPIVA16              Payment_tp
  TaxPIVA8               Payment_tp
  TaxPIVA0               Payment_tp
  invoices             []linv_tp
  firstInvoice           bool
  firstInvoTraslIVA16    bool
  firstInvoRetenIVA16    bool
  firstInvoTraslIVA8     bool
  firstInvoRetenIVA8     bool
  firstInvoTraslIVA0     bool
  firstInvoRetenIVA0     bool
}

func NewCalctax(s Settings_tp) *Calctax_tp {
  var c Calctax_tp
  c.firstInvoice        = true
  c.firstInvoTraslIVA16 = true
  c.firstInvoRetenIVA16 = true
  c.firstInvoTraslIVA8  = true
  c.firstInvoRetenIVA8  = true
  c.firstInvoTraslIVA0  = true
  c.firstInvoRetenIVA0  = true
  IMPUESTO              = s.Konst.IMPUESTO
  TIPOFACTOR            = s.Konst.TIPOFACTOR
  OBJETOIMPUESTO        = s.Const.Taxbl
  DEC                   = s.Konst.DEC
  return &c
}

func (c *Calctax_tp) ResetPaymentData() {
  c.resetCommonPaymentData().reset1TaxPaymentData().reset3TaxPaymentData()
}

func (c *Calctax_tp) StorePayment(lin Reader_tp) {
  c.PaymentData = lin
  if lin.Src.DocumentCurrency != "MXN" {
    c.EffExchangeRate = lin.EffExchangeRate
  } else {
    c.EffExchangeRate = 1.0
  }
  c.AmountDocCurr = lin.AmountDocCurr * c.EffExchangeRate
  c.AmountDocCurr = ut.Round(c.AmountDocCurr, 6)
}

func (c *Calctax_tp) StoreDocRel(lin Reader_tp, inv Docrel_tp) {
  var linv linv_tp
  linv.src = lin.Src
  linv.docrel.ObjetoImpDR               = inv.ObjetoImpDR
  // Setup data of each invoice for internal invoices'-payment table
  linv.docrel.TrasladoDR.BaseDR         = inv.TrasladoDR.BaseDR
  linv.docrel.TrasladoDR.ImpuestoDR     = inv.TrasladoDR.ImpuestoDR
  linv.docrel.TrasladoDR.TipoFactorDR   = inv.TrasladoDR.TipoFactorDR
  linv.docrel.TrasladoDR.TasaOCuotaDR   = inv.TrasladoDR.TasaOCuotaDR
  linv.docrel.TrasladoDR.ImporteDR      = inv.TrasladoDR.ImporteDR
  linv.docrel.RetncionDR.BaseDR         = inv.RetncionDR.BaseDR
  linv.docrel.RetncionDR.ImpuestoDR     = inv.RetncionDR.ImpuestoDR
  linv.docrel.RetncionDR.TipoFactorDR   = inv.RetncionDR.TipoFactorDR
  linv.docrel.RetncionDR.TasaOCuotaDR   = inv.RetncionDR.TasaOCuotaDR
  linv.docrel.RetncionDR.ImporteDR      = inv.RetncionDR.ImporteDR
  c.invoices = append(c.invoices, linv_tp{linv.src, linv.docrel})
  // Reset cumulative amounts of One-taxcode payments
  if c.firstInvoice {
    c.OneTaxPaym.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
    c.OneTaxPaym.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
    c.OneTaxPaym.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
    c.OneTaxPaym.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
    c.OneTaxPaym.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
    c.OneTaxPaym.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
    c.firstInvoice = false
  }
  // Reset cumulative amounts of Multiple-taxcode payments
  if inv.TrasladoDR.TasaOCuotaDR == 0.16 {
    if c.firstInvoTraslIVA16 {
      c.TaxPIVA16.TrasladoP.ImpuestoP   = inv.TrasladoDR.ImpuestoDR
      c.TaxPIVA16.TrasladoP.TipoFactorP = inv.TrasladoDR.TipoFactorDR
      c.TaxPIVA16.TrasladoP.TasaOCuotaP = inv.TrasladoDR.TasaOCuotaDR
      c.firstInvoTraslIVA16 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.16 {
    if c.firstInvoRetenIVA16 {
      c.TaxPIVA16.RetncionP.ImpuestoP   = inv.RetncionDR.ImpuestoDR
      c.TaxPIVA16.RetncionP.TipoFactorP = inv.RetncionDR.TipoFactorDR
      c.TaxPIVA16.RetncionP.TasaOCuotaP = inv.RetncionDR.TasaOCuotaDR
      c.firstInvoRetenIVA16 = false
    }
  }
  if inv.TrasladoDR.TasaOCuotaDR == 0.08 {
    if c.firstInvoTraslIVA8 {
      c.TaxPIVA8.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      c.TaxPIVA8.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      c.TaxPIVA8.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      c.firstInvoTraslIVA8 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.08 {
    if c.firstInvoRetenIVA8 {
      c.TaxPIVA8.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      c.TaxPIVA8.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      c.TaxPIVA8.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      c.firstInvoRetenIVA8 = false
    }
  }
  if inv.TrasladoDR.TasaOCuotaDR == 0.0 {
    if c.firstInvoTraslIVA0 {
      c.TaxPIVA0.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      c.TaxPIVA0.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      c.TaxPIVA0.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      c.firstInvoTraslIVA0 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.0 {
    if c.firstInvoRetenIVA0 {
      c.TaxPIVA0.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      c.TaxPIVA0.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      c.TaxPIVA0.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      c.firstInvoRetenIVA0 = false
    }
  }
}

func (c *Calctax_tp) FetchPaymentLine(w *Writer_tp) *Calctax_tp {
  // Setup output values of source fields
  ps[m["companyCode"]]             = c.PaymentData.Src.companyCode
  ps[m["customer"]]                = c.PaymentData.Src.customer
  ps[m["documentNumber"]]          = c.PaymentData.Src.documentNumber
  ps[m["documentType"]]            = c.PaymentData.Src.DocumentType
  ps[m["paymentDateTime"]]         = c.PaymentData.Src.paymentDateTime
  ps[m["clearingDocument"]]        = c.PaymentData.Src.ClearingDocument
  ps[m["amountDocCurr"]]           = c.PaymentData.Src.AmountDocCurr
  ps[m["documentCurrency"]]        = c.PaymentData.Src.DocumentCurrency
  ps[m["effExchangeRate"]]         = c.PaymentData.Src.EffExchangeRate
  ps[m["assignment"]]              = c.PaymentData.Src.assignment
  ps[m["formaPago"]]               = c.PaymentData.Src.formaPago
  ps[m["noParcialidad"]]           = c.PaymentData.Src.noParcialidad
  ps[m["importeSaldoAnterior"]]    = c.PaymentData.Src.importeSaldoAnterior
  ps[m["ImportePago"]]             = c.PaymentData.Src.ImportePago
  ps[m["importeSaldoInsoluto"]]    = c.PaymentData.Src.importeSaldoInsoluto
  ps[m["tipoRelacion"]]            = c.PaymentData.Src.tipoRelacion
  ps[m["pagoCanceladoDocNumber"]]  = c.PaymentData.Src.pagoCanceladoDocNumber
  ps[m["numOperacion"]]            = c.PaymentData.Src.numOperacion
  ps[m["rfcBancoOrdenente"]]       = c.PaymentData.Src.rfcBancoOrdenente
  ps[m["nombreBancoOrdenante"]]    = c.PaymentData.Src.nombreBancoOrdenante
  ps[m["cuentaOrdenante"]]         = c.PaymentData.Src.cuentaOrdenante
  ps[m["rfcBancoBeneficiario"]]    = c.PaymentData.Src.rfcBancoBeneficiario
  ps[m["cuentaBeneficiario"]]      = c.PaymentData.Src.cuentaBeneficiario
  ps[m["tipoCadenaPago"]]          = c.PaymentData.Src.tipoCadenaPago
  ps[m["certificadoPago"]]         = c.PaymentData.Src.certificadoPago
  ps[m["cadenaPago"]]              = c.PaymentData.Src.cadenaPago
  ps[m["selloPago"]]               = c.PaymentData.Src.selloPago
  ps[m["taxCode"]]                 = c.PaymentData.Src.TaxCode

  // Setup output values of common fields
  pf[m["retencionesIVA"]]          = round(c.Totales.RetencionesIVA)
  pf[m["trasladosBaseIVA16"]]      = round(c.Totales.TrasladosBaseIVA16)
  pf[m["trasladosImpuestoIVA16"]]  = round(c.Totales.TrasladosImpuestoIVA16)
  pf[m["trasladosBaseIVA8"]]       = round(c.Totales.TrasladosBaseIVA8)
  pf[m["trasladosImpuestoIVA8"]]   = round(c.Totales.TrasladosImpuestoIVA8)
  pf[m["trasladosBaseIVA0"]]       = round(c.Totales.TrasladosBaseIVA0)
  pf[m["trasladosImpuestoIVA0"]]   = round(c.Totales.TrasladosImpuestoIVA0)
  pf[m["montoTotalPagos"]]         = round(c.Totales.MontoTotalPagos)
  ps[m["objetoImpuesto"]]          = ""

  // Setup output values of One-taxcode payments
  pf[m["taxTrasladoBase"]]         = round(c.ImpuestosP.TrasladoP.BaseP)
  ps[m["taxTrasladoImpuesto"]]     =       c.OneTaxPaym.TrasladoP.ImpuestoP
  ps[m["taxTrasladoTipoFactor"]]   =       c.OneTaxPaym.TrasladoP.TipoFactorP
  pf[m["taxTrasladoTasaOCuota"]]   = round(c.OneTaxPaym.TrasladoP.TasaOCuotaP)
  pf[m["taxTrasladoImporte"]]      = round(c.ImpuestosP.TrasladoP.ImporteP)
  if c.ImpuestosP.RetncionP.ImporteP != 0.0 {
    pf[m["taxRetncionBase"]]       = round(c.ImpuestosP.RetncionP.BaseP)
    ps[m["taxRetncionImpuesto"]]   =       c.OneTaxPaym.RetncionP.ImpuestoP
    ps[m["taxRetncionTipoFactor"]] =       c.OneTaxPaym.RetncionP.TipoFactorP
    pf[m["taxRetncionTasaOCuota"]] = round(c.OneTaxPaym.RetncionP.TasaOCuotaP)
    pf[m["taxRetncionImporte"]]    = round(c.ImpuestosP.RetncionP.ImporteP)
  } else {
    pf[m["taxRetncionBase"]]       = 0.0
    ps[m["taxRetncionImpuesto"]]   = ""
    ps[m["taxRetncionTipoFactor"]] = ""
    pf[m["taxRetncionTasaOCuota"]] = 0.0
    pf[m["taxRetncionImporte"]]    = 0.0
  }
  importePagoCalc := c.ImpuestosP.TrasladoP.BaseP +
                     c.ImpuestosP.TrasladoP.ImporteP -
                     c.ImpuestosP.RetncionP.ImporteP
  importePagoCalc  = ut.Round(importePagoCalc, 2)
  importePago, _  := strconv.ParseFloat(c.PaymentData.Src.AmountDocCurr, 64)
  importePago = ut.Round(importePago, 2)
  pf[m["difImportePago1"]] = -1.0 * importePago - importePagoCalc
  if math.Abs(pf[m["difImportePago1"]]) < 0.0000015 {
    pf[m["difImportePago1"]] = 0.0
  }

  // Setup output values of Multiple-taxcode payments
  if c.TaxPIVA16.TrasladoP.BaseP != 0.0 {
    pf[m["trasladoBasePIVA16"]]       = round(c.TaxPIVA16.TrasladoP.BaseP)
    ps[m["trasladoImpuestoPIVA16"]]   =       c.TaxPIVA16.TrasladoP.ImpuestoP
    ps[m["trasladoTipoFactorPIVA16"]] =       c.TaxPIVA16.TrasladoP.TipoFactorP
    pf[m["trasladoTasaOCuotaPIVA16"]] = round(c.TaxPIVA16.TrasladoP.TasaOCuotaP)
    pf[m["trasladoImportePIVA16"]]    = round(c.TaxPIVA16.TrasladoP.ImporteP)
  } else {
    pf[m["trasladoBasePIVA16"]]       = 0.0
    ps[m["trasladoImpuestoPIVA16"]]   = ""
    ps[m["trasladoTipoFactorPIVA16"]] = ""
    pf[m["trasladoTasaOCuotaPIVA16"]] = 0.0
    pf[m["trasladoImportePIVA16"]]    = 0.0
  }
  if c.TaxPIVA16.RetncionP.BaseP != 0.0 {
    pf[m["retncionBasePIVA16"]]       = round(c.TaxPIVA16.RetncionP.BaseP)
    ps[m["retncionImpuestoPIVA16"]]   =       c.TaxPIVA16.RetncionP.ImpuestoP
    ps[m["retncionTipoFactorPIVA16"]] =       c.TaxPIVA16.RetncionP.TipoFactorP
    pf[m["retncionTasaOCuotaPIVA16"]] = round(c.TaxPIVA16.RetncionP.TasaOCuotaP)
    pf[m["retncionImportePIVA16"]]    = round(c.TaxPIVA16.RetncionP.ImporteP)
  } else {
    pf[m["retncionBasePIVA16"]]       = 0.0
    ps[m["retncionImpuestoPIVA16"]]   = ""
    ps[m["retncionTipoFactorPIVA16"]] = ""
    pf[m["retncionTasaOCuotaPIVA16"]] = 0.0
    pf[m["retncionImportePIVA16"]]    = 0.0
  }
  if c.TaxPIVA8.TrasladoP.BaseP != 0.0 {
    pf[m["trasladoBasePIVA8"]]        = round(c.TaxPIVA8.TrasladoP.BaseP)
    ps[m["trasladoImpuestoPIVA8"]]    =       c.TaxPIVA8.TrasladoP.ImpuestoP
    ps[m["trasladoTipoFactorPIVA8"]]  =       c.TaxPIVA8.TrasladoP.TipoFactorP
    pf[m["trasladoTasaOCuotaPIVA8"]]  = round(c.TaxPIVA8.TrasladoP.TasaOCuotaP)
    pf[m["trasladoImportePIVA8"]]     = round(c.TaxPIVA8.TrasladoP.ImporteP)
  } else {
    pf[m["trasladoBasePIVA8"]]        = 0.0
    ps[m["trasladoImpuestoPIVA8"]]    = ""
    ps[m["trasladoTipoFactorPIVA8"]]  = ""
    pf[m["trasladoTasaOCuotaPIVA8"]]  = 0.0
    pf[m["trasladoImportePIVA8"]]     = 0.0
  }
  if c.TaxPIVA8.RetncionP.BaseP != 0.0 {
    pf[m["retncionBasePIVA8"]]        = round(c.TaxPIVA8.RetncionP.BaseP)
    ps[m["retncionImpuestoPIVA8"]]    =       c.TaxPIVA8.RetncionP.ImpuestoP
    ps[m["retncionTipoFactorPIVA8"]]  =       c.TaxPIVA8.RetncionP.TipoFactorP
    pf[m["retncionTasaOCuotaPIVA8"]]  = round(c.TaxPIVA8.RetncionP.TasaOCuotaP)
    pf[m["retncionImportePIVA8"]]     = round(c.TaxPIVA8.RetncionP.ImporteP)
  } else {
    pf[m["retncionBasePIVA8"]]        = 0.0
    ps[m["retncionImpuestoPIVA8"]]    = ""
    ps[m["retncionTipoFactorPIVA8"]]  = ""
    pf[m["retncionTasaOCuotaPIVA8"]]  = 0.0
    pf[m["retncionImportePIVA8"]]     = 0.0
  }
  if c.TaxPIVA0.TrasladoP.BaseP != 0.0 {
    pf[m["trasladoBasePIVA0"]]        = round(c.TaxPIVA0.TrasladoP.BaseP)
    ps[m["trasladoImpuestoPIVA0"]]    =       c.TaxPIVA0.TrasladoP.ImpuestoP
    ps[m["trasladoTipoFactorPIVA0"]]  =       c.TaxPIVA0.TrasladoP.TipoFactorP
    pf[m["trasladoTasaOCuotaPIVA0"]]  = round(c.TaxPIVA0.TrasladoP.TasaOCuotaP)
    pf[m["trasladoImportePIVA0"]]     = round(c.TaxPIVA0.TrasladoP.ImporteP)
  } else {
    pf[m["trasladoBasePIVA0"]]        = 0.0
    ps[m["trasladoImpuestoPIVA0"]]    = ""
    ps[m["trasladoTipoFactorPIVA0"]]  = ""
    pf[m["trasladoTasaOCuotaPIVA0"]]  = 0.0
    pf[m["trasladoImportePIVA0"]]     = 0.0
  }
  if c.TaxPIVA0.RetncionP.BaseP != 0.0 {
    pf[m["retncionBasePIVA0"]]        = round(c.TaxPIVA0.RetncionP.BaseP)
    ps[m["retncionImpuestoPIVA0"]]    =       c.TaxPIVA0.RetncionP.ImpuestoP
    ps[m["retncionTipoFactorPIVA0"]]  =       c.TaxPIVA0.RetncionP.TipoFactorP
    pf[m["retncionTasaOCuotaPIVA0"]]  = round(c.TaxPIVA0.RetncionP.TasaOCuotaP)
    pf[m["retncionImportePIVA0"]]     = round(c.TaxPIVA0.RetncionP.ImporteP)
  } else {
    pf[m["retncionBasePIVA0"]]        = 0.0
    ps[m["retncionImpuestoPIVA0"]]    = ""
    ps[m["retncionTipoFactorPIVA0"]]  = ""
    pf[m["retncionTasaOCuotaPIVA0"]]  = 0.0
    pf[m["retncionImportePIVA0"]]     = 0.0
  }
  trasladoBaseP    := 0.0
  if c.TaxPIVA16.TrasladoP.BaseP != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = c.TaxPIVA16.TrasladoP.BaseP
    }
  }
  if c.TaxPIVA8.TrasladoP.BaseP != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = c.TaxPIVA8.TrasladoP.BaseP
    }
  }
  if c.TaxPIVA0.TrasladoP.BaseP != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = c.TaxPIVA0.TrasladoP.BaseP
    }
  }
  trasladoImporteP := c.TaxPIVA16.TrasladoP.ImporteP +
                       c.TaxPIVA8.TrasladoP.ImporteP +
                       c.TaxPIVA0.TrasladoP.ImporteP
  retncionImporteP := c.TaxPIVA16.RetncionP.ImporteP +
                       c.TaxPIVA8.RetncionP.ImporteP +
                       c.TaxPIVA0.RetncionP.ImporteP
  importePagoCalc = trasladoBaseP + trasladoImporteP - retncionImporteP
  importePago, _  = strconv.ParseFloat(c.PaymentData.Src.AmountDocCurr, 64)
  importePago = ut.Round(importePago, 2)
  pf[m["difImportePago3"]] = -1.0 * importePago - importePagoCalc
  if math.Abs(pf[m["difImportePago3"]]) < 0.0000015 {
    pf[m["difImportePago3"]] = 0.0
  }
  amountDocCurr, _ := strconv.ParseFloat(c.PaymentData.Src.AmountDocCurr, 64)
  amountDocCurr = ut.Round(amountDocCurr, 2)
  pf[m["difMontoTotalPagos"]] = -1.0 * amountDocCurr - c.Totales.MontoTotalPagos
  if math.Abs(pf[m["difMontoTotalPagos"]]) < 0.0000015 {
    pf[m["difMontoTotalPagos"]] = 0.0
  }
  w.PrintLineExcel()
  return c
}

func (c *Calctax_tp) FetchInvoiceLines(w *Writer_tp) *Calctax_tp {
  for _, i := range c.invoices {
  // Setup output values of source fields
  ps[m["companyCode"]]             = i.src.companyCode
  ps[m["customer"]]                = i.src.customer
  ps[m["documentNumber"]]          = i.src.documentNumber
  ps[m["documentType"]]            = i.src.DocumentType
  ps[m["paymentDateTime"]]         = i.src.paymentDateTime
  ps[m["clearingDocument"]]        = i.src.ClearingDocument
  ps[m["amountDocCurr"]]           = i.src.AmountDocCurr
  ps[m["documentCurrency"]]        = i.src.DocumentCurrency
  ps[m["effExchangeRate"]]         = i.src.EffExchangeRate
  ps[m["assignment"]]              = i.src.assignment
  ps[m["formaPago"]]               = i.src.formaPago
  ps[m["noParcialidad"]]           = i.src.noParcialidad
  ps[m["importeSaldoAnterior"]]    = i.src.importeSaldoAnterior
  ps[m["ImportePago"]]             = i.src.ImportePago
  ps[m["importeSaldoInsoluto"]]    = i.src.importeSaldoInsoluto
  ps[m["tipoRelacion"]]            = i.src.tipoRelacion
  ps[m["pagoCanceladoDocNumber"]]  = i.src.pagoCanceladoDocNumber
  ps[m["numOperacion"]]            = i.src.numOperacion
  ps[m["rfcBancoOrdenente"]]       = i.src.rfcBancoOrdenente
  ps[m["nombreBancoOrdenante"]]    = i.src.nombreBancoOrdenante
  ps[m["cuentaOrdenante"]]         = i.src.cuentaOrdenante
  ps[m["rfcBancoBeneficiario"]]    = i.src.rfcBancoBeneficiario
  ps[m["cuentaBeneficiario"]]      = i.src.cuentaBeneficiario
  ps[m["tipoCadenaPago"]]          = i.src.tipoCadenaPago
  ps[m["certificadoPago"]]         = i.src.certificadoPago
  ps[m["cadenaPago"]]              = i.src.cadenaPago
  ps[m["selloPago"]]               = i.src.selloPago
  ps[m["taxCode"]]                 = i.src.TaxCode

  // Setup output values of common fields
    pf[m["retencionesIVA"]]          = 0.0
    pf[m["trasladosBaseIVA16"]]      = 0.0
    pf[m["trasladosImpuestoIVA16"]]  = 0.0
    pf[m["trasladosBaseIVA8"]]       = 0.0
    pf[m["trasladosImpuestoIVA8"]]   = 0.0
    pf[m["trasladosBaseIVA0"]]       = 0.0
    pf[m["trasladosImpuestoIVA0"]]   = 0.0
    pf[m["montoTotalPagos"]]         = 0.0
    ps[m["objetoImpuesto"]]          = i.docrel.ObjetoImpDR

    // Setup output values of Invoices in One-taxcode payments
    pf[m["taxTrasladoBase"]]         = round(i.docrel.TrasladoDR.BaseDR)
    ps[m["taxTrasladoImpuesto"]]     = IMPUESTO
    ps[m["taxTrasladoTipoFactor"]]   = TIPOFACTOR
    pf[m["taxTrasladoTasaOCuota"]]   = round(i.docrel.TrasladoDR.TasaOCuotaDR)
    pf[m["taxTrasladoImporte"]]      = round(i.docrel.TrasladoDR.ImporteDR)
    if i.docrel.RetncionDR.ImporteDR != 0.0 {
      pf[m["taxRetncionBase"]]       = round(i.docrel.RetncionDR.BaseDR)
      ps[m["taxRetncionImpuesto"]]   =       i.docrel.RetncionDR.ImpuestoDR
      ps[m["taxRetncionTipoFactor"]] =       i.docrel.RetncionDR.TipoFactorDR
      pf[m["taxRetncionTasaOCuota"]] = round(i.docrel.RetncionDR.TasaOCuotaDR)
      pf[m["taxRetncionImporte"]]    = round(i.docrel.RetncionDR.ImporteDR)
    } else {
      pf[m["taxRetncionBase"]]       = 0.0
      ps[m["taxRetncionImpuesto"]]   = ""
      ps[m["taxRetncionTipoFactor"]] = ""
      pf[m["taxRetncionTasaOCuota"]] = 0.0
      pf[m["taxRetncionImporte"]]    = 0.0
    }
    // Setup output values of Invoices in Multiple-taxcode payments
    pf[m["trasladoBaseDR"]]         = round(i.docrel.TrasladoDR.BaseDR)
    ps[m["trasladoImpuestoDR"]]     = IMPUESTO
    ps[m["trasladoTipoFactorDR"]]   = TIPOFACTOR
    pf[m["trasladoTasaOCuotaDR"]]   = round(i.docrel.TrasladoDR.TasaOCuotaDR)
    pf[m["trasladoImporteDR"]]      = round(i.docrel.TrasladoDR.ImporteDR)
    if i.docrel.RetncionDR.ImporteDR != 0.0 {
      pf[m["retncionBaseDR"]]       = round(i.docrel.RetncionDR.BaseDR)
      ps[m["retncionImpuestoDR"]]   = i.docrel.RetncionDR.ImpuestoDR
      ps[m["retncionTipoFactorDR"]] = i.docrel.RetncionDR.TipoFactorDR
      pf[m["retncionTasaOCuotaDR"]] = round(i.docrel.RetncionDR.TasaOCuotaDR)
      pf[m["retncionImporteDR"]]    = round(i.docrel.RetncionDR.ImporteDR)
    } else {
      pf[m["retncionBaseDR"]]       = 0.0
      ps[m["retncionImpuestoDR"]]   = ""
      ps[m["retncionTipoFactorDR"]] = ""
      pf[m["retncionTasaOCuotaDR"]] = 0.0
      pf[m["retncionImporteDR"]]    = 0.0
    }
    importePagoCalc := i.docrel.TrasladoDR.BaseDR +
                       i.docrel.TrasladoDR.ImporteDR -
                       i.docrel.RetncionDR.ImporteDR
    importePago, _ := strconv.ParseFloat(i.src.ImportePago, 64)
    importePago = ut.Round(importePago, 2)
    pf[m["difImportePago1"]] = importePago - importePagoCalc
    if math.Abs(pf[m["difImportePago1"]]) < 0.0000015 {
      pf[m["difImportePago1"]] = 0.0
    }
    w.PrintLineExcel()
  }
  return c
}

// Linear assignments
func (c *Calctax_tp) resetCommonPaymentData() *Calctax_tp {
  c.Totales.RetencionesIVA           = 0.0
  c.Totales.TrasladosBaseIVA16       = 0.0
  c.Totales.TrasladosImpuestoIVA16   = 0.0
  c.Totales.TrasladosBaseIVA8        = 0.0
  c.Totales.TrasladosImpuestoIVA8    = 0.0
  c.Totales.TrasladosBaseIVA0        = 0.0
  c.Totales.TrasladosImpuestoIVA0    = 0.0
  c.Totales.MontoTotalPagos          = 0.0
  return c
}

func (c *Calctax_tp) reset1TaxPaymentData() *Calctax_tp {
  c.ImpuestosP.TrasladoP.BaseP       = 0.0
  c.ImpuestosP.TrasladoP.ImpuestoP   = ""
  c.ImpuestosP.TrasladoP.TipoFactorP = ""
  c.ImpuestosP.TrasladoP.TasaOCuotaP = 0.0
  c.ImpuestosP.TrasladoP.ImporteP    = 0.0
  c.ImpuestosP.RetncionP.BaseP       = 0.0
  c.ImpuestosP.RetncionP.ImpuestoP   = ""
  c.ImpuestosP.RetncionP.TipoFactorP = ""
  c.ImpuestosP.RetncionP.TasaOCuotaP = 0.0
  c.ImpuestosP.RetncionP.ImporteP    = 0.0
  c.OneTaxPaym.TrasladoP.ImpuestoP   = ""
  c.OneTaxPaym.TrasladoP.TipoFactorP = ""
  c.OneTaxPaym.TrasladoP.TasaOCuotaP = 0.0
  c.OneTaxPaym.RetncionP.ImpuestoP   = ""
  c.OneTaxPaym.RetncionP.TipoFactorP = ""
  c.OneTaxPaym.RetncionP.TasaOCuotaP = 0.0
  c.firstInvoice                     = true
  c.invoices = nil
  return c
}

func (c *Calctax_tp) reset3TaxPaymentData() *Calctax_tp {
  c.TaxPIVA16.TrasladoP.BaseP        = 0.0
  c.TaxPIVA16.TrasladoP.ImpuestoP    = ""
  c.TaxPIVA16.TrasladoP.TipoFactorP  = ""
  c.TaxPIVA16.TrasladoP.TasaOCuotaP  = 0.0
  c.TaxPIVA16.TrasladoP.ImporteP     = 0.0
  c.TaxPIVA16.RetncionP.BaseP        = 0.0
  c.TaxPIVA16.RetncionP.ImpuestoP    = ""
  c.TaxPIVA16.RetncionP.TipoFactorP  = ""
  c.TaxPIVA16.RetncionP.TasaOCuotaP  = 0.0
  c.TaxPIVA16.RetncionP.ImporteP     = 0.0
  c.TaxPIVA8.TrasladoP.BaseP         = 0.0
  c.TaxPIVA8.TrasladoP.ImpuestoP     = ""
  c.TaxPIVA8.TrasladoP.TipoFactorP   = ""
  c.TaxPIVA8.TrasladoP.TasaOCuotaP   = 0.0
  c.TaxPIVA8.TrasladoP.ImporteP      = 0.0
  c.TaxPIVA8.RetncionP.BaseP         = 0.0
  c.TaxPIVA8.RetncionP.ImpuestoP     = ""
  c.TaxPIVA8.RetncionP.TipoFactorP   = ""
  c.TaxPIVA8.RetncionP.TasaOCuotaP   = 0.0
  c.TaxPIVA8.RetncionP.ImporteP      = 0.0
  c.TaxPIVA0.TrasladoP.BaseP         = 0.0
  c.TaxPIVA0.TrasladoP.ImpuestoP     = ""
  c.TaxPIVA0.TrasladoP.TipoFactorP   = ""
  c.TaxPIVA0.TrasladoP.TasaOCuotaP   = 0.0
  c.TaxPIVA0.TrasladoP.ImporteP      = 0.0
  c.TaxPIVA0.RetncionP.BaseP         = 0.0
  c.TaxPIVA0.RetncionP.ImpuestoP     = ""
  c.TaxPIVA0.RetncionP.TipoFactorP   = ""
  c.TaxPIVA0.RetncionP.TasaOCuotaP   = 0.0
  c.TaxPIVA0.RetncionP.ImporteP      = 0.0
  c.firstInvoTraslIVA16              = true
  c.firstInvoRetenIVA16              = true
  c.firstInvoTraslIVA8               = true
  c.firstInvoRetenIVA8               = true
  c.firstInvoTraslIVA0               = true
  c.firstInvoRetenIVA0               = true
  c.invoices = nil
  return c
}

func round(f float64) float64 {
  return ut.Round(f, DEC)
}
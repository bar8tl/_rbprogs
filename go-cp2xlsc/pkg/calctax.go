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
  var o lout_tp
  o.src = c.PaymentData.Src
  o.retencionesIVA          = ut.Round(c.Totales.RetencionesIVA, DEC)
  o.trasladosBaseIVA16      = ut.Round(c.Totales.TrasladosBaseIVA16, DEC)
  o.trasladosImpuestoIVA16  = ut.Round(c.Totales.TrasladosImpuestoIVA16, DEC)
  o.trasladosBaseIVA8       = ut.Round(c.Totales.TrasladosBaseIVA8, DEC)
  o.trasladosImpuestoIVA8   = ut.Round(c.Totales.TrasladosImpuestoIVA8, DEC)
  o.trasladosBaseIVA0       = ut.Round(c.Totales.TrasladosBaseIVA0, DEC)
  o.trasladosImpuestoIVA0   = ut.Round(c.Totales.TrasladosImpuestoIVA0, DEC)
  o.montoTotalPagos         = ut.Round(c.Totales.MontoTotalPagos, DEC)
  o.objetoImpuesto          = ""
  // Setup output values of One-taxcode payments
  o.taxTrasladoBase         = ut.Round(c.ImpuestosP.TrasladoP.BaseP, DEC)
  o.taxTrasladoImpuesto     = c.OneTaxPaym.TrasladoP.ImpuestoP
  o.taxTrasladoTipoFactor   = c.OneTaxPaym.TrasladoP.TipoFactorP
  o.taxTrasladoTasaOCuota   = ut.Round(c.OneTaxPaym.TrasladoP.TasaOCuotaP, DEC)
  o.taxTrasladoImporte      = ut.Round(c.ImpuestosP.TrasladoP.ImporteP, DEC)
  if c.ImpuestosP.RetncionP.ImporteP != 0.0 {
    o.taxRetncionBase       = ut.Round(c.ImpuestosP.RetncionP.BaseP, DEC)
    o.taxRetncionImpuesto   = c.OneTaxPaym.RetncionP.ImpuestoP
    o.taxRetncionTipoFactor = c.OneTaxPaym.RetncionP.TipoFactorP
    o.taxRetncionTasaOCuota = ut.Round(c.OneTaxPaym.RetncionP.TasaOCuotaP, DEC)
    o.taxRetncionImporte    = ut.Round(c.ImpuestosP.RetncionP.ImporteP, DEC)
  } else {
    o.taxRetncionBase       = 0.0
    o.taxRetncionImpuesto   = ""
    o.taxRetncionTipoFactor = ""
    o.taxRetncionTasaOCuota = 0.0
    o.taxRetncionImporte    = 0.0
  }
  importePagoCalc := o.taxTrasladoBase + o.taxTrasladoImporte -
                     o.taxRetncionImporte
  importePago, _  := strconv.ParseFloat(c.PaymentData.Src.AmountDocCurr, 64)
  importePago = ut.Round(importePago, 2)
  o.difImportePago1 = -1.0 * importePago - importePagoCalc
  if math.Abs(o.difImportePago1) < 0.0000015 {
    o.difImportePago1 = 0.0
  }
  // Setup output values of Multiple-taxcode payments
  if c.TaxPIVA16.TrasladoP.BaseP != 0.0 {
    o.trasladoBasePIVA16       = ut.Round(c.TaxPIVA16.TrasladoP.BaseP, DEC)
    o.trasladoImpuestoPIVA16   =          c.TaxPIVA16.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA16 =          c.TaxPIVA16.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA16 = ut.Round(c.TaxPIVA16.TrasladoP.TasaOCuotaP,DEC)
    o.trasladoImportePIVA16    = ut.Round(c.TaxPIVA16.TrasladoP.ImporteP, DEC)
  } else {
    o.trasladoBasePIVA16       = 0.0
    o.trasladoImpuestoPIVA16   = ""
    o.trasladoTipoFactorPIVA16 = ""
    o.trasladoTasaOCuotaPIVA16 = 0.0
    o.trasladoImportePIVA16    = 0.0
  }
  if c.TaxPIVA16.RetncionP.BaseP != 0.0 {
    o.retncionBasePIVA16       = ut.Round(c.TaxPIVA16.RetncionP.BaseP, DEC)
    o.retncionImpuestoPIVA16   =          c.TaxPIVA16.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA16 =          c.TaxPIVA16.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA16 = ut.Round(c.TaxPIVA16.RetncionP.TasaOCuotaP,DEC)
    o.retncionImportePIVA16    = ut.Round(c.TaxPIVA16.RetncionP.ImporteP, DEC)
  } else {
    o.retncionBasePIVA16       = 0.0
    o.retncionImpuestoPIVA16   = ""
    o.retncionTipoFactorPIVA16 = ""
    o.retncionTasaOCuotaPIVA16 = 0.0
    o.retncionImportePIVA16    = 0.0
  }
  if c.TaxPIVA8.TrasladoP.BaseP != 0.0 {
    o.trasladoBasePIVA8        = ut.Round(c.TaxPIVA8.TrasladoP.BaseP, DEC)
    o.trasladoImpuestoPIVA8    =          c.TaxPIVA8.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA8  =          c.TaxPIVA8.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA8  = ut.Round(c.TaxPIVA8.TrasladoP.TasaOCuotaP, DEC)
    o.trasladoImportePIVA8     = ut.Round(c.TaxPIVA8.TrasladoP.ImporteP, DEC)
  } else {
    o.trasladoBasePIVA8        = 0.0
    o.trasladoImpuestoPIVA8    = ""
    o.trasladoTipoFactorPIVA8  = ""
    o.trasladoTasaOCuotaPIVA8  = 0.0
    o.trasladoImportePIVA8     = 0.0
  }
  if c.TaxPIVA8.RetncionP.BaseP != 0.0 {
    o.retncionBasePIVA8        = ut.Round(c.TaxPIVA8.RetncionP.BaseP, DEC)
    o.retncionImpuestoPIVA8    =          c.TaxPIVA8.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA8  =          c.TaxPIVA8.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA8  = ut.Round(c.TaxPIVA8.RetncionP.TasaOCuotaP, DEC)
    o.retncionImportePIVA8     = ut.Round(c.TaxPIVA8.RetncionP.ImporteP, DEC)
  } else {
    o.retncionBasePIVA8        = 0.0
    o.retncionImpuestoPIVA8    = ""
    o.retncionTipoFactorPIVA8  = ""
    o.retncionTasaOCuotaPIVA8  = 0.0
    o.retncionImportePIVA8     = 0.0
  }
  if c.TaxPIVA0.TrasladoP.BaseP != 0.0 {
    o.trasladoBasePIVA0        = ut.Round(c.TaxPIVA0.TrasladoP.BaseP, DEC)
    o.trasladoImpuestoPIVA0    =          c.TaxPIVA0.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA0  =          c.TaxPIVA0.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA0  = ut.Round(c.TaxPIVA0.TrasladoP.TasaOCuotaP, DEC)
    o.trasladoImportePIVA0     = ut.Round(c.TaxPIVA0.TrasladoP.ImporteP, DEC)
  } else {
    o.trasladoBasePIVA0        = 0.0
    o.trasladoImpuestoPIVA0    = ""
    o.trasladoTipoFactorPIVA0  = ""
    o.trasladoTasaOCuotaPIVA0  = 0.0
    o.trasladoImportePIVA0     = 0.0
  }
  if c.TaxPIVA0.RetncionP.BaseP != 0.0 {
    o.retncionBasePIVA0        = ut.Round(c.TaxPIVA0.RetncionP.BaseP, DEC)
    o.retncionImpuestoPIVA0    =          c.TaxPIVA0.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA0  =          c.TaxPIVA0.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA0  = ut.Round(c.TaxPIVA0.RetncionP.TasaOCuotaP, DEC)
    o.retncionImportePIVA0     = ut.Round(c.TaxPIVA0.RetncionP.ImporteP, DEC)
  } else {
    o.retncionBasePIVA0        = 0.0
    o.retncionImpuestoPIVA0    = ""
    o.retncionTipoFactorPIVA0  = ""
    o.retncionTasaOCuotaPIVA0  = 0.0
    o.retncionImportePIVA0     = 0.0
  }
  trasladoBaseP    := 0.0
  if o.trasladoBasePIVA16 != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = o.trasladoBasePIVA16
    }
  }
  if o.trasladoBasePIVA8 != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = o.trasladoBasePIVA8
    }
  }
  if o.trasladoBasePIVA0 != 0.0 {
    if trasladoBaseP == 0.0 {
      trasladoBaseP = o.trasladoBasePIVA0
    }
  }
  trasladoImporteP := o.trasladoImportePIVA16 + o.trasladoImportePIVA8 +
                      o.trasladoImportePIVA0
  retncionImporteP := o.retncionImportePIVA16 + o.retncionImportePIVA8 +
                      o.retncionImportePIVA0
  importePagoCalc = trasladoBaseP + trasladoImporteP - retncionImporteP
  importePago, _  = strconv.ParseFloat(c.PaymentData.Src.AmountDocCurr, 64)
  importePago = ut.Round(importePago, 2)
  o.difImportePago3 = -1.0 * importePago - importePagoCalc
  if math.Abs(o.difImportePago3) < 0.0000015 {
    o.difImportePago3 = 0.0
  }
  amountDocCurr, _ := strconv.ParseFloat(c.PaymentData.Src.AmountDocCurr, 64)
  amountDocCurr = ut.Round(amountDocCurr, 2)
  o.difMontoTotalPagos = -1.0 * amountDocCurr - c.Totales.MontoTotalPagos
  if math.Abs(o.difMontoTotalPagos) < 0.0000015 {
    o.difMontoTotalPagos = 0.0
  }
  w.PrintLineExcel(o)
  return c
}

func (c *Calctax_tp) FetchInvoiceLines(w *Writer_tp) *Calctax_tp {
  for _, i := range c.invoices {
    var o lout_tp
    o.src = i.src
    o.retencionesIVA         = 0.0
    o.trasladosBaseIVA16     = 0.0
    o.trasladosImpuestoIVA16 = 0.0
    o.trasladosBaseIVA8      = 0.0
    o.trasladosImpuestoIVA8  = 0.0
    o.trasladosBaseIVA0      = 0.0
    o.trasladosImpuestoIVA0  = 0.0
    o.montoTotalPagos        = 0.0
    o.objetoImpuesto         = i.docrel.ObjetoImpDR
    // Setup output values of Invoices in One-taxcode payments
    o.taxTrasladoBase         = ut.Round(i.docrel.TrasladoDR.BaseDR, DEC)
    o.taxTrasladoImpuesto     = IMPUESTO
    o.taxTrasladoTipoFactor   = TIPOFACTOR
    o.taxTrasladoTasaOCuota   = ut.Round(i.docrel.TrasladoDR.TasaOCuotaDR, DEC)
    o.taxTrasladoImporte      = ut.Round(i.docrel.TrasladoDR.ImporteDR, DEC)
    if i.docrel.RetncionDR.ImporteDR != 0.0 {
      o.taxRetncionBase       = ut.Round(i.docrel.RetncionDR.BaseDR, DEC)
      o.taxRetncionImpuesto   =          i.docrel.RetncionDR.ImpuestoDR
      o.taxRetncionTipoFactor =          i.docrel.RetncionDR.TipoFactorDR
      o.taxRetncionTasaOCuota = ut.Round(i.docrel.RetncionDR.TasaOCuotaDR, DEC)
      o.taxRetncionImporte    = ut.Round(i.docrel.RetncionDR.ImporteDR, DEC)
    } else {
      o.taxRetncionBase       = 0.0
      o.taxRetncionImpuesto   = ""
      o.taxRetncionTipoFactor = ""
      o.taxRetncionTasaOCuota = 0.0
      o.taxRetncionImporte    = 0.0
    }
    // Setup output values of Invoices in Multiple-taxcode payments
    o.trasladoBaseDR         = ut.Round(i.docrel.TrasladoDR.BaseDR, DEC)
    o.trasladoImpuestoDR     = IMPUESTO
    o.trasladoTipoFactorDR   = TIPOFACTOR
    o.trasladoTasaOCuotaDR   = ut.Round(i.docrel.TrasladoDR.TasaOCuotaDR, DEC)
    o.trasladoImporteDR      = ut.Round(i.docrel.TrasladoDR.ImporteDR, DEC)
    if i.docrel.RetncionDR.ImporteDR != 0.0 {
      o.retncionBaseDR       = ut.Round(i.docrel.RetncionDR.BaseDR, DEC)
      o.retncionImpuestoDR   = i.docrel.RetncionDR.ImpuestoDR
      o.retncionTipoFactorDR = i.docrel.RetncionDR.TipoFactorDR
      o.retncionTasaOCuotaDR = ut.Round(i.docrel.RetncionDR.TasaOCuotaDR, DEC)
      o.retncionImporteDR    = ut.Round(i.docrel.RetncionDR.ImporteDR, DEC)
    } else {
      o.retncionBaseDR       = 0.0
      o.retncionImpuestoDR   = ""
      o.retncionTipoFactorDR = ""
      o.retncionTasaOCuotaDR = 0.0
      o.retncionImporteDR    = 0.0
    }
    importePagoCalc := i.docrel.TrasladoDR.BaseDR +
                       i.docrel.TrasladoDR.ImporteDR -
                       i.docrel.RetncionDR.ImporteDR
    importePago, _ := strconv.ParseFloat(i.src.ImportePago, 64)
    importePago = ut.Round(importePago, 2)
    o.difImportePago1 = importePago - importePagoCalc
    if math.Abs(o.difImportePago1) < 0.0000015 {
      o.difImportePago1 = 0.0
    }
    w.PrintLineExcel(o)
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

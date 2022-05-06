// cpm2xlsp.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields. Version for Multiple-TaxRate Invoices per Payment - Auxiliary tools
package bbmwht

import ut "bar8tl/p/rblib"
import "github.com/xuri/excelize/v2"
import "log"
import "math"
import "strconv"

// Auxiliary funtions ----------------------------------------------------------
type Mtools_tp struct {
  F             *excelize.File
  g             *excelize.File
  index         int
  recn          int
  AmountDocCurr float64
  ImportePago   float64
  PaymentData   Line_tp
  Totales       Totales_tp
  TaxPIVA16     Payment_tp
  TaxPIVA8      Payment_tp
  TaxPIVA0      Payment_tp
  invoices      []linv_tp
  firstInvoTraslIVA16 bool
  firstInvoRetenIVA16 bool
  firstInvoTraslIVA8  bool
  firstInvoRetenIVA8  bool
  firstInvoTraslIVA0  bool
  firstInvoRetenIVA0  bool
}

func NewMtools() *Mtools_tp {
  var p Mtools_tp
  p.firstInvoTraslIVA16 = true
  p.firstInvoRetenIVA16 = true
  p.firstInvoTraslIVA8  = true
  p.firstInvoRetenIVA8  = true
  p.firstInvoTraslIVA0  = true
  p.firstInvoRetenIVA0  = true
  return &p
}

func (p *Mtools_tp) OpenInpExcel(file string) {
  var err error
  p.F, err = excelize.OpenFile(file)
  if err != nil {
    log.Fatal(err)
  }
}

func (p *Mtools_tp) CreateOutExcel(tabnam string) {
  p.g = excelize.NewFile()
  p.index = p.g.NewSheet(tabnam)
}

func (p *Mtools_tp) WriteOutExcel(fname string) {
  p.g.SetActiveSheet(p.index)
  if err := p.g.SaveAs(fname); err != nil {
    log.Fatal(err)
  }
}   

func (p *Mtools_tp) StoreDocRel(lin Line_tp, inv Docrel_tp) {
  var linv linv_tp
  linv.src = lin
  linv.docrel.ObjetoImpDR             = inv.ObjetoImpDR
  linv.docrel.TrasladoDR.BaseDR       = inv.TrasladoDR.BaseDR
  linv.docrel.TrasladoDR.ImpuestoDR   = inv.TrasladoDR.ImpuestoDR
  linv.docrel.TrasladoDR.TipoFactorDR = inv.TrasladoDR.TipoFactorDR
  linv.docrel.TrasladoDR.TasaOCuotaDR = inv.TrasladoDR.TasaOCuotaDR
  linv.docrel.TrasladoDR.ImporteDR    = inv.TrasladoDR.ImporteDR
  linv.docrel.RetncionDR.BaseDR       = inv.RetncionDR.BaseDR
  linv.docrel.RetncionDR.ImpuestoDR   = inv.RetncionDR.ImpuestoDR
  linv.docrel.RetncionDR.TipoFactorDR = inv.RetncionDR.TipoFactorDR
  linv.docrel.RetncionDR.TasaOCuotaDR = inv.RetncionDR.TasaOCuotaDR
  linv.docrel.RetncionDR.ImporteDR    = inv.RetncionDR.ImporteDR
  p.invoices = append(p.invoices, linv_tp{linv.src, linv.docrel})
  if inv.TrasladoDR.TasaOCuotaDR == 0.16 {
    if p.firstInvoTraslIVA16 {
      p.TaxPIVA16.TrasladoP.ImpuestoP   = inv.TrasladoDR.ImpuestoDR
      p.TaxPIVA16.TrasladoP.TipoFactorP = inv.TrasladoDR.TipoFactorDR
      p.TaxPIVA16.TrasladoP.TasaOCuotaP = inv.TrasladoDR.TasaOCuotaDR
      p.firstInvoTraslIVA16 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.16 {
    if p.firstInvoRetenIVA16 {
      p.TaxPIVA16.RetncionP.ImpuestoP   = inv.RetncionDR.ImpuestoDR
      p.TaxPIVA16.RetncionP.TipoFactorP = inv.RetncionDR.TipoFactorDR
      p.TaxPIVA16.RetncionP.TasaOCuotaP = inv.RetncionDR.TasaOCuotaDR
      p.firstInvoRetenIVA16 = false
    }
  }
  if inv.TrasladoDR.TasaOCuotaDR == 0.08 {
    if p.firstInvoTraslIVA8 {
      p.TaxPIVA8.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      p.TaxPIVA8.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      p.TaxPIVA8.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      p.firstInvoTraslIVA8 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.08 {
    if p.firstInvoRetenIVA8 {
      p.TaxPIVA8.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      p.TaxPIVA8.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      p.TaxPIVA8.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      p.firstInvoRetenIVA8 = false
    }
  }
  if inv.TrasladoDR.TasaOCuotaDR == 0.0 {
    if p.firstInvoTraslIVA0 {
      p.TaxPIVA0.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      p.TaxPIVA0.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      p.TaxPIVA0.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      p.firstInvoTraslIVA0 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.0 {
    if p.firstInvoRetenIVA0 {
      p.TaxPIVA0.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      p.TaxPIVA0.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      p.TaxPIVA0.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      p.firstInvoRetenIVA0 = false
    }
  }
}

func (p *Mtools_tp) WritePaymentLine() *Mtools_tp {
  var o lmout_tp
  o.src = p.PaymentData
  o.retencionesIVA          = ut.Round(p.Totales.RetencionesIVA, DEC)
  o.trasladosBaseIVA16      = ut.Round(p.Totales.TrasladosBaseIVA16, DEC)
  o.trasladosImpuestoIVA16  = ut.Round(p.Totales.TrasladosImpuestoIVA16, DEC)
  o.trasladosBaseIVA8       = ut.Round(p.Totales.TrasladosBaseIVA8, DEC)
  o.trasladosImpuestoIVA8   = ut.Round(p.Totales.TrasladosImpuestoIVA8, DEC)
  o.trasladosBaseIVA0       = ut.Round(p.Totales.TrasladosBaseIVA0, DEC)
  o.trasladosImpuestoIVA0   = ut.Round(p.Totales.TrasladosImpuestoIVA0, DEC)
  o.montoTotalPagos         = ut.Round(p.Totales.MontoTotalPagos, DEC)
  o.objetoImpuesto          = ""
  if p.TaxPIVA16.TrasladoP.BaseP != 0.0 {
    o.trasladoBasePIVA16       = ut.Round(p.TaxPIVA16.TrasladoP.BaseP, DEC)
    o.trasladoImpuestoPIVA16   =          p.TaxPIVA16.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA16 =          p.TaxPIVA16.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA16 = ut.Round(p.TaxPIVA16.TrasladoP.TasaOCuotaP,DEC)
    o.trasladoImportePIVA16    = ut.Round(p.TaxPIVA16.TrasladoP.ImporteP, DEC)
  } else {
    o.trasladoBasePIVA16       = 0.0
    o.trasladoImpuestoPIVA16   = ""
    o.trasladoTipoFactorPIVA16 = ""
    o.trasladoTasaOCuotaPIVA16 = 0.0
    o.trasladoImportePIVA16    = 0.0
  }
  if p.TaxPIVA16.RetncionP.BaseP != 0.0 {
    o.retncionBasePIVA16       = ut.Round(p.TaxPIVA16.RetncionP.BaseP, DEC)
    o.retncionImpuestoPIVA16   =          p.TaxPIVA16.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA16 =          p.TaxPIVA16.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA16 = ut.Round(p.TaxPIVA16.RetncionP.TasaOCuotaP,DEC)
    o.retncionImportePIVA16    = ut.Round(p.TaxPIVA16.RetncionP.ImporteP, DEC)
  } else {
    o.retncionBasePIVA16       = 0.0
    o.retncionImpuestoPIVA16   = ""
    o.retncionTipoFactorPIVA16 = ""
    o.retncionTasaOCuotaPIVA16 = 0.0
    o.retncionImportePIVA16    = 0.0
  }
  if p.TaxPIVA8.TrasladoP.BaseP != 0.0 {
    o.trasladoBasePIVA8        = ut.Round(p.TaxPIVA8.TrasladoP.BaseP, DEC)
    o.trasladoImpuestoPIVA8    =          p.TaxPIVA8.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA8  =          p.TaxPIVA8.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA8  = ut.Round(p.TaxPIVA8.TrasladoP.TasaOCuotaP, DEC)
    o.trasladoImportePIVA8     = ut.Round(p.TaxPIVA8.TrasladoP.ImporteP, DEC)
  } else {
    o.trasladoBasePIVA8        = 0.0
    o.trasladoImpuestoPIVA8    = ""
    o.trasladoTipoFactorPIVA8  = ""
    o.trasladoTasaOCuotaPIVA8  = 0.0
    o.trasladoImportePIVA8     = 0.0
  }
  if p.TaxPIVA8.RetncionP.BaseP != 0.0 {
    o.retncionBasePIVA8        = ut.Round(p.TaxPIVA8.RetncionP.BaseP, DEC)
    o.retncionImpuestoPIVA8    =          p.TaxPIVA8.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA8  =          p.TaxPIVA8.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA8  = ut.Round(p.TaxPIVA8.RetncionP.TasaOCuotaP, DEC)
    o.retncionImportePIVA8     = ut.Round(p.TaxPIVA8.RetncionP.ImporteP, DEC)
  } else {
    o.retncionBasePIVA8        = 0.0
    o.retncionImpuestoPIVA8    = ""
    o.retncionTipoFactorPIVA8  = ""
    o.retncionTasaOCuotaPIVA8  = 0.0
    o.retncionImportePIVA8     = 0.0
  }
  if p.TaxPIVA0.TrasladoP.BaseP != 0.0 {
    o.trasladoBasePIVA0        = ut.Round(p.TaxPIVA0.TrasladoP.BaseP, DEC)
    o.trasladoImpuestoPIVA0    =          p.TaxPIVA0.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA0  =          p.TaxPIVA0.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA0  = ut.Round(p.TaxPIVA0.TrasladoP.TasaOCuotaP, DEC)
    o.trasladoImportePIVA0     = ut.Round(p.TaxPIVA0.TrasladoP.ImporteP, DEC)
  } else {
    o.trasladoBasePIVA0        = 0.0
    o.trasladoImpuestoPIVA0    = ""
    o.trasladoTipoFactorPIVA0  = ""
    o.trasladoTasaOCuotaPIVA0  = 0.0
    o.trasladoImportePIVA0     = 0.0
  }
  if p.TaxPIVA0.RetncionP.BaseP != 0.0 {
    o.retncionBasePIVA0        = ut.Round(p.TaxPIVA0.RetncionP.BaseP, DEC)
    o.retncionImpuestoPIVA0    =          p.TaxPIVA0.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA0  =          p.TaxPIVA0.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA0  = ut.Round(p.TaxPIVA0.RetncionP.TasaOCuotaP, DEC)
    o.retncionImportePIVA0     = ut.Round(p.TaxPIVA0.RetncionP.ImporteP, DEC)
  } else {
    o.retncionBasePIVA0        = 0.0
    o.retncionImpuestoPIVA0    = ""
    o.retncionTipoFactorPIVA0  = ""
    o.retncionTasaOCuotaPIVA0  = 0.0
    o.retncionImportePIVA0     = 0.0
  }
  amountDocCurr, _ := strconv.ParseFloat(p.PaymentData.AmountDocCurr, 64)
  o.difMontoTotalPagos = -1.0 * amountDocCurr - p.Totales.MontoTotalPagos
  if math.Abs(o.difMontoTotalPagos) < 0.0000009 {
    o.difMontoTotalPagos = 0.0
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
  importePagoCalc := trasladoBaseP + trasladoImporteP - retncionImporteP
  importePago, _  := strconv.ParseFloat(p.PaymentData.AmountDocCurr, 64)
  o.difImportePago = -1.0 * importePago - importePagoCalc
  if math.Abs(o.difImportePago) < 0.0000009 {
    o.difImportePago = 0.0
  }
  p.buildLineExcel(TABNAME, o)
  return p
}

func (p *Mtools_tp) WriteInvoiceLines() *Mtools_tp {
  for _, i := range p.invoices {
    var o lmout_tp
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
    o.difImportePago = importePago - importePagoCalc
    if math.Abs(o.difImportePago) < 0.0000009 {
      o.difImportePago = 0.0
    }
    p.buildLineExcel(TABNAME, o)
  }
  return p
}
// ----------------------------- end of file -----------------------------------

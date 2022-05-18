// cps2xlsp.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields. Version for One-TaxCode Invoices per Payment - Auxiliary tools
package bbmwht

import ut "bar8tl/p/rblib"
import "github.com/xuri/excelize/v2"
import "log"
import "math"
import "strconv"

// Auxiliary funtions ----------------------------------------------------------
type Stools_tp struct {
  F               *excelize.File
  g               *excelize.File
  index           int
  recn            int
  AmountDocCurr   float64
  ImportePago     float64
  EffExchangeRate float64
  PaymentData     Line_tp
  Totales         Totales_tp
  ImpuestosP      Payment_tp
  invoices        []linv_tp
  OneTaxPaym      Payment_tp
  firstInvoice    bool
}

func NewStools() *Stools_tp {
  var p Stools_tp
  p.firstInvoice = true
  return &p
}

func (p *Stools_tp) OpenInpExcel(file string) {
  var err error
  p.F, err = excelize.OpenFile(file)
  if err != nil {
    log.Fatal(err)
  }
}

func (p *Stools_tp) CreateOutExcel(tabnam string) {
  p.g = excelize.NewFile()
  p.index = p.g.NewSheet(tabnam)
}

func (p *Stools_tp) WriteOutExcel(fname string) {
  p.g.SetActiveSheet(p.index)
  if err := p.g.SaveAs(fname); err != nil {
    log.Fatal(err)
  }
}

func (p *Stools_tp) StoreDocRel(lin Line_tp, inv Docrel_tp) {
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
  if p.firstInvoice {
    p.OneTaxPaym.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
    p.OneTaxPaym.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
    p.OneTaxPaym.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
    p.OneTaxPaym.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
    p.OneTaxPaym.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
    p.OneTaxPaym.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
    p.firstInvoice = false
  }
}

func (p *Stools_tp) WritePaymentLine() *Stools_tp {
  var o lsout_tp
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
  o.taxTrasladoBase         = ut.Round(p.ImpuestosP.TrasladoP.BaseP, DEC)
  o.taxTrasladoImpuesto     = p.OneTaxPaym.TrasladoP.ImpuestoP
  o.taxTrasladoTipoFactor   = p.OneTaxPaym.TrasladoP.TipoFactorP
  o.taxTrasladoTasaOCuota   = ut.Round(p.OneTaxPaym.TrasladoP.TasaOCuotaP, DEC)
  o.taxTrasladoImporte      = ut.Round(p.ImpuestosP.TrasladoP.ImporteP, DEC)
  if p.ImpuestosP.RetncionP.ImporteP != 0.0 {
    o.taxRetncionBase       = ut.Round(p.ImpuestosP.RetncionP.BaseP, DEC)
    o.taxRetncionImpuesto   = p.OneTaxPaym.RetncionP.ImpuestoP
    o.taxRetncionTipoFactor = p.OneTaxPaym.RetncionP.TipoFactorP
    o.taxRetncionTasaOCuota = ut.Round(p.OneTaxPaym.RetncionP.TasaOCuotaP, DEC)
    o.taxRetncionImporte    = ut.Round(p.ImpuestosP.RetncionP.ImporteP, DEC)
  } else {
    o.taxRetncionBase       = 0.0
    o.taxRetncionImpuesto   = ""
    o.taxRetncionTipoFactor = ""
    o.taxRetncionTasaOCuota = 0.0
    o.taxRetncionImporte    = 0.0
  }
  amountDocCurr, _ := strconv.ParseFloat(p.PaymentData.AmountDocCurr, 64)
  amountDocCurr = ut.Round(amountDocCurr, 2)
  o.difMontoTotalPagos = -1.0 * amountDocCurr - p.Totales.MontoTotalPagos
  if math.Abs(o.difMontoTotalPagos) < 0.0000015 {
    o.difMontoTotalPagos = 0.0
  }
  importePagoCalc := o.taxTrasladoBase + o.taxTrasladoImporte -
                     o.taxRetncionImporte
  importePago, _  := strconv.ParseFloat(p.PaymentData.AmountDocCurr, 64)
  importePago = ut.Round(importePago, 2)
  o.difImportePago = -1.0 * importePago - importePagoCalc
  if math.Abs(o.difImportePago) < 0.0000015 {
    o.difImportePago = 0.0
  }
  p.buildLineExcel(TABNAME, o)
  return p
}

func (p *Stools_tp) WriteInvoiceLines() *Stools_tp {
  for _, i := range p.invoices {
    var o lsout_tp
    o.src = i.src
    o.retencionesIVA          = 0.0
    o.trasladosBaseIVA16      = 0.0
    o.trasladosImpuestoIVA16  = 0.0
    o.trasladosBaseIVA8       = 0.0
    o.trasladosImpuestoIVA8   = 0.0
    o.trasladosBaseIVA0       = 0.0
    o.trasladosImpuestoIVA0   = 0.0
    o.montoTotalPagos         = 0.0
    o.objetoImpuesto          = i.docrel.ObjetoImpDR
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
    importePagoCalc := i.docrel.TrasladoDR.BaseDR +
                       i.docrel.TrasladoDR.ImporteDR -
                       i.docrel.RetncionDR.ImporteDR
    importePago, _ := strconv.ParseFloat(i.src.ImportePago, 64)
    importePago = ut.Round(importePago, 2)
    o.difImportePago = importePago - importePagoCalc
    if math.Abs(o.difImportePago) < 0.0000015 {
      o.difImportePago = 0.0
    }
    p.buildLineExcel(TABNAME, o)
  }
  return p
}
// ----------------------------- end of file -----------------------------------

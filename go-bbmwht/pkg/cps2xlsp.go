// cps2xlsp.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields. Version for One-TaxCode Invoices per Payment - Auxiliary tools
package bbmwht

import "fmt"
import "github.com/xuri/excelize/v2"
import "log"
import "strconv"

// Data types and definitions --------------------------------------------------
type lsout_tp struct {
  src                     Line_tp
  retencionesIVA          string
  trasladosBaseIVA16      string
  trasladosImpuestoIVA16  string
  trasladosBaseIVA8       string
  trasladosImpuestoIVA8   string
  trasladosBaseIVA0       string
  trasladosImpuestoIVA0   string
  montoTotalPagos         string
  objetoImpuesto          string
  taxTrasladoBase         string
  taxTrasladoImpuesto     string
  taxTrasladoTipoFactor   string
  taxTrasladoTasaOCuota   string
  taxTrasladoImporte      string
  taxRetncionBase         string
  taxRetncionImpuesto     string
  taxRetncionTipoFactor   string
  taxRetncionTasaOCuota   string
  taxRetncionImporte      string
}

// Auxiliary funtions ----------------------------------------------------------
type Stools_tp struct {
  F             *excelize.File
  g             *excelize.File
  index         int
  recn          int
  AmountDocCurr float64
  PaymentData   Line_tp
  Totales       Totales_tp
  ImpuestosP    Payment_tp
  invoices      []linv_tp
  OneTaxPaym    Payment_tp
  firstInvoice  bool
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

func (p *Stools_tp) ClearPaymentData() {
  p.Totales.RetencionesIVA           = 0.00
  p.Totales.TrasladosBaseIVA16       = 0.00
  p.Totales.TrasladosImpuestoIVA16   = 0.00
  p.Totales.TrasladosBaseIVA8        = 0.00
  p.Totales.TrasladosImpuestoIVA8    = 0.00
  p.Totales.TrasladosBaseIVA0        = 0.00
  p.Totales.TrasladosImpuestoIVA0    = 0.00
  p.Totales.MontoTotalPagos          = 0.00
  p.ImpuestosP.TrasladoP.BaseP       = 0.00
  p.ImpuestosP.TrasladoP.ImpuestoP   = ""
  p.ImpuestosP.TrasladoP.TipoFactorP = ""
  p.ImpuestosP.TrasladoP.TasaOCuotaP = 0.000000
  p.ImpuestosP.TrasladoP.ImporteP    = 0.00
  p.ImpuestosP.RetncionP.BaseP       = 0.00
  p.ImpuestosP.RetncionP.ImpuestoP   = ""
  p.ImpuestosP.RetncionP.TipoFactorP = ""
  p.ImpuestosP.RetncionP.TasaOCuotaP = 0.000000
  p.ImpuestosP.RetncionP.ImporteP    = 0.00
  p.OneTaxPaym.TrasladoP.ImpuestoP   = ""
  p.OneTaxPaym.TrasladoP.TipoFactorP = ""
  p.OneTaxPaym.TrasladoP.TasaOCuotaP = 0.000000
  p.OneTaxPaym.RetncionP.ImpuestoP   = ""
  p.OneTaxPaym.RetncionP.TipoFactorP = ""
  p.OneTaxPaym.RetncionP.TasaOCuotaP = 0.000000
  p.firstInvoice                     = true
  p.invoices = nil
}

func (p *Stools_tp) GetLineFields(row []string) (l Line_tp) {
  for i, _ := range row {
    switch i {
      case 0  : l.companyCode            = row[i]
      case 1  : l.customer               = row[i]
      case 2  : l.documentNumber         = row[i]
      case 3  : l.DocumentType           = row[i]
      case 4  : l.paymentDateTime        = row[i]
      case 5  : l.clearingDocument       = row[i]
      case 6  : l.AmountDocCurr          = row[i];
        p.AmountDocCurr, _ = strconv.ParseFloat(l.AmountDocCurr, 64)
      case 7  : l.documentCurrency       = row[i]
      case 8  : l.effExchangeRate        = row[i]
      case 9  : l.assignment             = row[i]
      case 10 : l.formaPago              = row[i]
      case 11 : l.noParcialidad          = row[i]
      case 12 : l.importeSaldoAnterior   = row[i]
      case 13 : l.importePago            = row[i]
      case 14 : l.importeSaldoInsoluto   = row[i]
      case 15 : l.tipoRelacion           = row[i]
      case 16 : l.pagoCanceladoDocNumber = row[i]
      case 17 : l.numOperacion           = row[i]
      case 18 : l.rfcBancoOrdenente      = row[i]
      case 19 : l.nombreBancoOrdenante   = row[i]
      case 20 : l.cuentaOrdenante        = row[i]
      case 21 : l.rfcBancoBeneficiario   = row[i]
      case 22 : l.cuentaBeneficiario     = row[i]
      case 23 : l.tipoCadenaPago         = row[i]
      case 24 : l.certificadoPago        = row[i]
      case 25 : l.cadenaPago             = row[i]
      case 26 : l.selloPago              = row[i]
      case 27 : l.TaxCode                = row[i]
    }
  }
  return l
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

func (p *Stools_tp) WriteTitle(lin Line_tp) {
  var o lsout_tp
  o.src = lin
  o.retencionesIVA         = "Retenciones IVA"
  o.trasladosBaseIVA16     = "Traslados Base IVA16"
  o.trasladosImpuestoIVA16 = "Traslados Impuesto IVA16"
  o.trasladosBaseIVA8      = "Traslados Base IVA8"
  o.trasladosImpuestoIVA8  = "Traslados Impuesto IVA8"
  o.trasladosBaseIVA0      = "Traslados Base IVA0"
  o.trasladosImpuestoIVA0  = "Traslados Impuesto IVA0"
  o.montoTotalPagos        = "Monto Total Pagos"
  o.objetoImpuesto         = "Objeto Impuesto"
  o.taxTrasladoBase        = "Tax Traslado Base"
  o.taxTrasladoImpuesto    = "Tax Traslado Impuesto"
  o.taxTrasladoTipoFactor  = "Tax Traslado TipoFactor"
  o.taxTrasladoTasaOCuota  = "Tax Traslado TasaOCuota"
  o.taxTrasladoImporte     = "Tax Traslado Importe"
  o.taxRetncionBase        = "Tax Retencion Base"
  o.taxRetncionImpuesto    = "Tax Retencion Impuesto"
  o.taxRetncionTipoFactor  = "Tax Retencion TipoFactor"
  o.taxRetncionTasaOCuota  = "Tax Retencion TasaOCuota"
  o.taxRetncionImporte     = "Tax Retencion Importe"
  p.buildLineExcel(TABNAME, o)
}

func (p *Stools_tp) WritePaymentLine() *Stools_tp {
  var o lsout_tp
  o.src = p.PaymentData
  o.retencionesIVA          = fmt.Sprintf("%.2f", p.Totales.RetencionesIVA)
  o.trasladosBaseIVA16      = fmt.Sprintf("%.2f", p.Totales.TrasladosBaseIVA16)
  o.trasladosImpuestoIVA16  = fmt.Sprintf("%.2f", p.Totales.TrasladosImpuestoIVA16)
  o.trasladosBaseIVA8       = fmt.Sprintf("%.2f", p.Totales.TrasladosBaseIVA8)
  o.trasladosImpuestoIVA8   = fmt.Sprintf("%.2f", p.Totales.TrasladosImpuestoIVA8)
  o.trasladosBaseIVA0       = fmt.Sprintf("%.2f", p.Totales.TrasladosBaseIVA0)
  o.trasladosImpuestoIVA0   = fmt.Sprintf("%.2f", p.Totales.TrasladosImpuestoIVA0)
  o.montoTotalPagos         = fmt.Sprintf("%.2f", p.Totales.MontoTotalPagos)
  o.objetoImpuesto          = ""
  o.taxTrasladoBase         = fmt.Sprintf("%.2f", p.ImpuestosP.TrasladoP.BaseP)
  o.taxTrasladoImpuesto     = p.OneTaxPaym.TrasladoP.ImpuestoP
  o.taxTrasladoTipoFactor   = p.OneTaxPaym.TrasladoP.TipoFactorP
  o.taxTrasladoTasaOCuota   = fmt.Sprintf("%.2f", p.OneTaxPaym.TrasladoP.TasaOCuotaP)
  o.taxTrasladoImporte      = fmt.Sprintf("%.2f", p.ImpuestosP.TrasladoP.ImporteP)
  if p.ImpuestosP.RetncionP.ImporteP != 0.00 {
    o.taxRetncionBase       = fmt.Sprintf("%.2f", p.ImpuestosP.RetncionP.BaseP)
    o.taxRetncionImpuesto   = p.OneTaxPaym.RetncionP.ImpuestoP
    o.taxRetncionTipoFactor = p.OneTaxPaym.RetncionP.TipoFactorP
    o.taxRetncionTasaOCuota = fmt.Sprintf("%.2f", p.OneTaxPaym.RetncionP.TasaOCuotaP)
    o.taxRetncionImporte    = fmt.Sprintf("%.2f", p.ImpuestosP.RetncionP.ImporteP)
  } else {
    o.taxRetncionBase       = ""
    o.taxRetncionImpuesto   = ""
    o.taxRetncionTipoFactor = ""
    o.taxRetncionTasaOCuota = ""
    o.taxRetncionImporte    = ""
  }
  p.buildLineExcel(TABNAME, o)
  return p
}

func (p *Stools_tp) WriteInvoiceLines() *Stools_tp {
  for _, i := range p.invoices {
    var o lsout_tp
    o.src = i.src
    o.retencionesIVA          = ""
    o.trasladosBaseIVA16      = ""
    o.trasladosImpuestoIVA16  = ""
    o.trasladosBaseIVA8       = ""
    o.trasladosImpuestoIVA8   = ""
    o.trasladosBaseIVA0       = ""
    o.trasladosImpuestoIVA0   = ""
    o.montoTotalPagos         = ""
    o.objetoImpuesto          = i.docrel.ObjetoImpDR
    o.taxTrasladoBase         = fmt.Sprintf("%.2f", i.docrel.TrasladoDR.BaseDR)
    o.taxTrasladoImpuesto     = IMPUESTO
    o.taxTrasladoTipoFactor   = TIPOFACTOR
    o.taxTrasladoTasaOCuota   = fmt.Sprintf("%.2f", i.docrel.TrasladoDR.TasaOCuotaDR)
    o.taxTrasladoImporte      = fmt.Sprintf("%.2f", i.docrel.TrasladoDR.ImporteDR)
    if i.docrel.RetncionDR.ImporteDR != 0.00 {
      o.taxRetncionBase       = fmt.Sprintf("%.2f", i.docrel.RetncionDR.BaseDR)
      o.taxRetncionImpuesto   =                     i.docrel.RetncionDR.ImpuestoDR
      o.taxRetncionTipoFactor =                     i.docrel.RetncionDR.TipoFactorDR
      o.taxRetncionTasaOCuota = fmt.Sprintf("%.2f", i.docrel.RetncionDR.TasaOCuotaDR)
      o.taxRetncionImporte    = fmt.Sprintf("%.2f", i.docrel.RetncionDR.ImporteDR)
    } else {
      o.taxRetncionBase       = ""
      o.taxRetncionImpuesto   = ""
      o.taxRetncionTipoFactor = ""
      o.taxRetncionTasaOCuota = ""
      o.taxRetncionImporte    = ""
    }
    p.buildLineExcel(TABNAME, o)
  }
  return p
}

func (p *Stools_tp) buildLineExcel(tab string, o lsout_tp) {
  p.recn++
  p.g.SetCellValue(tab, fmt.Sprintf("A%d",  p.recn), o.src.companyCode)
  p.g.SetCellValue(tab, fmt.Sprintf("B%d",  p.recn), o.src.customer)
  p.g.SetCellValue(tab, fmt.Sprintf("C%d",  p.recn), o.src.documentNumber)
  p.g.SetCellValue(tab, fmt.Sprintf("D%d",  p.recn), o.src.DocumentType)
  p.g.SetCellValue(tab, fmt.Sprintf("F%d",  p.recn), o.src.paymentDateTime)
  p.g.SetCellValue(tab, fmt.Sprintf("F%d",  p.recn), o.src.clearingDocument)
  p.g.SetCellValue(tab, fmt.Sprintf("G%d",  p.recn), o.src.AmountDocCurr)
  p.g.SetCellValue(tab, fmt.Sprintf("H%d",  p.recn), o.src.documentCurrency)
  p.g.SetCellValue(tab, fmt.Sprintf("I%d",  p.recn), o.src.effExchangeRate)
  p.g.SetCellValue(tab, fmt.Sprintf("J%d",  p.recn), o.src.assignment)
  p.g.SetCellValue(tab, fmt.Sprintf("K%d",  p.recn), o.src.formaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("L%d",  p.recn), o.src.noParcialidad)
  p.g.SetCellValue(tab, fmt.Sprintf("M%d",  p.recn), o.src.importeSaldoAnterior)
  p.g.SetCellValue(tab, fmt.Sprintf("N%d",  p.recn), o.src.importePago)
  p.g.SetCellValue(tab, fmt.Sprintf("O%d",  p.recn), o.src.importeSaldoInsoluto)
  p.g.SetCellValue(tab, fmt.Sprintf("P%d",  p.recn), o.src.tipoRelacion)
  p.g.SetCellValue(tab, fmt.Sprintf("Q%d",  p.recn), o.src.pagoCanceladoDocNumber)
  p.g.SetCellValue(tab, fmt.Sprintf("R%d",  p.recn), o.src.numOperacion)
  p.g.SetCellValue(tab, fmt.Sprintf("S%d",  p.recn), o.src.rfcBancoOrdenente)
  p.g.SetCellValue(tab, fmt.Sprintf("T%d",  p.recn), o.src.nombreBancoOrdenante)
  p.g.SetCellValue(tab, fmt.Sprintf("U%d",  p.recn), o.src.cuentaOrdenante)
  p.g.SetCellValue(tab, fmt.Sprintf("V%d",  p.recn), o.src.rfcBancoBeneficiario)
  p.g.SetCellValue(tab, fmt.Sprintf("W%d",  p.recn), o.src.cuentaBeneficiario)
  p.g.SetCellValue(tab, fmt.Sprintf("X%d",  p.recn), o.src.tipoCadenaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("Y%d",  p.recn), o.src.certificadoPago)
  p.g.SetCellValue(tab, fmt.Sprintf("Z%d",  p.recn), o.src.cadenaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("AA%d", p.recn), o.src.selloPago)
  p.g.SetCellValue(tab, fmt.Sprintf("AB%d", p.recn), o.src.TaxCode)
  p.g.SetCellValue(tab, fmt.Sprintf("AC%d", p.recn), o.retencionesIVA)
  p.g.SetCellValue(tab, fmt.Sprintf("AD%d", p.recn), o.trasladosBaseIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("AE%d", p.recn), o.trasladosImpuestoIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("AF%d", p.recn), o.trasladosBaseIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("AG%d", p.recn), o.trasladosImpuestoIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("AH%d", p.recn), o.trasladosBaseIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("AI%d", p.recn), o.trasladosImpuestoIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("AJ%d", p.recn), o.montoTotalPagos)
  p.g.SetCellValue(tab, fmt.Sprintf("AK%d", p.recn), o.objetoImpuesto)
  p.g.SetCellValue(tab, fmt.Sprintf("AL%d", p.recn), o.taxTrasladoBase)
  p.g.SetCellValue(tab, fmt.Sprintf("AM%d", p.recn), o.taxTrasladoImpuesto)
  p.g.SetCellValue(tab, fmt.Sprintf("AN%d", p.recn), o.taxTrasladoTipoFactor)
  p.g.SetCellValue(tab, fmt.Sprintf("AO%d", p.recn), o.taxTrasladoTasaOCuota)
  p.g.SetCellValue(tab, fmt.Sprintf("AP%d", p.recn), o.taxTrasladoImporte)
  p.g.SetCellValue(tab, fmt.Sprintf("AQ%d", p.recn), o.taxRetncionBase)
  p.g.SetCellValue(tab, fmt.Sprintf("AR%d", p.recn), o.taxRetncionImpuesto)
  p.g.SetCellValue(tab, fmt.Sprintf("AS%d", p.recn), o.taxRetncionTipoFactor)
  p.g.SetCellValue(tab, fmt.Sprintf("AT%d", p.recn), o.taxRetncionTasaOCuota)
  p.g.SetCellValue(tab, fmt.Sprintf("AU%d", p.recn), o.taxRetncionImporte)
}
// ----------------------------- end of file -----------------------------------

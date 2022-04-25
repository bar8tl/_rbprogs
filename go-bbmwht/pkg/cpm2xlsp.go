// cpm2xlsp.go [2022-04-06 BAR8TL] Extend Pagos1.0 EDICOM-file with Pagos2.0
// fields. Version for Multiple-TaxRate Invoices per Payment - Auxiliary tools
package bbmwht

import "fmt"
import "github.com/xuri/excelize/v2"
import "log"
import "strconv"

// Data types and definitions --------------------------------------------------
type lmout_tp struct {
  src                      Line_tp
  retencionesIVA           string
  trasladosBaseIVA16       string
  trasladosImpuestoIVA16   string
  trasladosBaseIVA8        string
  trasladosImpuestoIVA8    string
  trasladosBaseIVA0        string
  trasladosImpuestoIVA0    string
  montoTotalPagos          string
  objetoImpuesto           string
  trasladoBaseDR           string
  trasladoImpuestoDR       string
  trasladoTipoFactorDR     string
  trasladoTasaOCuotaDR     string
  trasladoImporteDR        string
  retncionBaseDR           string
  retncionImpuestoDR       string
  retncionTipoFactorDR     string
  retncionTasaOCuotaDR     string
  retncionImporteDR        string
  trasladoBasePIVA16       string
  trasladoImpuestoPIVA16   string
  trasladoTipoFactorPIVA16 string
  trasladoTasaOCuotaPIVA16 string
  trasladoImportePIVA16    string
  retncionBasePIVA16       string
  retncionImpuestoPIVA16   string
  retncionTipoFactorPIVA16 string
  retncionTasaOCuotaPIVA16 string
  retncionImportePIVA16    string
  trasladoBasePIVA8        string
  trasladoImpuestoPIVA8    string
  trasladoTipoFactorPIVA8  string
  trasladoTasaOCuotaPIVA8  string
  trasladoImportePIVA8     string
  retncionBasePIVA8        string
  retncionImpuestoPIVA8    string
  retncionTipoFactorPIVA8  string
  retncionTasaOCuotaPIVA8  string
  retncionImportePIVA8     string
  trasladoBasePIVA0        string
  trasladoImpuestoPIVA0    string
  trasladoTipoFactorPIVA0  string
  trasladoTasaOCuotaPIVA0  string
  trasladoImportePIVA0     string
  retncionBasePIVA0        string
  retncionImpuestoPIVA0    string
  retncionTipoFactorPIVA0  string
  retncionTasaOCuotaPIVA0  string
  retncionImportePIVA0     string
}

// Auxiliary funtions ----------------------------------------------------------
type Mtools_tp struct {
  F             *excelize.File
  g             *excelize.File
  index         int
  recn          int
  AmountDocCurr float64
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

func (p *Mtools_tp) ClearPaymentData() {
  p.Totales.RetencionesIVA           = 0.00
  p.Totales.TrasladosBaseIVA16       = 0.00
  p.Totales.TrasladosImpuestoIVA16   = 0.00
  p.Totales.TrasladosBaseIVA8        = 0.00
  p.Totales.TrasladosImpuestoIVA8    = 0.00
  p.Totales.TrasladosBaseIVA0        = 0.00
  p.Totales.TrasladosImpuestoIVA0    = 0.00
  p.Totales.MontoTotalPagos          = 0.00
  p.TaxPIVA16.TrasladoP.BaseP        = 0.00
  p.TaxPIVA16.TrasladoP.ImpuestoP    = ""
  p.TaxPIVA16.TrasladoP.TipoFactorP  = ""
  p.TaxPIVA16.TrasladoP.TasaOCuotaP  = 0.000000
  p.TaxPIVA16.TrasladoP.ImporteP     = 0.00
  p.TaxPIVA16.RetncionP.BaseP        = 0.00
  p.TaxPIVA16.RetncionP.ImpuestoP    = ""
  p.TaxPIVA16.RetncionP.TipoFactorP  = ""
  p.TaxPIVA16.RetncionP.TasaOCuotaP  = 0.000000
  p.TaxPIVA16.RetncionP.ImporteP     = 0.00
  p.TaxPIVA8.TrasladoP.BaseP         = 0.00
  p.TaxPIVA8.TrasladoP.ImpuestoP     = ""
  p.TaxPIVA8.TrasladoP.TipoFactorP   = ""
  p.TaxPIVA8.TrasladoP.TasaOCuotaP   = 0.000000
  p.TaxPIVA8.TrasladoP.ImporteP      = 0.00
  p.TaxPIVA8.RetncionP.BaseP         = 0.00
  p.TaxPIVA8.RetncionP.ImpuestoP     = ""
  p.TaxPIVA8.RetncionP.TipoFactorP   = ""
  p.TaxPIVA8.RetncionP.TasaOCuotaP   = 0.000000
  p.TaxPIVA8.RetncionP.ImporteP      = 0.00
  p.TaxPIVA0.TrasladoP.BaseP         = 0.00
  p.TaxPIVA0.TrasladoP.ImpuestoP     = ""
  p.TaxPIVA0.TrasladoP.TipoFactorP   = ""
  p.TaxPIVA0.TrasladoP.TasaOCuotaP   = 0.000000
  p.TaxPIVA0.TrasladoP.ImporteP      = 0.00
  p.TaxPIVA0.RetncionP.BaseP         = 0.00
  p.TaxPIVA0.RetncionP.ImpuestoP     = ""
  p.TaxPIVA0.RetncionP.TipoFactorP   = ""
  p.TaxPIVA0.RetncionP.TasaOCuotaP   = 0.000000
  p.TaxPIVA0.RetncionP.ImporteP      = 0.00
  p.firstInvoTraslIVA16              = true
  p.firstInvoRetenIVA16              = true
  p.firstInvoTraslIVA8               = true
  p.firstInvoRetenIVA8               = true
  p.firstInvoTraslIVA0               = true
  p.firstInvoRetenIVA0               = true
  p.invoices = nil
}

func (p *Mtools_tp) GetLineFields(row []string) (l Line_tp) {
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
  if inv.TrasladoDR.TasaOCuotaDR == 0.160000 {
    if p.firstInvoTraslIVA16 {
      p.TaxPIVA16.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      p.TaxPIVA16.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      p.TaxPIVA16.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      p.firstInvoTraslIVA16 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.160000 {
    if p.firstInvoRetenIVA16 {
      p.TaxPIVA16.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      p.TaxPIVA16.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      p.TaxPIVA16.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      p.firstInvoRetenIVA16 = false
    }
  }
  if inv.TrasladoDR.TasaOCuotaDR == 0.080000 {
    if p.firstInvoTraslIVA8 {
      p.TaxPIVA8.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      p.TaxPIVA8.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      p.TaxPIVA8.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      p.firstInvoTraslIVA8 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.080000 {
    if p.firstInvoRetenIVA8 {
      p.TaxPIVA8.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      p.TaxPIVA8.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      p.TaxPIVA8.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      p.firstInvoRetenIVA8 = false
    }
  }
  if inv.TrasladoDR.TasaOCuotaDR == 0.000000 {
    if p.firstInvoTraslIVA0 {
      p.TaxPIVA0.TrasladoP.ImpuestoP    = inv.TrasladoDR.ImpuestoDR
      p.TaxPIVA0.TrasladoP.TipoFactorP  = inv.TrasladoDR.TipoFactorDR
      p.TaxPIVA0.TrasladoP.TasaOCuotaP  = inv.TrasladoDR.TasaOCuotaDR
      p.firstInvoTraslIVA0 = false
    }
  }
  if inv.RetncionDR.TasaOCuotaDR == 0.000000 {
    if p.firstInvoRetenIVA0 {
      p.TaxPIVA0.RetncionP.ImpuestoP    = inv.RetncionDR.ImpuestoDR
      p.TaxPIVA0.RetncionP.TipoFactorP  = inv.RetncionDR.TipoFactorDR
      p.TaxPIVA0.RetncionP.TasaOCuotaP  = inv.RetncionDR.TasaOCuotaDR
      p.firstInvoRetenIVA0 = false
    }
  }
}

func (p *Mtools_tp) WriteTitle(lin Line_tp) {
  var o lmout_tp
  o.src = lin
  o.retencionesIVA           = "Retenciones IVA"
  o.trasladosBaseIVA16       = "Traslados Base IVA16"
  o.trasladosImpuestoIVA16   = "Traslados Impuesto IVA16"
  o.trasladosBaseIVA8        = "Traslados Base IVA8"
  o.trasladosImpuestoIVA8    = "Traslados Impuesto IVA8"
  o.trasladosBaseIVA0        = "Traslados Base IVA0"
  o.trasladosImpuestoIVA0    = "Traslados Impuesto IVA0"
  o.montoTotalPagos          = "Monto Total Pagos"
  o.objetoImpuesto           = "Objeto Impuesto"
  o.trasladoBaseDR           = "DR Traslado Base"
  o.trasladoImpuestoDR       = "DR Traslado Impuesto"
  o.trasladoTipoFactorDR     = "DR Traslado TipoFactor"
  o.trasladoTasaOCuotaDR     = "DR Traslado TasaOCuota"
  o.trasladoImporteDR        = "DR Traslado Importe"
  o.retncionBaseDR           = "DR Retencion Base"
  o.retncionImpuestoDR       = "DR Retencion Impuesto"
  o.retncionTipoFactorDR     = "DR Retencion TipoFactor"
  o.retncionTasaOCuotaDR     = "DR Retencion TasaOCuota"
  o.retncionImporteDR        = "DR Retencion Importe"
  o.trasladoBasePIVA16       = "P Traslado Base IVA16"
  o.trasladoImpuestoPIVA16   = "P Traslado Impuesto IVA16"
  o.trasladoTipoFactorPIVA16 = "P Traslado TipoFactor IVA16"
  o.trasladoTasaOCuotaPIVA16 = "P Traslado TasaOCuota IVA16"
  o.trasladoImportePIVA16    = "P Traslado Importe IVA16"
  o.retncionBasePIVA16       = "P Retencion Base IVA16"
  o.retncionImpuestoPIVA16   = "P Retencion Impuesto IVA16"
  o.retncionTipoFactorPIVA16 = "P Retencion TipoFactor IVA16"
  o.retncionTasaOCuotaPIVA16 = "P Retencion TasaOCuota IVA16"
  o.retncionImportePIVA16    = "P Retencion Importe IVA16"
  o.trasladoBasePIVA8        = "P Traslado Base IVA8"
  o.trasladoImpuestoPIVA8    = "P Traslado Impuesto IVA8"
  o.trasladoTipoFactorPIVA8  = "P Traslado TipoFactor IVA8"
  o.trasladoTasaOCuotaPIVA8  = "P Traslado TasaOCuota IVA8"
  o.trasladoImportePIVA8     = "P Traslado Importe IVA8"
  o.retncionBasePIVA8        = "P Retencion Base IVA8"
  o.retncionImpuestoPIVA8    = "P Retencion Impuesto IVA8"
  o.retncionTipoFactorPIVA8  = "P Retencion TipoFactor IVA8"
  o.retncionTasaOCuotaPIVA8  = "P Retencion TasaOCuota IVA8"
  o.retncionImportePIVA8     = "P Retencion Importe IVA8"
  o.trasladoBasePIVA0        = "P Traslado Base IVA0"
  o.trasladoImpuestoPIVA0    = "P Traslado Impuesto IVA0"
  o.trasladoTipoFactorPIVA0  = "P Traslado TipoFactor IVA0"
  o.trasladoTasaOCuotaPIVA0  = "P Traslado TasaOCuota IVA0"
  o.trasladoImportePIVA0     = "P Traslado Importe IVA0"
  o.retncionBasePIVA0        = "P Retencion Base IVA0"
  o.retncionImpuestoPIVA0    = "P Retencion Impuesto IVA0"
  o.retncionTipoFactorPIVA0  = "P Retencion TipoFactor IVA0"
  o.retncionTasaOCuotaPIVA0  = "P Retencion TasaOCuota IVA0"
  o.retncionImportePIVA0     = "P Retencion Importe IVA0"
  p.buildLineExcel(TABNAME, o)
}

func (p *Mtools_tp) WritePaymentLine() *Mtools_tp {
  var o lmout_tp
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
  if p.TaxPIVA16.TrasladoP.BaseP != 0.00 {
    o.trasladoBasePIVA16       = fmt.Sprintf("%.2f", p.TaxPIVA16.TrasladoP.BaseP)
    o.trasladoImpuestoPIVA16   = p.TaxPIVA16.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA16 = p.TaxPIVA16.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA16 = fmt.Sprintf("%.2f", p.TaxPIVA16.TrasladoP.TasaOCuotaP)
    o.trasladoImportePIVA16    = fmt.Sprintf("%.2f", p.TaxPIVA16.TrasladoP.ImporteP)
  } else {
    o.trasladoBasePIVA16       = ""
    o.trasladoImpuestoPIVA16   = ""
    o.trasladoTipoFactorPIVA16 = ""
    o.trasladoTasaOCuotaPIVA16 = ""
    o.trasladoImportePIVA16    = ""
  }
  if p.TaxPIVA16.RetncionP.BaseP != 0.00 {
    o.retncionBasePIVA16       = fmt.Sprintf("%.2f", p.TaxPIVA16.RetncionP.BaseP)
    o.retncionImpuestoPIVA16   = p.TaxPIVA16.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA16 = p.TaxPIVA16.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA16 = fmt.Sprintf("%.2f", p.TaxPIVA16.RetncionP.TasaOCuotaP)
    o.retncionImportePIVA16    = fmt.Sprintf("%.2f", p.TaxPIVA16.RetncionP.ImporteP)
  } else {
    o.retncionBasePIVA16       = ""
    o.retncionImpuestoPIVA16   = ""
    o.retncionTipoFactorPIVA16 = ""
    o.retncionTasaOCuotaPIVA16 = ""
    o.retncionImportePIVA16    = ""
  }
  if p.TaxPIVA8.TrasladoP.BaseP != 0.00 {
    o.trasladoBasePIVA8        = fmt.Sprintf("%.2f", p.TaxPIVA8.TrasladoP.BaseP)
    o.trasladoImpuestoPIVA8    = p.TaxPIVA8.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA8  = p.TaxPIVA8.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA8  = fmt.Sprintf("%.2f", p.TaxPIVA8.TrasladoP.TasaOCuotaP)
    o.trasladoImportePIVA8     = fmt.Sprintf("%.2f", p.TaxPIVA8.TrasladoP.ImporteP)
  } else {
    o.trasladoBasePIVA8        = ""
    o.trasladoImpuestoPIVA8    = ""
    o.trasladoTipoFactorPIVA8  = ""
    o.trasladoTasaOCuotaPIVA8  = ""
    o.trasladoImportePIVA8     = ""
  }
  if p.TaxPIVA8.RetncionP.BaseP != 0.00 {
    o.retncionBasePIVA8        = fmt.Sprintf("%.2f", p.TaxPIVA8.RetncionP.BaseP)
    o.retncionImpuestoPIVA8    = p.TaxPIVA8.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA8  = p.TaxPIVA8.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA8  = fmt.Sprintf("%.2f", p.TaxPIVA8.RetncionP.TasaOCuotaP)
    o.retncionImportePIVA8     = fmt.Sprintf("%.2f", p.TaxPIVA8.RetncionP.ImporteP)
  } else {
    o.retncionBasePIVA8        = ""
    o.retncionImpuestoPIVA8    = ""
    o.retncionTipoFactorPIVA8  = ""
    o.retncionTasaOCuotaPIVA8  = ""
    o.retncionImportePIVA8     = ""
  }
  if p.TaxPIVA0.TrasladoP.BaseP != 0.00 {
    o.trasladoBasePIVA0        = fmt.Sprintf("%.2f", p.TaxPIVA0.TrasladoP.BaseP)
    o.trasladoImpuestoPIVA0    = p.TaxPIVA0.TrasladoP.ImpuestoP
    o.trasladoTipoFactorPIVA0  = p.TaxPIVA0.TrasladoP.TipoFactorP
    o.trasladoTasaOCuotaPIVA0  = fmt.Sprintf("%.2f", p.TaxPIVA0.TrasladoP.TasaOCuotaP)
    o.trasladoImportePIVA0     = fmt.Sprintf("%.2f", p.TaxPIVA0.TrasladoP.ImporteP)
  } else {
    o.trasladoBasePIVA0        = ""
    o.trasladoImpuestoPIVA0    = ""
    o.trasladoTipoFactorPIVA0  = ""
    o.trasladoTasaOCuotaPIVA0  = ""
    o.trasladoImportePIVA0     = ""
  }
  if p.TaxPIVA0.RetncionP.BaseP != 0.00 {
    o.retncionBasePIVA0        = fmt.Sprintf("%.2f", p.TaxPIVA0.RetncionP.BaseP)
    o.retncionImpuestoPIVA0    = p.TaxPIVA0.RetncionP.ImpuestoP
    o.retncionTipoFactorPIVA0  = p.TaxPIVA0.RetncionP.TipoFactorP
    o.retncionTasaOCuotaPIVA0  = fmt.Sprintf("%.2f", p.TaxPIVA0.RetncionP.TasaOCuotaP)
    o.retncionImportePIVA0     = fmt.Sprintf("%.2f", p.TaxPIVA0.RetncionP.ImporteP)
  } else {
    o.retncionBasePIVA0        = ""
    o.retncionImpuestoPIVA0    = ""
    o.retncionTipoFactorPIVA0  = ""
    o.retncionTasaOCuotaPIVA0  = ""
    o.retncionImportePIVA0     = ""
  }
  p.buildLineExcel(TABNAME, o)
  return p
}

func (p *Mtools_tp) WriteInvoiceLines() *Mtools_tp {
  for _, i := range p.invoices {
    var o lmout_tp
    o.src = i.src
    o.retencionesIVA         = ""
    o.trasladosBaseIVA16     = ""
    o.trasladosImpuestoIVA16 = ""
    o.trasladosBaseIVA8      = ""
    o.trasladosImpuestoIVA8  = ""
    o.trasladosBaseIVA0      = ""
    o.trasladosImpuestoIVA0  = ""
    o.montoTotalPagos        = ""
    o.objetoImpuesto         = i.docrel.ObjetoImpDR
    o.trasladoBaseDR         = fmt.Sprintf("%.2f", i.docrel.TrasladoDR.BaseDR)
    o.trasladoImpuestoDR     = IMPUESTO
    o.trasladoTipoFactorDR   = TIPOFACTOR
    o.trasladoTasaOCuotaDR   = fmt.Sprintf("%.2f", i.docrel.TrasladoDR.TasaOCuotaDR)
    o.trasladoImporteDR      = fmt.Sprintf("%.2f", i.docrel.TrasladoDR.ImporteDR)
    if i.docrel.RetncionDR.ImporteDR != 0.00 {
      o.retncionBaseDR       = fmt.Sprintf("%.2f", i.docrel.RetncionDR.BaseDR)
      o.retncionImpuestoDR   =                     i.docrel.RetncionDR.ImpuestoDR
      o.retncionTipoFactorDR =                     i.docrel.RetncionDR.TipoFactorDR
      o.retncionTasaOCuotaDR = fmt.Sprintf("%.2f", i.docrel.RetncionDR.TasaOCuotaDR)
      o.retncionImporteDR    = fmt.Sprintf("%.2f", i.docrel.RetncionDR.ImporteDR)
    } else {
      o.retncionBaseDR       = ""
      o.retncionImpuestoDR   = ""
      o.retncionTipoFactorDR = ""
      o.retncionTasaOCuotaDR = ""
      o.retncionImporteDR    = ""
    }
    p.buildLineExcel(TABNAME, o)
  }
  return p
}

func (p *Mtools_tp) buildLineExcel(tab string, o lmout_tp) {
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
  p.g.SetCellValue(tab, fmt.Sprintf("AL%d", p.recn), o.trasladoBaseDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AM%d", p.recn), o.trasladoImpuestoDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AN%d", p.recn), o.trasladoTipoFactorDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AO%d", p.recn), o.trasladoTasaOCuotaDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AP%d", p.recn), o.trasladoImporteDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AQ%d", p.recn), o.retncionBaseDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AR%d", p.recn), o.retncionImpuestoDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AS%d", p.recn), o.retncionTipoFactorDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AT%d", p.recn), o.retncionTasaOCuotaDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AU%d", p.recn), o.retncionImporteDR)
  p.g.SetCellValue(tab, fmt.Sprintf("AV%d", p.recn), o.trasladoBasePIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("AW%d", p.recn), o.trasladoImpuestoPIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("AX%d", p.recn), o.trasladoTipoFactorPIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("AY%d", p.recn), o.trasladoTasaOCuotaPIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("AZ%d", p.recn), o.trasladoImportePIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("BA%d", p.recn), o.retncionBasePIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("BB%d", p.recn), o.retncionImpuestoPIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("BC%d", p.recn), o.retncionTipoFactorPIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("BD%d", p.recn), o.retncionTasaOCuotaPIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("BE%d", p.recn), o.retncionImportePIVA16)
  p.g.SetCellValue(tab, fmt.Sprintf("BF%d", p.recn), o.trasladoBasePIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BG%d", p.recn), o.trasladoImpuestoPIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BH%d", p.recn), o.trasladoTipoFactorPIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BI%d", p.recn), o.trasladoTasaOCuotaPIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BJ%d", p.recn), o.trasladoImportePIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BK%d", p.recn), o.retncionBasePIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BL%d", p.recn), o.retncionImpuestoPIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BM%d", p.recn), o.retncionTipoFactorPIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BN%d", p.recn), o.retncionTasaOCuotaPIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BO%d", p.recn), o.retncionImportePIVA8)
  p.g.SetCellValue(tab, fmt.Sprintf("BP%d", p.recn), o.trasladoBasePIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BQ%d", p.recn), o.trasladoImpuestoPIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BR%d", p.recn), o.trasladoTipoFactorPIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BS%d", p.recn), o.trasladoTasaOCuotaPIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BT%d", p.recn), o.trasladoImportePIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BU%d", p.recn), o.retncionBasePIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BV%d", p.recn), o.retncionImpuestoPIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BW%d", p.recn), o.retncionTipoFactorPIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BX%d", p.recn), o.retncionTasaOCuotaPIVA0)
  p.g.SetCellValue(tab, fmt.Sprintf("BY%d", p.recn), o.retncionImportePIVA0)
}
// ----------------------------- end of file -----------------------------------

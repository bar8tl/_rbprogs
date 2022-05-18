// dtypes.go [2022-04-06 BAR8TL] Data Types to Extend Pagos1.0 EDICOM-file with
// Pagos2.0 fields. Auxiliary tools
package bbmwht

import ut "bar8tl/p/rblib"
import "fmt"
import "strconv"

// Data types and definitions --------------------------------------------------
const IMPUESTO       = "002"
const TIPOFACTOR     = "Tasa"
const OBJETOIMPUESTO = "02"
const TABNAME        = "Sheet1"
const DEC            = 6

type Line_tp struct {
  companyCode              string
  customer                 string
  documentNumber           string
  DocumentType             string
  paymentDateTime          string
  clearingDocument         string
  AmountDocCurr            string // Payment total amount
  documentCurrency         string
  EffExchangeRate          string // Exchange rate
  assignment               string
  formaPago                string
  noParcialidad            string
  importeSaldoAnterior     string
  ImportePago              string // Payment amount corresponding to an invoice
  importeSaldoInsoluto     string
  tipoRelacion             string
  pagoCanceladoDocNumber   string
  numOperacion             string
  rfcBancoOrdenente        string
  nombreBancoOrdenante     string
  cuentaOrdenante          string
  rfcBancoBeneficiario     string
  cuentaBeneficiario       string
  tipoCadenaPago           string
  certificadoPago          string
  cadenaPago               string
  selloPago                string
  TaxCode                  string
}

type linv_tp struct {
  src                      Line_tp
  docrel                   Docrel_tp
}

type Totales_tp struct {
  RetencionesIVA           float64
  TrasladosBaseIVA16       float64
  TrasladosImpuestoIVA16   float64
  TrasladosBaseIVA8        float64
  TrasladosImpuestoIVA8    float64
  TrasladosBaseIVA0        float64
  TrasladosImpuestoIVA0    float64
  MontoTotalPagos          float64
}

type Docrel_tp struct {
  ObjetoImpDR              string
  TrasladoDR               TaxesDR_tp
  RetncionDR               TaxesDR_tp
}

type TaxesDR_tp struct {
  BaseDR                   float64
  ImpuestoDR               string
  TipoFactorDR             string
  TasaOCuotaDR             float64
  ImporteDR                float64
}

type Payment_tp struct {
  TrasladoP                TaxesP_tp
  RetncionP                TaxesP_tp
}

type TaxesP_tp struct {
  BaseP                    float64
  ImpuestoP                string
  TipoFactorP              string
  TasaOCuotaP              float64
  ImporteP                 float64
}

type lsout_tp struct {
  src                      Line_tp
  retencionesIVA           float64
  trasladosBaseIVA16       float64
  trasladosImpuestoIVA16   float64
  trasladosBaseIVA8        float64
  trasladosImpuestoIVA8    float64
  trasladosBaseIVA0        float64
  trasladosImpuestoIVA0    float64
  montoTotalPagos          float64
  objetoImpuesto           string
  taxTrasladoBase          float64
  taxTrasladoImpuesto      string
  taxTrasladoTipoFactor    string
  taxTrasladoTasaOCuota    float64
  taxTrasladoImporte       float64
  taxRetncionBase          float64
  taxRetncionImpuesto      string
  taxRetncionTipoFactor    string
  taxRetncionTasaOCuota    float64
  taxRetncionImporte       float64
  difMontoTotalPagos       float64
  difImportePago           float64
}

type lmout_tp struct {
  src                      Line_tp
  retencionesIVA           float64
  trasladosBaseIVA16       float64
  trasladosImpuestoIVA16   float64
  trasladosBaseIVA8        float64
  trasladosImpuestoIVA8    float64
  trasladosBaseIVA0        float64
  trasladosImpuestoIVA0    float64
  montoTotalPagos          float64
  objetoImpuesto           string
  trasladoBaseDR           float64
  trasladoImpuestoDR       string
  trasladoTipoFactorDR     string
  trasladoTasaOCuotaDR     float64
  trasladoImporteDR        float64
  retncionBaseDR           float64
  retncionImpuestoDR       string
  retncionTipoFactorDR     string
  retncionTasaOCuotaDR     float64
  retncionImporteDR        float64
  trasladoBasePIVA16       float64
  trasladoImpuestoPIVA16   string
  trasladoTipoFactorPIVA16 string
  trasladoTasaOCuotaPIVA16 float64
  trasladoImportePIVA16    float64
  retncionBasePIVA16       float64
  retncionImpuestoPIVA16   string
  retncionTipoFactorPIVA16 string
  retncionTasaOCuotaPIVA16 float64
  retncionImportePIVA16    float64
  trasladoBasePIVA8        float64
  trasladoImpuestoPIVA8    string
  trasladoTipoFactorPIVA8  string
  trasladoTasaOCuotaPIVA8  float64
  trasladoImportePIVA8     float64
  retncionBasePIVA8        float64
  retncionImpuestoPIVA8    string
  retncionTipoFactorPIVA8  string
  retncionTasaOCuotaPIVA8  float64
  retncionImportePIVA8     float64
  trasladoBasePIVA0        float64
  trasladoImpuestoPIVA0    string
  trasladoTipoFactorPIVA0  string
  trasladoTasaOCuotaPIVA0  float64
  trasladoImportePIVA0     float64
  retncionBasePIVA0        float64
  retncionImpuestoPIVA0    string
  retncionTipoFactorPIVA0  string
  retncionTasaOCuotaPIVA0  float64
  retncionImportePIVA0     float64
  difMontoTotalPagos       float64
  difImportePago           float64
}

type lstit_tp struct {
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
  taxTrasladoBase          string
  taxTrasladoImpuesto      string
  taxTrasladoTipoFactor    string
  taxTrasladoTasaOCuota    string
  taxTrasladoImporte       string
  taxRetncionBase          string
  taxRetncionImpuesto      string
  taxRetncionTipoFactor    string
  taxRetncionTasaOCuota    string
  taxRetncionImporte       string
  difMontoTotalPagos       string
  difImportePago           string
}

type lmtit_tp struct {
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
  difMontoTotalPagos       string
  difImportePago           string
}

// Functions to perform linear procedures --------------------------------------
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
        p.AmountDocCurr = ut.Round(p.AmountDocCurr, 6)
      case 7  : l.documentCurrency       = row[i]
      case 8  : l.EffExchangeRate        = row[i]
      case 9  : l.assignment             = row[i]
      case 10 : l.formaPago              = row[i]
      case 11 : l.noParcialidad          = row[i]
      case 12 : l.importeSaldoAnterior   = row[i]
      case 13 : l.ImportePago            = row[i];
        p.ImportePago, _ = strconv.ParseFloat(l.ImportePago, 64)
        p.ImportePago = ut.Round(p.ImportePago, 6)
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
        p.AmountDocCurr = ut.Round(p.AmountDocCurr, 6)
      case 7  : l.documentCurrency       = row[i]
      case 8  : l.EffExchangeRate        = row[i]
      case 9  : l.assignment             = row[i]
      case 10 : l.formaPago              = row[i]
      case 11 : l.noParcialidad          = row[i]
      case 12 : l.importeSaldoAnterior   = row[i]
      case 13 : l.ImportePago            = row[i];
        p.ImportePago, _ = strconv.ParseFloat(l.ImportePago, 64)
        p.ImportePago = ut.Round(p.ImportePago, 6)
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

func (p *Stools_tp) ClearPaymentData() {
  p.Totales.RetencionesIVA           = 0.0
  p.Totales.TrasladosBaseIVA16       = 0.0
  p.Totales.TrasladosImpuestoIVA16   = 0.0
  p.Totales.TrasladosBaseIVA8        = 0.0
  p.Totales.TrasladosImpuestoIVA8    = 0.0
  p.Totales.TrasladosBaseIVA0        = 0.0
  p.Totales.TrasladosImpuestoIVA0    = 0.0
  p.Totales.MontoTotalPagos          = 0.0
  p.ImpuestosP.TrasladoP.BaseP       = 0.0
  p.ImpuestosP.TrasladoP.ImpuestoP   = ""
  p.ImpuestosP.TrasladoP.TipoFactorP = ""
  p.ImpuestosP.TrasladoP.TasaOCuotaP = 0.0
  p.ImpuestosP.TrasladoP.ImporteP    = 0.0
  p.ImpuestosP.RetncionP.BaseP       = 0.0
  p.ImpuestosP.RetncionP.ImpuestoP   = ""
  p.ImpuestosP.RetncionP.TipoFactorP = ""
  p.ImpuestosP.RetncionP.TasaOCuotaP = 0.0
  p.ImpuestosP.RetncionP.ImporteP    = 0.0
  p.OneTaxPaym.TrasladoP.ImpuestoP   = ""
  p.OneTaxPaym.TrasladoP.TipoFactorP = ""
  p.OneTaxPaym.TrasladoP.TasaOCuotaP = 0.0
  p.OneTaxPaym.RetncionP.ImpuestoP   = ""
  p.OneTaxPaym.RetncionP.TipoFactorP = ""
  p.OneTaxPaym.RetncionP.TasaOCuotaP = 0.0
  p.firstInvoice                     = true
  p.invoices = nil
}

func (p *Mtools_tp) ClearPaymentData() {
  p.Totales.RetencionesIVA           = 0.0
  p.Totales.TrasladosBaseIVA16       = 0.0
  p.Totales.TrasladosImpuestoIVA16   = 0.0
  p.Totales.TrasladosBaseIVA8        = 0.0
  p.Totales.TrasladosImpuestoIVA8    = 0.0
  p.Totales.TrasladosBaseIVA0        = 0.0
  p.Totales.TrasladosImpuestoIVA0    = 0.0
  p.Totales.MontoTotalPagos          = 0.0
  p.TaxPIVA16.TrasladoP.BaseP        = 0.0
  p.TaxPIVA16.TrasladoP.ImpuestoP    = ""
  p.TaxPIVA16.TrasladoP.TipoFactorP  = ""
  p.TaxPIVA16.TrasladoP.TasaOCuotaP  = 0.0
  p.TaxPIVA16.TrasladoP.ImporteP     = 0.0
  p.TaxPIVA16.RetncionP.BaseP        = 0.0
  p.TaxPIVA16.RetncionP.ImpuestoP    = ""
  p.TaxPIVA16.RetncionP.TipoFactorP  = ""
  p.TaxPIVA16.RetncionP.TasaOCuotaP  = 0.0
  p.TaxPIVA16.RetncionP.ImporteP     = 0.0
  p.TaxPIVA8.TrasladoP.BaseP         = 0.0
  p.TaxPIVA8.TrasladoP.ImpuestoP     = ""
  p.TaxPIVA8.TrasladoP.TipoFactorP   = ""
  p.TaxPIVA8.TrasladoP.TasaOCuotaP   = 0.0
  p.TaxPIVA8.TrasladoP.ImporteP      = 0.0
  p.TaxPIVA8.RetncionP.BaseP         = 0.0
  p.TaxPIVA8.RetncionP.ImpuestoP     = ""
  p.TaxPIVA8.RetncionP.TipoFactorP   = ""
  p.TaxPIVA8.RetncionP.TasaOCuotaP   = 0.0
  p.TaxPIVA8.RetncionP.ImporteP      = 0.0
  p.TaxPIVA0.TrasladoP.BaseP         = 0.0
  p.TaxPIVA0.TrasladoP.ImpuestoP     = ""
  p.TaxPIVA0.TrasladoP.TipoFactorP   = ""
  p.TaxPIVA0.TrasladoP.TasaOCuotaP   = 0.0
  p.TaxPIVA0.TrasladoP.ImporteP      = 0.0
  p.TaxPIVA0.RetncionP.BaseP         = 0.0
  p.TaxPIVA0.RetncionP.ImpuestoP     = ""
  p.TaxPIVA0.RetncionP.TipoFactorP   = ""
  p.TaxPIVA0.RetncionP.TasaOCuotaP   = 0.0
  p.TaxPIVA0.RetncionP.ImporteP      = 0.0
  p.firstInvoTraslIVA16              = true
  p.firstInvoRetenIVA16              = true
  p.firstInvoTraslIVA8               = true
  p.firstInvoRetenIVA8               = true
  p.firstInvoTraslIVA0               = true
  p.firstInvoRetenIVA0               = true
  p.invoices = nil
}

func (p *Stools_tp) WriteTitle(lin Line_tp) {
  var o lstit_tp
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
  o.taxTrasladoBase          = "Tax Traslado Base"
  o.taxTrasladoImpuesto      = "Tax Traslado Impuesto"
  o.taxTrasladoTipoFactor    = "Tax Traslado TipoFactor"
  o.taxTrasladoTasaOCuota    = "Tax Traslado TasaOCuota"
  o.taxTrasladoImporte       = "Tax Traslado Importe"
  o.taxRetncionBase          = "Tax Retencion Base"
  o.taxRetncionImpuesto      = "Tax Retencion Impuesto"
  o.taxRetncionTipoFactor    = "Tax Retencion TipoFactor"
  o.taxRetncionTasaOCuota    = "Tax Retencion TasaOCuota"
  o.taxRetncionImporte       = "Tax Retencion Importe"
  o.difMontoTotalPagos       = "Diff Monto Total Pagos"
  o.difImportePago           = "Diff Importe Pago"
  p.buildTitleExcel(TABNAME, o)
}

func (p *Mtools_tp) WriteTitle(lin Line_tp) {
  var o lmtit_tp
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
  o.difMontoTotalPagos       = "Diff Monto Total Pagos"
  o.difImportePago           = "Diff Importe Pago"
  p.buildTitleExcel(TABNAME, o)
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
  p.g.SetCellValue(tab, fmt.Sprintf("I%d",  p.recn), o.src.EffExchangeRate)
  p.g.SetCellValue(tab, fmt.Sprintf("J%d",  p.recn), o.src.assignment)
  p.g.SetCellValue(tab, fmt.Sprintf("K%d",  p.recn), o.src.formaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("L%d",  p.recn), o.src.noParcialidad)
  p.g.SetCellValue(tab, fmt.Sprintf("M%d",  p.recn), o.src.importeSaldoAnterior)
  p.g.SetCellValue(tab, fmt.Sprintf("N%d",  p.recn), o.src.ImportePago)
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
  p.g.SetCellValue(tab, fmt.Sprintf("AV%d", p.recn), o.difMontoTotalPagos)
  p.g.SetCellValue(tab, fmt.Sprintf("AW%d", p.recn), o.difImportePago)
}

func (p *Stools_tp) buildTitleExcel(tab string, o lstit_tp) {
  p.recn++
  p.g.SetCellValue(tab, fmt.Sprintf("A%d",  p.recn), o.src.companyCode)
  p.g.SetCellValue(tab, fmt.Sprintf("B%d",  p.recn), o.src.customer)
  p.g.SetCellValue(tab, fmt.Sprintf("C%d",  p.recn), o.src.documentNumber)
  p.g.SetCellValue(tab, fmt.Sprintf("D%d",  p.recn), o.src.DocumentType)
  p.g.SetCellValue(tab, fmt.Sprintf("F%d",  p.recn), o.src.paymentDateTime)
  p.g.SetCellValue(tab, fmt.Sprintf("F%d",  p.recn), o.src.clearingDocument)
  p.g.SetCellValue(tab, fmt.Sprintf("G%d",  p.recn), o.src.AmountDocCurr)
  p.g.SetCellValue(tab, fmt.Sprintf("H%d",  p.recn), o.src.documentCurrency)
  p.g.SetCellValue(tab, fmt.Sprintf("I%d",  p.recn), o.src.EffExchangeRate)
  p.g.SetCellValue(tab, fmt.Sprintf("J%d",  p.recn), o.src.assignment)
  p.g.SetCellValue(tab, fmt.Sprintf("K%d",  p.recn), o.src.formaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("L%d",  p.recn), o.src.noParcialidad)
  p.g.SetCellValue(tab, fmt.Sprintf("M%d",  p.recn), o.src.importeSaldoAnterior)
  p.g.SetCellValue(tab, fmt.Sprintf("N%d",  p.recn), o.src.ImportePago)
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
  p.g.SetCellValue(tab, fmt.Sprintf("AV%d", p.recn), o.difMontoTotalPagos)
  p.g.SetCellValue(tab, fmt.Sprintf("AW%d", p.recn), o.difImportePago)
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
  p.g.SetCellValue(tab, fmt.Sprintf("I%d",  p.recn), o.src.EffExchangeRate)
  p.g.SetCellValue(tab, fmt.Sprintf("J%d",  p.recn), o.src.assignment)
  p.g.SetCellValue(tab, fmt.Sprintf("K%d",  p.recn), o.src.formaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("L%d",  p.recn), o.src.noParcialidad)
  p.g.SetCellValue(tab, fmt.Sprintf("M%d",  p.recn), o.src.importeSaldoAnterior)
  p.g.SetCellValue(tab, fmt.Sprintf("N%d",  p.recn), o.src.ImportePago)
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
  p.g.SetCellValue(tab, fmt.Sprintf("BZ%d", p.recn), o.difMontoTotalPagos)
  p.g.SetCellValue(tab, fmt.Sprintf("CA%d", p.recn), o.difImportePago)
}

func (p *Mtools_tp) buildTitleExcel(tab string, o lmtit_tp) {
  p.recn++
  p.g.SetCellValue(tab, fmt.Sprintf("A%d",  p.recn), o.src.companyCode)
  p.g.SetCellValue(tab, fmt.Sprintf("B%d",  p.recn), o.src.customer)
  p.g.SetCellValue(tab, fmt.Sprintf("C%d",  p.recn), o.src.documentNumber)
  p.g.SetCellValue(tab, fmt.Sprintf("D%d",  p.recn), o.src.DocumentType)
  p.g.SetCellValue(tab, fmt.Sprintf("F%d",  p.recn), o.src.paymentDateTime)
  p.g.SetCellValue(tab, fmt.Sprintf("F%d",  p.recn), o.src.clearingDocument)
  p.g.SetCellValue(tab, fmt.Sprintf("G%d",  p.recn), o.src.AmountDocCurr)
  p.g.SetCellValue(tab, fmt.Sprintf("H%d",  p.recn), o.src.documentCurrency)
  p.g.SetCellValue(tab, fmt.Sprintf("I%d",  p.recn), o.src.EffExchangeRate)
  p.g.SetCellValue(tab, fmt.Sprintf("J%d",  p.recn), o.src.assignment)
  p.g.SetCellValue(tab, fmt.Sprintf("K%d",  p.recn), o.src.formaPago)
  p.g.SetCellValue(tab, fmt.Sprintf("L%d",  p.recn), o.src.noParcialidad)
  p.g.SetCellValue(tab, fmt.Sprintf("M%d",  p.recn), o.src.importeSaldoAnterior)
  p.g.SetCellValue(tab, fmt.Sprintf("N%d",  p.recn), o.src.ImportePago)
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
  p.g.SetCellValue(tab, fmt.Sprintf("BZ%d", p.recn), o.difMontoTotalPagos)
  p.g.SetCellValue(tab, fmt.Sprintf("CA%d", p.recn), o.difImportePago)
}
// ----------------------------- end of file -----------------------------------

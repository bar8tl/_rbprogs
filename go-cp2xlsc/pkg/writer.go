// writer.go [2022-04-06 BAR8TL] Excel file EDICOM writer functions
package cp2xlsc

import "github.com/xuri/excelize/v2"
import "fmt"
import "log"

var f1    *excelize.File
var f2    *excelize.File
var recn   int
var flfil  string
var flnam  string
var flext  string
var outpt  string
var ONE    string
var MANY   string
var FULL   string
var TAB    string

type ltit_tp struct {
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
  taxTrasladoBase          string // One tax data group
  taxTrasladoImpuesto      string
  taxTrasladoTipoFactor    string
  taxTrasladoTasaOCuota    string
  taxTrasladoImporte       string
  taxRetncionBase          string
  taxRetncionImpuesto      string
  taxRetncionTipoFactor    string
  taxRetncionTasaOCuota    string
  taxRetncionImporte       string
  trasladoBaseDR           string // Three tax data group
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
  difMontoTotalPagos       string // Common fields for differences
  difImportePago1          string
  difImportePago3          string
}

type lout_tp struct {
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
  taxTrasladoBase          float64 // One tax data group
  taxTrasladoImpuesto      string
  taxTrasladoTipoFactor    string
  taxTrasladoTasaOCuota    float64
  taxTrasladoImporte       float64
  taxRetncionBase          float64
  taxRetncionImpuesto      string
  taxRetncionTipoFactor    string
  taxRetncionTasaOCuota    float64
  taxRetncionImporte       float64
  trasladoBaseDR           float64 // Three tax data group
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
  difMontoTotalPagos       float64 // Common fields for differences
  difImportePago1          float64
  difImportePago3          float64
}

type Writer_tp struct {
  index1 int
  index2 int
}

func NewWriter(s Settings_tp) *Writer_tp {
  var w Writer_tp
  recn  = 0
  flfil = s.Flfil
  flnam = s.Flnam
  flext = s.Flext
  outpt = s.Outpt
  ONE   = s.Konst.ONE
  MANY  = s.Konst.MANY
  FULL  = s.Konst.FULL
  TAB   = s.Konst.TAB
  return &w
}

func (w *Writer_tp) CreateOutExcel() {
  f1 = excelize.NewFile()
  f2 = excelize.NewFile()
  w.index1 = f1.NewSheet(TAB)
  w.index2 = f2.NewSheet(TAB)
}

func (w *Writer_tp) ProduceExcelOutput(dir string) {
  if outpt == ONE  || outpt == FULL {
    f1.SetActiveSheet(w.index1)
    if err := f1.SaveAs(dir+flnam+"-s"+flext); err != nil {
      log.Fatal(err)
    }
    RenameOutFile(dir, flnam+"-s", flext)
  }
  if outpt == MANY || outpt == FULL {
    f2.SetActiveSheet(w.index2)
    if err := f2.SaveAs(dir+flnam+"-m"+flext); err != nil {
      log.Fatal(err)
    }
    RenameOutFile(dir, flnam+"-m", flext)
  }
  RenameInpFile(dir, flnam, flext)
}

func (w *Writer_tp) FetchTitle(lin Line_tp) {
  var o ltit_tp
  recn++
  o.src = lin
  w.setupCommonTitle(&o)
  w.setup1TaxcdTitle(&o)
  w.setup3TaxcdTitle(&o)
  if outpt == ONE  || outpt == FULL {
    w.printCommonTitleExcel(f1, o)
    w.print1TaxcdTitleExcel(f1, o)
  }
  if outpt == MANY || outpt == FULL {
    w.printCommonTitleExcel(f2, o)
    w.print3TaxcdTitleExcel(f2, o)
  }
}

func (w *Writer_tp) PrintLineExcel(o lout_tp) {
  recn++
  if outpt == ONE  || outpt == FULL {
    w.printCommonLineExcel(f1, o)
    w.print1TaxcdLineExcel(f1, o)
  }
  if outpt == MANY || outpt == FULL {
    w.printCommonLineExcel(f2, o)
    w.print3TaxcdLineExcel(f2, o)
  }
}

func (w *Writer_tp) setupCommonTitle(o *ltit_tp) {
  o.retencionesIVA           = "Retenciones IVA"
  o.trasladosBaseIVA16       = "Traslados Base IVA16"
  o.trasladosImpuestoIVA16   = "Traslados Impuesto IVA16"
  o.trasladosBaseIVA8        = "Traslados Base IVA8"
  o.trasladosImpuestoIVA8    = "Traslados Impuesto IVA8"
  o.trasladosBaseIVA0        = "Traslados Base IVA0"
  o.trasladosImpuestoIVA0    = "Traslados Impuesto IVA0"
  o.montoTotalPagos          = "Monto Total Pagos"
  o.objetoImpuesto           = "Objeto Impuesto"
}

func (w *Writer_tp) setup1TaxcdTitle(o *ltit_tp) {
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
  o.difImportePago1          = "Diff Importe Pago1"
}

func (w *Writer_tp) setup3TaxcdTitle(o *ltit_tp) {
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
  o.difImportePago3          = "Diff Importe Pago3"
}

func (w *Writer_tp) printCommonTitleExcel(f *excelize.File, o ltit_tp) {
  f.SetCellValue(TAB, fmt.Sprintf("A%d",  recn), o.src.companyCode)
  f.SetCellValue(TAB, fmt.Sprintf("B%d",  recn), o.src.customer)
  f.SetCellValue(TAB, fmt.Sprintf("C%d",  recn), o.src.documentNumber)
  f.SetCellValue(TAB, fmt.Sprintf("D%d",  recn), o.src.DocumentType)
  f.SetCellValue(TAB, fmt.Sprintf("F%d",  recn), o.src.paymentDateTime)
  f.SetCellValue(TAB, fmt.Sprintf("F%d",  recn), o.src.ClearingDocument)
  f.SetCellValue(TAB, fmt.Sprintf("G%d",  recn), o.src.AmountDocCurr)
  f.SetCellValue(TAB, fmt.Sprintf("H%d",  recn), o.src.DocumentCurrency)
  f.SetCellValue(TAB, fmt.Sprintf("I%d",  recn), o.src.EffExchangeRate)
  f.SetCellValue(TAB, fmt.Sprintf("J%d",  recn), o.src.assignment)
  f.SetCellValue(TAB, fmt.Sprintf("K%d",  recn), o.src.formaPago)
  f.SetCellValue(TAB, fmt.Sprintf("L%d",  recn), o.src.noParcialidad)
  f.SetCellValue(TAB, fmt.Sprintf("M%d",  recn), o.src.importeSaldoAnterior)
  f.SetCellValue(TAB, fmt.Sprintf("N%d",  recn), o.src.ImportePago)
  f.SetCellValue(TAB, fmt.Sprintf("O%d",  recn), o.src.importeSaldoInsoluto)
  f.SetCellValue(TAB, fmt.Sprintf("P%d",  recn), o.src.tipoRelacion)
  f.SetCellValue(TAB, fmt.Sprintf("Q%d",  recn), o.src.pagoCanceladoDocNumber)
  f.SetCellValue(TAB, fmt.Sprintf("R%d",  recn), o.src.numOperacion)
  f.SetCellValue(TAB, fmt.Sprintf("S%d",  recn), o.src.rfcBancoOrdenente)
  f.SetCellValue(TAB, fmt.Sprintf("T%d",  recn), o.src.nombreBancoOrdenante)
  f.SetCellValue(TAB, fmt.Sprintf("U%d",  recn), o.src.cuentaOrdenante)
  f.SetCellValue(TAB, fmt.Sprintf("V%d",  recn), o.src.rfcBancoBeneficiario)
  f.SetCellValue(TAB, fmt.Sprintf("W%d",  recn), o.src.cuentaBeneficiario)
  f.SetCellValue(TAB, fmt.Sprintf("X%d",  recn), o.src.tipoCadenaPago)
  f.SetCellValue(TAB, fmt.Sprintf("Y%d",  recn), o.src.certificadoPago)
  f.SetCellValue(TAB, fmt.Sprintf("Z%d",  recn), o.src.cadenaPago)
  f.SetCellValue(TAB, fmt.Sprintf("AA%d", recn), o.src.selloPago)
  f.SetCellValue(TAB, fmt.Sprintf("AB%d", recn), o.src.TaxCode)
  f.SetCellValue(TAB, fmt.Sprintf("AC%d", recn), o.retencionesIVA)
  f.SetCellValue(TAB, fmt.Sprintf("AD%d", recn), o.trasladosBaseIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AE%d", recn), o.trasladosImpuestoIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AF%d", recn), o.trasladosBaseIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("AG%d", recn), o.trasladosImpuestoIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("AH%d", recn), o.trasladosBaseIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("AI%d", recn), o.trasladosImpuestoIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("AJ%d", recn), o.montoTotalPagos)
  f.SetCellValue(TAB, fmt.Sprintf("AK%d", recn), o.objetoImpuesto)
}

func (w *Writer_tp) print1TaxcdTitleExcel(f *excelize.File, o ltit_tp) {
  f.SetCellValue(TAB, fmt.Sprintf("AL%d", recn), o.taxTrasladoBase)
  f.SetCellValue(TAB, fmt.Sprintf("AM%d", recn), o.taxTrasladoImpuesto)
  f.SetCellValue(TAB, fmt.Sprintf("AN%d", recn), o.taxTrasladoTipoFactor)
  f.SetCellValue(TAB, fmt.Sprintf("AO%d", recn), o.taxTrasladoTasaOCuota)
  f.SetCellValue(TAB, fmt.Sprintf("AP%d", recn), o.taxTrasladoImporte)
  f.SetCellValue(TAB, fmt.Sprintf("AQ%d", recn), o.taxRetncionBase)
  f.SetCellValue(TAB, fmt.Sprintf("AR%d", recn), o.taxRetncionImpuesto)
  f.SetCellValue(TAB, fmt.Sprintf("AS%d", recn), o.taxRetncionTipoFactor)
  f.SetCellValue(TAB, fmt.Sprintf("AT%d", recn), o.taxRetncionTasaOCuota)
  f.SetCellValue(TAB, fmt.Sprintf("AU%d", recn), o.taxRetncionImporte)
  f.SetCellValue(TAB, fmt.Sprintf("AV%d", recn), o.difMontoTotalPagos)
  f.SetCellValue(TAB, fmt.Sprintf("AW%d", recn), o.difImportePago1)
}

func (w *Writer_tp) print3TaxcdTitleExcel(f *excelize.File, o ltit_tp) {
  f.SetCellValue(TAB, fmt.Sprintf("AL%d", recn), o.trasladoBaseDR)
  f.SetCellValue(TAB, fmt.Sprintf("AM%d", recn), o.trasladoImpuestoDR)
  f.SetCellValue(TAB, fmt.Sprintf("AN%d", recn), o.trasladoTipoFactorDR)
  f.SetCellValue(TAB, fmt.Sprintf("AO%d", recn), o.trasladoTasaOCuotaDR)
  f.SetCellValue(TAB, fmt.Sprintf("AP%d", recn), o.trasladoImporteDR)
  f.SetCellValue(TAB, fmt.Sprintf("AQ%d", recn), o.retncionBaseDR)
  f.SetCellValue(TAB, fmt.Sprintf("AR%d", recn), o.retncionImpuestoDR)
  f.SetCellValue(TAB, fmt.Sprintf("AS%d", recn), o.retncionTipoFactorDR)
  f.SetCellValue(TAB, fmt.Sprintf("AT%d", recn), o.retncionTasaOCuotaDR)
  f.SetCellValue(TAB, fmt.Sprintf("AU%d", recn), o.retncionImporteDR)
  f.SetCellValue(TAB, fmt.Sprintf("AV%d", recn), o.trasladoBasePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AW%d", recn), o.trasladoImpuestoPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AX%d", recn), o.trasladoTipoFactorPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AY%d", recn), o.trasladoTasaOCuotaPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AZ%d", recn), o.trasladoImportePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BA%d", recn), o.retncionBasePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BB%d", recn), o.retncionImpuestoPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BC%d", recn), o.retncionTipoFactorPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BD%d", recn), o.retncionTasaOCuotaPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BE%d", recn), o.retncionImportePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BF%d", recn), o.trasladoBasePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BG%d", recn), o.trasladoImpuestoPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BH%d", recn), o.trasladoTipoFactorPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BI%d", recn), o.trasladoTasaOCuotaPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BJ%d", recn), o.trasladoImportePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BK%d", recn), o.retncionBasePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BL%d", recn), o.retncionImpuestoPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BM%d", recn), o.retncionTipoFactorPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BN%d", recn), o.retncionTasaOCuotaPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BO%d", recn), o.retncionImportePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BP%d", recn), o.trasladoBasePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BQ%d", recn), o.trasladoImpuestoPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BR%d", recn), o.trasladoTipoFactorPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BS%d", recn), o.trasladoTasaOCuotaPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BT%d", recn), o.trasladoImportePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BU%d", recn), o.retncionBasePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BV%d", recn), o.retncionImpuestoPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BW%d", recn), o.retncionTipoFactorPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BX%d", recn), o.retncionTasaOCuotaPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BY%d", recn), o.retncionImportePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BZ%d", recn), o.difMontoTotalPagos)
  f.SetCellValue(TAB, fmt.Sprintf("CA%d", recn), o.difImportePago3)
}

func (w *Writer_tp) printCommonLineExcel(f *excelize.File, o lout_tp) {
  f.SetCellValue(TAB, fmt.Sprintf("A%d",  recn), o.src.companyCode)
  f.SetCellValue(TAB, fmt.Sprintf("B%d",  recn), o.src.customer)
  f.SetCellValue(TAB, fmt.Sprintf("C%d",  recn), o.src.documentNumber)
  f.SetCellValue(TAB, fmt.Sprintf("D%d",  recn), o.src.DocumentType)
  f.SetCellValue(TAB, fmt.Sprintf("F%d",  recn), o.src.paymentDateTime)
  f.SetCellValue(TAB, fmt.Sprintf("F%d",  recn), o.src.ClearingDocument)
  f.SetCellValue(TAB, fmt.Sprintf("G%d",  recn), o.src.AmountDocCurr)
  f.SetCellValue(TAB, fmt.Sprintf("H%d",  recn), o.src.DocumentCurrency)
  f.SetCellValue(TAB, fmt.Sprintf("I%d",  recn), o.src.EffExchangeRate)
  f.SetCellValue(TAB, fmt.Sprintf("J%d",  recn), o.src.assignment)
  f.SetCellValue(TAB, fmt.Sprintf("K%d",  recn), o.src.formaPago)
  f.SetCellValue(TAB, fmt.Sprintf("L%d",  recn), o.src.noParcialidad)
  f.SetCellValue(TAB, fmt.Sprintf("M%d",  recn), o.src.importeSaldoAnterior)
  f.SetCellValue(TAB, fmt.Sprintf("N%d",  recn), o.src.ImportePago)
  f.SetCellValue(TAB, fmt.Sprintf("O%d",  recn), o.src.importeSaldoInsoluto)
  f.SetCellValue(TAB, fmt.Sprintf("P%d",  recn), o.src.tipoRelacion)
  f.SetCellValue(TAB, fmt.Sprintf("Q%d",  recn), o.src.pagoCanceladoDocNumber)
  f.SetCellValue(TAB, fmt.Sprintf("R%d",  recn), o.src.numOperacion)
  f.SetCellValue(TAB, fmt.Sprintf("S%d",  recn), o.src.rfcBancoOrdenente)
  f.SetCellValue(TAB, fmt.Sprintf("T%d",  recn), o.src.nombreBancoOrdenante)
  f.SetCellValue(TAB, fmt.Sprintf("U%d",  recn), o.src.cuentaOrdenante)
  f.SetCellValue(TAB, fmt.Sprintf("V%d",  recn), o.src.rfcBancoBeneficiario)
  f.SetCellValue(TAB, fmt.Sprintf("W%d",  recn), o.src.cuentaBeneficiario)
  f.SetCellValue(TAB, fmt.Sprintf("X%d",  recn), o.src.tipoCadenaPago)
  f.SetCellValue(TAB, fmt.Sprintf("Y%d",  recn), o.src.certificadoPago)
  f.SetCellValue(TAB, fmt.Sprintf("Z%d",  recn), o.src.cadenaPago)
  f.SetCellValue(TAB, fmt.Sprintf("AA%d", recn), o.src.selloPago)
  f.SetCellValue(TAB, fmt.Sprintf("AB%d", recn), o.src.TaxCode)
  f.SetCellValue(TAB, fmt.Sprintf("AC%d", recn), o.retencionesIVA)
  f.SetCellValue(TAB, fmt.Sprintf("AD%d", recn), o.trasladosBaseIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AE%d", recn), o.trasladosImpuestoIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AF%d", recn), o.trasladosBaseIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("AG%d", recn), o.trasladosImpuestoIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("AH%d", recn), o.trasladosBaseIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("AI%d", recn), o.trasladosImpuestoIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("AJ%d", recn), o.montoTotalPagos)
  f.SetCellValue(TAB, fmt.Sprintf("AK%d", recn), o.objetoImpuesto)
}

func (w *Writer_tp) print1TaxcdLineExcel(f *excelize.File, o lout_tp) {
  f.SetCellValue(TAB, fmt.Sprintf("AL%d", recn), o.taxTrasladoBase)
  f.SetCellValue(TAB, fmt.Sprintf("AM%d", recn), o.taxTrasladoImpuesto)
  f.SetCellValue(TAB, fmt.Sprintf("AN%d", recn), o.taxTrasladoTipoFactor)
  f.SetCellValue(TAB, fmt.Sprintf("AO%d", recn), o.taxTrasladoTasaOCuota)
  f.SetCellValue(TAB, fmt.Sprintf("AP%d", recn), o.taxTrasladoImporte)
  f.SetCellValue(TAB, fmt.Sprintf("AQ%d", recn), o.taxRetncionBase)
  f.SetCellValue(TAB, fmt.Sprintf("AR%d", recn), o.taxRetncionImpuesto)
  f.SetCellValue(TAB, fmt.Sprintf("AS%d", recn), o.taxRetncionTipoFactor)
  f.SetCellValue(TAB, fmt.Sprintf("AT%d", recn), o.taxRetncionTasaOCuota)
  f.SetCellValue(TAB, fmt.Sprintf("AU%d", recn), o.taxRetncionImporte)
  f.SetCellValue(TAB, fmt.Sprintf("AV%d", recn), o.difMontoTotalPagos)
  f.SetCellValue(TAB, fmt.Sprintf("AW%d", recn), o.difImportePago1)
}

func (w *Writer_tp) print3TaxcdLineExcel(f *excelize.File, o lout_tp) {
  f.SetCellValue(TAB, fmt.Sprintf("AL%d", recn), o.trasladoBaseDR)
  f.SetCellValue(TAB, fmt.Sprintf("AM%d", recn), o.trasladoImpuestoDR)
  f.SetCellValue(TAB, fmt.Sprintf("AN%d", recn), o.trasladoTipoFactorDR)
  f.SetCellValue(TAB, fmt.Sprintf("AO%d", recn), o.trasladoTasaOCuotaDR)
  f.SetCellValue(TAB, fmt.Sprintf("AP%d", recn), o.trasladoImporteDR)
  f.SetCellValue(TAB, fmt.Sprintf("AQ%d", recn), o.retncionBaseDR)
  f.SetCellValue(TAB, fmt.Sprintf("AR%d", recn), o.retncionImpuestoDR)
  f.SetCellValue(TAB, fmt.Sprintf("AS%d", recn), o.retncionTipoFactorDR)
  f.SetCellValue(TAB, fmt.Sprintf("AT%d", recn), o.retncionTasaOCuotaDR)
  f.SetCellValue(TAB, fmt.Sprintf("AU%d", recn), o.retncionImporteDR)
  f.SetCellValue(TAB, fmt.Sprintf("AV%d", recn), o.trasladoBasePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AW%d", recn), o.trasladoImpuestoPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AX%d", recn), o.trasladoTipoFactorPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AY%d", recn), o.trasladoTasaOCuotaPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("AZ%d", recn), o.trasladoImportePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BA%d", recn), o.retncionBasePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BB%d", recn), o.retncionImpuestoPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BC%d", recn), o.retncionTipoFactorPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BD%d", recn), o.retncionTasaOCuotaPIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BE%d", recn), o.retncionImportePIVA16)
  f.SetCellValue(TAB, fmt.Sprintf("BF%d", recn), o.trasladoBasePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BG%d", recn), o.trasladoImpuestoPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BH%d", recn), o.trasladoTipoFactorPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BI%d", recn), o.trasladoTasaOCuotaPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BJ%d", recn), o.trasladoImportePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BK%d", recn), o.retncionBasePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BL%d", recn), o.retncionImpuestoPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BM%d", recn), o.retncionTipoFactorPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BN%d", recn), o.retncionTasaOCuotaPIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BO%d", recn), o.retncionImportePIVA8)
  f.SetCellValue(TAB, fmt.Sprintf("BP%d", recn), o.trasladoBasePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BQ%d", recn), o.trasladoImpuestoPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BR%d", recn), o.trasladoTipoFactorPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BS%d", recn), o.trasladoTasaOCuotaPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BT%d", recn), o.trasladoImportePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BU%d", recn), o.retncionBasePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BV%d", recn), o.retncionImpuestoPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BW%d", recn), o.retncionTipoFactorPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BX%d", recn), o.retncionTasaOCuotaPIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BY%d", recn), o.retncionImportePIVA0)
  f.SetCellValue(TAB, fmt.Sprintf("BZ%d", recn), o.difMontoTotalPagos)
  f.SetCellValue(TAB, fmt.Sprintf("CA%d", recn), o.difImportePago3)
}

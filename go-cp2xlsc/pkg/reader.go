// reader.go [2022-04-06 BAR8TL] Excel file EDICOM reader functions
package cp2xlsc

import ut "bar8tl/p/rblib"
import "github.com/xuri/excelize/v2"
import "log"
import "strconv"

var F *excelize.File

type Line_tp struct {
  companyCode            string
  customer               string
  documentNumber         string
  DocumentType           string // Document type [DZ|RV]
  paymentDateTime        string
  ClearingDocument       string // Clearing document (payment reference number)
  AmountDocCurr          string // Payment total amount
  DocumentCurrency       string
  EffExchangeRate        string // Exchange rate
  assignment             string
  formaPago              string
  noParcialidad          string
  importeSaldoAnterior   string
  ImportePago            string // Payment amount corresponding to an invoice
  importeSaldoInsoluto   string
  tipoRelacion           string
  pagoCanceladoDocNumber string
  numOperacion           string
  rfcBancoOrdenente      string
  nombreBancoOrdenante   string
  cuentaOrdenante        string
  rfcBancoBeneficiario   string
  cuentaBeneficiario     string
  tipoCadenaPago         string
  certificadoPago        string
  cadenaPago             string
  selloPago              string
  TaxCode                string
}

type Reader_tp struct {
  Src                    Line_tp
  AmountDocCurr          float64
  EffExchangeRate        float64
  ImportePago            float64
}

func NewReader() *Reader_tp {
  var r Reader_tp
  return &r
}

func (r *Reader_tp) OpenInpExcel(dir, fname string) {
  var err error
  F, err = excelize.OpenFile(dir+fname)
  if err != nil {
    log.Fatal(err)
  }
}

func (r *Reader_tp) GetLineFields(row []string) {
  for i, _ := range row {
    switch i {
      case 0  : r.Src.companyCode            = row[i]
      case 1  : r.Src.customer               = row[i]
      case 2  : r.Src.documentNumber         = row[i]
      case 3  : r.Src.DocumentType           = row[i]
      case 4  : r.Src.paymentDateTime        = row[i]
      case 5  : r.Src.ClearingDocument       = row[i]
      case 6  : r.Src.AmountDocCurr          = row[i];
        r.AmountDocCurr, _ = strconv.ParseFloat(r.Src.AmountDocCurr, 64)
        r.AmountDocCurr = ut.Round(r.AmountDocCurr, 6)
      case 7  : r.Src.DocumentCurrency       = row[i]
      case 8  : r.Src.EffExchangeRate        = row[i];
        r.EffExchangeRate, _ = strconv.ParseFloat(r.Src.EffExchangeRate, 64)
        r.EffExchangeRate = ut.Round(r.EffExchangeRate, 7)
      case 9  : r.Src.assignment             = row[i]
      case 10 : r.Src.formaPago              = row[i]
      case 11 : r.Src.noParcialidad          = row[i]
      case 12 : r.Src.importeSaldoAnterior   = row[i]
      case 13 : r.Src.ImportePago            = row[i];
        r.ImportePago, _ = strconv.ParseFloat(r.Src.ImportePago, 64)
        r.ImportePago = ut.Round(r.ImportePago, 6)
      case 14 : r.Src.importeSaldoInsoluto   = row[i]
      case 15 : r.Src.tipoRelacion           = row[i]
      case 16 : r.Src.pagoCanceladoDocNumber = row[i]
      case 17 : r.Src.numOperacion           = row[i]
      case 18 : r.Src.rfcBancoOrdenente      = row[i]
      case 19 : r.Src.nombreBancoOrdenante   = row[i]
      case 20 : r.Src.cuentaOrdenante        = row[i]
      case 21 : r.Src.rfcBancoBeneficiario   = row[i]
      case 22 : r.Src.cuentaBeneficiario     = row[i]
      case 23 : r.Src.tipoCadenaPago         = row[i]
      case 24 : r.Src.certificadoPago        = row[i]
      case 25 : r.Src.cadenaPago             = row[i]
      case 26 : r.Src.selloPago              = row[i]
      case 27 : r.Src.TaxCode                = row[i]
    }
  }
}

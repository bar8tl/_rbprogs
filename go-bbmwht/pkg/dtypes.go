// dtypes.go [2022-04-06 BAR8TL] Data Types to Extend Pagos1.0 EDICOM-file with
// Pagos2.0 fields. Auxiliary tools
package bbmwht

// Data types and definitions --------------------------------------------------
const IMPUESTO       = "002"
const TIPOFACTOR     = "Tasa"
const OBJETOIMPUESTO = "02"
const TABNAME        = "Sheet1"

type Line_tp struct {
  companyCode              string
  customer                 string
  documentNumber           string
  DocumentType             string
  paymentDateTime          string
  clearingDocument         string
  AmountDocCurr            string
  documentCurrency         string
  effExchangeRate          string
  assignment               string
  formaPago                string
  noParcialidad            string
  importeSaldoAnterior     string
  importePago              string
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


type Fbl5n_tp struct {        // Cleared Items
  CompanyCode          string // Company Code
  Customer             string // Sold-to number
  DocumentNumber       string // Document number (Payment / Invoice)
  DocumentType         string // Document type (Payment / Invoice)
  PaymentDate          string // Payment date
  ClearingDocument     string // Clearing document number
  AmountDocCurr        string // Importe en la moneda del documento de referencia
  DocumentCurrency     string // Tipo de moneda del documento de referencia
  EffExchangeRate      string // Fecha de efectividad del tipo de cambio
  Assignment           string // Asignacion
  FormaPago            string // Payment form
  NoParcialidad        string // Partiallity number
  ImporteSaldoAnterior string // Prior balance amount
  ImportePago          string // Payment amount
  ImporteSaldoInsoluto string //   
  DiferenciaPago       string // Difference in payment
  Comentario           string // Comment
  Status               string // Status
}

type Fb03_tp struct {         // Documents DZ

}
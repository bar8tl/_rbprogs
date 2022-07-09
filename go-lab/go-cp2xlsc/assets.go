// assets.go - Define valuable global data types and functions
// 2022-05-17 BAR8TL Version1.0 - In progress
package main

import rb "bar8tl/p/cp2xlsc"
import ut "bar8tl/p/rblib"

var IMPUESTO   string
var TIPOFACTOR string
var OBJETOIMP  string
var DEC        int
var m      map[string]int
var n      map[string]int
var drs    [91]string
var drf    [91]float64
var pms    [91]string
var pmf    [91]float64
var tt     [91]string
var cc     [37]string
var c1     [49]string
var c3     [91]string
var outpt      string
var recn       int
var numers   []int
var alphac   []int
var alpha1   []int
var alpha3   []int

// Indexes maps
func loadIndexMaps() {
  // Source fields
  m["companyCode"]              = 0
  m["customer"]                 = 1
  m["documentNumber"]           = 2
  m["documentType"]             = 3
  m["paymentDateTime"]          = 4
  m["clearingDocument"]         = 5
  m["amountDocCurr"]            = 6
  m["documentCurrency"]         = 7
  m["effExchangeRate"]          = 8
  m["assignment"]               = 9
  m["formaPago"]                = 10
  m["noParcialidad"]            = 11
  m["importeSaldoAnterior"]     = 12
  m["importePago"]              = 13
  m["importeSaldoInsoluto"]     = 14
  m["tipoRelacion"]             = 15
  m["pagoCanceladoDocNumber"]   = 16
  m["numOperacion"]             = 17
  m["rfcBancoOrdenente"]        = 18
  m["nombreBancoOrdenante"]     = 19
  m["cuentaOrdenante"]          = 20
  m["rfcBancoBeneficiario"]     = 21
  m["cuentaBeneficiario"]       = 22
  m["tipoCadenaPago"]           = 23
  m["certificadoPago"]          = 24
  m["cadenaPago"]               = 25
  m["selloPago"]                = 26
  m["taxCode"]                  = 27
  // Common new fields
  m["retencionesIVA"]           = 28
  m["trasladosBaseIVA16"]       = 29
  m["trasladosImpuestoIVA16"]   = 30
  m["trasladosBaseIVA8"]        = 31
  m["trasladosImpuestoIVA8"]    = 32
  m["trasladosBaseIVA0"]        = 33
  m["trasladosImpuestoIVA0"]    = 34
  m["montoTotalPagos"]          = 35
  m["objetoImpuesto"]           = 36
  // One tax code fields
  m["taxTrasladoBase"]          = 37
  m["taxTrasladoImpuesto"]      = 38
  m["taxTrasladoTipoFactor"]    = 39
  m["taxTrasladoTasaOCuota"]    = 40
  m["taxTrasladoImporte"]       = 41
  m["taxRetncionBase"]          = 42
  m["taxRetncionImpuesto"]      = 43
  m["taxRetncionTipoFactor"]    = 44
  m["taxRetncionTasaOCuota"]    = 45
  m["taxRetncionImporte"]       = 46
  m["difMontoTotalPagos1"]      = 47
  m["difImportePago1"]          = 48
  // Multiple tax code fields
  m["trasladoBaseDR"]           = 49
  m["trasladoImpuestoDR"]       = 50
  m["trasladoTipoFactorDR"]     = 51
  m["trasladoTasaOCuotaDR"]     = 52
  m["trasladoImporteDR"]        = 53
  m["retncionBaseDR"]           = 54
  m["retncionImpuestoDR"]       = 55
  m["retncionTipoFactorDR"]     = 56
  m["retncionTasaOCuotaDR"]     = 57
  m["retncionImporteDR"]        = 58
  m["trasladoBasePIVA16"]       = 59
  m["trasladoImpuestoPIVA16"]   = 60
  m["trasladoTipoFactorPIVA16"] = 61
  m["trasladoTasaOCuotaPIVA16"] = 62
  m["trasladoImportePIVA16"]    = 63
  m["retncionBasePIVA16"]       = 64
  m["retncionImpuestoPIVA16"]   = 65
  m["retncionTipoFactorPIVA16"] = 66
  m["retncionTasaOCuotaPIVA16"] = 67
  m["retncionImportePIVA16"]    = 68
  m["trasladoBasePIVA8"]        = 69
  m["trasladoImpuestoPIVA8"]    = 70
  m["trasladoTipoFactorPIVA8"]  = 71
  m["trasladoTasaOCuotaPIVA8"]  = 72
  m["trasladoImportePIVA8"]     = 73
  m["retncionBasePIVA8"]        = 74
  m["retncionImpuestoPIVA8"]    = 75
  m["retncionTipoFactorPIVA8"]  = 76
  m["retncionTasaOCuotaPIVA8"]  = 77
  m["retncionImportePIVA8"]     = 78
  m["trasladoBasePIVA0"]        = 79
  m["trasladoImpuestoPIVA0"]    = 80
  m["trasladoTipoFactorPIVA0"]  = 81
  m["trasladoTasaOCuotaPIVA0"]  = 82
  m["trasladoImportePIVA0"]     = 83
  m["retncionBasePIVA0"]        = 84
  m["retncionImpuestoPIVA0"]    = 85
  m["retncionTipoFactorPIVA0"]  = 86
  m["retncionTasaOCuotaPIVA0"]  = 87
  m["retncionImportePIVA0"]     = 88
  m["difMontoTotalPagos3"]      = 89
  m["difImportePago3"]          = 90
  // Fisrt invoice indicators
  n["firstInvoice"]             = 0
  n["firstInvoTraslIVA16"]      = 1
  n["firstInvoRetenIVA16"]      = 2
  n["firstInvoTraslIVA8"]       = 3
  n["firstInvoRetenIVA8"]       = 4
  n["firstInvoTraslIVA0"]       = 5
  n["firstInvoRetenIVA0"]       = 6
}

// Title table
func loadTitleTable() {
  // Source fields titles
  tt[0]  = "Company Code"
  tt[1]  = "Customer"
  tt[2]  = "Document Number"
  tt[3]  = "Document Type"
  tt[4]  = "Payment Date - Time"
  tt[5]  = "Clearing Document"
  tt[6]  = "Amount in Doc. Curr"
  tt[7]  = "Document Currency"
  tt[8]  = "Eff.exchange rate"
  tt[9]  = "Assignment"
  tt[10] = "Forma de Pago"
  tt[11] = "No. de Parcialidad"
  tt[12] = "Importe Saldo Anterior"
  tt[13] = "Importe Pago"
  tt[14] = "Importe Saldo Insoluto"
  tt[15] = "Tipo Relacion (04)"
  tt[16] = "Pago Cancelado (Doc Number)"
  tt[17] = "Num Operacion"
  tt[18] = "RFC Banco Ordenente"
  tt[19] = "Nombre Banco Ordenante"
  tt[20] = "Cuenta Ordenante"
  tt[21] = "RFC Banco Beneficiario"
  tt[22] = "Cuenta Beneficiario"
  tt[23] = "Tipo Cadena Pago (01)"
  tt[24] = "Certificado Pago"
  tt[25] = "Cadena Pago"
  tt[26] = "Sello Pago"
  tt[27] = "Tax Code"
  // Common new fields titles
  tt[28] = "Retenciones IVA"
  tt[29] = "Traslados Base IVA16"
  tt[30] = "Traslados Impuesto IVA16"
  tt[31] = "Traslados Base IVA8"
  tt[32] = "Traslados Impuesto IVA8"
  tt[33] = "Traslados Base IVA0"
  tt[34] = "Traslados Impuesto IVA0"
  tt[35] = "Monto Total Pagos"
  tt[36] = "Objeto Impuesto"
  // One tax code fields titles
  tt[37] = "Tax Traslado Base"
  tt[38] = "Tax Traslado Impuesto"
  tt[39] = "Tax Traslado TipoFactor"
  tt[40] = "Tax Traslado TasaOCuota"
  tt[41] = "Tax Traslado Importe"
  tt[42] = "Tax Retencion Base"
  tt[43] = "Tax Retencion Impuesto"
  tt[44] = "Tax Retencion TipoFactor"
  tt[45] = "Tax Retencion TasaOCuota"
  tt[46] = "Tax Retencion Importe"
  tt[47] = "Diff Monto Total Pagos"
  tt[48] = "Diff Importe Pago1"
  // Multiple tax code fields titles
  tt[49] = "DR Traslado Base"
  tt[50] = "DR Traslado Impuesto"
  tt[51] = "DR Traslado TipoFactor"
  tt[52] = "DR Traslado TasaOCuota"
  tt[53] = "DR Traslado Importe"
  tt[54] = "DR Retencion Base"
  tt[55] = "DR Retencion Impuesto"
  tt[56] = "DR Retencion TipoFactor"
  tt[57] = "DR Retencion TasaOCuota"
  tt[58] = "DR Retencion Importe"
  tt[59] = "P Traslado Base IVA16"
  tt[60] = "P Traslado Impuesto IVA16"
  tt[61] = "P Traslado TipoFactor IVA16"
  tt[62] = "P Traslado TasaOCuota IVA16"
  tt[63] = "P Traslado Importe IVA16"
  tt[64] = "P Retencion Base IVA16"
  tt[65] = "P Retencion Impuesto IVA16"
  tt[66] = "P Retencion TipoFactor IVA16"
  tt[67] = "P Retencion TasaOCuota IVA16"
  tt[68] = "P Retencion Importe IVA16"
  tt[69] = "P Traslado Base IVA8"
  tt[70] = "P Traslado Impuesto IVA8"
  tt[71] = "P Traslado TipoFactor IVA8"
  tt[72] = "P Traslado TasaOCuota IVA8"
  tt[73] = "P Traslado Importe IVA8"
  tt[74] = "P Retencion Base IVA8"
  tt[75] = "P Retencion Impuesto IVA8"
  tt[76] = "P Retencion TipoFactor IVA8"
  tt[77] = "P Retencion TasaOCuota IVA8"
  tt[78] = "P Retencion Importe IVA8"
  tt[79] = "P Traslado Base IVA0"
  tt[80] = "P Traslado Impuesto IVA0"
  tt[81] = "P Traslado TipoFactor IVA0"
  tt[82] = "P Traslado TasaOCuota IVA0"
  tt[83] = "P Traslado Importe IVA0"
  tt[84] = "P Retencion Base IVA0"
  tt[85] = "P Retencion Impuesto IVA0"
  tt[86] = "P Retencion TipoFactor IVA0"
  tt[87] = "P Retencion TasaOCuota IVA0"
  tt[88] = "P Retencion Importe IVA0"
  tt[89] = "Diff Monto Total Pagos"
  tt[90] = "Diff Importe Pago3"
}

// Column names table
func loadColumnTable() {
  // Source fields columns
  cc[0]  = "A"
  cc[1]  = "B"
  cc[2]  = "C"
  cc[3]  = "D"
  cc[4]  = "E"
  cc[5]  = "F"
  cc[6]  = "G"
  cc[7]  = "H"
  cc[8]  = "I"
  cc[9]  = "J"
  cc[10] = "K"
  cc[11] = "L"
  cc[12] = "M"
  cc[13] = "N"
  cc[14] = "O"
  cc[15] = "P"
  cc[16] = "Q"
  cc[17] = "R"
  cc[18] = "S"
  cc[19] = "T"
  cc[20] = "U"
  cc[21] = "V"
  cc[22] = "W"
  cc[23] = "X"
  cc[24] = "Y"
  cc[25] = "Z"
  cc[26] = "AA"
  cc[27] = "AB"
  // Common fields columns
  cc[28] = "AC"
  cc[29] = "AD"
  cc[30] = "AE"
  cc[31] = "AF"
  cc[32] = "AG"
  cc[33] = "AH"
  cc[34] = "AI"
  cc[35] = "AJ"
  cc[36] = "AK"
  // One tax code fields columns
  c1[37] = "AL"
  c1[38] = "AM"
  c1[39] = "AN"
  c1[40] = "AO"
  c1[41] = "AP"
  c1[42] = "AQ"
  c1[43] = "AR"
  c1[44] = "AS"
  c1[45] = "AT"
  c1[46] = "AU"
  c1[47] = "AV"
  c1[48] = "AW"
  // Multiple tax code fields columns
  c3[49] = "AL"
  c3[50] = "AM"
  c3[51] = "AN"
  c3[52] = "AO"
  c3[53] = "AP"
  c3[54] = "AQ"
  c3[55] = "AR"
  c3[56] = "AS"
  c3[57] = "AT"
  c3[58] = "AU"
  c3[59] = "AV"
  c3[60] = "AW"
  c3[61] = "AX"
  c3[62] = "AY"
  c3[63] = "AZ"
  c3[64] = "BA"
  c3[65] = "BB"
  c3[66] = "BC"
  c3[67] = "BD"
  c3[68] = "BE"
  c3[69] = "BF"
  c3[70] = "BG"
  c3[71] = "BH"
  c3[72] = "BI"
  c3[73] = "BJ"
  c3[74] = "BK"
  c3[75] = "BL"
  c3[76] = "BM"
  c3[77] = "BN"
  c3[78] = "BO"
  c3[79] = "BP"
  c3[80] = "BQ"
  c3[81] = "BR"
  c3[82] = "BS"
  c3[83] = "BT"
  c3[84] = "BU"
  c3[85] = "BV"
  c3[86] = "BW"
  c3[87] = "BX"
  c3[88] = "BY"
  c3[89] = "BZ"
  c3[90] = "CA"
}

func loadAssets(s rb.Settings_tp) {
  IMPUESTO   = s.Konst.IMPUESTO
  TIPOFACTOR = s.Konst.TIPOFACTOR
  OBJETOIMP  = s.Const.Taxbl
  DEC        = s.Konst.DEC
  outpt      = s.Outpt
  m = make(map[string]int)
  n = make(map[string]int)
  loadIndexMaps()
  loadTitleTable()
  loadColumnTable()
  numers = []int{
    m["amountDocCurr"],
    m["importePago"],
    m["importeSaldoAnterior"],
    m["importeSaldoInsoluto"],
  }
  alphac = []int{
    m["objetoImpuesto"],
  }
  alpha1 = []int{
    m["taxTrasladoImpuesto"],
    m["taxTrasladoTipoFactor"],
    m["taxRetncionImpuesto"],
    m["taxRetncionTipoFactor"],
  }
  alpha3 = []int{
    m["trasladoImpuestoDR"],
    m["trasladoTipoFactorDR"],
    m["retncionImpuestoDR"],
    m["retncionTipoFactorDR"],
    m["trasladoImpuestoPIVA16"],
    m["trasladoTipoFactorPIVA16"],
    m["retncionImpuestoPIVA16"],
    m["retncionTipoFactorPIVA16"],
    m["trasladoImpuestoPIVA8"],
    m["trasladoTipoFactorPIVA8"],
    m["retncionImpuestoPIVA8"],
    m["retncionTipoFactorPIVA8"],
    m["trasladoImpuestoPIVA0"],
    m["trasladoTipoFactorPIVA0"],
    m["retncionImpuestoPIVA0"],
    m["retncionTipoFactorPIVA0"],
  }
}

func round(n float64) float64 {
  return ut.Round(n, DEC)
}

func contains(s []int, e int) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}

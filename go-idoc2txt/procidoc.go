// rtxtidoc.go [2017-05-24/BAR8TL]
// SAP-IDoc command processor
package main

import rb "bar8tl/p/rbidoc"

func main() {
  s := rb.NewSettings("_config.json", "_deflts.json")
  for _, parm := range s.Cmdpr {
           if parm.Optn == "cdb" { // Create reference IDoc-definition database
      cdb := rb.NewCrtdb()
      cdb.CrtTable(parm, s)
    } else if parm.Optn == "upl" { // Read and upload IDoc-definition files
      upl := rb.NewUplddefs()
      upl.UpldDefs(parm, s)
    } else if parm.Optn == "usa" { // Upload segment-definition alias names
      usa := rb.NewUpldsgma()
      usa.Upldsgma(parm, s)
    } else if parm.Optn == "unf" { // Convert IDOC-data parser-fmt SAP->Flat-TXT
      unf := rb.NewUnfold()
      unf.UfldData(parm, s)
    } else if parm.Optn == "dat" { // Convert IDOC-data Flat-TXT->Intern Struct
      dat := rb.NewDidoc()
      dat.ReadDire(parm, s)
    } else if parm.Optn == "qry" { //
      qry := rb.NewQuery()
      qry.UpldQuery(parm, s)
    }
  }
}

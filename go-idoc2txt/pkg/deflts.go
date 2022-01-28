// defaults.go [2017-05-24 BAR8TL]
// Reads defaults file and gets run settings
package rbidoc

import "encoding/json"
import "io/ioutil"
import "log"
import "os"

type Dflt_tp struct {
  CNNS_SQLIT3    string `json:"CNNS_SQLIT3"`
  CONTROL_CODE   string `json:"CONTROL_CODE"`
  CLIENT_CODE    string `json:"CLIENT_CODE"`
  DB_NAME        string `json:"DB_NAME"`
  DB_DIR         string `json:"DB_DIR"`
  INPUTS_DIR     string `json:"INPUTS_DIR"`
  OUTPUTS_DIR    string `json:"OUTPUTS_DIR"`
  INPUTS_FILTER  string `json:"INPUTS_FILTER"`
  INPUTS_NAMING  string `json:"INPUTS_NAMING"`
  OUTPUTS_NAMING string `json:"OUTPUTS_NAMING"`
}

type Konst_tp struct {
  BEGIN   string `json:"BEGIN"`
  END     string `json:"END"`
  CONTROL string `json:"CONTROL"`
  RECORD  string `json:"RECORD"`
  IDOC    string `json:"IDOC"`
  GROUP   string `json:"GROUP"`
  SEGMENT string `json:"SEGMENT"`
  FIELDS  string `json:"FIELDS"`
  ITM     string `json:"ITM"`
  GRP     string `json:"GRP"`
  SGM     string `json:"SGM"`
}

type Sqlcr_tp struct {
  Activ bool   `json:"activ"`
  Table string `json:"table"`
  Sqlst string `json:"sqlst"`
}

type Sqlst_tp struct {
  Sqlcr []Sqlcr_tp `json:"sqlcr"`
}

type Errs_tp struct {
}

type Deflts_tp struct {
  Dflt  Dflt_tp  `json:"dflt"`
  Konst Konst_tp `json:"konst"`
  Sqlst Sqlst_tp `json:"sqlst"`
  Errs  Errs_tp  `json:"errs"`
  Sqlcr []Sqlcr_tp
}

func (d *Deflts_tp) NewDeflts(fname string) {
  f, err := os.Open(fname)
  if err != nil {
    log.Fatalf("File %s open error: %s\n", fname, err)
  }
  defer f.Close()
  jsonv, _ := ioutil.ReadAll(f)
  err = json.Unmarshal(jsonv, &d)
  if err != nil {
    log.Fatalf("File %s read error: %s\n", fname, err)
  }
  d.getActivSqlcr()
}

func (d *Deflts_tp) getActivSqlcr() {
  for _, sqc := range d.Sqlst.Sqlcr {
    if sqc.Activ {
      d.Sqlcr = append(d.Sqlcr, Sqlcr_tp{sqc.Activ, sqc.Table, sqc.Sqlst})
    }
  }
}

// defaults.go - Upload program settings to be used as default values
// 2022-05-17 BAR8TL Version1.0 - Released
package cp2xlsc

import "encoding/json"
import "io/ioutil"
import "log"
import "os"

type Dflt_tp struct {
  INPUTS_DIR     string `json:"INPUTS_DIR"`
  OUTPUTS_DIR    string `json:"OUTPUTS_DIR"`
  FILES_DIR      string `json:"FILES_DIR"`
  INPUTS_FILTER  string `json:"INPUTS_FILTER"`
  INPUTS_NAMING  string `json:"INPUTS_NAMING"`
  OUTPUTS_NAMING string `json:"OUTPUTS_NAMING"`
}

type Konst_tp struct {
  IMPUESTO       string `json:"IMPUESTO"`
  TIPOFACTOR     string `json:"TIPOFACTOR"`
  OBJETOIMPUESTO string `json:"OBJETOIMPUESTO"`
  TAB            string `json:"TAB"`
  DEC            int    `json:"DEC"`
  ONE            string `json:"ONE"`
  MANY           string `json:"MANY"`
  FULL           string `json:"FULL"`
  BATCH          string `json:"BATCH"`
  INDIV          string `json:"INDIV"`
}

type Deflts_tp struct {
  Dflt  Dflt_tp  `json:"dflt"`
  Konst Konst_tp `json:"konst"`
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
}

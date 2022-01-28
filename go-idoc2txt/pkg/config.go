// config.go [2017-05-24 BAR8TL]
// Reads config file and gets run parameters
package rbidoc

import "encoding/json"
import "io/ioutil"
import "log"
import "os"

type Constant_tp struct {
  Cntrl string `json:"contrl"`
  Clien string `json:"client"`
}

type Program_tp struct {
  Dbonm string `json:"dboNam"`
  Dbodr string `json:"dboDir"`
  Inpdr string `json:"inpDir"`
  Outdr string `json:"outDir"`
  Ifilt string `json:"inFilt"`
  Ifnam string `json:"inName"`
  Ofnam string `json:"ouName"`
}

type Run_tp struct {
  Optcd string `json:"option"`
  Objnm string `json:"objNam"`
  Qrynm string `json:"qryNam"`
  Dbonm string `json:"dboNam"`
  Dbodr string `json:"dboDir"`
  Inpdr string `json:"inpDir"`
  Outdr string `json:"outDir"`
  Qrydr string `json:"qryDir"`
  Ifilt string `json:"inFilt"`
  Ifnam string `json:"inName"`
  Ofnam string `json:"ouName"`
  Rcvpf string `json:"rcPrnF"`
}

type Cdb_tp struct {
  Id    string `json:"id"`
  Table string `json:"table"`
  Cr    bool   `json:"cr"`
}

type Config_tp struct {
  Const Constant_tp `json:"constants"`
  Progm Program_tp  `json:"program"`
  Run   []Run_tp    `json:"run"`
  Cdb   []Cdb_tp    `json:"cdb"`
}

func (c *Config_tp) NewConfig(fname string) {
  f, err := os.Open(fname)
  if err != nil {
    log.Fatalf("File %s open error: %s\n", fname, err)
  }
  defer f.Close()
  jsonv, _ := ioutil.ReadAll(f)
  err = json.Unmarshal(jsonv, &c)
  if err != nil {
    log.Fatalf("File %s read error: %s\n", fname, err)
  }
}

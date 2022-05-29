// config.go - Upload program/run-option configuration settings
// 2022-05-17 BAR8TL Version1.0 - Released
package cp2xlsc

import "encoding/json"
import "io/ioutil"
import "log"
import "os"

type Constant_tp struct {
  Taxbl string `json:"OBJETOIMPUESTO"`
}

type Program_tp struct {
  Inpdr string `json:"inpDir"`
  Outdr string `json:"outDir"`
  Fildr string `json:"filDir"`
  Ifilt string `json:"inFilt"`
  Ifnam string `json:"inName"`
  Ofnam string `json:"ouName"`
}

type Run_tp struct {
  Optcd string `json:"option"`
  Objnm string `json:"objNam"`
  Modep string `json:"mode"`
  Outpt string `json:"output"`
  Inpdr string `json:"inpDir"`
  Outdr string `json:"outDir"`
  Fildr string `json:"filDir"`
  Infil string `json:"inpFil"`
  Ifilt string `json:"inFilt"`
  Ifnam string `json:"inName"`
  Ofnam string `json:"ouName"`
}

type Config_tp struct {
  Const   Constant_tp `json:"constants"`
  Progm   Program_tp  `json:"program"`
  Run   []Run_tp      `json:"run"`
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

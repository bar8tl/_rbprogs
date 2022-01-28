// upldsgma.go [2017-05-24 BAR8TL] 
// Upload _segma.json file for segment aliases
package rbidoc

import lib "bar8tl/p/rblib"
import "code.google.com/p/go-sqlite/go1/sqlite3"
import "encoding/json"
import "log"
import "io/ioutil"

type Sgma_tp struct {
  Type string
  Defn string
}
type Segma_tp struct {
  Idoc string
  Segm []Sgma_tp
}
type Segal_tp struct {
  Segma []Segma_tp
}

type Upldsgma_tp struct {
  Db    *sqlite3.Conn
  Cnnst string
  Segal Segal_tp
}

func NewUpldsgma() *Upldsgma_tp {
  var u Upldsgma_tp
  return &u
}

func (u *Upldsgma_tp) Upldsgma(parm lib.Param_tp, s Settings_tp) {
  s.SetRunVars(parm)
  u.Cnnst = s.Cnnst
  var err error
  u.Db, err = sqlite3.Open(u.Cnnst)
  if err != nil {
    log.Fatalf("Open SQLite database error: %v\n", err)
  }
  defer u.Db.Close()
  u.Db.Exec(`DELETE FROM segma;`)
  fa, _ := ioutil.ReadFile("idoctypes\\_segma.json") // <-- change for sttgs
  json.Unmarshal(fa, &u.Segal)
  for _, sa := range u.Segal.Segma {
    for _, ss := range sa.Segm {
      err := u.Db.Exec(`INSERT INTO segma VALUES(?,?,?)`, sa.Idoc, ss.Type,
        ss.Defn)
      if err != nil {
        log.Fatalf("Insert SEGMA SQL Table error: %v\n", err)
      }
    }
  }
}

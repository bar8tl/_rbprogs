// tables.go [2020-01-22 BAR8TL]
// Create Sqlite DB Tables based in a list
package rblib

import "code.google.com/p/go-sqlite/go1/sqlite3"
import "errors"
import "fmt"
import "log"

type Tlist_tp struct {
  Table string
  Sqlst string
}

type Tables_tp struct {
  Tlist []Tlist_tp
  Messg string
}

func NewTables() *Tables_tp {
  var t Tables_tp
  return &t
}

func (t *Tables_tp) CrtTables(cnnst, dbnam string, tlist []Tlist_tp) error {
  db, _ := sqlite3.Open(cnnst)
  defer db.Close()
  t.Tlist = tlist
  for _, tabl := range tlist {
    db.Exec(`DROP TABLE IF EXISTS ?;`, tabl.Table)
    err := db.Exec(tabl.Sqlst)
    if err != nil {
      t.Messg = fmt.Sprintf("Table %s creation error: %s", tabl.Table, err)
      return errors.New(t.Messg)
    }
    log.Printf("Table %-8s created...\r\n", tabl.Table)
  }
  log.Printf("Creation of dabatase %s completed.\r\n", dbnam)
  return nil
}

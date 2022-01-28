// crtdb.go [2017-05-24 BAR8TL]
// Create Sqlite DB and Selected tables
package rbidoc

import lib "bar8tl/p/rblib"

type Crtdb_tp struct {
  Tlist []lib.Tlist_tp
}

func NewCrtdb() *Crtdb_tp {
  var d Crtdb_tp
  return &d
}

func (d *Crtdb_tp) CrtTable(parm lib.Param_tp, s Settings_tp) {
  s.SetRunVars(parm)
  for _, cdb := range s.Cdb {
    for _, sq := range s.Sqlcr {
      if cdb.Table == sq.Table && cdb.Cr && sq.Activ {
        d.Tlist = append(d.Tlist, lib.Tlist_tp{sq.Table, sq.Sqlst})
        break
      }
    }
  }
  tbl := lib.NewTables()
  tbl.CrtTables(s.Cnnst, s.Dbonm, d.Tlist)
}

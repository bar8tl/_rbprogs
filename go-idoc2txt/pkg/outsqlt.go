// outsqlt.go [2017-05-24 BAR8TL]
// Addressing outputs to sqlite3 database
package rbidoc

import "code.google.com/p/go-sqlite/go1/sqlite3"
import "fmt"
import "log"
import "strconv"

type Outsqlt_tp struct {
  Db    *sqlite3.Conn
  Cnnst string
}

func (o *Outsqlt_tp) NewOutsqlt(s Settings_tp) {
  o.Cnnst = s.Cnnst
}

// ITEMS DB Options
func (o *Outsqlt_tp) ClearItems(idocn string) {
  var err error
  o.Db, err = sqlite3.Open(o.Cnnst)
  if err != nil {
    log.Fatalf("Open SQLite DB error: %v\n", err)
  }
  o.Db.Exec(`DELETE FROM items WHERE idocn=?;`, idocn)
}

func (o *Outsqlt_tp) IsrtItems(w Items_tp) {
  err := o.Db.Exec(`INSERT INTO items VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
    w.Idocn, w.Rname, w.Dname, w.Rclas, w.Rtype, w.Dtype, w.Dtext, w.Level,
    w.Stats, w.Minlp, w.Maxlp, w.Lngth, w.Seqno, w.Strps, w.Endps)
  if err != nil {
    log.Fatalf("Insert ITEMS SQL Table error: %v\n", err)
  }
}

// STRUC DB Options
type Struc_tp struct { // IDoc-Structure Descr (*=key field in DB record)
//Field:       //  GROUP                   SEGMENT
//-----------------------------------------------------------------------
  Idocn string //* Ex/Ba-Name              Ex/Ba-Name
  Strtp string //* 'GRP'                   'SGM'
  Level int    //  auto-gen                auto-gen
  // PARENT
  Prnam string //* p.rname='IDOC'/'GROUP'  p.rname='SEGMENT'
  Pseqn int    //* p.pseqn=autogen         p.pseqn=autogen
  Pdnam string //* p.dname=Group#          p.dname=Segm-ID
  Pdtyp string //  ''                      p.dtype=Segm-Type
  Pdqlf string //  ''                      'QUAL'
  // CHILD
  Crnam string //* c.rname='GROUP          c.rname*=Segm-ID
  Cseqn int    //* p.seqno=Group-Seq       p.seqno*=Seqno
  Cdnam string //* c.dname=Group#          c.dname*=Segm/Field-Name
  Cdtyp string //  ''                      c.dtype =Segm/Field-Type
  Cdqlf string //  ''                      'QUAL'
}

func (o *Outsqlt_tp) ClearStruc(idocn, strtp string) {
  var err error
  o.Db, err = sqlite3.Open(o.Cnnst)
  if err != nil {
    log.Fatalf("Open SQLite DB error: %v\n", err)
  }
  o.Db.Exec(`DELETE FROM struc WHERE idocn=? and strtp=?;`, idocn, strtp)
}

func (o *Outsqlt_tp) IsrtStruc(idocn, strtp string, pnode, cnode Keyst_tp) {
  if strtp == "GRP" {
    pd, _ := strconv.Atoi(pnode.Dname)
    pnode.Dname = fmt.Sprintf("%02d", pd)
    cd, _ := strconv.Atoi(cnode.Dname)
    cnode.Dname = fmt.Sprintf("%02d", cd)
  }
  var w Struc_tp
  w.Idocn = idocn
  w.Strtp = strtp
  w.Level = pnode.Level
  w.Prnam = pnode.Rname
  w.Pseqn = pnode.Pseqn
  w.Pdnam = pnode.Dname
  w.Pdtyp = pnode.Dtype
  w.Pdqlf = pnode.Dqual
  w.Crnam = cnode.Rname
  w.Cseqn = pnode.Seqno
  w.Cdnam = cnode.Dname
  w.Cdtyp = cnode.Dtype
  w.Cdqlf = cnode.Dqual
  err := o.Db.Exec(`INSERT INTO struc VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`,
    w.Idocn, w.Strtp, w.Level, w.Prnam, w.Pseqn, w.Pdnam, w.Pdtyp, w.Pdqlf,
    w.Crnam, w.Cseqn, w.Cdnam, w.Cdtyp, w.Cdqlf)
  if err != nil {
    log.Fatalf("Insert STRUC SQL Table error: %v\n", err)
  }
}

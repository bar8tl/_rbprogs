// idocdata.go [2017-05-24/BAR8TL]
// Read SAP-Idoc content in standard flat TXT format and upload data into
// internal structures
package rbidoc

import lib "bar8tl/p/rblib"
import "bufio"
import "code.google.com/p/go-sqlite/go1/sqlite3"
import "encoding/json"
import "io"
import "io/ioutil"
import "log"
import "os"
import "path/filepath"
import "strings"

const TRACE = true

type Field_tp struct {
  Key string
  Val string
}

type Rctrl_tp struct {
  Instn int
  Field []Field_tp
}

type Lctrl_tp struct {
  Rctrl []Rctrl_tp
}

type Rdata_tp struct {
  Segmn string
  Qualf string
  Level int
  Recno int
  Field []Field_tp
}

type Sdata_tp struct {
  Instn int
  Rdata  []Rdata_tp
}

type Ldata_tp struct {
  Sdata []Sdata_tp
}

type Rsegm_tp struct {
  Segmn string
  Recno int
  Level int
  Qlkey string
  Qlval string
  Instn int
  Field []Field_tp
  Child []Rsegm_tp
}

type Ssegm_tp struct {
  Instn int
  Rsegm []Rsegm_tp
}

type Lsegm_tp struct {
  Ssegm []Ssegm_tp
}

type Sfild_tp struct {
  Segmn string
  Recno int
  Level int
  Qlkey string
  Qlval string
  Field []Field_tp
}

type Count_tp struct {
  Segmn string
  Instn int
}

type Didoc_tp struct {
  Db     *sqlite3.Conn
  Rdr    *bufio.Reader
  Cnnst  string
  Idocn  string
  Recnf  int
  Setno  int
  Recno  int
  Lctrl  Lctrl_tp // Control list
  Sdata  Sdata_tp // Dataset
  Ldata  Ldata_tp // Dataset list
  Rsegm  Rsegm_tp // Segment record
  Ssegm  Ssegm_tp // Segmentset
  Lsegm  Lsegm_tp // Segmentset list
  Sfild  Sfild_tp
  Count  [9][]Count_tp
  L      int
  c1, c2 int
  c3, c4 int
  c5, c6 int
  c7, c8 int
}

func NewDidoc() *Didoc_tp {
  var d Didoc_tp
  return &d
}

func (d *Didoc_tp) ReadDire(parm lib.Param_tp, s Settings_tp) {
  s.SetRunVars(parm)
  d.Cnnst = s.Cnnst
  d.Idocn = strings.ToUpper(s.Objnm)
  files, _ := ioutil.ReadDir(s.Inpdr)
  for _, f := range files {
    d.ProcIndivFile(s, f)
  }
}

//******************************************************************************
// Process Input IDOC File
//******************************************************************************
func (d *Didoc_tp) ProcIndivFile(s Settings_tp, f os.FileInfo) {
  var err error
  d.Setno = -1 // Initialize Instance of data sets in the file
  d.Recnf =  0 // Initialize Number of data records in the file
  d.Db, err = sqlite3.Open(d.Cnnst)
  if err != nil {
    log.Fatalf("SQLite data base %s not found: %s\r\n", d.Cnnst, err)
  }
  defer d.Db.Close()
  ifile, err := os.Open(s.Inpdr + f.Name())
  if err != nil {
    log.Fatalf("Input IDOC file %s not found: %s\r\n", f.Name(), err)
  }
  defer ifile.Close()
  rdr := bufio.NewReader(ifile)
  d.ProcStartOfFile(rdr, f.Name())
  for iline, err := rdr.ReadString(byte('\n')); err != io.EOF;
    iline, err = rdr.ReadString(byte('\n')) {
    if strings.TrimSpace(iline[0:10]) == "EDI_DC40" {
      d.ReadControl(iline, d.Idocn, "CONTROL", false)
    } else {
      d.ReadData(iline, d.Idocn, "DATA")
    }
  }
  d.ProcEndOfFile(err, s.Outdr, f)
}

// Open input IDOC file and check first record is a Control Record
func (d *Didoc_tp) ProcStartOfFile(rdr *bufio.Reader, fname string) {
  iline, err := rdr.ReadString(byte('\n'))
  if err == io.EOF {
    log.Fatalf("Input IDOC file %s is empty: %s\r\n", fname, err)
  }
  if err != nil {
    log.Fatalf("Input IDOC file %s cannot be read: %s\r\n", fname, err)
  }
  if strings.TrimSpace(iline[0:10]) == "EDI_DC40" {
    d.ReadControl(iline, d.Idocn, "CONTROL", true)
  } else {
    log.Fatalf("IDOC File %s should start with Control Record\r\n", fname)
  }
}

// Fetch last records in structure to complete data detail in memory
func (d *Didoc_tp) ProcEndOfFile(err error, outdr string, f os.FileInfo) {
  ffile := f.Name()
  flext := filepath.Ext(ffile)
  flnam := strings.TrimRight(ffile, flext)
  if err != io.EOF && err != nil {
    log.Fatalf("Error during reading inout IDOC file %s %s:\r\n", flnam, err)
  }
  d.Ldata.Sdata = append(d.Ldata.Sdata, Sdata_tp{d.Setno, d.Sdata.Rdata})
  d.Ssegm.Rsegm = append(d.Ssegm.Rsegm, d.Rsegm)
  d.Lsegm.Ssegm = append(d.Lsegm.Ssegm, Ssegm_tp{d.Setno, d.Ssegm.Rsegm})
  if TRACE {
    fc, _ := json.MarshalIndent(d.Lctrl, "", " ")
    _ = ioutil.WriteFile(outdr + flnam + "-control.json", fc, 0644)
    fd, _ := json.MarshalIndent(d.Ldata, "", " ")
    _ = ioutil.WriteFile(outdr + flnam + "-data.json", fd, 0644)
    fs, _ := json.MarshalIndent(d.Lsegm, "", " ")
    _ = ioutil.WriteFile(outdr + flnam + "-segment.json", fs, 0644)
  }
}

//******************************************************************************
// Process Control Record
//******************************************************************************
const SELITEMS = `SELECT dname, strps, endps FROM items WHERE idocn=? and
  rname=? order by seqno;`

func (d *Didoc_tp) ReadControl(iline, idocn, rname string, first bool) {
  var f     Items_tp
  var rctrl Rctrl_tp
  var cdval string
  if !first {
    d.Ldata.Sdata = append(d.Ldata.Sdata, Sdata_tp{d.Setno, d.Sdata.Rdata})
    d.Sdata.Rdata = nil
    d.Ssegm.Rsegm = append(d.Ssegm.Rsegm, d.Rsegm)
    d.Lsegm.Ssegm = append(d.Lsegm.Ssegm, Ssegm_tp{d.Setno, d.Ssegm.Rsegm})
    d.Ssegm.Rsegm = nil
  }
  d.Recno = 0                             // Inits at Control Record level
  d.L = -1                                //
  d.c1, d.c2, d.c3, d.c4 = -1, -1, -1, -1 //
  d.c5, d.c6, d.c7, d.c8 = -1, -1, -1, -1 //
  d.Setno++
  d.Recnf++
  for dbo, err := d.Db.Query(SELITEMS, idocn, rname); err == nil;
    err = dbo.Next() {
    dbo.Scan(&f.Dname, &f.Strps, &f.Endps)
    cdval = strings.TrimSpace(iline[f.Strps-1:f.Endps])
    if len(cdval) == 0 || cdval == "" {
      continue
    }
    rctrl.Field = append(rctrl.Field, Field_tp{f.Dname, cdval})
  }
  rctrl.Instn = d.Setno
  d.Lctrl.Rctrl = append(d.Lctrl.Rctrl, rctrl)
  d.AddRoot(idocn)
}

// Define root node in segment structure
func (d *Didoc_tp) AddRoot(idocn string) {
  d.Rsegm = Rsegm_tp{idocn, 0, 0, "", "", 0, nil, nil}
}

//******************************************************************************
// Process Data Record
//******************************************************************************
const SELITEMF = `SELECT dname, dtype, dtext, level FROM items WHERE idocn=?
  and dname=? and rname=?;`

func (d *Didoc_tp) ReadData(iline, idocn, rname string) {
  var f, g  Items_tp
  var rdata Rdata_tp
  var cdval string
  d.Recnf++
  d.Recno++
  for dbo, err := d.Db.Query(SELITEMS, idocn, rname); err == nil;
    err = dbo.Next() {
    dbo.Scan(&f.Dname, &f.Strps, &f.Endps)
    if f.Endps >= len(iline) {
      f.Endps = len(iline)
    }
    cdval = strings.TrimSpace(iline[f.Strps-1:f.Endps])
    if len(cdval) == 0 || cdval == "" {
      continue
    }
    if f.Dname == "SEGNAM" {
      dbs, err := d.Db.Query(SELITEMF, idocn, d.alias(idocn, cdval), "SEGMENT")
      if err != nil {
        log.Printf("Select ITEMS table error: %v\n", err)
      }
      err = dbs.Scan(&g.Dname, &g.Dtype, &g.Dtext, &g.Level)
      if err != nil {
        log.Printf("Scan ITEMS table error: %v\n", err)
      }
      rdata.Segmn = g.Dtype
      rdata.Qualf = g.Dtext
      rdata.Level = g.Level
      rdata.Recno = d.Recno
    }
    if f.Dname == "SDATA" {
      d.ProcSegment(iline, idocn, "SGM", g.Dname, rdata.Level)
      continue
    }
    rdata.Field = append(rdata.Field, Field_tp{f.Dname, cdval})
  }
  d.Sdata.Rdata = append(d.Sdata.Rdata, rdata)
}

// Get the Segment Type from a Segment Alias
const SELALIAS = `select segtp from segma where idocn=? and segdf=?;`

func (d *Didoc_tp) alias(idocn, aname string) string {
  dname := aname
  dbo, err := d.Db.Query(SELALIAS, idocn, aname)
  if err != nil {
    return dname
  }
  dbo.Scan(&dname)
  return dname
}

//******************************************************************************
// Process Segment Data
//******************************************************************************
// Determines segment Qualifier and Instance Number
func (d *Didoc_tp) ProcSegment(iline, idocn, strtp, cdnam string, level int) {
  instn := -1
  ident := ""
  if level == d.L {
    instn = d.UpdateCounter(cdnam, d.L)
    ident = "SAME"
  } else if level > d.L {
    d.L = level
    d.Count[d.L] = append(d.Count[d.L], Count_tp{cdnam, 1})
    instn = d.RetrieveCounter(cdnam, d.L)
    ident = "LOWER"
  } else if level < d.L {
    goupl := d.L - level
    for i := 0;  i < goupl; i++ {
      d.Count[d.L] = nil
      d.L--
    }
    instn = d.UpdateCounter(cdnam, d.L)
    ident = "UPPER"
  }
  d.AddToStruct(iline, idocn, ident, cdnam, d.L, instn)
}

// Update counter of segment with equal segment ID in the current struct level
func (d *Didoc_tp) UpdateCounter(segmn string, l int) int {
  for j := 0; j < len(d.Count[l]); j++ {
    if d.Count[l][j].Segmn == segmn {
      d.Count[l][j].Instn += 1
      return d.Count[l][j].Instn
    }
  }
  d.Count[l] = append(d.Count[l], Count_tp{segmn, 1})
  return 1
}

// Retrieve last counter of segment with equal segm ID in the current struct lvl
func (d *Didoc_tp) RetrieveCounter(segmn string, l int) int {
  for j := 0; j < len(d.Count[l]); j++ {
    if d.Count[l][j].Segmn == segmn {
      return d.Count[l][j].Instn
    }
  }
  return -1
}

// Build segment structure into an non-linked segment node
func (d *Didoc_tp) AddToStruct(iline, idocn, ident, segmn string, l, instn int){
  if d.Recno <= 9999 {
  d.Sfild.Qlkey = ""
  d.Sfild.Qlval = ""
  d.Sfild.Field = nil
  d.GetSegmData(iline, idocn, "SGM", segmn, l)
  if l == 1 {
    d.Rsegm.Child = append(d.Rsegm.Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c2, d.c3, d.c4, d.c5, d.c6, d.c7, d.c8 = -1, -1, -1, -1, -1, -1, -1
    d.c1++
  } else if l == 2 {
    d.Rsegm.Child[d.c1].Child = append(
      d.Rsegm.Child[d.c1].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c3, d.c4, d.c5, d.c6, d.c7, d.c8 = -1, -1, -1, -1, -1, -1
    d.c2++
  } else if l == 3 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c4, d.c5, d.c6, d.c7, d.c8 = -1, -1, -1, -1, -1
    d.c3++
  } else if l == 4 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c5, d.c6, d.c7, d.c8 = -1, -1, -1, -1
    d.c4++
  } else if l == 5 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c6, d.c7, d.c8 = -1, -1, -1
    d.c5++
  } else if l == 6 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c7, d.c8 = -1, -1
    d.c6++
  } else if l == 7 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child[d.c6].Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child[d.c6].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c8 = -1
    d.c7++
  } else if l == 8 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child[d.c6].Child[d.c7].Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child[d.c6].Child[d.c7].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
    d.c8++
  } else if l == 9 {
    d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child[d.c6].Child[d.c7].Child[d.c8].Child = append(
      d.Rsegm.Child[d.c1].Child[d.c2].Child[d.c3].Child[d.c4].Child[d.c5].
      Child[d.c6].Child[d.c7].Child[d.c8].Child, Rsegm_tp{
      segmn, d.Recno, l, d.Sfild.Qlkey, d.Sfild.Qlval, instn,
      d.Sfild.Field, nil})
  }
  }
}

//Get field values of a segment into the IDOC structure
const SELSTRUC = `SELECT a.idocn, a.level, a.pseqn, a.pdnam, a.pdtyp, a.pdqlf,
  a.cseqn, a.cdnam, a.cdtyp, a.cdqlf, b.dname, b.seqno, b.strps, b.endps
  FROM struc a LEFT JOIN items b
  ON (a.idocn = b.idocn and a.cdnam = b.rname)
  WHERE a.idocn=? and a.strtp=? and a.cdnam=?
  ORDER BY a.idocn, a.strtp, a.pseqn, a.prnam, a.pdnam, b.seqno;`

func (d *Didoc_tp) GetSegmData(iline, idocn, strtp, cdnam string, level int) {
  var f Items_tp
  var e Struc_tp
  var cdval string
  fitem := true
  for dbo, err := d.Db.Query(SELSTRUC, idocn, strtp, cdnam); err == nil;
    err = dbo.Next() {
    dbo.Scan( &e.Idocn, &e.Level, &e.Pseqn, &e.Pdnam, &e.Pdtyp, &e.Pdqlf,
      &e.Cseqn, &e.Cdnam, &e.Cdtyp, &e.Cdqlf, &f.Dname, &f.Seqno, &f.Strps,
      &f.Endps)
    if f.Endps >= len(iline) {
      break
    }
    cdval = strings.TrimSpace(iline[f.Strps-1:f.Endps])
    if len(cdval) == 0 || cdval == "" {
      continue
    }
    if fitem {
      d.Sfild.Segmn = e.Cdtyp
      d.Sfild.Recno = d.Recno
      d.Sfild.Level = e.Level
      if e.Cdqlf == "QUAL" {
        d.Sfild.Qlkey = f.Dname
        d.Sfild.Qlval = cdval
      } else {
        d.Sfild.Qlkey = ""
        d.Sfild.Qlval = ""
      }
      fitem = false
    }
    d.Sfild.Field = append(d.Sfild.Field, Field_tp{f.Dname, cdval})
  }
}

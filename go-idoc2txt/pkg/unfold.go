// unfold.go [2017-05-24 BAR8TL]
// Convert IDoc classic hierarchical format to flat text file format
package rbidoc

import lib "bar8tl/p/rblib"
import "bufio"
import "code.google.com/p/go-sqlite/go1/sqlite3"
import "fmt"
import "io/ioutil"
import "log"
import "io"
import "os"
import "strconv"
import "strings"

type Hstruc_tp struct {
  Sgnum string
  Sgnam string
  Sglvl string
}

type Unfold_tp struct {
  Idocx string
  Idocn string
  Idocb string
  Sectn string
  Secnb string
  Sgnum string
  Sgnam string
  Sgdsc string
  Sgnbk string
  Sghnb string
  Sglvl string
  Serie string
  Nsegm int
  Dirty bool
  Lctrl [524]byte
  Lsegm [1063]byte
  Lstat [562]byte
  Ifile *os.File
  Ofile *os.File
  Db    *sqlite3.Conn
  Parnt []Hstruc_tp
  L     int
}

// Constructor of object Dunf: Define input/output file and database location
// folders, database full connection string as well
func NewUnfold() *Unfold_tp {
  var u Unfold_tp
  return &u
}

// Public option UNF: Unfold data IDocs based on specific IDoc-type. Produces
// system readeable flat text files
func (u *Unfold_tp) UfldData(parm lib.Param_tp, s Settings_tp) {
  s.SetRunVars(parm)
  fmt.Println("entry", s.Inpdr, s.Outdr)
  u.Idocx = strings.ToUpper(s.Objnm)
  files, _ := ioutil.ReadDir(s.Inpdr)
  for _, f := range files {
    fmt.Println(f)
    u.ProcDataLines(s, f)
    if len(s.Ifilt) == 0 || (len(s.Ifilt) > 0 && PassFilter(s, f)) {
      u.ProcDataLines(s, f)
    }
  }
}

// Function to process IDoc data files, reading line by line and determining
// measures for format conversion
func (u *Unfold_tp) ProcDataLines(s Settings_tp, f os.FileInfo) {
  u.OpenProgStreams(s, f).DetermIdocProps()
  u.Idocn = ""
  u.Nsegm = 0
  u.L = -1
  u.Parnt = u.Parnt[:u.L+1]
  rdr := bufio.NewReader(u.Ifile)
  wtr := bufio.NewWriter(u.Ofile)
  for l, err := rdr.ReadString(byte('\n')); err != io.EOF;
    l, err = rdr.ReadString(byte('\n')) {
    l = strings.TrimSpace(l)
    t := strings.Split(l, "\t")
    if len(l) == 0 { // ignores lines in blank
      continue
    }
    log.Printf("%s\r\n", l)

    // Gets IDoc number
    if len(u.Idocn) == 0 && len(t) == 1 && l[0:11] == "IDoc Number" {
      i := strings.Split(l, " : ")
      u.Idocn = strings.TrimSpace(i[1])
      continue
    }

    // Ignores lines no containing tabulators (after to have gotten IDoc number)
    if len(t) <= 1 {
      continue
    }

    // Determines data section to analyze
    if t[0] == "EDIDC" || t[0] == "EDIDD" || t[0] == "EDIDS" {
      u.SetupSection(s, t, wtr)
      continue
    }

    // Checks in segment number to analize
    if t[0] == "SEGNUM" && len(t) == 3 {
      u.Sgnbk = u.Sgnum
      u.Sgnum = t[2]
      continue
    }

    // Checks in segment name to analize
    if t[0] == "SEGNAM" && len(t) == 3 {
      u.SetupSegment(s, t, wtr)
      continue
    }

    // Process fields of each data section
    if u.Sectn == "EDIDC" {
      u.procEdidc(t)
    } else if u.Sectn == "EDIDD" {
      u.procEdidd(u.Sgnum, u.Sgnam, t)
    } else if u.Sectn == "EDIDS" {
      u.procEdids(u.Secnb, t)
    }
  }
  u.CloseProgStreams(s, f)
}

// Function to setup measures to take for each data section. Each new section
// causes dumping data from previous one
func (u *Unfold_tp) SetupSection(s Settings_tp, t []string, wtr *bufio.Writer) {
  u.Sectn = t[0]
  if u.Sectn == "EDIDC" {
    for i := 0; i < len(u.Lctrl); i++ {
      u.Lctrl[i] = ' '
    }
  }
  if u.Sectn == "EDIDD" {
    u.DumpControlLine(s, wtr)
  }
  if u.Sectn == "EDIDS" {
    u.Sgnbk = u.Sgnum
    u.DumpSegmentLine(s, wtr)
    for i := 0; i < len(u.Lstat); i++ {
      u.Lstat[i] = ' '
    }
    if len(t) == 3 {
      u.Secnb = t[2]
    }
  }
}

// Function to setup measures to take for each data segment in Data Idoc being
// converted
func (u *Unfold_tp) SetupSegment(s Settings_tp, t []string, wtr *bufio.Writer) {
  u.Nsegm++
  if u.Nsegm > 1 {
    u.DumpSegmentLine(s, wtr)
  }
  u.Sgnam = t[2]
  for i := 0; i < len(u.Lsegm); i++ {
    u.Lsegm[i] = ' '
  }
  rdb, err := u.Db.Query(
    `SELECT dname, level FROM items WHERE idocn=? and rname=? and dtype=?;`,
    u.Idocx, "SEGMENT", u.Sgnam)
  if err != nil {
    log.Fatalf("Error during searching segment description: %s %s\r\n",
      u.Sgnam, err)
  }
  var level int
  rdb.Scan(&u.Sgdsc, &level)
  u.Sglvl = fmt.Sprintf("%02d", level)

  if u.Nsegm == 1 {
    u.Parnt = append(u.Parnt, Hstruc_tp{u.Sgnum, u.Sgnam, u.Sglvl})
    u.L++
    u.Sghnb = "000000"
  } else {
    if u.Sglvl > u.Parnt[u.L].Sglvl {
      u.Parnt = append(u.Parnt, Hstruc_tp{u.Sgnum, u.Sgnam, u.Sglvl})
      u.L++
      u.Sghnb = u.Parnt[u.L-1].Sgnum
    } else if u.Sglvl == u.Parnt[u.L].Sglvl {
      u.Parnt[u.L].Sgnum = u.Sgnum
      u.Parnt[u.L].Sgnam = u.Sgnam
      u.Parnt[u.L].Sglvl = u.Sglvl
      u.Sghnb = u.Parnt[u.L-1].Sgnum
    } else {
      prvlv, _ := strconv.Atoi(u.Parnt[u.L].Sglvl)
      curlv, _ := strconv.Atoi(u.Sglvl)
      nstep := prvlv - curlv
      for i := 1; i <= nstep; i++ {
        u.L--
        u.Parnt = u.Parnt[:u.L+1]
      }
      u.Parnt[u.L].Sgnum = u.Sgnum
      u.Parnt[u.L].Sgnam = u.Sgnam
      u.Parnt[u.L].Sglvl = u.Sglvl
      u.Sghnb = u.Parnt[u.L-1].Sgnum
    }
  }
  rdb.Close()
}

// Functions to process format conversion to fields in control record
func (u *Unfold_tp) procEdidc(t []string) {
  flkey := t[0]
  if flkey == "RVCPRN" {
    flkey = "RCVPRN"
  }
  flval := ""
  if len(t) == 3 {
    f := strings.Split(t[2], " :")
    flval = strings.TrimSpace(f[0])
  }
  if flkey == "CREDAT" {
    u.Serie = flval
  }
  if flkey == "CRETIM" {
    u.Serie += flval
  }
  if len(flval) > 0 {
    u.Dirty = true
    u.SetControlField(flkey, flval)
  }
}

func (u *Unfold_tp) SetControlField(flkey, flval string) {
  var strps int
  rdb, err := u.Db.Query(
    `SELECT strps FROM items WHERE idocn=? and rname=? and dname=?;`,
    u.Idocx, "CONTROL", flkey)
  if err != nil {
    log.Fatalf("Error during reading database for control data: %v\r\n", err)
  }
  rdb.Scan(&strps)
  rdb.Close()
  if flkey == "IDOCTYP" && flval == "14" {
    flval = u.Idocb
  }
  if flkey == "CIMTYP"  && flval == "14" {
    flval = u.Idocx
  }
  k := strps - 1
  for i := 0; i < len(flval); i++ {
    u.Lctrl[k] = flval[i]
    k++
  }
}

func (u *Unfold_tp) DumpControlLine(s Settings_tp, wtr *bufio.Writer) {
  if u.Dirty {
    u.SetControlField("TABNAM", s.Cntrl)
    u.SetControlField("MANDT",  s.Clien)
    u.SetControlField("DOCNUM", u.Idocn)
    u.SetControlField("RCVPFC", s.Rcvpf)
    u.SetControlField("SERIAL", u.Serie)
    fmt.Fprintf(wtr, "%s\r\n", u.Lctrl)
    wtr.Flush()
    u.Dirty = false
  }
}

// Functions to process format conversion to fields in data records
func (u *Unfold_tp) procEdidd(sgnum, sgnam string, t []string) {
  flkey := t[0]
  flval := ""
  if len(t) == 3 {
    f := strings.Split(t[2], " :")
    flval = strings.TrimSpace(f[0])
  }
  if len(flval) > 0 {
    u.Dirty = true
    u.SetSegmentField(u.Sgdsc, flkey, flval)
  }
}

func (u *Unfold_tp) SetSegmentField(sgdsc, flkey, flval string) {
  var strps int
  rdb, err := u.Db.Query(
    `SELECT strps FROM items WHERE idocn=? and rname=? and dname=?;`,
    u.Idocx, sgdsc, flkey)
  if err != nil {
    log.Fatalf(
      "Error during reading database for segment data: %s %s %s %v\r\n",
      u.Idocx, u.Sgdsc, flkey, err)
  }
  rdb.Scan(&strps)
  rdb.Close()
  k := strps - 1
  for i := 0; i < len(flval); i++ {
    u.Lsegm[k] = flval[i]
    k++
  }
}

func (u *Unfold_tp) DumpSegmentLine(s Settings_tp, wtr *bufio.Writer) {
  if u.Dirty {
    u.SetSegmentField("DATA", "SEGNAM", u.Sgdsc)
    u.SetSegmentField("DATA", "MANDT",  s.Clien)
    u.SetSegmentField("DATA", "DOCNUM", u.Idocn)
    u.SetSegmentField("DATA", "SEGNUM", u.Sgnbk)
    u.SetSegmentField("DATA", "PSGNUM", u.Sghnb)
    u.SetSegmentField("DATA", "HLEVEL", u.Sglvl)
    fmt.Fprintf(wtr, "%s\r\n", u.Lsegm)
    fmt.Println(wtr)
    wtr.Flush()
    u.Dirty = false
  }
}

func (u *Unfold_tp) procEdids(secnb string, t []string) {}

func (u *Unfold_tp) OpenProgStreams(s Settings_tp, f os.FileInfo) *Unfold_tp {
  var err error
  u.Db, err = sqlite3.Open(s.Cnnst)
  if err != nil {
    log.Fatalf("Open SQLite database error: %s\n", err)
  }
  u.Ifile, err = os.Open(s.Inpdr + f.Name())
  if err != nil {
    log.Fatalf("Input file not found: %s\r\n", err)
  }
  u.Ofile, err = os.Create(s.Outdr + f.Name())
  if err != nil {
    log.Fatalf("Error during output file creation: %s\r\n", err)
  }
  return u
}

func (u *Unfold_tp) CloseProgStreams(s Settings_tp, f os.FileInfo) *Unfold_tp {
  u.Ifile.Close()
  u.Ofile.Close()
  u.Db.Close()
  RanameInpFile(s.Inpdr, f)
  RanameOutFile(s.Outdr, f)
  return u
}

func (u *Unfold_tp) DetermIdocProps() *Unfold_tp {
  rdb, err := u.Db.Query(
    "SELECT dname FROM items WHERE idocn=? and rname=?;", u.Idocx, "IDOC")
  if err != nil {
    log.Fatalf("Error in searching Idoc properties: %s %v\r\n", u.Idocx, err)
  }
  rdb.Scan(&u.Idocb)
  rdb.Close()
  return u
}

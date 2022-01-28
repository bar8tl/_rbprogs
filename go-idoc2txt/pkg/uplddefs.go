// uplddefs.go [2017-05-24 BAR8TL]
// Read SAP IDoc parser file and to upload IDoc definition detail and structure
// into an internal reference database
package rbidoc

import lib "bar8tl/p/rblib"
import "bufio"
import "io"
import "log"
import "os"
import "strings"

type Uplddefs_tp struct {
  ui Upldmitm_tp
  ug Upldsgrp_tp
  us Upldssgm_tp
}

func NewUplddefs() *Uplddefs_tp {
  var u Uplddefs_tp
  return &u
}

func (u *Uplddefs_tp) UpldDefs(parm lib.Param_tp, s Settings_tp) {
  s.SetRunVars(parm)
  ifile, err := os.Open(s.Inpdr + s.Objnm)
  if err != nil {
    log.Fatalf("Input file %s not found: %s\r\n", s.Inpdr+s.Objnm, err)
  }
  defer ifile.Close()
  u.ProcStartOfFile(s)
  rdr := bufio.NewReader(ifile)
  for l, _, err := rdr.ReadLine(); err != io.EOF; l, _, err = rdr.ReadLine() {
    if line := strings.TrimSpace(string(l)); len(line) > 0 {
      sline := lib.ScanTextIdocLine(line)
      u.ProcLinesOfFile(s, sline)
    }
  }
  u.ProcEndOfFile(s)
}

func (u *Uplddefs_tp) ProcStartOfFile(s Settings_tp) {
  if s.Mitm {
    u.ui.NewUpldmitm(s)
  }
  if s.Sgrp {
    u.ug.NewUpldsgrp(s, s.Konst.GRP)
  }
  if s.Ssgm {
    u.us.NewUpldssgm(s, s.Konst.SGM)
  }
}

func (u *Uplddefs_tp) ProcLinesOfFile(s Settings_tp, sline lib.Parsl_tp) {
  if s.Mitm {
    u.ui.GetData(sline)
  }
  if s.Sgrp {
    u.ug.GetData(sline)
  }
  if s.Ssgm {
    u.us.GetData(sline)
  }
}

func (u *Uplddefs_tp) ProcEndOfFile(s Settings_tp) {
  if s.Mitm {
    u.ui.IsrtData(s)
  }
}

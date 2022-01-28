// envmnt.go [2017-05-24/BAR8TL]
// Global environment variables
package rbidoc

import lib "bar8tl/p/rblib"
import "time"

type Envmnt_tp struct {
  Cnnsq string
  Cnnst string
  Cntrl string
  Clien string
  Dbonm string
  Dbodr string
  Inpdr string
  Outdr string
  Qrydr string
  Ifilt string
  Ifnam string
  Ofnam string
  Objnm string
  Qrynm string
  Rcvpf string
  Found bool
  Mitm  bool
  Sgrp  bool
  Ssgm  bool
  Dtsys time.Time
  Dtcur time.Time
  Dtnul time.Time
}

func (e *Envmnt_tp) NewEnvmnt(s Settings_tp) {
  e.Cnnsq = s.Dflt.CNNS_SQLIT3
  e.Cntrl =
    lib.Ternary_op(len(s.Const.Cntrl) > 0, s.Const.Cntrl, s.Dflt.CONTROL_CODE)
  e.Clien =
    lib.Ternary_op(len(s.Const.Clien) > 0, s.Const.Clien, s.Dflt.CLIENT_CODE)
  e.Dbonm =
    lib.Ternary_op(len(s.Progm.Dbonm) > 0, s.Progm.Dbonm, s.Dflt.DB_NAME)
  e.Dbodr =
    lib.Ternary_op(len(s.Progm.Dbodr) > 0, s.Progm.Dbodr, s.Dflt.DB_DIR)
  e.Inpdr =
    lib.Ternary_op(len(s.Progm.Inpdr) > 0, s.Progm.Inpdr, s.Dflt.INPUTS_DIR)
  e.Outdr =
    lib.Ternary_op(len(s.Progm.Outdr) > 0, s.Progm.Outdr, s.Dflt.OUTPUTS_DIR)
  e.Ifilt =
    lib.Ternary_op(len(s.Progm.Ifilt) > 0, s.Progm.Ifilt, s.Dflt.INPUTS_FILTER)
  e.Ifnam =
    lib.Ternary_op(len(s.Progm.Ifnam) > 0, s.Progm.Ifnam, s.Dflt.INPUTS_NAMING)
  e.Ofnam =
    lib.Ternary_op(len(s.Progm.Ofnam) > 0, s.Progm.Ofnam, s.Dflt.OUTPUTS_NAMING)
  e.Dtsys = time.Now()
  e.Dtcur = time.Now()
  e.Dtnul = time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
}

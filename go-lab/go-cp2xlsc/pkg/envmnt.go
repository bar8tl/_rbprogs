// envmnt.go - Definition of program environment variables
// 2022-05-17 BAR8TL Version1.0 - Released
package cp2xlsc

import lib "bar8tl/p/rblib"
import "time"

type Envmnt_tp struct {
  Inpdr string
  Outdr string
  Fildr string
  Ifilt string
  Ifnam string
  Ofnam string
  Objnm string
  Modep string
  Outpt string
  Infil string
  Flfil string
  Flnam string
  Flext string
  Found bool
  Dtsys time.Time
  Dtcur time.Time
  Dtnul time.Time
}

func (e *Envmnt_tp) NewEnvmnt(s Settings_tp) {
  e.Inpdr =
    lib.Ternary_op(len(s.Progm.Inpdr) > 0, s.Progm.Inpdr, s.Dflt.INPUTS_DIR)
  e.Outdr =
    lib.Ternary_op(len(s.Progm.Outdr) > 0, s.Progm.Outdr, s.Dflt.OUTPUTS_DIR)
  e.Fildr =
    lib.Ternary_op(len(s.Progm.Fildr) > 0, s.Progm.Fildr, s.Dflt.FILES_DIR)
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

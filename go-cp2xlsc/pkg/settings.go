// settings.go - Container of program/run-option level settings
// 2022-05-17 BAR8TL Version1.0 - Released
package cp2xlsc

import ut "bar8tl/p/rblib"
import "log"

type Settings_tp struct {
  Config_tp
  ut.Parms_tp
  Deflts_tp
  Envmnt_tp
}

func NewSettings(cfnam, dfnam string) Settings_tp {
  var s Settings_tp
  s.NewParms()
  s.NewConfig(cfnam)
  s.NewDeflts(dfnam)
  s.NewEnvmnt(s)
  return s
}

func (s *Settings_tp) SetRunVars(p ut.Param_tp) {
  if len(p.Prm1) > 0 {
    s.Objnm = p.Prm1
  } else {
    log.Fatalf("Error: Not possible to determine EDICOM Type name.\r\n")
  }
  s.Found = false
  for _, run := range s.Run {
    if p.Optn == run.Optcd && p.Prm1 == run.Objnm {
      if p.Optn == "txc" {
        s.Objnm = ut.Ternary_op(len(run.Objnm) > 0, run.Objnm, s.Objnm)
        s.Modep = ut.Ternary_op(len(run.Modep) > 0, run.Modep, s.Modep)
        s.Outpt = ut.Ternary_op(len(run.Outpt) > 0, run.Outpt, s.Outpt)
        s.Infil = ut.Ternary_op(len(run.Infil) > 0, run.Infil, s.Infil)
      }
      s.Found = true
      break
    }
  }
}

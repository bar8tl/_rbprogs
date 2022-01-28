// settings.go [2017-05-24 BAR8TL]
// Container of pgm-level & run-level settings
package rbidoc

import lib "bar8tl/p/rblib"
import "log"
import "strings"

type Settings_tp struct {
  Config_tp
  lib.Parms_tp
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

func (s *Settings_tp) SetRunVars(p lib.Param_tp) {
  if len(p.Prm1) > 0 {
    s.Objnm = p.Prm1
  } else {
    log.Fatalf("Error: Not possible to determine IDOC-Type name.\r\n")
  }
  s.Found = false
  for _, run := range s.Run {
    if p.Optn == run.Optcd && p.Prm1 == run.Objnm {
      if p.Optn == "cdb" || p.Optn == "upl" || p.Optn == "unf" ||
        p.Optn == "dat" || p.Optn == "usa" || p.Optn == "qry" {
        s.Objnm = lib.Ternary_op(len(run.Objnm) > 0, run.Objnm, s.Objnm)
        s.Dbonm = lib.Ternary_op(len(run.Dbonm) > 0, run.Dbonm, s.Dbonm)
        s.Dbodr = lib.Ternary_op(len(run.Dbodr) > 0, run.Dbodr, s.Dbodr)
      }
      if p.Optn == "upl" || p.Optn == "unf" || p.Optn == "dat" ||
        p.Optn == "usa" || p.Optn == "qry" {
        s.Inpdr = lib.Ternary_op(len(run.Inpdr) > 0, run.Inpdr, s.Inpdr)
        s.Outdr = lib.Ternary_op(len(run.Outdr) > 0, run.Outdr, s.Outdr)
      }
      if p.Optn == "unf" || p.Optn == "dat" || p.Optn == "usa" {
        s.Ifilt = lib.Ternary_op(len(run.Ifilt) > 0, run.Ifilt, s.Ifilt)
        s.Ifnam = lib.Ternary_op(len(run.Ifnam) > 0, run.Ifnam, s.Ifnam)
        s.Ofnam = lib.Ternary_op(len(run.Ofnam) > 0, run.Ofnam, s.Ofnam)
        s.Rcvpf = lib.Ternary_op(len(run.Rcvpf) > 0, run.Rcvpf, s.Rcvpf)
      }
      if p.Optn == "qry" {
        s.Qrydr = lib.Ternary_op(len(run.Qrydr) > 0, run.Qrydr, s.Qrydr)
        s.Qrynm = lib.Ternary_op(len(run.Qrynm) > 0, run.Qrynm, s.Qrynm)
      }
      s.Found = true
      break
    }
  }
  if p.Optn == "upl" {
    s.Mitm = true
    s.Sgrp = false
    s.Ssgm = false
    if len(p.Prm2) > 0 {
      mflds := strings.Split(p.Prm2, ".")
      for i := 0; i < len(mflds); i++ {
        switch strings.ToLower(mflds[i]) {
        case s.Konst.ITM:
          s.Mitm = true
        case s.Konst.GRP:
          s.Sgrp = true
        case s.Konst.SGM:
          s.Ssgm = true
        default:
          s.Mitm = true
          s.Sgrp = false
          s.Ssgm = false
        }
      }
    }
  }
  s.Cnnst = strings.Replace(s.Cnnsq, "@", s.Dbodr+s.Dbonm, 1)
}

// upldsgrp.go [2017-05-24 BAR8TL]
// Get IDoc groups structure data and to create corresponding structure records
// in ref database
package rbidoc

import lib "bar8tl/p/rblib"
import "strings"

type Upldsgrp_tp struct {
  Out   Outsqlt_tp
  Stack []Keyst_tp
  Idocn string
  Strtp string
  L     int
  Gseqn int
}

func (u *Upldsgrp_tp) NewUpldsgrp(s Settings_tp, strtp string) {
  u.Strtp = strings.ToUpper(strtp)
  u.L = -1
  u.Out.NewOutsqlt(s)
}

func (u *Upldsgrp_tp) GetData(sline lib.Parsl_tp) {
  if sline.Label.Ident == "BEGIN" {
    if sline.Label.Recnm == "IDOC" {
      u.Stack = append(u.Stack, Keyst_tp{sline.Label.Recnm, sline.Value, "", "",
        0, 0, 0})
      u.L++
      u.Idocn = sline.Value
      u.Out.ClearStruc(u.Idocn, u.Strtp)
    } else if sline.Label.Recnm == "GROUP" {
      u.Stack[u.L].Seqno += 1
      u.Stack = append(u.Stack, Keyst_tp{sline.Label.Recnm, sline.Value, "", "",
        0, 0, 0})
      u.L++
    }
    return
  }
  if sline.Label.Ident == "END" {
    if sline.Label.Recnm == "IDOC" {
      u.Stack = u.Stack[:u.L]
      u.L--
    } else if sline.Label.Recnm == "GROUP" {
      u.Gseqn++
      u.Stack[u.L-1].Pseqn = u.Gseqn
      u.Out.IsrtStruc(u.Idocn, u.Strtp, u.Stack[u.L-1], u.Stack[u.L])
      u.Stack = u.Stack[:u.L]
      u.L--
    }
    return
  }
  if u.L >= 0 && u.Stack[u.L].Rname == "IDOC" {
    if sline.Label.Ident == "EXTENSION" {
      u.Idocn = sline.Value
      u.Out.ClearStruc(u.Idocn, u.Strtp)
    }
    return
  }
}

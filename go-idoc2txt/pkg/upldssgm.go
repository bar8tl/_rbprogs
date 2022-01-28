// upldssgm.go [2017-05-24 BAR8TL]
// Get IDoc segments structure data and to create corresponding structure
// records in ref database
package rbidoc

import lib "bar8tl/p/rblib"
import "strconv"
import "strings"

type Keyst_tp struct { // Structure Node Attributes
//Field:       // IDOC        GROUP      SEGMENT
//-------------------------------------------------
  Rname string // 'IDOC'      'GROUP'    'SEGMENT'
  Dname string // Basic-IDoc  Group#     Segm-ID
  Dtype string // ''          ''         Segm-Type
  Dqual string // ''          ''         'QUAL'
  Level int    // 0           Level      Level
  Pseqn int    // 0           auto-gen   auto-gen
  Seqno int    // 0           Group-Seq  Segm-Seq
}

type Upldssgm_tp struct {
  Out   Outsqlt_tp
  Stack []Keyst_tp
  Tnode Keyst_tp
  Fnode Keyst_tp
  Snode Keyst_tp
  Idocn string
  Strtp string
  L     int
  Sseqn int
}

func (u *Upldssgm_tp) NewUpldssgm(s Settings_tp, strtp string) {
  u.Strtp = strings.ToUpper(strtp)
  u.L = -1
  u.Out.NewOutsqlt(s)
}

func (u *Upldssgm_tp) GetData(sline lib.Parsl_tp) {
  if sline.Label.Ident == "BEGIN" {
    if sline.Label.Recnm == "IDOC" {
      u.Stack = append(u.Stack, Keyst_tp{sline.Label.Recnm, sline.Value, "", "",
        0, 0, 0})
      u.L++
      u.Tnode.Rname = sline.Label.Recnm
      u.Tnode.Dname = sline.Value
      u.Tnode.Dqual = ""
      u.Tnode.Pseqn = 0
      u.Idocn       = sline.Value
      u.Out.ClearStruc(u.Idocn, u.Strtp)
    } else if sline.Label.Recnm == "SEGMENT" && len(sline.Label.Rectp) == 0 {
      u.Sseqn++
      u.Tnode.Rname = sline.Label.Recnm
      u.Tnode.Dname = sline.Value
      u.Tnode.Dqual = ""
      u.Tnode.Pseqn = u.Sseqn
    }
    return
  }

  if sline.Label.Ident == "END" && u.L >= 0 {
    if sline.Label.Recnm == "IDOC" {
      u.Stack = u.Stack[:u.L]
      u.L--
    } else if sline.Label.Recnm == "SEGMENT" && len(sline.Label.Rectp) == 0 {
      if u.L == 0 {
        u.Stack[u.L].Seqno += 1
        u.Stack = append(u.Stack, Keyst_tp{
          u.Tnode.Rname, u.Tnode.Dname, u.Tnode.Dtype, u.Tnode.Dqual,
          u.Tnode.Level, u.Tnode.Pseqn, 0})
        u.L++
      } else if u.Tnode.Level <= u.Stack[u.L].Level {
        for u.Tnode.Level <= u.Stack[u.L].Level {
          u.Out.IsrtStruc(u.Idocn, u.Strtp, u.Stack[u.L-1], u.Stack[u.L])
          u.Stack = u.Stack[:u.L]
          u.L--
        }
        u.Stack[u.L].Seqno += 1
        u.Stack = append(u.Stack, Keyst_tp{
          u.Tnode.Rname, u.Tnode.Dname, u.Tnode.Dtype, u.Tnode.Dqual,
          u.Tnode.Level, u.Tnode.Pseqn, 0})
        u.L++
      } else if u.Tnode.Level > u.Stack[u.L].Level {
        u.Stack[u.L].Seqno += 1
        u.Stack = append(u.Stack, Keyst_tp{
          u.Tnode.Rname, u.Tnode.Dname, u.Tnode.Dtype, u.Tnode.Dqual,
          u.Tnode.Level, u.Tnode.Pseqn, 0})
        u.L++
      }
    } else if sline.Label.Recnm == "FIELDS" && u.L >= 0 {
      u.Fnode.Rname = ""
      u.Fnode.Dname = ""
      u.Fnode.Dqual = ""
    }
    return
  }

  if u.Tnode.Rname == "SEGMENT" && len(u.Tnode.Dname) > 0 {
    if sline.Label.Ident == "SEGMENTTYPE" {
      u.Tnode.Dtype = sline.Value
    }
    if sline.Label.Ident == "QUALIFIED" {
      u.Tnode.Dqual = "QUAL"
    }
    if sline.Label.Ident == "LEVEL" {
      l, _ := strconv.Atoi(sline.Value)
      u.Tnode.Level = l
    }
    return
  }

  if u.Tnode.Rname == "IDOC" {
    if sline.Label.Ident == "EXTENSION" {
      u.Idocn = sline.Value
      u.Out.ClearStruc(u.Idocn, u.Strtp)
    }
    return
  }
}

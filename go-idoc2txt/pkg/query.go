// query.go [2017-05-24 BAR8TL]
// Request and Response to IDOC Query
package rbidoc

import lib "bar8tl/p/rblib"
import "encoding/json"
import "io/ioutil"
import "fmt"
//import "os"
import "path/filepath"
import "strings"

type Rqury_tp struct {
  Field []Field_tp
}

type Query_tp struct {
  Cntrl Rctrl_tp
  Segm  Rsegm_tp
  Qries Rqury_tp
  Token []lib.Qtokn_tp
  Field string
}

func NewQuery() *Query_tp {
  var d Query_tp
  return &d
}

func (d *Query_tp) UpldQuery(parm lib.Param_tp, s Settings_tp) {
  s.SetRunVars(parm)
  fq, _ := ioutil.ReadFile(s.Qrydr + s.Qrynm)
  json.Unmarshal(fq, &d.Qries)
  files, _ := ioutil.ReadDir(s.Inpdr)
  for _, f := range files {
    ffile := f.Name()
    flext := filepath.Ext(ffile)
    flnam := strings.TrimRight(ffile, flext)
    if strings.Index(flnam, "_segment") != -1 {
      d.ProcIndivFile(s)
    }  
  }
}

func (d *Query_tp) ProcIndivFile(s Settings_tp) {
  for i := 0; i < len(d.Qries.Field); i++ {
    tokn := strings.Split(d.Qries.Field[i].Key, "\\")
    if len(tokn) == 1 {
      continue
    }
    if len(tokn) == 2 && tokn[0] == "CONTROL" {
      d.Qries.Field[i].Val = d.QueryControl(tokn[1])
      continue
    }
    for i, t := range tokn {
      if i < len(tokn)-1 {
        c := lib.SplitQueryKey(t)
        fmt.Printf("%v|", c)
        d.Token = append(d.Token, c)
      } else {
        d.Field = tokn[len(tokn)-1]
        fmt.Printf("%s\r\n", d.Field)

        if len(d.Token) == 1 {
          d.Qries.Field[i].Val = d.QuerySegment(d.Token[0], d.Field)
          continue
        }
      }
    }
  }
  fr, _ := json.MarshalIndent(d.Qries, "", " ")
  ioutil.WriteFile("_queries.json", fr, 0644)
}

// Read specific field into Control Record
func (d *Query_tp) QueryControl(key string) (string) {
  fc, _ := ioutil.ReadFile("control.json")
  json.Unmarshal(fc, &d.Cntrl)
  for _, c := range d.Cntrl.Field {
    if c.Key == key {
      return c.Val
    }
  }
  return ""
}

func (d *Query_tp) QuerySegment(sgkey lib.Qtokn_tp, key string) (string) {
  fs, _ := ioutil.ReadFile("segment.json")
  json.Unmarshal(fs, &d.Segm)
  if d.MatchSegmL0(0, sgkey) {
    for _, f := range d.Segm.Child[0].Field {
      if f.Key == key {
        return f.Val
      }
    }
  }
  return ""
}

func (d *Query_tp) MatchSegmL0(l0 int, sgkey lib.Qtokn_tp) (bool) {
  if d.Segm.Child[l0].Segmn == sgkey.Segmn {
    if sgkey.Instn != 0 && d.Segm.Child[l0].Instn == sgkey.Instn {
      if sgkey.Qlkey != "" && d.Segm.Child[l0].Qlkey == sgkey.Qlkey {
        if sgkey.Qlval != "" && d.Segm.Child[l0].Qlval == sgkey.Qlval {
          return true
        }
      }
    } else {
      if sgkey.Qlkey != "" && d.Segm.Child[l0].Qlkey == sgkey.Qlkey {
        if sgkey.Qlval != "" && d.Segm.Child[l0].Qlval == sgkey.Qlval {
          return true
        }
      } else {
        return true
      }
    }
  } else {
    return true
  }
  return false
}
/*
func (d *Query_tp) MatchSegmL1(l0 int, sgkey lib.Qtokn_tp) (bool) {
  MatchSegmL0(0, sgkey)

  if d.Segm.Child[l0].Segmn == sgkey.Segmn {
    if sgkey.Instn != 0 && d.Segm.Child[l0].Child[l1].Instn == sgkey.Instn {
      if sgkey.Qlkey != "" && d.Segm.Child[l0].Child[l1].Qlkey == sgkey.Qlkey {
        if sgkey.Qlval != "" && d.Segm.Child[l0].Child[l1].Qlval == sgkey.Qlval {
          return true
        }
      }
    } else {
      if sgkey.Qlkey != "" && d.Segm.Child[l0].Child[l1].Qlkey == sgkey.Qlkey {
        if sgkey.Qlval != "" && d.Segm.Child[l0].Child[l1].Qlval == sgkey.Qlval {
          return true
        }
      } else {
        return true
      }
    }
  } else {
    return true
  }
  return false
}
*/
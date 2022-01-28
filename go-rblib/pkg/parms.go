// parms.go [2020-01-22 BAR8TL]
// Gets a list of command-line parameters
package rblib

import "errors"
import "os"
import "strings"

type Param_tp struct {
  Optn string
  Prm1 string
  Prm2 string
}

type Parms_tp struct {
  Cmdpr []Param_tp
  Messg string
}

func NewParms() Parms_tp {
  var p Parms_tp
  return p
}

func (p *Parms_tp) NewParms() ([]Param_tp, error) {
  if len(os.Args) == 1 {
    p.Messg = "Run option missing"
    return nil, errors.New(p.Messg)
  }
  for _, curarg := range os.Args {
    if curarg[0:1] == "-" || curarg[0:1] == "/" {
      optn := strings.ToLower(curarg[1:len(curarg)])
      prm1 := ""
      prm2 := ""
      if optn != "" {
        if strings.Index(optn, ":") != -1 {
          prm1 = optn[strings.Index(optn, ":")+1 : len(optn)]
          optn = strings.TrimSpace(optn[0:strings.Index(optn, ":")])
          if strings.Index(prm1, ":") != -1 {
            prm2 = strings.TrimSpace(prm1[strings.Index(prm1, ":")+1:len(prm1)])
            prm1 = strings.TrimSpace(prm1[0:strings.Index(prm1, ":")])
          }
        }
        p.Cmdpr = append(p.Cmdpr, Param_tp{optn, prm1, prm2})
      } else {
        p.Messg = "Run option missing"
        return nil, errors.New(p.Messg)
      }
    }
  }
  return p.Cmdpr, nil
}

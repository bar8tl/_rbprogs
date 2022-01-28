// folio.go [2020-01-22 BAR8TL]
// Return a current counter number
package main

import "encoding/json"
import "errors"
import "fmt"
import "github.com/atotto/clipboard"
import "io/ioutil"
import "os"
import "strings"
import "time"

const FPATH = "\\\\bosch.com\\dfsrb\\DfsUS\\loc\\Mx\\ILM\\Projects\\CFA\\" +
  "Edifcack\\db\\folios.json"

func main() {
  if len(os.Args) == 1 {
    fmt.Printf("Missing Counter ID.\r\n")
    time.Sleep(1 * time.Second)
    return
  }
  if len(os.Args) > 1 {
    clipboard.WriteAll("")
    f := NewCounters()
    err := f.ProcCounter(FPATH, strings.ToUpper(os.Args[1]))
    if err != nil {
      fmt.Println(err)
      time.Sleep(1 * time.Second)
      return
    }
    fmt.Printf("Counter assigned (Copied to clipboard): %s\r\n", f.Fcont)
    time.Sleep(2 * time.Second)
    clipboard.WriteAll(f.Fcont)
  }
}

// counters.go [2020-01-22/BAR8TL]
// Encapsulated data type to keep current counters of objects/tasks in a project
type Cline_tp struct {
  Id    string `json:"ID"`
  Prjct string `json:"Project"`
  Formt string `json:"Format"`
  Step  int    `json:"Step"`
  Count int    `json:"Counter"`
}

type Clist_tp struct {
  Contr []Cline_tp `json:"Counters"`
}

type Counters_tp struct {
  Clist Clist_tp
  Index int
  Ccntr Cline_tp
  Fcont string
}

func NewCounters() *Counters_tp {
  var c Counters_tp
  return &c
}

func (c *Counters_tp) ProcCounter(fpath, id string) error {
  var err error
  err = c.GetCounter(fpath, id)
  if err != nil {
    return err
  }
  c.StepCounter()
  err = c.PutCounter(fpath)
  if err != nil {
    return err
  }
  return nil
}

func (c *Counters_tp) GetCounter(fpath, id string) error {
  fl, err := ioutil.ReadFile(fpath)
  if err != nil {
    return err
  }
  err = json.Unmarshal(fl, &c.Clist)
  if err != nil {
    return err
  }
  for i := 0; i < len(c.Clist.Contr); i++ {
    if c.Clist.Contr[i].Id == id {
      c.Ccntr = c.Clist.Contr[i]
      if c.Ccntr.Step == 0 {
        c.Ccntr.Step = 1
      }
      c.Index = i
      return nil
    }
  }
  return errors.New("Counter ID " + id + " not valid.")
}

func (c *Counters_tp) StepCounter() {
  c.Clist.Contr[c.Index].Count = c.Clist.Contr[c.Index].Count + c.Ccntr.Step
  c.Fcont = fmt.Sprintf(c.Ccntr.Formt, c.Clist.Contr[c.Index].Count)
}

func (c *Counters_tp) PutCounter(fpath string) error {
  fl, err := json.MarshalIndent(c.Clist, "", " ")
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(fpath, fl, 0644)
  if err != nil {
    return err
  }
  return nil
}

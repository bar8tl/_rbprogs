// utils.go [2019-11-19 BAR8TL]
// Utility functions
package rblib

import "archive/zip"
import "bytes"
import "fmt"
import "io"
import "log"
import "math"
import "os"
import "strconv"
import "strings"

// Function to simulate C ternary operator
func Ternary_op(statement bool, tcond, fcond string) string {
  if statement {
    return tcond
  }
  return fcond
}

// Function to identify individual tokens in SAP IDoc parser file
type Reclb_tp struct {
  Ident string
  Recnm string
  Rectp string
}

type Parsl_tp struct {
  Label Reclb_tp
  Value string
}

func ScanTextIdocLine(s string) (p Parsl_tp) {
  var key string
  var val string
  flds := strings.Fields(s)
  if len(flds) > 0 {
    key = flds[0]
    if (len(key) >= 6 && key[0:6] == "BEGIN_") ||
      (len(key) >= 4 && key[0:4] == "END_") {
      tokn := strings.Split(key, "_")
      if len(tokn) == 2 {
        p.Label.Ident = tokn[0]
        p.Label.Recnm = tokn[1]
        p.Label.Rectp = ""
      } else if len(tokn) == 3 {
        p.Label.Ident = tokn[0]
        p.Label.Recnm = tokn[1]
        p.Label.Rectp = tokn[2]
      }
    } else {
      p.Label.Ident = key
      p.Label.Recnm = ""
      p.Label.Rectp = ""
    }
  }
  if len(flds) > 1 {
    val = flds[1]
    for i := 2; i < len(flds); i++ {
      val += " " + flds[i]
    }
    p.Value = val
  }
  return p
}

// Function to identify individual tokens in IDoc query key
type Qtokn_tp struct {
  Segmn string
  Instn int
  Qlkey string
  Qlval string
}

func SplitQueryKey(key string) (q Qtokn_tp) {
  atokn := strings.SplitN(key, "[", 2)
  if len(atokn) == 2 {
    q.Segmn = atokn[0]
    btokn := strings.SplitN(atokn[1], "]", 2)
    if len(btokn) == 2 {
      q.Instn, _ = strconv.Atoi(btokn[0])
      ctokn := strings.SplitN(btokn[1], ".", 2)
      if len(ctokn) == 2 {
        q.Segmn = ctokn[0]
        dtokn := strings.SplitN(ctokn[1], ":", 2)
        if len(dtokn) == 2 {
          q.Qlkey = dtokn[0]
          q.Qlval = dtokn[1]
        }
      }
    }
  } else {
    btokn := strings.SplitN(key, ".", 2)
    if len(btokn) == 2 {
      q.Segmn = btokn[0]
      ctokn := strings.SplitN(btokn[1], ":", 2)
      if len(ctokn) == 2 {
        q.Qlkey = ctokn[0]
        q.Qlval = ctokn[1]
      }
    } else {
      q.Segmn = key
    }
  }
  return q
}

// Funtions to create/display ZIP compressed data files
func Zipgen(ifnam, ofnam string) {
  outf, err := os.Create(ofnam)
  if err != nil {
    log.Fatalf("Create: %v\n", err)
  }
  w := zip.NewWriter(outf)
  f, err := w.Create(ifnam)
  if err != nil {
    log.Fatal(err)
  }
  inf, err := os.Open(ifnam)
  if err != nil {
    log.Fatalf("Open: %v\n", err)
  }
  fs, _ := inf.Stat()
  ibuf := make([]byte, fs.Size())
  _, err = inf.Read(ibuf)
  if err != nil {
    log.Fatal(err)
  }
  inf.Close()
  _, err = f.Write(ibuf)
  if err != nil {
    log.Fatal(err)
  }
  err = w.Close()
  if err != nil {
    log.Fatal(err)
  }
}

func Zipdsp(fname string) {
  rc, err := zip.OpenReader(fname)
  if err != nil {
    log.Fatal(err)
  }
  defer rc.Close()
  for _, f := range rc.File {
    d, err := f.Open()
    if err != nil {
      log.Fatal(err)
    }
    defer d.Close()
    buf := new(bytes.Buffer)
    buf.ReadFrom(d)
    for iline, err := buf.ReadString(byte('\n')); err != io.EOF; iline,
      err = buf.ReadString(byte('\n')) {
      fmt.Print(iline)
    }
  }
}

// Functions to perform rounding over numbers type float64
func Round(n float64, d int) float64 {
  sg := 1.0
  if n < 0 {
    sg = -1.0
  }
  soutc := "%." + strconv.Itoa(d) + "f"
  sintr := fmt.Sprintf(soutc, math.Trunc((n + sg * (0.5 * math.Pow10(-d))) *
    math.Pow10(d)) * math.Pow10(-d))
  fintr, _ := strconv.ParseFloat(sintr, 64)
  return fintr
}

func Roundup(n float64, d int) float64 {
  sg := 1.0
  if n < 0 {
    sg = -1.0
  }
  soutc := "%." + strconv.Itoa(d) + "f"
  sintr := fmt.Sprintf(soutc, math.Trunc((n + sg * (0.9 * math.Pow10(-d))) *
    math.Pow10(d)) * math.Pow10(-d))
  fintr, _ := strconv.ParseFloat(sintr, 64)
  return fintr
}

func Truncate(n float64, d int) float64 {
  soutc := "%." + strconv.Itoa(d) + "f"
  sintr := fmt.Sprintf(soutc, math.Trunc(n * math.Pow10(d)) * math.Pow10(-d))
  fintr, _ := strconv.ParseFloat(sintr, 64)
  return fintr
}

func Ffloor(n, r float64, d int) float64 {
  return Truncate((n - (math.Pow10(-d) / 2)) * r, d)
}

func Fceil(n, r float64, d int) float64 {
  return Roundup(((n + (math.Pow10(-d)) / 2) - (math.Pow10(-12))) * r, d)
}

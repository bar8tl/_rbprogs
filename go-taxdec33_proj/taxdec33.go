package main

import rb "bar8tl/p/taxdec33"
import "bufio"
import "fmt"
import "io"
import "log"
import "os"

var t *rb.It_conc
var g *rb.It_invo
var m *rb.It_matr

func main() {
  t = rb.New_It_conc()
  g = rb.New_It_invo(os.Args[1])
  m = rb.New_It_matr()
	cargaTablaDeConceptos()
  t.EvaluayEjecutaCorreccion(g, m)
  descargaResultado()
}

func cargaTablaDeConceptos() {
	ifile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer ifile.Close()
	rdr := bufio.NewReader(ifile)
	for line, err := rdr.ReadString(byte('\n')); err != io.EOF; line, 
    err = rdr.ReadString(byte('\n')) {
    t.AdicionaEntradaATabla(g, line)
  }
}

func descargaResultado() {
  t.MuestraTabla(g)
  ofile, _ := os.Create(os.Args[3])
  defer ofile.Close()
  w := bufio.NewWriter(ofile)
  for i, _ := range t.Conc {
    line := t.ObtieneEntradaDeTabla(i, g)
    fmt.Fprintf(w, line)
  }
  w.Flush()
}

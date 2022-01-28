// replxls.go [2012-11-08 BAR8TL]
// Perform the process to build and to include Excel attachments (base64 coded) 
// within Mex XML Invoices
package rb

import "bar8tl/p/xmlx"
import "bufio"
import "encoding/base64"
import "io"
import "os"
import "path"
import "strings"

func ReplXls(xmlnam, xlsnam, newxml string) error {
  docm := xmlx.New()
  docm.LoadFile(xmlnam, nil)
  node := docm.SelectNode("", "cargosCreditos")
  node.SetAttr("archivo", encodeXls(&xlsnam))
  docm.SaveFile(path.Dir(newxml)+"\\~intern.xml")
  dlteAddedElem(path.Dir(newxml)+"\\~intern.xml", newxml)

  var err error
  docm.LoadFile(newxml, nil)
  node = docm.SelectNode("", "cargosCreditos")
  for i,_ := range node.Attributes {
    if node.Attributes[i].Name.Local == "archivo" {
      err = decodeXls(newxml, &node.Attributes[i].Value)
      break
    }
  }
  return err
}

func encodeXls(fname *string) string {
  ifile, _ := os.Open(*fname); defer ifile.Close()
  bstat, _ := ifile.Stat()
  var bsize int64 = bstat.Size()
  bbyte := make([]byte, bsize)
  bbuff := bufio.NewReader(ifile)
  bbuff.Read(bbyte)
  return base64.StdEncoding.EncodeToString(bbyte)
}

func decodeXls(newxml string, datos *string) error {
  oftxt, _ := os.Create(path.Dir(newxml)+"\\~intern.txt")
  defer oftxt.Close()
  oftxt.WriteString(*datos)
  ofpdf, _ := os.Create(path.Dir(newxml)+"\\~intern.xlsx")
  defer ofpdf.Close()
  bbuff := bufio.NewWriter(ofpdf)
  bbyte, err := base64.StdEncoding.DecodeString(*datos)
  if err != nil {
    return err
  }
  bbuff.Write(bbyte)
  return nil
}

func dlteAddedElem(ifile, ofile string) {
  var word  [256]byte
  var sword []byte
  sword = make([]byte, 256)
  inword, removed := false, false
  i := 0
  ifxml, _ := os.Open  (ifile)
  defer ifxml.Close()
  ibuff := bufio.NewReader(ifxml)
  ofxml, _ := os.Create(ofile)
  defer ofxml.Close()
  obuff := bufio.NewWriter(ofxml)
  for c, err := ibuff.ReadByte(); err != io.EOF; c, err = ibuff.ReadByte() {
    if removed {
      obuff.WriteByte(c)
    } else {
      if c == ' ' || c == '\n' || c == '\t' || c == '<' || c == '>' ||
         c == '?' {
        if inword {
          copy(sword[:], word[0:i])
          if strings.Contains(string(sword), "standalone=") {
            removed = true
          } else {
            for j := 0; j < i; j++ {
              obuff.WriteByte(word[j])
            }
          }
        }
        obuff.WriteByte(c)
        i = 0
      } else {
        word[i] = c
        i++
        inword = true
      }
    }
  }
  obuff.Flush()
}
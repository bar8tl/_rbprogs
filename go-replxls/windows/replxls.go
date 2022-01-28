// replxls.go [2012-11-08 BAR8TL]
// Starts user dialog box to build and include Excel attachments base64
// encoded within Mex XML e-invoices
package main

import rb "bar8tl/p/replxls"
import "github.com/lxn/walk"
import . "github.com/lxn/walk/declarative"
import "log"
import "regexp"

type MyMainWindow struct {
  *walk.MainWindow
  xmlnam, xlsnam *walk.LineEdit
}

func main() {
  mw := new(MyMainWindow)
  if err := ProcMainDialog(*mw); err != nil {
    log.Fatal(err)
  }
}

func (mw *MyMainWindow) replaceAttachment() {
  if len(mw.xmlnam.Text()) == 0 {
    walk.MsgBox(mw, "Error",
      "File 'Initial invoice' not specified.\nSpecify a valid XML file.",
      walk.MsgBoxIconError)
    return
  }
  if len(mw.xlsnam.Text()) == 0 {
    walk.MsgBox(mw, "Error",
      "File 'XLSX file' not specified.\nSpecify a valid XLSX file.",
      walk.MsgBoxIconError)
    return
  }
  re := regexp.MustCompile("[.]([^.]+)$")
  newxml := re.ReplaceAllString(mw.xmlnam.Text(), "_new.$1")
  if err := rb.ReplXls(mw.xmlnam.Text(), mw.xlsnam.Text(), newxml); err != nil {
    walk.MsgBox(mw, "Error",
      "One error ocurred during XLSX base64 production.",
      walk.MsgBoxIconError)
  } else {
    walk.MsgBox(mw, "Success",
      "XLSX portion has been replaced successfully in XML invoice.\n" +
      "Find output in file " + newxml,
      walk.MsgBoxIconInformation)
  }
}

func ProcMainDialog(mw MyMainWindow) (err error) {
   _, err = (MainWindow{
    AssignTo: &mw.MainWindow,
    Title: "Replace XLSX File in FCA eInvoices",
    MinSize: Size{650, 100},
    Layout: VBox{},
    Children: [] Widget{
      Composite{
        Layout: Grid{Columns: 3},
        Children: [] Widget{
          Label{
            Text: "Initial Invoice:",
          },
          LineEdit{
            AssignTo: &(mw.xmlnam),
            ReadOnly: true,
          },
          PushButton{
            Text: "Browse...",
            OnClicked: mw.openInvoice_Triggered,
          },
          Label{
            Text: "XLSX File:",
          },
          LineEdit{
            AssignTo: &(mw.xlsnam),
            ReadOnly: true,
          },
          PushButton{
            Text: "Browse...",
            OnClicked: mw.openAttachment_Triggered,
          },
        },
      },
      Composite{
        Layout: HBox{},
        Children: [] Widget{
          HSpacer{},
          PushButton{
            Text: "Close",
            OnClicked: func() { mw.Close() },
          },
          PushButton{
            Text: "Replace",
            OnClicked: mw.replaceAttachment,
          },
        },
      },
    },
  }.Run())
  return
}

func (mw *MyMainWindow) openInvoice_Triggered() {
  if err := mw.openXml(); err != nil {
    log.Print(err)
  }
}

func (mw *MyMainWindow) openXml() error {
  dlg := new(walk.FileDialog)
  dlg.Filter = "XML Files (*.xml)|*.xml"
  dlg.Title = "Select an Invoice"
  if ok, err := dlg.ShowOpen(mw); err != nil {
    return err
  } else if !ok {
    return nil
  }
  mw.xmlnam.SetText(dlg.FilePath)
  return nil
}

func (mw *MyMainWindow) openAttachment_Triggered() {
  if err := mw.openXls(); err != nil {
    log.Print(err)
  }
}

func (mw *MyMainWindow) openXls() error {
  dlg := new(walk.FileDialog)
  dlg.Filter = "XLSX Files (*.xlsx)|*.xlsx|XLS Files (*.xls)|*.xls"
  dlg.Title = "Select an XLSX file"
  if ok, err := dlg.ShowOpen(mw); err != nil {
    return err
  } else if !ok {
    return nil
  }
  mw.xlsnam.SetText(dlg.FilePath)
  return nil
}

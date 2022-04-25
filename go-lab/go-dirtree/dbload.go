package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "io"
import "encoding/csv"
import "log"
import "os"

func main() {
  cdb := NewCrtdb()
  if os.Args[1] == "-c" {
    cdb.CrtTable(os.Args[2])
  } else if os.Args[1] == "-u" {
    cdb.UplData(os.Args[2], os.Args[3])
  } else {
    log.Printf("Run option invalid.\r\n")
  }
}

type Dbo_tp struct {
}

func NewCrtdb() *Dbo_tp {
  var d Dbo_tp
  return &d
}

// Create DB table
func (d *Dbo_tp) CrtTable(drnam string) {
	os.Remove(drnam + "\\" + "dirtree.db")
	file, err := os.Create(drnam + "\\" + "dirtree.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Creating Sqlite database dirtree.db...")
	file.Close()
	log.Println("Sqlite database dirtree.db created")
  db, err := sql.Open("sqlite3", drnam + "\\" + "dirtree.db")
	defer db.Close()
  if err != nil {
    log.Fatalf("Error opening DB: %v\n", err)
  }
  sqlst, err := db.Prepare(`DROP TABLE IF EXISTS dirtree`)
	sqlst.Exec()
	sqlst, err  = db.Prepare(`CREATE TABLE IF NOT EXISTS dirtree (ldate TEXT,
    mtype TEXT, dire0 TEXT, dire1 TEXT, dire2 TEXT, dire3 TEXT, dire4 TEXT,
    dire5 TEXT, dire6 TEXT, dire7 TEXT, dire8 TEXT, dire9 TEXT, filen TEXT,
    msize TEXT);`)
	sqlst.Exec()
	log.Println("Table created.")
}

// Upload data to DB table
func (d *Dbo_tp) UplData(drnam, icsvfl string) {
  ifile, err := os.Open(drnam + "\\" + icsvfl)
  if err != nil {
    log.Fatal(err.Error())
  }
  defer ifile.Close()
  rdr := csv.NewReader(ifile)
  db, err := sql.Open("sqlite3", drnam + "\\" + "dirtree.db")
	defer db.Close()
  if err != nil {
    log.Fatalf("Error opening DB: %v\n", err)
  }
  log.Println("Inserting dirtree record ...")
  for {
    rec, err := rdr.Read()
    if err == io.EOF {
      break
    }
    mtype  := rec[0]
    dire0  := rec[1]
    dire1  := rec[2]
    dire2  := rec[3]
    dire3  := rec[4]
    dire4  := rec[5]
    dire5  := rec[6]
    dire6  := rec[7]
    dire7  := rec[8]
    dire8  := rec[9]
    dire9  := rec[10]
    dire10 := rec[11]
    dire11 := rec[12]
    sqlst, err := db.Prepare(`INSERT INTO dirtree(ldate, mtype, dire0, dire1,
      dire2, dire3, dire4, dire5, dire6, dire7, dire8, dire9, filen, msize)
      VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?);`)
    if err != nil {
      log.Fatal(err.Error())
    }
  	sqlst.Exec("20220204", mtype, dire0, dire1, dire2, dire3, dire4, dire5,
      dire6, dire7, dire8, dire9, dire10, dire11)
  }
}

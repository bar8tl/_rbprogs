// upldmitm.go [2017-05-24 BAR8TL]
// Get IDoc item data (records, groups, segments and fields) and to create
// corresponding item records in a reference database
package rbidoc

import lib "bar8tl/p/rblib"
import "strconv"

type Idcdf_tp struct {
  Name string
  Type string
  Cols [2]string // Name, Extn
}

type Grpdf_tp struct {
  Name string
  Type string
  Seqn int
  Cols [5]string // Numb, Levl, Stat, Mnlp, Mxlp
}

type Segdf_tp struct {
  Name string
  Type string
  Seqn int
  Cols [7]string // Name, Type, Qual, Levl, Stat, Mnlp, Mxlp
}

type Flddf_tp struct {
  Name string
  Type string
  Clas string
  Cols [7]string // Name, Text, Type, Lgth, Seqn, Strp, Endp
}

type Upldmitm_tp struct {
  Icol  []string
  Gcol  []string
  Scol  []string
  Fcol  []string
  Stack []lib.Parsl_tp // List of Parsl_tp: Levels stack
  Lidoc []Idcdf_tp // List of Idcdf_tp: Idoc
  Lgrup []Grpdf_tp // List of Grpdf_tp: Grup
  Lsegm []Segdf_tp // List of Segdf_tp: Segm
  Lfild []Flddf_tp // List of Flddf_tp: Fild
  Lrecd []Flddf_tp // List of Flddf_tp: Fild
  Colsi [2]string  // Name, Extn
  Colsg [5]string  // Numb, Levl, Stat, Mnlp, Mxlp
  Colss [7]string  // Name, Type, Qual, Levl, Stat, Mnlp, Mxlp
  Colsf [7]string  // Name, Text, Type, Lgth, Seqn, Strp, Endp
  Colsr [7]string  // Name, Text, Type, Lgth, Seqn, Strp, Endp
  Out   Outsqlt_tp
  L     int        // Stack level
  Gseqn int        // Group counter
  Sseqn int        // Segment counter
}

func (u *Upldmitm_tp) NewUpldmitm(s Settings_tp) {
  u.Out.NewOutsqlt(s)
  u.Icol = []string{"EXTENSION"}
  u.Gcol = []string{"LEVEL", "STATUS", "LOOPMIN", "LOOPMAX"}
  u.Scol = []string{"SEGMENTTYPE", "QUALIFIED", "LEVEL", "STATUS", "LOOPMIN",
    "LOOPMAX"}
  u.Fcol = []string{"NAME", "TEXT", "TYPE", "LENGTH", "FIELD_POS",
    "CHARACTER_FIRST", "CHARACTER_LAST"}
  u.L = -1
}

// Scan SAP parser file to identify IDoc elements
func (u *Upldmitm_tp) GetData(sline lib.Parsl_tp) {
  if sline.Label.Ident == "BEGIN" {
    u.L++
    u.Stack = append(u.Stack, lib.Parsl_tp{
      lib.Reclb_tp{sline.Label.Ident, sline.Label.Recnm, sline.Label.Rectp},
      sline.Value})
    if sline.Value != "" {
      if sline.Label.Recnm == "IDOC" {
        u.Colsi[0] = sline.Value
        u.Colsi[1] = sline.Value
        u.Lidoc = append(u.Lidoc, Idcdf_tp{
          u.Colsi[0], u.Stack[u.L].Label.Recnm, u.Colsi})
      } else if sline.Label.Recnm == "GROUP" {
        u.Colsg[0] = sline.Value
      } else if sline.Label.Recnm == "SEGMENT" {
        u.Colss[0] = sline.Value
        u.Colss[2] = ""
      }
    }
    return
  }

  if sline.Label.Ident == "END" {
    u.L--
    u.Stack = u.Stack[:u.L+1]
    return
  }

  if u.Stack[u.L].Label.Recnm == "IDOC" {
    for i := 0; i < len(u.Icol); i++ {
      if sline.Label.Ident == u.Icol[i] {
        u.Colsi[i+1] = sline.Value
        if i == (len(u.Icol) - 1) {
          u.Lidoc[0].Cols[1] = u.Colsi[i+1]
        }
        break
      }
    }
  }

  if u.Stack[u.L].Label.Recnm == "GROUP" {
    for i := 0; i < len(u.Gcol); i++ {
      if sline.Label.Ident == u.Gcol[i] {

        u.Colsg[i+1] = sline.Value
        if i == (len(u.Gcol) - 1) {
          u.Gseqn++
          u.Lgrup = append(u.Lgrup, Grpdf_tp{
            u.Colsg[0], u.Stack[u.L].Label.Recnm, u.Gseqn, u.Colsg})
        }
        break
      }
    }
  }

  if u.Stack[u.L].Label.Recnm == "SEGMENT" {
    for i := 0; i < len(u.Scol); i++ {
      if sline.Label.Ident == u.Scol[i] {
        if sline.Label.Ident == "QUALIFIED" {
          u.Colss[i+1] = "QUAL"
        } else {
          u.Colss[i+1] = sline.Value
        }
        if i == (len(u.Scol) - 1) {
          u.Sseqn++
          u.Lsegm = append(u.Lsegm, Segdf_tp{
            u.Colss[0], u.Stack[u.L].Label.Recnm, u.Sseqn, u.Colss})
        }
        break
      }
    }
  }

  if u.Stack[u.L].Label.Recnm == "FIELDS" {
    match := false
    for i := 0; i < len(u.Fcol) && !match; i++ {
      if sline.Label.Ident == u.Fcol[i] {
        u.Colsf[i] = sline.Value
        match = true
      }
      if i == (len(u.Fcol) - 1) {
        if u.Stack[u.L-1].Label.Rectp == "RECORD" {
          u.Lrecd = append(u.Lrecd, Flddf_tp{
            u.Stack[u.L-1].Label.Recnm, u.Stack[u.L].Label.Recnm,
            u.Stack[u.L-1].Label.Rectp, u.Colsf})
        } else if u.Stack[u.L-1].Label.Recnm == "SEGMENT" {
          u.Lfild = append(u.Lfild, Flddf_tp{
            u.Colss[0], u.Stack[u.L].Label.Recnm, u.Stack[u.L-1].Label.Recnm,
            u.Colsf})
        }
      }
    }
  }
}

type Items_tp struct { // ITEMS fields description (*=key field in DB record)
//Field:       //  IDOC        GROUP       SEGMENT     SGM-FIELD   RECRD-FIELD
//-----------------------------------------------------------------------------
  Idocn string //* Ex/Ba-Name  Ex/Ba-Name  Ex/Ba-Name  Ex/Ba-Name  Ex/Ba-Name
  Rname string //* 'IDOC'      'GROUP'     'SEGMENT'   Segm-ID     'CONTROL'...
  Dname string //* Basic-IDoc  Group#      Segm-ID     Field-Name  Field-Name
  Rclas string //  Basic-IDoc  Group#      Segm-ID     'SEGMENT'   'RECORD'
  Rtype string //  'IDOC'      'GROUP'     'SEGMENT'   'FIELDS'    'FIELDS'
  Dtype string //  ''          ''          Segm-Type    Data-Type   Data-Type
  Dtext string //  Extsn-name  Group#      Qualified   Field-Desc  Field-Desc
  Level int    //  0           Level       Level       0           0
  Stats string //  ''          Status      Status      ''          ''
  Minlp int    //  0           Loop-Min    Loop-Min    0           0
  Maxlp int    //  0           Loop-Max    Loop-Max    0           0
  Lngth int    //  0           0           0           Length      Length
  Seqno int    //  0           auto-gen    Auto-gen    Field-Seqn  Field-Seqn
  Strps int    //  0           0           0           Start-Pos   Start-Pos
  Endps int    //  0           0           0           End-Pos     End-Pos
}

// Functions to upload IDoc data elements into a reference definition database
func (u *Upldmitm_tp) IsrtData(s Settings_tp) {
  u.Out.ClearItems(u.Lidoc[0].Cols[1])
  u.UplRecd()
  u.UplIdoc()
  u.UplGrup()
  u.UplSegm()
  u.UplFlds()
}

// /RB04/YP3_DELVRY_RBNA|IDOC|DELVRY07|DELVRY07|IDOC|||/RB04/YP3_DELVRY_RBNA|0|
// 0||0|0|0|0|0|0
func (u *Upldmitm_tp) UplIdoc() *Upldmitm_tp { // Upload IDoc idoc data
  var w Items_tp
  for _, lidoc := range u.Lidoc {
    w.Idocn = u.Lidoc[0].Cols[1] // EXTENSION/BASIC /RB04/YP3_DELVRY_RBNA
    w.Rname = lidoc.Type         // B…_IDOC         IDOC
    w.Dname = lidoc.Cols[0]      // BEGIN_IDOC      DELVRY07
    w.Rclas = lidoc.Name         // BEGIN_IDOC      DELVRY07
    w.Rtype = lidoc.Type         // B…_IDOC         IDOC
    w.Dtype = ""
    w.Dtext = lidoc.Cols[1]      // EXTENSION       /RB04/YP3_DELVRY_RBNA
    w.Level = 0
    w.Stats = ""
    w.Minlp = 0
    w.Maxlp = 0
    w.Lngth = 0
    w.Seqno = 0
    w.Strps = 0
    w.Endps = 0
    u.Out.IsrtItems(w)
  }
  return u
}

// /RB04/YP3_DELVRY_RBNA|GROUP|1|1|GROUP||||1|2|MANDATORY|1|9999|0|0|0|0
func (u *Upldmitm_tp) UplGrup() *Upldmitm_tp { // Upload IDoc groups data
  var w Items_tp
  for _, lgrup := range u.Lgrup {
    u.Gseqn++
    w.Idocn = u.Lidoc[0].Cols[1] // EXTENSION/BASIC /RB04/YP3_DELVRY_RBNA
    w.Rname = lgrup.Type         // B…_GROUP        GROUP
    w.Dname = lgrup.Cols[0]      // BEGIN_GROUP     1
    w.Rclas = lgrup.Name         // BEGIN_GROUP     1
    w.Rtype = lgrup.Type         // B…_GROUP        GROUP
    w.Dtype = ""
    w.Dtext = lgrup.Cols[0]      // BEGIN_GROUP     1
    w.Level, _ = strconv.Atoi(lgrup.Cols[1]) // LEVEL           02
    w.Stats = lgrup.Cols[2]                  // STATUS          MANDATORY
    w.Minlp, _ = strconv.Atoi(lgrup.Cols[3]) // LOOPMIN         0000000001
    w.Maxlp, _ = strconv.Atoi(lgrup.Cols[4]) // LOOPMAX         0000009999
    w.Lngth = 0
    w.Seqno = lgrup.Seqn
    w.Strps = 0
    w.Endps = 0
    u.Out.IsrtItems(w)
  }
  return u
}

// /RB04/YP3_DELVRY_RBNA|SEGMENT|E2EDL20004|E2EDL20004|SEGMENT|E1EDL20|QUAL||0|
// 2|MANDATORY|1|1|0|0|0|0
func (u *Upldmitm_tp) UplSegm() *Upldmitm_tp { // Upload IDoc segments data
  var w Items_tp
  for _, lsegm := range u.Lsegm {
    u.Sseqn++
    w.Idocn = u.Lidoc[0].Cols[1] // EXTENSION/BASIC /RB04/YP3_DELVRY_RBNA
    w.Rname = lsegm.Type         // B…_SEGMENT      SEGMENT
    w.Dname = lsegm.Cols[0]      // BEGIN_SEGMENT   E2EDL20004
    w.Rclas = lsegm.Name         // BEGIN_SEGMENT   E2EDL20004
    w.Rtype = lsegm.Type         // B…_SEGMENT      SEGMENT
    w.Dtype = lsegm.Cols[1]      // SEGMENTTYPE     E1EDL20
    w.Dtext = lsegm.Cols[2]      // QUALIFIED       QUAL
    w.Level, _ = strconv.Atoi(lsegm.Cols[3]) // LEVEL           02
    w.Stats = lsegm.Cols[4]                  // STATUS          MANDATORY
    w.Minlp, _ = strconv.Atoi(lsegm.Cols[5]) // LOOPMIN         0000000001
    w.Maxlp, _ = strconv.Atoi(lsegm.Cols[6]) // LOOPMAX         0000000001
    w.Lngth = 0
    w.Seqno = lsegm.Seqn
    w.Strps = 0
    w.Endps = 0
    u.Out.IsrtItems(w)
  }
  return u
}

// /RB04/YP3_DELVRY_RBNA|E2EDL20004|VKBUR|SEGMENT|FIELDS|CHARACTER|Sales Office|
// |0|0||0|0|4|5|84|87
func (u *Upldmitm_tp) UplFlds() *Upldmitm_tp { // Upload IDoc fields data
  var w Items_tp
  for _, lfild := range u.Lfild {
    w.Idocn = u.Lidoc[0].Cols[1] // EXTENSION/BASIC /RB04/YP3_DELVRY_RBNA
    w.Rname = lfild.Name         // BEGIN_SEGMENT   E2EDL20004
    w.Dname = lfild.Cols[0]      // NAME            VKBUR
    w.Rclas = lfild.Clas         // B…_SEGMENT      SEGMENT
    w.Rtype = lfild.Type         // B…_FIELDS       FIELDS
    w.Dtype = lfild.Cols[2]      // TYPE            CHARACTER
    w.Dtext = lfild.Cols[1]      // TEXT            Sales Office
    w.Level = 0
    w.Stats = ""
    w.Minlp = 0
    w.Maxlp = 0
    w.Lngth, _ = strconv.Atoi(lfild.Cols[3]) // LENGTH          000004
    w.Seqno, _ = strconv.Atoi(lfild.Cols[4]) // FIELD_POS       0005
    w.Strps, _ = strconv.Atoi(lfild.Cols[5]) // CHARACTER_FIRST 000084
    w.Endps, _ = strconv.Atoi(lfild.Cols[6]) // CHARACTER_LAST  000087
    u.Out.IsrtItems(w)
  }
  return u
}

// /RB04/YP3_DELVRY_RBNA|CONTROL|TABNAM|RECORD|FIELDS|CHARACTER|
// Name of Table Structure||0|0||0|0|10|1|1|10
func (u *Upldmitm_tp) UplRecd() *Upldmitm_tp { // Upload IDoc records data
  var w Items_tp
  for _, lrecd := range u.Lrecd {
    w.Idocn = u.Lidoc[0].Cols[1] // EXTENSION/BASIC /RB04/YP3_DELVRY_RBNA
    w.Rname = lrecd.Name         // B…_CONTROL_R…   CONTROL
    w.Dname = lrecd.Cols[0]      // NAME            TABNAM
    w.Rclas = lrecd.Clas         // B…_C…_RECORD    RECORD
    w.Rtype = lrecd.Type         // B…_FIELDS       FIELDS
    w.Dtype = lrecd.Cols[2]      // TYPE            CHARACTER
    w.Dtext = lrecd.Cols[1]      // TEXT            Name of Table Stru...
    w.Level = 0
    w.Stats = ""
    w.Minlp = 0
    w.Maxlp = 0
    w.Lngth, _ = strconv.Atoi(lrecd.Cols[3]) // LENGTH          000010
    w.Seqno, _ = strconv.Atoi(lrecd.Cols[4]) // FIELD_POS       0001
    w.Strps, _ = strconv.Atoi(lrecd.Cols[5]) // CHARACTER_FIRST 000001
    w.Endps, _ = strconv.Atoi(lrecd.Cols[6]) // CHARACTER_LAST  000010
    u.Out.IsrtItems(w)
  }
  return u
}

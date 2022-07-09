// parms.rs [2020-01-22 BAR8TL]
// Gets a list of command-line parameters
use std::env;

#[derive(Debug)]

struct ParamTp {
  optn: String,
  prm1: String,
  prm2: String,
}

struct ParmsTp {
  cmdpr: [ParamTp; 2],
  messg: String,
}

impl ParmsTp {
  fn NewParms() -> ParmsTp {
    let par = ParamTp { 
      optn: String::from(""),
      prm1: String::from(""),
      prm2: String::from(""),
    };
    let prm: ParmsTp; 
    prm.cmdpr[0] = par;
    prm.cmdpr[1] = par;
    prm.messg = "".to_string();
    prm
  }
}

fn main() {
  let args: Vec<String> = env::args().collect();

  let query = &args[1];
  let filename = &args[2];

  println!("Searching for {}", query);
  println!("In file {}", filename);
}

/*
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
*/
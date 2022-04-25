package main

import "fmt"
import "log"
import "os"
import "path/filepath"

func main() {
  err := filepath.Walk(os.Args[1],
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    fmt.Printf("%s\t%d\r\n", path, info.Size())
    return nil
  })
  if err != nil {
    log.Println(err)
  }
}

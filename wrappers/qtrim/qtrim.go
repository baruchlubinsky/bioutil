package qtrim

import (
  //"io"
  "os"
  //"io/ioutil"
  "path"
  "os/exec"
  "log"
)

func init() {
  cmd := exec.Command("which", "qtrim")
  err := cmd.Run()
  if err != nil {
    log.Print(err)
    log.Fatal("QTrim is not installed on the system and I don't know how to install it.")
  }
}

func Run(arg ...string) (*exec.Cmd, error) {
  var output string
  for index, a := range arg {
    if a == "-output" {
      output = arg[index + 1]
    }
  }
  outputDir := path.Dir(output)
  outputErr, err := os.Create(path.Join(outputDir, "qtrim.err"))
  if err != nil {
    log.Fatal(err)
  }
  outputOut, err := os.Create(path.Join(outputDir, "qtrim.out"))
  if err != nil {
    log.Fatal(err)
  }
  cmd := exec.Command("qtrim", arg...)
  cmd.Stdout = outputOut
  cmd.Stderr = outputErr
  cmdError := cmd.Start()
  return cmd, cmdError
}

func Illumina2sanger(input []byte) {
  for i := range input {
    if i != len(input) - 1 {
      input[i] += 31
    }
  }
}

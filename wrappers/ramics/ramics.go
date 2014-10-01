package ramics

import (
  "io"
  "os"
  //"io/ioutil"
  "os/exec"
  "log"
  "text/template"
)

const installScript = `#! /bin/sh
mkdir -p {{.}}
cd {{.}}
git clone git@github.com:SANBIHIV/ramics-final.git
cd ramics-final
git checkout with-poi-for-seq2res
autoreconf --install
./configure CXXFLAGS='-O4'
make 
`

var path = "/usr/local/etc/gobiotools/ramics"

var exe string

func init()  {
  cmd := exec.Command("which", "ramics")
  err := cmd.Run()
  if err == nil {
    exe = "ramics"
    return
  }
  if os.Getenv("BIOTOOLS_RAMICS") != "" {
    path = os.Getenv("BIOTOOLS_RAMICS")
  }
  exe = path + "/ramics-final/ramics"
  _, err = os.Stat(exe)
  if err != nil {
    install(path)
  }
}

func install(path string) {
  t := template.Must(template.New("install").Parse(installScript))
  scriptFile, _ := os.Create("install-ramics.sh")
  scriptFile.Chmod(8660)
  t.Execute(scriptFile, path)
  scriptFile.Close()
  cmd := exec.Command("./install-ramics.sh")
  err := cmd.Run()
  if err != nil {
    log.Print(err)
    log.Fatal("Unable to install RAMICS")
  }
}

func Run(arg ...string) (*exec.Cmd, error) {
  cmd := exec.Command(exe, arg...)
  err := cmd.Start()
  return cmd, err
}

func RunWithOutput(output *io.Writer, arg ...string) (*exec.Cmd, error) {
  cmd := exec.Command(exe, arg...)
  cmd.Stdout = *output
  cmd.Stderr = *output
  err := cmd.Start()
  return cmd, err
}

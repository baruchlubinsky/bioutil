package pear

import (
	"io"
	"os"
	//"io/ioutil"
	"log"
	"os/exec"
	"text/template"
)

const installScript = `#! /bin/sh
mkdir -p {{.}}
cd {{.}}
git clone https://github.com/xflouris/PEAR.git .
./autogen.sh
./configure
make 
`

var path = "/usr/local/etc/gobiotools/pear"

var exe string

func init() {
	cmd := exec.Command("which", "pear")
	err := cmd.Run()
	if err == nil {
		exe = "pear"
		return
	}
	if os.Getenv("BIOTOOLS_PEAR") != "" {
		path = os.Getenv("BIOTOOLS_PEAR")
	}
	exe = path + "/pear"
	_, err := os.Stat(exe)
	if err != nil {
		install(path)
	}
}

func install(path string) {
	t := template.Must(template.New("install").Parse(installScript))
	scriptFile, _ := os.Create("install-pear.sh")
	scriptFile.Chmod(8660)
	t.Execute(scriptFile, path)
	scriptFile.Close()
	cmd := exec.Command("./install-pear.sh")
	err := cmd.Run()
	if err != nil {
		log.Print(err)
		log.Fatal("Unable to install PEAR")
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

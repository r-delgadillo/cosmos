package cmd

import (
	"os"

	"github.com/princjef/mageutil/shellcmd"
)

func init() {
	os.Chdir("cmd/mycommand") // Change directory to your command directory
}

func Build() error {
	return shellcmd.Command("docker build -t myapp .").Run()
}

func Run() error {
	Build()
	return shellcmd.Command("docker run -p 8080:8080 myapp").Run()
}

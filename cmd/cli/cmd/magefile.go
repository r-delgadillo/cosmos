package cmd

import (
	"os"

	"github.com/princjef/mageutil/shellcmd"
)

func init() {
	os.Chdir("cmd/mycommand") // Change directory to your command directory
}

func DockerBuild() error {
	return shellcmd.Command("docker build -t myapp .").Run()
}

func DockerRun() error {
	DockerBuild()
	return shellcmd.Command("docker run -p 8080:8080 myapp").Run()
}

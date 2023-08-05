package cmd

import (
	"fmt"
	"os"

	"github.com/princjef/mageutil/shellcmd"
)

func init() {
	os.Chdir("cmd/mycommand") // Change directory to your command directory
}

func GoBuild(tag string) error {
	return shellcmd.Command("go build ./cmd/cli/cli.go").Run()
}

func DockerBuildWithTag(tag string) error {
	// shellcmd.Command("docker build -t cosmos .").Run()
	return shellcmd.Command("minikube image build -t cosmos .").Run()
}

func DockerBuild() error {
	return shellcmd.Command("docker build -t myapp .").Run()
}

func DockerRun() error {
	DockerBuild()
	return shellcmd.Command("docker run -p 8080:8080 myapp").Run()
}

func RestartCosmos() error {
	return shellcmd.Command("kubectl delete pod -l app=cosmosapp").Run()
}

func Portforward() error {
	podName := shellcmd.Command("kubectl get pod -l app=cosmosapp -o name")
	val, _ := podName.Output()
	return shellcmd.Command(fmt.Sprintf("kubectl port-forward %s 8080", string(val))).Run()
}

func Debug() error {
	podName := shellcmd.Command("kubectl get pod -l app=cosmosapp -o name")
	val, _ := podName.Output()
	return shellcmd.Command(fmt.Sprintf("kubectl port-forward %s 32768", string(val))).Run()
}

func Start() error {
	// podName := shellcmd.Command("kubectl get pod -l app=cosmosapp -o name")
	// val, _ := podName.Output()
	return shellcmd.Command("minikube start --mount --mount-string=/home/rodelga/github/cosmos:/cosmos").Run()
}

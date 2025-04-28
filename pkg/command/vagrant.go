package command

import (
	"os"
	"strings"
	"os/exec"
	"d4do/pkg/config"
)

func VagrantCommand(command string, vmName string) {
	VagrantCommandWithPassword(command, vmName, "")
}

func VagrantCommandWithPassword(command string, vmName string, password string) {
	commandSlice := strings.Fields(command)

	vagrantDir, _ := os.MkdirTemp(config.AppConfigDir, "temp")
	CreateVagrantFile(
		vagrantDir,
		vmName,
		password,
	)
	defer os.RemoveAll(vagrantDir)
	
	os.Setenv("VAGRANT_CWD", config.AppConfigDir + "/vagrant/")
	os.Setenv("VAGRANT_VAGRANTFILE", vagrantDir + "/Vagrantfile")
	
	Exec(commandSlice[0], commandSlice[1:]...)
}

func VagrantListVms() []string {
	shOut, _ := exec.Command("vagrant", "global-status").Output()
	lines := strings.Split(strings.ReplaceAll(string(shOut), "\r\n", "\n"), "\n")[2:]

	vms := make([]string, 0)
	for i := range lines {
		if len(strings.Fields(lines[i])) > 1 {
			vms = append(vms, strings.Fields(lines[i])[1])
		} else {
			break
		}
	}
	return vms
}

func CreateVagrantFile(vagrantDir, machineName string, password string) {
	networkInterface, _ := exec.Command("bash", "-c", `ip route | awk '/^default/ {printf "%s", $5; exit 0}'`).Output()
	vagrantConfig := config.VagrantConfig(machineName, string(networkInterface), password)
	os.WriteFile(vagrantDir + "/Vagrantfile", ([]byte)(vagrantConfig), os.ModePerm)	
}

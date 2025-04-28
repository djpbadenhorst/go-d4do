package config

func VagrantConfig(machineName string, networkInterface string, password string) string {
	return `
Vagrant.configure("2") do |config|
  config.vm.define "` + machineName + `"
  config.vm.box = "bento/ubuntu-24.04"
  config.vm.network "public_network", bridge: "` + string(networkInterface) + `"
  config.vm.hostname = "` + machineName + `.local"
  config.vm.synced_folder "~/.d4do/vagrant/` + machineName + `", "/home/vagrant/Workspace/"
  config.vm.provision "shell", inline: <<-SHELL
    echo "vagrant:` + password + `" | sudo chpasswd
    sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/g' /etc/ssh/sshd_config
    useradd -mg vagrant djpb
    usermod -aG sudo djpb
    echo "djpb:` + password + `" | sudo chpasswd
    usermod -s /bin/bash djpb
SHELL
end`
}

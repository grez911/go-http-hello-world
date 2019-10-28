# -*- mode: ruby -*-
# vi: set ft=ruby :

$script = <<SCRIPT
#echo "Installing GitLab..."
#sudo apt-get update
#sudo apt-get install -y curl openssh-server ca-certificates
#sudo DEBIAN_FRONTEND=noninteractive apt-get install -y postfix
#curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ee/script.deb.sh | sudo bash
#sudo EXTERNAL_URL="http://gitlab.local" apt-get install gitlab-ee
SCRIPT

Vagrant.configure(2) do |config|
  config.vm.box = "bento/ubuntu-18.04" # 18.04 LTS
  #config.vm.hostname = "gitlab.local"
  config.vm.provision "shell", inline: $script, privileged: false

  # Expose GitLab api and ui to the host.
  config.vm.network "forwarded_port", guest: 80, host: 8080, auto_correct: true

  # Increase memory for Virtualbox.
  config.vm.provider "virtualbox" do |vb|
        vb.memory = "6144"
        vb.cpus = "4"
  end
end

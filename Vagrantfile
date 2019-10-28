# -*- mode: ruby -*-
# vi: set ft=ruby :

$script = <<SCRIPT
echo "Installing GitLab..."
apt-get update
apt-get install -y curl openssh-server ca-certificates apt-transport-https
DEBIAN_FRONTEND=noninteractive apt-get install -y postfix
curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ee/script.deb.sh | sudo bash
EXTERNAL_URL="http://gitlab.local" apt-get install gitlab-ee

echo "Installing Docker..."
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
apt update
apt install -y docker-ce

echo "Installing Minikube..."
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube
install minikube /usr/local/bin/

echo "Installing kubectl..."
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl

echo "Starting Minikube..."
minikube start --vm-driver=none
SCRIPT

Vagrant.configure(2) do |config|
  config.vm.box = "bento/ubuntu-18.04" # 18.04 LTS
  #config.vm.hostname = "gitlab.local"
  config.vm.provision "shell", inline: $script, privileged: true

  # Expose GitLab api and ui to the host.
  #config.vm.network "forwarded_port", guest: 80, host: 8080, auto_correct: true

  # Increase memory for Virtualbox.
  config.vm.provider "virtualbox" do |vb|
        vb.memory = "6144"
        vb.cpus = "4"
  end
end

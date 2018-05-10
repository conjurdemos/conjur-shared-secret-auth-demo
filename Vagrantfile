# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"

  config.vm.provision "shell", inline: <<-SHELL

    # Install packages to allow apt to use a repository over HTTPS
    apt-get update -y
    apt-get install -y \
      apt-transport-https \
      ca-certificates \
      curl \
      software-properties-common

    # Add Docker's official GPG key
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

    # Set up the stable Docker repository
    add-apt-repository \
      "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) \
      stable"

    # Install the latest version of Docker CE
    apt-get update -y
    apt-get install -y docker-ce

    # Install docker-compose
    sudo curl -L https://github.com/docker/compose/releases/download/1.21.1/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose

    # Setup the user for Docker
    groupadd docker
    usermod -aG docker vagrant

  SHELL
end

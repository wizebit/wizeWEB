#!/usr/bin/env bash

sudo apt-get -y update
sudo apt-get install wget zip unzip -y
sudo apt-get install gcc g++ -y
#Install python for Ansible
test -e /usr/bin/python || (sudo apt -qqy update && sudo apt install -qqy python-minimal python-setuptools)
#Install golang compiler
cd ~/
wget https://dl.google.com/go/go1.9.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.9.5.linux-amd64.tar.gz    #INFO:  change to right distro
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
mkdir /home/ubuntu/go; mkdir /home/ubuntu/go/src; mkdir /home/ubuntu/go/bin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
echo 'export GOPATH=/home/ubuntu/go' >> ~/.profile
echo 'export GOBIN=/home/ubuntu/go/bin' >> ~/.profile
echo 'export PATH=$PATH:$GOROOT/bin' >> ~/.profile

#ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts
#chmod 600 ~/wize_web
# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # All Vagrant configuration is done here. The most common configuration
  # options are documented and commented below. For a complete reference,
  # please see the online documentation at vagrantup.com.

  # Every Vagrant virtual environment requires a box to build off of.
  config.vm.box = "precise64"

  # The url from where the 'config.vm.box' box will be fetched if it
  # doesn't already exist on the user's system.
  config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/precise/current/precise-server-cloudimg-amd64-vagrant-disk1.box"
  
  config.vm.provider :virtualbox do |vb, override|
    # Use VBoxManage to customize the VM. For example to change memory:
    vb.customize ["modifyvm", :id, "--memory", "1024"]

    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine. In the example below,
    # accessing "localhost:8080" will access port 80 on the guest machine.
    override.vm.network :forwarded_port, guest: 8087, host: 8087
    override.vm.network :forwarded_port, guest: 8098, host: 8098
  end

  config.vm.provider :digital_ocean do |digitalocean, override|
    override.ssh.private_key_path = 'C:/Users/Matthew/SkyDrive/Documents/ssh/digital_ocean'
    override.vm.box = 'digital_ocean'
    override.vm.box_url = "https://github.com/smdahlen/vagrant-digitalocean/raw/master/box/digital_ocean.box"

    digitalocean.client_id = 'f9cb8da54b9c3689e24eb4fbb922a646'
    digitalocean.api_key = '401f0bc58c73e9e68feed6b9bec7014a'

    digitalocean.image = 'Ubuntu 12.04.3 x64'
    digitalocean.region = 'New York 2'
    digitalocean.size = '512MB'
    digitalocean.private_networking = false
    digitalocean.ssh_key_name = 'DigitalOcean'
  end

  # Enable shell provisioning
  config.vm.provision :shell, :path => "vagrant/bootstrap.sh"
end

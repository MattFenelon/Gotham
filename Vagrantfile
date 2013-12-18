# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # Every Vagrant virtual environment requires a box to build off of.
  config.vm.box = "precise64"

  # The url from where the 'config.vm.box' box will be fetched if it
  # doesn't already exist on the user's system.
  config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/precise/current/precise-server-cloudimg-amd64-vagrant-disk1.box"

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  config.vm.network :forwarded_port, guest: 8087, host: 8087
  config.vm.network :forwarded_port, guest: 8098, host: 8098
  
  # Default synched folder
  config.vm.synced_folder "./", "/vagrant"

  # Sync SSH keys to connect with DigitalOcean
  config.vm.synced_folder "C:/Users/Matthew/SkyDrive/Documents/ssh", "/ssh_keys",
    mount_options: ["dmode=700,fmode=700"]

  config.vm.provider :virtualbox do |vb|
    # Use VBoxManage to customize the VM. For example to change memory:
    vb.customize ["modifyvm", :id, "--memory", "1024"]
  end

  # Enable shell provisioning
  config.vm.provision :shell, :path => "vagrant/bootstrap.sh"
  # Install and configure riak
  config.vm.provision :shell, :path => "vagrant/riak.sh"
  # Set up for API development (install Go and set GOPATH)
  config.vm.provision :shell, :path => "vagrant/golang.sh"
end

hostname = 'GoApi'
ram      = 1024

Vagrant.configure("2") do |config|
    config.vm.box = "chef/debian-7.6"

    config.vm.provider :virtualbox do |vb|
        vb.customize ["modifyvm", :id, "--name", hostname, "--memory", ram]
    end

    # Synced folders
    config.vm.synced_folder ".", "/vagrant", group: "www-data", owner: "www-data"

    # Network settings
    config.vm.network "private_network", ip: "192.168.10.1"

    # Provision
    config.vm.provision "ansible" do |ansible|
        ansible.playbook = "provision/site.yml"
    end
end

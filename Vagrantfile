# -*- mode: ruby -*-
# # vi: set ft=ruby :

Vagrant.require_version ">= 1.6.0"
VAGRANTFILE_API_VERSION = "2"

SHARED_DIR="/home/core/share"

$update_channel = "stable"
$image_version = "current"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "coreos-%s" % $update_channel
  config.vm.box_url = "http://%s.release.core-os.net/amd64-usr/%s/coreos_production_vagrant.json" % [$update_channel, $image_version]
  config.vm.define "local" do |cfg|
    cfg.ssh.insert_key = false
    cfg.vm.host_name = "local"
    cfg.vm.network :private_network, ip: "172.17.11.101"
    cfg.vm.synced_folder ".", SHARED_DIR, id: "core", :nfs => true, :mount_options => ['nolock,vers=3,udp']
    cfg.vm.provision "postgres", type: "docker" do |d|
      d.build_image "#{SHARED_DIR}/container/postgres", args: "-t takasing/postgres:0.0.1"
      d.run "postgres",
        image: "takasing/postgres:0.0.1",
        args: "-p 5432:5432",
        auto_assign_name: true,
        daemonize: true
    end

    cfg.vm.provision "stop-container", type: "shell", inline: "docker rm -f golang-webapp nginx"

    cfg.vm.provision "golang", type: "docker" do |d|
      d.build_image "#{SHARED_DIR}", args: "-t takasing/golang-webapp:0.0.1"
      d.run "golang-webapp",
        image: "takasing/golang-webapp:0.0.1",
        args: "-p 3000:3000 --link postgres:db",
        auto_assign_name: true,
        daemonize: true
    end
    cfg.vm.provision "nginx", type: "docker" do |d|
      d.build_image "#{SHARED_DIR}/container/nginx", args: "-t takasing/nginx:0.0.1"
      d.run "nginx",
        image: "takasing/nginx:0.0.1",
        args: "-p 80:80 --link golang-webapp:api",
        auto_assign_name: true,
        daemonize: true
    end
  end
end


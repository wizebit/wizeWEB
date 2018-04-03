provider "google" {
  region = "${var.region}"
  project = "${var.project_name}"
  credentials = "${file(var.account_file_path)}"
}

resource "google_compute_address" "www" {
  name = "tf-www-address"
}

resource "google_compute_target_pool" "www" {
  name = "tf-www-target-pool"
  instances = ["${google_compute_instance.www.*.self_link}"]
  health_checks = ["${google_compute_http_health_check.http.name}"]
}

resource "google_compute_forwarding_rule" "http" {
  name = "tf-www-http-forwarding-rule"
  target = "${google_compute_target_pool.www.self_link}"
  ip_address = "${google_compute_address.www.address}"
  port_range = "80"
}

resource "google_compute_forwarding_rule" "https" {
  name = "tf-www-https-forwarding-rule"
  target = "${google_compute_target_pool.www.self_link}"
  ip_address = "${google_compute_address.www.address}"
  port_range = "443"
}

resource "google_compute_http_health_check" "http" {
  name = "tf-www-http-basic-check"
  request_path = "/"
  check_interval_sec = 1
  healthy_threshold = 1
  unhealthy_threshold = 10
  timeout_sec = 1
}

//# Salt master server is used to configure and manage the minions
//resource "google_compute_instance" "master" {
//  count = 1
//  name = "master-node"
//  machine_type = "f1-micro"
//  zone = "${var.region_zone}"
//  tags = ["master", "letsencrypt"]
//
//  boot_disk {
//    initialize_params {
//      image = "ubuntu-1604-lts"
//    }
//  }
//  network_interface {
//    network = "default"
//    access_config {
//      # Ephemeral
//    }
//  }
//  service_account {
//    scopes = ["https://www.googleapis.com/auth/compute.readonly"]
//  }
//
//  metadata_startup_script = <<SCRIPT
//# Update system
//apt -y update && apt -y upgrade
//
//# Install python
//apt -y install python2.7 python
//#echo -e "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC/bdJXEehLaDR+n794QByH/JKsqCuTc2ZrY9V19xtze6g4CNMnd0QcEAbi47zTvzjFv9HE08g/nf9YhBgKIUW8fgU6Qd1T7f1jBSOsImLhK1uN01T3kcN3LD5+mSoktCDxX6mvs8lCi8K5F0tu+1CurWVWxGSK6hX56k2s6qXLCm/QV3vmYlMKUTYXCGoIFQvcdXkCeliHuRYLWBcvOVMUOhtO4f81ab09q5qFvicxK2YjHDdGi8WxK5g8/MO3zLR3wHWN09NR8D6vfnEkvKOkfocPsvh5DS5wDS3VxXnXPRbCyFBi8aAUjOXO1L57JK3RTj3mxLaNz+ZADAQERhWL kolyvayko@MacBook-Pro-Alexey.local" > /root/.ssh/known_hosts
//#apt -y update
//#aptitude -y safe-upgrade
//#apt -y install salt-master salt-minion salt-ssh salt-cloud salt-doc
//#echo -e "master: $HOSTNAME" > /etc/salt/minion.d/master.conf
//#echo -e "grains:\n  roles:\n    - salt\n    - letsencrypt" > /etc/salt/minion.d/grains.conf
//#echo -e "file_roots:\n  base:\n    - /srv/gwadeloop-states/salt\npillar_roots:\n  base:\n    - /srv/gwadeloop-states/pillar" > /etc/salt/master.d/path_roots.conf
//#mkdir /srv/gwadeloop-states/
//#chown -R ubuntu:root /srv/gwadeloop-states/
//SCRIPT
//
////  provisioner "file" {
////    connection {
////      user = "ubuntu"
////    }
////
////    source = "/srv/gwadeloop-states"
////    destination = "/home/ubuntu/"
////  }
//
////  provisioner "remote-exec" {
////    connection {
////      user = "ubuntu"
////    }
////    inline = [
////      "sudo cp -r /home/ubuntu/gwadeloop-states /srv/gwadeloop-states",
////      "sudo chown -R ubuntu:root /srv/gwadeloop-states"
////    ]
////  }
//
//  metadata {
//    sshKeys = "ubuntu:${file("~/.ssh/id_rsa.pub")}"
//  }
//}

# Salt minions (2) part of the www cluster
resource "google_compute_instance" "www" {
  count = 1
  name = "tf-www-${count.index}"
  machine_type = "f1-micro"
  zone = "${var.region_zone}"
  tags = ["www-node"]

  boot_disk {
    initialize_params {
      image = "ubuntu-1604-lts"
    }
  }
  network_interface {
    network = "default"
    access_config {
      # Ephemeral
    }
  }
  service_account {
    scopes = ["https://www.googleapis.com/auth/compute.readonly"]
  }

  metadata_startup_script = <<SCRIPT
# Update system
apt -y update && apt -y upgrade

# Install python
apt -y install python2.7 python

echo 'ubuntu ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

SCRIPT


  metadata {
    sshKeys = "ubuntu:${file("~/.ssh/id_rsa.pub")}"
  }

//  provisioner "local-exec" {
//    command = "sleep 120; cd ../ansible; ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u ubuntu -i '${google_compute_address.www.address},' wizebit_web.yml"
//  }
}



resource "google_compute_firewall" "www" {
  name = "tf-www-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports = ["80", "443"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags = ["www-node"]
}


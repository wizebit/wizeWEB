output "public_ip" {
  value = "${google_compute_address.www.address}"
}

output "slave_ips" {
  value = "${join(" ", google_compute_instance.slave.*.network_interface.0.access_config.0.assigned_nat_ip)}"
}

output "master_ip" {
  value = "${join(" ", google_compute_instance.www.*.network_interface.0.access_config.0.assigned_nat_ip)}"
}
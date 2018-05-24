variable "region" {
  default = "us-east1"
}

variable "region_zone" {
  default = "us-east1-b"
}

variable "project_name" {
  description = "The ID of the Google Cloud project"
}

variable "account_file_path" {
  description = "Path to the JSON file used to describe your account credentials"
}

variable "number_of_servers" {
  default = "2"
}

terraform {
  required_providers {
    elice = {
      source = "elice-dev/elice"
      version = "0.1.0-alpha"
    }
  }
}

provider "elice" {
  host         = var.host
  token        = var.token
  organization = var.organization
}

resource "elice_instance" "example" {
  title            = "ex-instance"
  image_id         = "c64935ca-563e-4a05-9034-7469213f8204"
  instance_type_id = "77f3b59a-6cae-49e8-8978-4b92b3e7a060"
  disk             = 128
}
terraform {
  required_providers {
    ctoa = {
      source  = "vmware.com/edu/ctoa"
    }
  }
}

provider "ctoa" {
  host = "127.0.0.1"
}


resource "ctoa_people" "johnsmith" {
  first_name = "John"
  last_name = "Smith"
}

/*
resource "ctoa_people" "adeleporte" {
  first_name = "Antoine"
  last_name = "Deleporte"
}


resource "ctoa_people" "nvibert" {
  first_name = "Nico"
  last_name = "Vibert"
}
*/
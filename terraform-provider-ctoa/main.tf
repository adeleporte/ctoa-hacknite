terraform {
  required_providers {
    ctoa = {
      source  = "adeleporte.com/edu/ctoa"
    }
  }
}

provider "ctoa" {
  host = "127.0.0.1"
}


resource "ctoa_people" "deleporte" {
  first_name = "Antoine"
  last_name = "Deleporte"
}

/*
resource "ctoa_people" "nvibert" {
  first_name = "Nico"
  last_name = "Vibert"
}
*/
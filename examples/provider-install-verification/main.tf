terraform {
  required_providers {
    servicenowtable = {
        source = "hashicorp.com/edu/servicenowtable"
    }
  }
}

provider "servicenowtable" {}

data "servicenowtable_query" "query" {}

terraform {
  required_providers {
    servicenowtable = {
      source = "ckkannan/servicenowtable"
    }
  }
}

provider "servicenowtable" {
  sn_url  = "https://dev161016.service-now.com"
  sn_user = "admin"
  sn_pass = "@6zeg6GETWFiTze/"
  sn_auth = "Basic"
}

data "servicenowtable_queryorg" "query" {

}


output "myout" {
  value = data.servicenowtable_queryorg.query
}

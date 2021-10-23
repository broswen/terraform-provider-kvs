terraform {
  required_providers {
    kvs = {
      source  = "broswen.com/kvs/kvs"
      version = "~>0.1.0"
    }
  }
}

provider "kvs" {
  host = "http://kvs.broswen.com"
}

data "kvs_pair" "test" {
  key = "test"
}

resource "kvs_pair" "some_pair" {
  key   = "some-pair"
  value = "this is the value of a pair resource"
}


output "some_pair" {
  value = resource.kvs_pair.some_pair
}

output "datasource_test" {
  value = data.kvs_pair.test
}

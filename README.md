# terraform-provider-kvs

This is a Terraform provider for my [key-value store service (KVS)](https://github.com/broswen/kvs).

This provider exposes a resource to manage key-value pairs and a data source to read the value from a pair.


This sets the value of "pair1" to "value1"
```
resource "kvs_pair" "pair1"{
  key = "pair1"
  value = "value1"
}
```

This reads the value of "pair1"
```
data "kvs_pair" "pair1" {
  key = "pair1"
}
```
terraform {
  required_providers {
    flintlock = {
      version = "0.3.1"
      source = "weave.works/liquidmetal/flintlock"
    }
  }
}

provider "flintlock" {
  hosts = [ "127.0.0.1:9090" ]
}

resource "flintlock_microvm" "tst" {
    name = "test1"
    kernel_image = "myimagehere"
    root_volume_image = "myrootimage"
}
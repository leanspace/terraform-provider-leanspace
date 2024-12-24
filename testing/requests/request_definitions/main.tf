terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "8.1.5"
    }
  }
}

provider "leanspace" {
  env           = "develop"
  tenant        = "yuri"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

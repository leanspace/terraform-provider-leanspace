terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
      version = "8.5.27"
    }
  }
}

provider "leanspace" {
  tenant        = "yuri"
  env           = "develop"
  client_id     = "nlbja2p65j8kj7of0tfs29rf4"
  client_secret = "d762kk9862jn0j1qr4c2u3o8bjkv70o45pld3200ek89qtul6kg"
}

resource "leanspace_leaf_space_integrations" "integration_Connection" {
  name="Leafspace Connection Terraform"
  username= "leanspace01"
  password= "vVVy5syuUuSk9MVkRfvYm8Xp6Nhy6qAs"
  domain_url= "sandbox.api.leaf.space"
}
/*
data "leanspace_leaf_space_integrations" "all" {
}

output "all" {
  value = data.leanspace_leaf_space_integrations.all
}*/
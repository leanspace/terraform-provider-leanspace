terraform {
  required_providers {
    leanspace = {
      version = "0.3"
      source  = "leanspace.io/io/leanspace"
    }
  }
}

data "leanspace_plugins" "all" {}

variable "path" {
  type        = string
  description = "The path to the plugin's source code file (a .jar)."
}

resource "leanspace_plugins" "test" {
  plugin {
    file_path                            = var.path
    type                                 = "COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE"
    implementation_class_name            = "io.myplugin.SimpleCommandTransformer"
    name                                 = "Terraform Command Transformer Plugin"
    description                          = "This is a plugin created through terraform!"
    source_code_file_download_authorized = true
  }
}

output "test_plugin" {
  value = leanspace_plugins.test
}

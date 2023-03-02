terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

variable "path" {
  type        = string
  description = "The path to the plugin's source code file (a .jar)."
}

data "leanspace_plugins" "all" {
  filters {
    ids = []
    types = [
      "COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE",
      "COMMANDS_PROTOCOL_TRANSFORMER_PLUGIN_TYPE",
    ]
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

resource "leanspace_plugins" "test" {
  file_path                            = var.path
  type                                 = "COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE"
  implementation_class_name            = "io.myplugin.SimpleCommandTransformer"
  name                                 = "Terraform Command Transformer Plugin"
  description                          = "This is a plugin created through terraform!"
  source_code_file_download_authorized = true
  sdk_version                          = "2.1.2"
}

output "all_plugins" {
  value = data.leanspace_plugins.all
}

output "test_plugin" {
  value = leanspace_plugins.test
}

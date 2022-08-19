resource "leanspace_plugins" "plugin" {
  plugin {
    file_path                            = var.path
    type                                 = "COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE"
    implementation_class_name            = "io.myplugin.SimpleCommandTransformer"
    name                                 = "Terraform Command Transformer Plugin"
    description                          = "This is a plugin created through terraform!"
    source_code_file_download_authorized = true
  }
}

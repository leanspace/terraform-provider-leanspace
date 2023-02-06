data "leanspace_plugins" "all" {
  filters {
    types = [
      "COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE",
      "COMMANDS_PROTOCOL_TRANSFORMER_PLUGIN_TYPE",
    ]
    ids   = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

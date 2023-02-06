data "leanspace_command_definitions" "all" {
  filters {
    node_ids                    = [var.node_id]
    node_types                  = ["ASSET"]
    node_kinds                  = ["SATELLITE"]
    with_arguments_and_metadata = true
    ids                         = []
    created_by                  = null
    query                       = ""
    page                        = 0
    size                        = 10
    sort                        = ["name,asc"]
  }
}

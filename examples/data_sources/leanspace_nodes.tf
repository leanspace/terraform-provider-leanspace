data "leanspace_nodes" "all" {
  filters {
    parent_node_ids = []
    types           = ["ASSET"]
    is_root_node    = true
    created_by            = ""
    from_created_at       = ""
    last_modified_by      = ""
    to_created_at         = ""
    from_last_modified_at = ""
    to_last_modified_at   = ""
    ids                   = []
    kinds                 = []
    query                 = ""
    tags                  = []
    page                  = 0
    size                  = 10
    sort                  = ["name,asc"]
  }
}

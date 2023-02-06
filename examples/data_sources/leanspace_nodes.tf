data "leanspace_nodes" "all" {
  filters {
    parent_node_ids = []
    types           = ["ASSET"]
    is_root_node    = true
    created_by            = null
    from_created_at       = null
    last_modified_by      = null
    to_created_at         = null
    from_last_modified_at = null
    to_last_modified_at   = null
    ids                   = []
    kinds                 = []
    query                 = ""
    tags                  = []
    page                  = 0
    size                  = 10
    sort                  = ["name,asc"]
  }
}

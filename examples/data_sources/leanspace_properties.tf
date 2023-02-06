data "leanspace_properties" "all" {
  filters {
    node_ids              = [var.node_id]
    category              = ""
    created_by            = null
    from_created_at       = null
    last_modified_by      = null
    to_created_at         = null
    from_last_modified_at = null
    to_last_modified_at   = null
    ids                   = []
    kinds                 = []
    node_types            = []
    query                 = ""
    tags                  = []
    page                  = 0
    size                  = 10
    sort                  = ["name,asc"]
  }
}

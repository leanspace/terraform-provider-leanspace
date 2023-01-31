data "leanspace_properties" "all" {
  filters {
    node_ids              = [var.node_id]
    category              = ""
    created_by            = ""
    from_created_at       = ""
    last_modified_by      = ""
    to_created_at         = ""
    from_last_modified_at = ""
    to_last_modified_at   = ""
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

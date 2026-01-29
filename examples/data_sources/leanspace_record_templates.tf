data "leanspace_record_templates" "all" {
  filters {
    ids                   = []
    names                 = []
    node_ids              = []
    metric_ids            = []
    query                 = ""
    created_by            = []
    last_modified_by      = []
    from_created_at       = []
    to_created_at         = []
    from_last_modified_at = []
    to_last_modified_at   = []
    tags                  = []
    page                  = 0
    size                  = 10
    sort                  = ["name,asc"]
  }
}

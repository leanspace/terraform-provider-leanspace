data "leanspace_passive_resource_functions" "all" {
  filters {
    ids                     = []
    resource_ids            = []
    tags                    = []
    query                   = ""
    page                    = 0
    size                    = 10
    sort                    = ["name,asc"]
  }
}

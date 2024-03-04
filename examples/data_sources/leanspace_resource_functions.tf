data "leanspace_resource_functions" "all" {
  filters {
    ids                     = []
    activity_definition_ids = []
    resource_ids            = []
    page                    = 0
    size                    = 10
    sort                    = ["name,asc"]
  }
}

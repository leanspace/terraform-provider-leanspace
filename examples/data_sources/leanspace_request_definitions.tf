data "leanspace_request_definitions" "all" {
  filters {
    plan_template_ids                     = []
    feasibility_constraint_definition_ids = []
    ids                                   = []
    query                                 = ""
    page                                  = 0
    size                                  = 10
    sort                                  = ["name,asc"]
  }
}

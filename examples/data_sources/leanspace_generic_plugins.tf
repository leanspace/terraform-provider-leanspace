data "leanspace_generic_plugins" "all" {
  filters {
    types = [
      "CHECKSUM_FUNCTION",
    ]
    ids   = []
    query = ""
    page  = 0
    size  = 10
    sort  = ["name,asc"]
  }
}

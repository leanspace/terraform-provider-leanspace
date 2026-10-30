resource "leanspace_units" "custom_units" {
  security_tags = (each.key == "k" ? ["terraform_mission"] : [])
  for_each = {
    "k" : "Kilo"
    "h" : "Hecto"
    "da" : "Deca"
    "" : ""
    "d" : "Deci"
    "c" : "Centi"
    "m" : "Milli"
  }
  symbol       = "${each.key}Cus"
  display_name = "${each.value}Customium"
}
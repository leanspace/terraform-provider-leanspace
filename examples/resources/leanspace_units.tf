resource "leanspace_units" "custom_units" {
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
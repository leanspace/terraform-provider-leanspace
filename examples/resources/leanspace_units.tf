resource "leanspace_units" "custom_units" {
  for_each = {
    "k" : { value = "Kilo", security_tags = true }
    "h" : { value = "Hecto", security_tags = false }
    "da" : { value = "Deca", security_tags = false }
    "" : { value = "", security_tags = false }
    "d" : { value = "Deci", security_tags = false }
    "c" : { value = "Centi", security_tags = false }
    "m" : { value = "Milli", security_tags = false }
  }
  security_tags = each.value.security_tags ? ["terraform_mission"] : []
  symbol        = "${each.key}Cus"
  display_name  = "${each.value.value}Customium"
}
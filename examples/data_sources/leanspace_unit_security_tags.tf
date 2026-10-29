resource "leanspace_unit_security_tags" "mission" {
  unit_id       = leanspace_units.custom_units["k"].id
  security_tags = ["terraform_mission"]
}
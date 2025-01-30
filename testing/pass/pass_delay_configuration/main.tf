terraform {
  required_providers {
    leanspace = {
      source = "leanspace/leanspace"
    }
  }
}

resource "leanspace_pass_delay_configuration" "delay" {
  name = "test pass delay"
  aos_delay_in_millisecond=1000
  los_delay_in_millisecond=2000
}
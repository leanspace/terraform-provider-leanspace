resource "leanspace_pass_delay_configuration" "delay" {
  name = "test"
  aos_delay_in_millisecond=1000
  los_delay_in_millisecond=2000
}
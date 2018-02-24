provider splunk {
  insecure_skip_verify = true
  url                  = "https://upside.splunkcloud.com:8089"
}

resource "splunk_saved_search" "test1" {
  name   = "dennis's test"
  search = "what"

  acl {
    owner = "terraform"
    sharing = "global"
    read = ["*"]
    write= ["admin","user"]
  }
}

// resource "splunk_saved_search" "test2" {
//   name                       = "dennis's test alarm"
//   search                     = "sisyphus"
//   actions                    = "slack"
//   action_email_to            = "dennis@bluesentryit.com"
//   action_email_subject       = "Test Alert"
//   action_email_message_alert = "Hey is this thing on!"
//   action_slack_channel       = "@dennis"
//   action_slack_message       = "Hey dude!"
//   alert_comparator           = "greater than"
//   alert_threshold            = "1"
//   dispatch_earliest_time     = "-15m"
//   cron_schedule              = "5 5 5 5 *"
//   is_scheduled               = true
// }
// 
// resource "splunk_saved_search" "test3" {
//   name                = "Sisyphus Kubectl Errors"
//   search              = "I need to build this"
//   action_email_inline = true
// }
// 
// resource "splunk_saved_search" "test4" {
//   name                = "Sisyphus Kubectl "
//   search              = "I need to build this"
//   action_email_inline = true
// }


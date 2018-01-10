provider splunk {
  insecure_skip_verify = true
  url                  = "https://upside.splunkcloud.com:8089"
}

resource "splunk_saved_search" "test1" {
  name   = "dennis's test now with ++ signs"
  search = "kubernetes.labels.app=sisyphus index=dev"
}

resource "splunk_saved_search" "test2" {
  name                       = "dennis's test alarm"
  search                     = "sisyphus"
  actions                    = "slack"
  action_email_to            = "dennis@bluesentryit.com"
  action_email_subject       = "Test Alert"
  action_email_message_alert = "Hey is this thing on!"
  action_slack_channel       = "@dennis"
  action_slack_message       = "Hey dude!"
  alert_comparator           = "greater than"
  alert_threshold            = "1"
  dispatch_earliest_time     = "-15m"
  cron_schedule              = "*/5 * * * *"
  is_scheduled               = true
}

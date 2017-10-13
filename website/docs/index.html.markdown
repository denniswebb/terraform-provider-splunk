---
layout: "splunk"
page_title: "Provider: Splunk"
sidebar_current: "docs-splunk-index"
description: |-
  The Splunk provider is used to manage the configuration of Splunk. The provider needs to be configured with the proper credentials before it can be used.
---

# Splunk Provider

The Splunk provider is used to interact with the
configuration of Splunk. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Splunk provider
provider "splunk" {
  username = "${var.splunk_email}"
  password = "${var.splunk_token}"
  api_url      = "${var.splunk_api_url}"

}

# Create a saved search
resource "splunk_saved_search" "login-attempts" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) Username used to authenticate to Splunk
* `password` - (Required) Password used to authenticate to Splunk
* `api_url` - (Required) URL to Splunk API. Example: `https://myorg.splunkcloud.com:8089`

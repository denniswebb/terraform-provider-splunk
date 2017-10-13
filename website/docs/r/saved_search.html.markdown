---
layout: "splunk"
page_title: "Splunk: splunk_saved_search"
sidebar_current: "docs-splunk-resource-record"
description: |-
  Provides a Splunk record resource.
---

# splunk_saved_search

Provides a Splunk saved search or alert.

## Example Usage

```hcl
# Add a saved search
resource "splunk_saved_search" "foobar" {
}
```

## Argument Reference

The following arguments are supported:

* `search` - (Required) The search string used for the saved search.
* `name` - (Required) The name of the saved search

## Attributes Reference

The following attributes are exported:

* `id` - The saved search ID
* `name` - The name of the saved search

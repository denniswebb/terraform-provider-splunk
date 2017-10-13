package main

import (
	"github.com/denniswebb/terraform-provider-splunk/splunk"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: splunk.Provider})
}

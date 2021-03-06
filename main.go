package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/zzxwill/terraform-provider-alicloud/alicloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: alicloud.Provider})
}

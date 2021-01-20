package main

import (
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gorillastack.Provider})
}

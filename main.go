package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-gorillastack/gorillastack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gorillastack.Provider})
}

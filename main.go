package main

import (
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gorillastack.Provider})
}

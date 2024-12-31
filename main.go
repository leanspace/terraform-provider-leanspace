package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	services.AddDataTypes()
	provider.Serve(&provider.ServeOpts{
		Debug:        debug,
		ProviderAddr: "registry.terraform.io/leanspace/leanspace",
		ProviderFunc: provider.Provider,
	})
}

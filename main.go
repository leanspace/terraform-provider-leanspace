package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	services.RegisterDataTypes()

	err := providerserver.Serve(context.Background(), provider.New(), providerserver.ServeOpts{
		Address: "registry.terraform.io/leanspace/leanspace",
		Debug:   debug,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

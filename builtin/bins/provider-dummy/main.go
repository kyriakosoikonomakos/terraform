package main

import (
	"github.com/hashicorp/terraform/builtin/providers/dummy"
	"github.com/hashicorp/terraform/plugin"
	"log"
)

func main() {
	log.Printf("[DEBUG] LLALALALA")
	plugin.Serve(new(dummy.ResourceProvider))
}

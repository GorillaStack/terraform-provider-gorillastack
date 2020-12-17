terraform {
  required_providers {
    gorillastack = {
      source = "terraform.gorillastack.com/gorillastack/gorillastack"
      #
      # For more information, see the provider source documentation:
      #
      # https://www.terraform.io/docs/configuration/providers.html#provider-source
    }
  }
  required_version = ">= 0.13"
}

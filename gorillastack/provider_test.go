package gorillastack

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"gorillastack": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GORILLASTACK_API_KEY"); v == "" {
		t.Fatal("GORILLASTACK_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("GORILLASTACK_TEAM_ID"); v == "" {
		t.Fatal("GORILLASTACK_TEAM_ID must be set for acceptance tests")
	}
}

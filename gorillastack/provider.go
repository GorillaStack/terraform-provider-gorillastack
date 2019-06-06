package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GORILLASTACK_API_KEY", nil),
				Description: "Key for accessing the GorillaStack API.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"gorillastack_rule": resourceRule(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	client, err := NewClient(apiKey)
	if (err != nil) {
		return nil, err
	}

	return client, nil
}
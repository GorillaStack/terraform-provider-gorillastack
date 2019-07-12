package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GORILLASTACK_API_KEY", nil),
				Description: "Key for accessing the GorillaStack API.",
			},
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GORILLASTACK_TEAM_ID", nil),
				Description: "TeamId to select tenant context for the GorillaStack API.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"gorillastack_rule":      resourceRule(),
			"gorillastack_tag_group": resourceTagGroup(),
		},

		ConfigureFunc: configureProvider,
	}
}

type ProviderContext struct {
	Client *Client
	TeamId string
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	teamId := d.Get("team_id").(string)
	client, err := NewClient(apiKey)
	if err != nil {
		return nil, err
	}

	return &ProviderContext{Client: client, TeamId: teamId}, nil
}

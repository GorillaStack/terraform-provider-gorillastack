package gorillastack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func tagGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tag_expression": {
			Type:     schema.TypeString,
			Required: true,
		},
		"team_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

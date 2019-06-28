package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func contextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aws": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: awsContextSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"azure": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: azureContextSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
	}
}

func awsContextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"platform": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"account_ids": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"regions": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"account_group_ids": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func azureContextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"platform": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subscription_ids": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			MinItems: 1,
			Optional: true,
		},
	}
}

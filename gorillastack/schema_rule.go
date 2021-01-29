package gorillastack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ruleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_id": {
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
		"team_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"user_group": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"labels": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			MinItems: 0,
			MaxItems: 100,
			Optional: true,
		},
		"context": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: contextSchema()},
			MinItems: 1,
			MaxItems: 1,
			Required: true,
		},
		"trigger": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: triggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Required: true,
		},
		"actions": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: actionsSchema()},
			MinItems: 1,
			MaxItems: 100,
			Required: true,
		},
	}
}

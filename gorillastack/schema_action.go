package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func actionsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delete_detached_volumes_action": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: deleteDetachedVolumesActionSchema()},
			Optional: true,
		},
		"delay_pause": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: delayPauseSchema()},
			Optional: true,
		},
	}
}

func deleteDetachedVolumesActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"dry_run": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"days_detached": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"tag_targeted": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"tag_groups": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 100,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag_group_combiner": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func delayPauseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"wait_duration": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

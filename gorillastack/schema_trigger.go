package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func triggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"schedule": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: scheduleTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"manual_trigger": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: manualTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
	}
}

func scheduleTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cron": {
			Type:     schema.TypeString,
			Required: true,
		},
		"timezone": {
			Type:     schema.TypeString,
			Required: true,
		},
		"notification_offset": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"default_snooze_duration": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"notifications": {
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: notificationSchema()},
		},
	}
}

func manualTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func notificationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"slack_webhook": {
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"room_id": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"email": {
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"send_to_team": {
						Type:     schema.TypeBool,
						Required: true,
					},
					"send_to_user_group": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"email_addresses": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
	}
}

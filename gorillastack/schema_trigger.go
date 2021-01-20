package gorillastack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func triggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloudtrail_event": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: cloudtrailEventTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"cost_threshold": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: costThresholdTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"detached_volumes_detected": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: detachedVolumesDetectedTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"incoming_webhook": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: incomingWebhookTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		// "manual": {
		// 	Type:     schema.TypeList,
		// 	Elem:     &schema.Resource{Schema: manualTriggerSchema()},
		// 	MinItems: 1,
		// 	MaxItems: 1,
		// 	Optional: true,
		// },
		"number_of_instances_threshold": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: numberOfInstancesThresholdTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"schedule": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: scheduleTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
		"sns_push": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: snsPushTriggerSchema()},
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
		},
	}
}

/* Shared subschemas */

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

/* Trigger Schemas */

// Supporting schemas for cloudtrailEventTriggerSchema
func cloudtrailEventMatchFieldsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"event_name": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			MaxItems: 100,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func cloudtrailEventMatchExpressionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"expression": {
			Type:     schema.TypeString,
			Required: true,
		},
		"expression_language": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func cloudtrailEventTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"match_fields": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 1,
			Required: true,
			Elem:     &schema.Resource{Schema: cloudtrailEventMatchFieldsSchema()},
		},
		"match_expression": {
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 1,
			Optional: true,
			Elem:     &schema.Resource{Schema: cloudtrailEventMatchExpressionSchema()},
		},
	}
}

func costThresholdTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"service": {
			Type:     schema.TypeString,
			Required: true,
		},
		"threshold": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"exceeded": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}

func detachedVolumesDetectedTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"days_detached": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func incomingWebhookTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"webhook_id": {
			Type:     schema.TypeString,
			Required: true,
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

func numberOfInstancesThresholdTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"threshold": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"exceeded": {
			Type:     schema.TypeBool,
			Computed: true,
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

func snsPushTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"trigger": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sns_subscription_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

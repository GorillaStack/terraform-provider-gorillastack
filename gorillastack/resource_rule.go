package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRuleCreate(d *schema.ResourceData, m interface{}) error {
	return resourceRuleRead(d, m)
}

func resourceRuleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRuleUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRuleRead(d, m)
}

func resourceRuleDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleCreate,
		Read:   resourceRuleRead,
		Update: resourceRuleUpdate,
		Delete: resourceRuleDelete,
		
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
			"slug": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type: schema.TypeInt,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type: schema.TypeInt,
				Computed: true,
			},
			"created_by": &schema.Schema{
				Type:			schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:			schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:			schema.TypeBool,
				Optional: true,
			},
			"user_group": &schema.Schema{
				Type:			schema.TypeString,
				Optional: true,
			},
			"labels": &schema.Schema{
				Type:			schema.TypeList,
				Elem      &schema.Schema{Type: schema.TypeString},
				MinItems: 0,
				MaxItems: 100,
				Optional: true,
			}
			"aws_context": &schema.Schema{
				Type: schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "platform": &schema.Schema{
							Type: schema.TypeString,
							Computed: true,
						},
						"account_ids": &schema.Schema{
							Type: schema.TypeList,
							MinItems: 1,
							Optional: true,
						},
						"regions": &schema.Schema{
							Type: schema.TypeList,
							MinItems: 1,
							Optional: true,
							Elem: &schema.Resource{Type: schema.TypeString},
						},
						"account_group_ids": &schema.Schema{
							Type: schema.TypeList,
							MinItems: 1,
							Optional: true,
							Elem: &schema.Resource{Type: schema.TypeString},
						},
					},
				},
			},
			"trigger": &schema.Schema{
				Required: true,
			},
			"actions": &schema.Schema{
				Type:			schema.TypeList,
				Elem      &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				MaxItems: 100,
				Required: true,
			},
		},
	}
}


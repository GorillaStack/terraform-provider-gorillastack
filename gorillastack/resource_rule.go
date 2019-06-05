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

		Schema: ruleSchema(),
	}
}

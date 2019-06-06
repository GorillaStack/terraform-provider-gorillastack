package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type CreateRuleInput struct {
	teamId	string
}

func constructRuleFromResourceData(d *schema.ResourceData) Rule {
	return Rule{
		_id: 				d.Get("_id").(string),
		teamId: 		d.Get("team_id").(string),
		name: 			d.Get("name").(string),
		slug:     	d.Get("slug").(string),
		enabled:   	d.Get("enabled").(bool),
	}
}

func resourceRuleCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	teamId := d.Get("team_id").(string)
	rule, err := client.CreateRule(teamId, constructRuleFromResourceData(d))

	if err != nil {
		return err
	}
	
	d.SetId(rule._id)
	return resourceRuleRead(d, m)
}

func resourceRuleRead(d *schema.ResourceData, m interface{}) error {
	teamId := d.Get("team_id").(string)
	ruleId := d.Id()
	client := m.(*Client)
	rule, err := client.GetRule(teamId, ruleId)

	if err != nil {
		return err
	}

	d.Set("name", rule.name)
	// TODO: set more attributes/whole rule

	return nil
}

func resourceRuleUpdate(d *schema.ResourceData, m interface{}) error {
	teamId := d.Get("team_id").(string)
	ruleId := d.Id()
	client := m.(*Client)
	_, err := client.UpdateRule(teamId, ruleId, constructRuleFromResourceData(d))

	if err != nil {
		return err
	}
	return resourceRuleRead(d, m)
}

func resourceRuleDelete(d *schema.ResourceData, m interface{}) error {
	teamId := d.Get("team_id").(string)
	ruleId := d.Id()
	client := m.(*Client)

	err := client.DeleteRule(teamId, ruleId)
	
	if err != nil {
		return err
	}

	d.SetId("")
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

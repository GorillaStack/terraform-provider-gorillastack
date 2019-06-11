package gorillastack

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
)

type CreateRuleInput struct {
	teamId string
}

func constructContextFromResourceData(d *schema.ResourceData) *Context {
	return nil
}

// func constructTriggerFromResourceData(d *schema.ResourceData) Context {
// }
// func constructActionsFromResourceData(d *schema.ResourceData) Context {
// }

func constructRuleFromResourceData(d *schema.ResourceData) *Rule {
	// context := constructContextFromResourceData(d)
	// trigger := constructTriggerFromResourceData(d)
	// actions := constructActionsFromResourceData(d)
	return &Rule{
		TeamId:  util.StringAddress(d.Get("team_id").(string)),
		Name:    util.StringAddress(d.Get("name").(string)),
		Slug:    util.StringAddress(d.Get("slug").(string)),
		Enabled: util.BoolAddress(d.Get("enabled").(bool)),
	}
}

func resourceRuleCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	teamId := d.Get("team_id").(string)
	log.Printf("[WARN] Attempting to create rule for team: %s", teamId)
	rule, err := client.CreateRule(teamId, constructRuleFromResourceData(d))

	if err != nil {
		return err
	}

	d.SetId(*rule.ID)
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

	d.Set("name", rule.Name)
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

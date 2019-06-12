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
	var context Context
	rawContext := d.Get("context").([]interface{})[0].(map[string]interface{})
	
	if rawAws := rawContext["aws"].([]interface{}); rawAws != nil {
		aws := rawAws[0].(map[string]interface{})
		context = Context{
			Platform: util.StringAddress("aws"),
		  AccountIds: &StringArrayOrNull{StringArray: util.StringArrayOrNil(aws["account_ids"].([]interface{}))},
		  Regions: &StringArrayOrNull{StringArray: util.StringArrayOrNil(aws["regions"].([]interface{}))},
			AccountGroupIds: util.StringArrayOrNil(aws["account_group_ids"].([]interface{})),
		}
	} else {
		if rawAzure := rawContext["azure"].([]interface{}); rawAzure != nil {
			azure := rawAzure[0].(map[string]interface{})
			context = Context{
				Platform: util.StringAddress("azure"),
				SubscriptionIds: &StringArrayOrNull{StringArray: util.StringArrayOrNil(azure["subscription_ids"].([]interface{}))},
			}
		} else {
			// TODO
			// not a supported platform
		}
	}

	return &context
}

func constructNotifications(rawNotifications []interface{}) []*Notification {
	return nil
}

func constructTriggerFromResourceData(d *schema.ResourceData) *Trigger {
	var trigger Trigger
	rawTrigger := d.Get("trigger").([]interface{})[0].(map[string]interface{})
	for k, rawV := range rawTrigger {
		triggerTypeArr := rawV.([]interface{})
		if len(triggerTypeArr) == 0 {
			continue
		}
		log.Printf("[WARN][GorillaStack][constructTriggerFromResourceData] %v", rawTrigger)
		v := triggerTypeArr[0].(map[string]interface{})
		switch k {
		case "schedule":
			trigger = Trigger{
				Trigger: 								util.StringAddress("schedule"),
				Cron:		 								util.StringAddress(v["cron"].(string)),
				Timezone:								util.StringAddress(v["timezone"].(string)),
				NotificationOffset:			util.IntAddress(v["notification_offset"].(int)),
				DefaultSnoozeDuration:	util.IntAddress(v["default_snooze_duration"].(int)),
				Notifications:					constructNotifications(v["notifications"].([]interface{})),
			}
		default:
			break
		}
	}

	return &trigger
}

func constructAction() *Action {
	return nil
}

func constructActionsFromResourceData(d *schema.ResourceData) []*Action {
	// rawActions := d.Get("trigger").([]interface{})[0].(map[string]interface{})
	rawActions := d.Get("actions").([]interface{})[0].(map[string]interface{})
	log.Printf("[WARN][GorillaStack][constructActionsFromResourceData] rawActions: %v", rawActions)
	log.Printf("[WARN][GorillaStack][constructActionsFromResourceData] len(rawActions) = %d", len(rawActions))
	actions := make([]*Action, len(rawActions))

	return actions
}

func constructRuleFromResourceData(d *schema.ResourceData) *Rule {
	return &Rule{
		TeamId:  util.StringAddress(d.Get("team_id").(string)),
		Name:    util.StringAddress(d.Get("name").(string)),
		Slug:    util.StringAddress(d.Get("slug").(string)),
    Enabled: util.BoolAddress(d.Get("enabled").(bool)),
		Labels:  util.ArrayOfStringPointers(d.Get("labels").([]interface{})),

		Context: constructContextFromResourceData(d),
		Trigger: constructTriggerFromResourceData(d),
		Actions: constructActionsFromResourceData(d),
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

	d.SetId(*rule.Id)
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

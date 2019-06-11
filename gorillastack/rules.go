package gorillastack

import (
	"log"

	"time"
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
)

type Rule struct {
	ID        *string
	Name      *string
	Slug      *string
	TeamId    *string
	Enabled   *bool
	CreatedBy *string
	UserGroup *string
	Labels    []*string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// context			*Context
	// trigger			*Trigger
	// actions			[]*Action
}

type RuleApiInput struct {
	Rule		*Rule
}

type RuleApiOutput struct {
	Rule		*Rule
}

func (r RuleApiInput) String() string {
	log.Printf("[WARN][GorillaStack] in RuleApiInput.String()")
	log.Printf("[WARN][GorillaStack] calling util.Prettify")
	return util.Prettify(r)
}

func (r RuleApiInput) GoString() string {
	return r.String()
}

func (r Rule) String() string {
	log.Printf("[WARN][GorillaStack] in Rule.String()")
	log.Printf("[WARN][GorillaStack] calling util.Prettify")
	return util.Prettify(r)
}

func (r Rule) GoString() string {
	return r.String()
}

/* START Rule Client Functions */
func (c *Client) ListRules() ([]*Rule, error) {
	req, err := c.newRequest("GET", "/rules", "")
	if err != nil {
		return nil, err
	}
	var rules []*Rule
	_, err = c.do(req, &rules)
	return rules, err
}

func (c *Client) GetRule(teamId string, ruleId string) (*Rule, error) {
	req, err := c.newRequest("GET", "/teams/" + teamId + "/rules/byId/" + ruleId, "")
	if err != nil {
		return nil, err
	}
	var response RuleApiOutput
	_, err = c.do(req, &response)
	return response.Rule, err
}

func (c *Client) CreateRule(teamId string, rule *Rule) (*Rule, error) {
	request := RuleApiInput{Rule: rule}
	log.Printf("[WARN][GorillaStack] request: %s", request)
	req, err := c.newRequest("POST", "/teams/" + teamId + "/rules", request)
	if err != nil {
		return nil, err
	}
	var response RuleApiOutput

	log.Printf("[WARN][GorillaStack] attempting to do")

	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}
	return response.Rule, err
}

func (c *Client) UpdateRule(teamId string, ruleId string, rule *Rule) (*Rule, error) {
	request := RuleApiInput{Rule: rule}
	req, err := c.newRequest("PUT", "/teams/" + teamId + "/rules/byId/" + ruleId, request)
	if err != nil {
		return nil, err
	}
	var response RuleApiOutput
	_, err = c.do(req, &response)
	return response.Rule, err
}

func (c *Client) DeleteRule(teamId string, ruleId string) error {
	req, err := c.newRequest("DELETE", "/teams/" + teamId + "/rules/byId/" + ruleId, "")

	if err != nil {
		return err
	}

	var result string
	_, err = c.do(req, &result)

	return nil
}

/* END Rule Client Functions */
package gorillastack

import (
	"time"
)

type Rule struct {
	_id				string
	name      string
	slug      string
	teamId    string
	enabled   bool
	createdBy string
	userGroup string
	labels    []string
	createdAt time.Time
	updatedAt time.Time
	// context			Context
	// trigger			Trigger
	// actions			[]Action
}

func (c *Client) ListRules() ([]Rule, error) {
	req, err := c.newRequest("GET", "/rules", nil)
	if err != nil {
		return nil, err
	}
	var rules []Rule
	_, err = c.do(req, &rules)
	return rules, err
}

type RulesWriteRequest struct {
	rule	Rule
}

type RulesWriteResponse struct {
	rule	Rule
}

func (c *Client) GetRule(teamId string, ruleId string) (Rule, error) {
	req, err := c.newRequest("GET", "/teams/" + teamId + "/rules/byId/" + ruleId, nil)
	if err != nil {
		return Rule{}, err
	}
	var response RulesWriteResponse
	_, err = c.do(req, &response)
	return response.rule, err
} 

func (c *Client) CreateRule(teamId string, rule Rule) (Rule, error) {
	request := RulesWriteRequest{rule: rule}
	req, err := c.newRequest("POST", "/teams/" + teamId + "/rules", request)
	if err != nil {
		return Rule{}, err
	}
	var response RulesWriteResponse
	_, err = c.do(req, &response)
	return response.rule, err
} 

func (c *Client) UpdateRule(teamId string, ruleId string, rule Rule) (Rule, error) {
	request := RulesWriteRequest{rule: rule}
	req, err := c.newRequest("PUT", "/teams/" + teamId + "/rules/byId/" + ruleId, request)
	if err != nil {
		return Rule{}, err
	}
	var response RulesWriteResponse
	_, err = c.do(req, &response)
	return response.rule, err
}

func (c *Client) DeleteRule(teamId string, ruleId string) error {
	req, err := c.newRequest("DELETE", "/teams/" + teamId + "/rules/byId/" + ruleId, nil)

	if err != nil {
		return err
	}

	var result string
	_, err = c.do(req, &result)

	return nil
}

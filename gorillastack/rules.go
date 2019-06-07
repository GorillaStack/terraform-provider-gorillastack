package gorillastack

import (
	"log"

	"encoding/json"
	"time"
)

type Rule struct {
	_id       string
	name      string
	slug      string
	teamId    string
	enabled   bool
	createdBy string
	userGroup string
	labels    []string
	createdAt time.Time
	updatedAt time.Time
	context			*Context
	// trigger			*Trigger
	// actions			[]*Action
}

/* START Rule Serialisation */
func (r *Rule) MarshalJSON() ([]byte, error) {
	rule := map[string]interface{}{}
	rule["_id"] = (*r)._id
	rule["name"] = (*r).name
	rule["slug"] = (*r).slug
	rule["teamId"] = (*r).teamId
	rule["enabled"] = (*r).enabled
	rule["createdBy"] = (*r).createdBy
	rule["userGroup"] = (*r).userGroup
	rule["labels"] = (*r).labels
	rule["createdAt"] = (*r).createdAt
	rule["updatedAt"] = (*r).updatedAt
	rule["context"], _ = (*r).context.MarshalJSON()
	// TODO - call trigger and actions .Marshal, 

	return json.Marshal(rule)
}

func (r *Rule) UnmarshalJSON(data []byte) error {
	var smap map[string]interface{}
	err := json.Unmarshal(data, &smap)
	if err != nil {
		return err
	}

	(*r)._id = smap["_id"].(string)
	(*r).name = smap["name"].(string)
	(*r).slug = smap["slug"].(string)
	(*r).teamId = smap["teamId"].(string)
	(*r).enabled = smap["enabled"] == "true"
	(*r).createdBy = smap["createdBy"].(string)
	(*r).userGroup = smap["userGroup"].(string)
	(*r).labels = smap["labels"].([]string)
	(*r).createdAt = smap["createdAt"].(time.Time)
	(*r).updatedAt = smap["updatedAt"].(time.Time)
	return nil
}
/* END Rule Serialisation */

/* START Rule Client Functions */
func (c *Client) ListRules() ([]Rule, error) {
	req, err := c.newRequest("GET", "/rules", "")
	if err != nil {
		return nil, err
	}
	var rules []Rule
	_, err = c.do(req, &rules)
	return rules, err
}

func (c *Client) GetRule(teamId string, ruleId string) (Rule, error) {
	req, err := c.newRequest("GET", "/teams/" + teamId + "/rules/byId/" + ruleId, "")
	if err != nil {
		return Rule{}, err
	}
	var response map[string]Rule
	_, err = c.do(req, &response)
	return response["rule"], err
}

func (c *Client) CreateRule(teamId string, rule Rule) (Rule, error) {
	ruleJSON, err := rule.MarshalJSON()
	if err != nil {
		return Rule{}, err
	}
	request := "{\"rule\": " + string(ruleJSON) + "}"
	log.Printf("[WARN][GorillaStack] request: %s", request)
	req, err := c.newRequest("POST", "/teams/" + teamId + "/rules", request)
	if err != nil {
		return Rule{}, err
	}
	var response map[string]string

	log.Printf("[WARN][GorillaStack] attempting to do")

	_, err = c.do(req, &response)
	result := Rule{}
	if err != nil {
		return Rule{}, err
	}
	err = result.UnmarshalJSON([]byte(response["rule"]))
	return result, err
}

func (c *Client) UpdateRule(teamId string, ruleId string, rule Rule) (Rule, error) {
	ruleJSON, err := rule.MarshalJSON()
	if err != nil {
		return Rule{}, err
	}
	request := "{\"rule\": " + string(ruleJSON) + "}"
	req, err := c.newRequest("PUT", "/teams/" + teamId + "/rules/byId/" + ruleId, request)
	if err != nil {
		return Rule{}, err
	}
	var response map[string]Rule
	_, err = c.do(req, &response)
	return response["rule"], err
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
package gorillastack

import (
	"time"
)

type UserService struct {
	Client 		*c
}

type Rule struct {
	name				string
	slug				string
	teamId  		string
	enabled			bool
	createdBy		string
	userGroup		string
	labels			[]string
	createdAt		Time
	updatedAt		Time
	context			Context
	trigger			Trigger
	actions			[]Action
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

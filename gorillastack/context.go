package gorillastack

import (
	"encoding/json"
)

type Context struct {
	platform 					string
	accountIds				[]string
	regions						[]string
	accountGroupIds 	[]string
	subscriptionIds		[]string
}

func (c *Context) MarshalJSON() ([]byte, error) {
	context := map[string]interface{}{}
	// TODO - logic to marshal
	context["platform"] = (*c).platform
	if context["platform"] == "aws" {
		context["accountIds"] = (*c).accountIds
		context["regions"] = (*c).regions
		context["accountGroupIds"] = (*c).accountGroupIds
	} else {
		context["subscriptionIds"] = (*c).subscriptionIds
	}

	return json.Marshal(context)
}

func (c *Context) UnmarshalJSON(data []byte) error {
	var smap map[string]interface{}
	err := json.Unmarshal(data, &smap)
	if err != nil {
		return err
	}

	(*c).platform = smap["platform"].(string)
	if (*c).platform == "aws" {
		(*c).accountIds = smap["accountIds"].([]string)
		(*c).regions = smap["regions"].([]string)
		(*c).accountGroupIds = smap["accountGroupIds"].([]string)
	} else {
		(*c).subscriptionIds = smap["subscriptionIds"].([]string)
	}

	return nil
}
package gorillastack

import (
	"encoding/json"
	"testing"

	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getRuleResourceData(t *testing.T, rawAwsContext map[string]interface{}) *schema.ResourceData {
	rawListOfAwsContexts := []interface{}{rawAwsContext}
	rawContext := map[string]interface{}{
		"aws": rawListOfAwsContexts,
	}
	rawListOfContexts := []interface{}{rawContext}
	rawRule := map[string]interface{}{"context": rawListOfContexts, "name": "testes"}
	return schema.TestResourceDataRaw(t, ruleSchema(), rawRule)
}

func TestConstructContextAllAccountsAllRegions(t *testing.T) {
	rawAwsContext := map[string]interface{}{
		"platform": "aws",
	}
	resourceData := getRuleResourceData(t, rawAwsContext)
	context := constructContext(resourceData)
	result := util.StringValue(context)
	expected := "{\"platform\":\"aws\",\"accountIds\":null,\"regions\":null}"

	if result != expected {
		t.Errorf("constructContextFromResourceData is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

// Want to be sure that we set an empty array for the accountIds rather than null
func TestConstructContextAccountGroupNoAccountsSelected(t *testing.T) {
	rawAwsContext := map[string]interface{}{
		"platform":          "aws",
		"account_group_ids": []string{"fake_account_group_id"},
	}
	resourceData := getRuleResourceData(t, rawAwsContext)
	context := constructContext(resourceData)
	result := util.StringValue(context)
	expected := "{\"platform\":\"aws\",\"accountIds\":[],\"regions\":null,\"accountGroupIds\":[\"fake_account_group_id\"]}"

	if result != expected {
		t.Errorf("constructContextFromResourceData is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

// Want to be sure that we send an array with the selected accounts
func TestConstructContextAccountGroupWithAccountsSelected(t *testing.T) {
	rawAwsContext := map[string]interface{}{
		"platform":          "aws",
		"account_ids":       []string{"123"},
		"account_group_ids": []string{"fake_account_group_id"},
	}
	resourceData := getRuleResourceData(t, rawAwsContext)
	context := constructContext(resourceData)
	result := util.StringValue(context)
	expected := "{\"platform\":\"aws\",\"accountIds\":[\"123\"],\"regions\":null,\"accountGroupIds\":[\"fake_account_group_id\"]}"

	if result != expected {
		t.Errorf("constructContextFromResourceData is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

// Test unmarshalling the JSON response to a StringArrayOrNull
func TestStringArrayOrNullUnmarshalJSONAccountIdsNull(t *testing.T) {
	var context Context
	if err := json.Unmarshal([]byte(`{ "platform": "aws", "accountIds": null, "regions": ["ap-southeast-2"] }`), &context); err != nil {
		t.Errorf("StringArrayOrNull.UnmarshalJSON is wrong. Caught error: %s", err)
	}
}

func TestStringArrayOrNullUnmarshalJSONAccountIdsEmptyArray(t *testing.T) {
	var context Context
	if err := json.Unmarshal([]byte(`{ "platform": "aws", "accountIds": [], "regions": ["ap-southeast-2"], "accountGroupIds": ["fake_account_group_id"] }`), &context); err != nil {
		t.Errorf("StringArrayOrNull.UnmarshalJSON is wrong. Caught error: %s", err)
	}
	result := util.StringValue(context)
	expected := "{\"platform\":\"aws\",\"accountIds\":[],\"regions\":[\"ap-southeast-2\"],\"accountGroupIds\":[\"fake_account_group_id\"]}"
	if result != expected {
		t.Errorf("StringArrayOrNull.UnmarshalJSON is wrong. Got: '%s' Expected: '%s'", result, expected)
	}
}

package gorillastack

import (
	"log"

	"time"
	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
)

// This type is for fields like accountIds and regions that can either be an array of strings or null
type StringArrayOrNull struct {
	StringArray 			[]*string
}

func (s StringArrayOrNull) String() string {
	if len(s.StringArray) == 0 {
		return "null"
	}
	return util.StringValue(s)
}

func (s StringArrayOrNull) GoString() string {
	return s.String()
}

type SlackNotificationConfig struct {
	RoomId			*string
}

type EmailNotificationConfig struct {
	SendToTeam		 		*bool
	SendToUserGroup		*bool
	EmailAddresses		[]*string 
}

type Notification struct {
	Slack			*SlackNotificationConfig
	Email			*EmailNotificationConfig
}

type Context struct {
	// Common fields
	Platform 					*string
	// AWS context fields
	AccountIds				*StringArrayOrNull
	Regions						*StringArrayOrNull
	AccountGroupIds 	[]*string
	// Azure context fields
	SubscriptionIds		*StringArrayOrNull
}

type Trigger struct {
	// Common fields
	Trigger 							*string
	Notifications					[]*Notification
	// schedule trigger fields
	Cron									*string
	Timezone							*string
	NotificationOffset		*int
	DefaultSnoozeDuration *int
}

type Wait struct {
	InstanceState		*bool
	InstanceStatus  *bool
	SystemStatus  	*bool
}

type AutoscalingParams struct {
	Min					*int
	Max					*int
	Desired			*int
}

type IntOrString struct {
	IVal						*int
	SVal						*string
}

func (s IntOrString) String() string {
	if len(*s.SVal) == 0 {
		return util.StringValue(s.IVal)
	}
	return util.StringValue(s.SVal)
}

func (s IntOrString) GoString() string {
	return s.String()
}

type SGMatchFields struct {
	Protocol							*IntOrString
	Port									*int
	Endpoint							*string
	EndpointDescription 	*string
}

type SGMatch struct {
	Type				*string
	Direction		*string
	Fields			*SGMatchFields
}

type SGChange struct {
	Operation		*string
}

type SGRuleChanges struct {
	Match				*SGMatch
	Change			*SGChange
}

type Action struct {
	// Common fields
	Action							*string
	ActionId						*string
	DryRun 							*bool
	TagTargeted					*bool
	TagGroups						[]*string
	TagGroupCombiner		[]*string
	// Copy (DB)?Snapshots
	Operator						*string
	Value								*int
  DestinationRegion		*string
	Mode								*string
 	KeyMapping					*map[string]string						
	UseDefaultKmsKey		*bool
	// Create DB/Images/Snapshots
	CopyDbInstanceTags	*bool
	MultiAzOnly					*bool
	AdditionalTags			*[]map[string]string
	NoReboot						*bool
	CopyVolumeTags			*bool
	CopyInstanceTags		*bool
	ExcludeBootVolume		*bool
	UseAdditionalTags		*bool
	// Delete detached volumes action
	DaysDetached				*int
	// Delete images/orphaned? snapshots
	KeepLatest					*bool
	KeepByVolume				*bool
	// EC2 Run Commands
	Commands						[]*string
	WorkingDirectory		*string
	ExecutionTimeout		*int
	// Invoke Lambdas
	FunctionName				*string
	InvocationType			*string
	Payload							*string
	ReplaceConflictingVars *bool
	EnvironmentVariables 	*[]map[string]string
	// Notify*
	Service							*string
	Notifications				[]*Notification
	// Release Disassociated IPs
	DaysDisassociated		*int
	// Suspend/Resume ASG Procs
	Processes						[]*string
	Wait								*Wait
	// Update ASG
	Params	 						*AutoscalingParams
	StoreExistingAsgSettings			*bool
	RestoreToPreviousAsgSettings	*bool
	IgnoreIfNoCachedAsgSettings		*bool
	// Update Dynamo
	ReadUnits						*int
	WriteUnits					*int
	// ECS Service Scale
	DesiredCount										*int
  StoredExistingDesiredCount			*bool
  RestoreToPreviousDesiredCount		*bool
	IgnoreIfNoCachedDesiredCount		*bool
	// Provisioned Iops
	Iops							*int
	// Update SecurityGroups
	RuleChanges				[]*SGRuleChanges
	// Update Scale Sets
	Capacity					*int
	// Delay pause schema
	WaitDuration			*int
}

type Rule struct {
	Id        *string
	Name      *string
	Slug      *string
	TeamId    *string
	Enabled   *bool
	CreatedBy *string
	UserGroup *string
	Labels    []*string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Context			*Context
	Trigger			*Trigger
	Actions			[]*Action
}

type RuleApiInput struct {
	Rule		*Rule
}

type RuleApiOutput struct {
	Rule		*Rule
}

func (r RuleApiInput) String() string {
	return util.StringValue(r)
}

func (r RuleApiInput) GoString() string {
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
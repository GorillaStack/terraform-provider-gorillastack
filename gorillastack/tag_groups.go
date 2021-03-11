package gorillastack

import (
	"log"
	"time"

	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
)

type TagGroup struct {
	Id            *string `json:"_id"`
	Name          *string
	TagExpression *string
	TeamId        *string
	Slug          *string
	CreatedBy     *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type TagGroupApiInput struct {
	TagGroup *TagGroup
}

type TagGroupApiOutput struct {
	TagGroup *TagGroup
}

func (tg TagGroupApiInput) String() string {
	return util.StringValue(tg)
}

func (tg TagGroupApiInput) GoString() string {
	return tg.String()
}

/* START TagGroup Client Functions */
func (c *Client) ListTagGroups(teamId string) ([]*TagGroup, error) {
	req, err := c.newRequest("GET", "/teams/"+teamId+"/tagGroups", "")
	if err != nil {
		return nil, err
	}
	var tagGroups []*TagGroup
	_, err = c.do(req, &tagGroups)
	return tagGroups, err
}

func (c *Client) GetTagGroup(teamId string, tagGroupId string) (*TagGroup, error) {
	req, err := c.newRequest("GET", "/teams/"+teamId+"/tagGroups/byId/"+tagGroupId, "")
	if err != nil {
		return nil, err
	}
	var response TagGroupApiOutput
	_, err = c.do(req, &response)
	return response.TagGroup, err
}

func (c *Client) CreateTagGroup(teamId string, tagGroup *TagGroup) (*TagGroup, error) {
	request := TagGroupApiInput{TagGroup: tagGroup}
	log.Printf("[WARN][GorillaStack] request: %s", request)
	req, err := c.newRequest("POST", "/teams/"+teamId+"/tagGroups", request)
	if err != nil {
		return nil, err
	}
	var response TagGroupApiOutput

	log.Printf("[WARN][GorillaStack] attempting to do")

	_, err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	log.Printf("[WARN][GorillaStack] response: %v", response)
	return response.TagGroup, err
}

func (c *Client) UpdateTagGroup(teamId string, tagGroupId string, tagGroup *TagGroup) (*TagGroup, error) {
	request := TagGroupApiInput{TagGroup: tagGroup}
	req, err := c.newRequest("PUT", "/teams/"+teamId+"/tagGroups/byId/"+tagGroupId, request)
	if err != nil {
		return nil, err
	}
	var response TagGroupApiOutput
	_, err = c.do(req, &response)
	return response.TagGroup, err
}

func (c *Client) DeleteTagGroup(teamId string, tagGroupId string) error {
	req, err := c.newRequest("DELETE", "/teams/"+teamId+"/tagGroups/byId/"+tagGroupId, "")

	if err != nil {
		return err
	}

	var result TagGroupApiOutput
	_, err = c.do(req, &result)

	if err != nil {
		return err
	}

	return nil
}

/* END TagGroup Client Functions */

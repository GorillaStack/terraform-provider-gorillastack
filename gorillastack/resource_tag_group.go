package gorillastack

import (
	"log"

	"github.com/gorillastack/terraform-provider-gorillastack/gorillastack/util"
	"github.com/hashicorp/terraform/helper/schema"
)

func constructTagGroupFromResourceData(d *schema.ResourceData) *TagGroup {
	return &TagGroup{
		Name:          util.StringAddress(d.Get("name").(string)),
		TagExpression: util.StringAddress(d.Get("tag_expression").(string)),
	}
}

func resourceTagGroupCreate(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	log.Printf("[WARN] Attempting to create TagGroup for team: %s", teamId)
	tagGroup, err := client.CreateTagGroup(teamId, constructTagGroupFromResourceData(d))

	if err != nil {
		return err
	}

	log.Printf("[WARN][GorillaStack][resourceTagGroupCreate] %+v", *tagGroup)

	d.SetId(*tagGroup.Id)
	return resourceTagGroupRead(d, m)
}

func resourceTagGroupRead(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	tagGroupId := d.Id()
	log.Printf("[WARN][GorillaStack][resourceTagGroupRead] %s", tagGroupId)
	tagGroup, err := client.GetTagGroup(teamId, tagGroupId)

	if err != nil {
		return err
	}

	d.Set("_id", tagGroupId)
	d.Set("name", tagGroup.Name)
	d.Set("slug", tagGroup.Slug)
	d.Set("team_id", tagGroup.TeamId)
	d.Set("created_by", tagGroup.CreatedBy)
	// d.Set("rootNode", tagGroup.RootNode)
	d.Set("created_at", tagGroup.CreatedAt)
	d.Set("updated_at", tagGroup.UpdatedAt)
	return nil
}

func resourceTagGroupUpdate(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	tagGroupId := d.Id()
	_, err := client.UpdateTagGroup(teamId, tagGroupId, constructTagGroupFromResourceData(d))

	if err != nil {
		return err
	}
	return resourceTagGroupRead(d, m)
}

func resourceTagGroupDelete(d *schema.ResourceData, m interface{}) error {
	providerContext := m.(*ProviderContext)
	client := providerContext.Client
	teamId := providerContext.TeamId
	tagGroupId := d.Id()

	err := client.DeleteTagGroup(teamId, tagGroupId)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceTagGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTagGroupCreate,
		Read:   resourceTagGroupRead,
		Update: resourceTagGroupUpdate,
		Delete: resourceTagGroupDelete,

		Schema: tagGroupSchema(),
	}
}

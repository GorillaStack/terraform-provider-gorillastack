package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTagGroupCreate(d *schema.ResourceData, m interface{}) error {
	return resourceTagGroupRead(d, m)
}

func resourceTagGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTagGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceTagGroupRead(d, m)
}

func resourceTagGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTagGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTagGroupCreate,
		Read:   resourceTagGroupRead,
		Update: resourceTagGroupUpdate,
		Delete: resourceTagGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

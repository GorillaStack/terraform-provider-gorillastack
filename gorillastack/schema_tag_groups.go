package gorillastack

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// func tagSchema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		"key": {
// 			Type:     schema.TypeString,
// 			Required: true,
// 		},
// 		"value": {
// 			Type:     schema.TypeString,
// 			Required: true,
// 		},
// 		"match_type": {
// 			Type:         schema.TypeString,
// 			Required:     true,
// 			ValidateFunc: validation.StringInSlice(constants.TagMatchTypes, false),
// 		},
// 	}
// }

// func tagGroupNodeSchema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		"tag": {
// 			Type:          schema.TypeList,
// 			Elem:          &schema.Resource{Schema: tagSchema()},
// 			MinItems:      1,
// 			ConfigMode:    schema.SchemaConfigModeAttr,
// 			ConflictsWith: []string{"and", "or", "not"},
// 		},
// 		"and": {
// 			Type:          schema.TypeList,
// 			Elem:          &schema.Resource{Schema: tagGroupNodeSchema()},
// 			MinItems:      1,
// 			ConflictsWith: []string{"tag"},
// 		},
// 		"or": {
// 			Type:          schema.TypeList,
// 			Elem:          &schema.Resource{Schema: tagGroupNodeSchema()},
// 			MinItems:      1,
// 			ConflictsWith: []string{"tag"},
// 		},
// 		"not": {
// 			Type:          schema.TypeList,
// 			Elem:          &schema.Resource{Schema: tagGroupNodeSchema()},
// 			MinItems:      1,
// 			ConflictsWith: []string{"tag"},
// 		},
// 	}
// }

func tagGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tag_expression": {
			Type:     schema.TypeString,
			Required: true,
		},
		// "root_node": {
		// 	Type:     schema.TypeList,
		// 	Elem:     &schema.Resource{Schema: tagGroupNodeSchema()},
		// 	MinItems: 1,
		// 	MaxItems: 1,
		// 	Required: true,
		// },
		"team_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

package dummy

import (
//	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/config"
	"github.com/hashicorp/terraform/helper/diff"
	"github.com/hashicorp/terraform/terraform"
)

func resource_dummy_record_create(
s *terraform.ResourceState,
	d *terraform.ResourceDiff,
	meta interface{}) (*terraform.ResourceState, error) {

	rs := s.MergeDiff(d)
	rs.ID = "1"
	log.Printf("[DEBUG] record create configuration")

	return rs, nil
}

func resource_dummy_record_update(
s *terraform.ResourceState,
	d *terraform.ResourceDiff,
	meta interface{}) (*terraform.ResourceState, error) {

	rs := s.MergeDiff(d)
	log.Printf("[DEBUG] record update configuration")
	return rs, nil
}

func resource_dummy_record_destroy(
s *terraform.ResourceState,
	meta interface{}) error {
	log.Printf("[INFO] Deleting record")

	return nil
}

func resource_dummy_record_refresh(
s *terraform.ResourceState,
	meta interface{}) (*terraform.ResourceState, error) {

	log.Printf("[INFO] refreshing record")

	return s, nil

}

func resource_dummy_record_diff(
s *terraform.ResourceState,
	c *terraform.ResourceConfig,
	meta interface{}) (*terraform.ResourceDiff, error) {

	b := &diff.ResourceBuilder{
		Attrs: map[string]diff.AttrType{
			"r_setting": diff.AttrTypeUpdate,
		},

		ComputedAttrsUpdate: []string{},
	}

	return b.Diff(s, c)
}


func resource_dummy_record_validation() *config.Validator {
	return &config.Validator{
		Required: []string{
			"r_setting",
		},
		Optional: []string{

		},
	}
}

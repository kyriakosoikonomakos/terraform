package dummy

import (
	"github.com/hashicorp/terraform/helper/resource"
)

// resourceMap is the mapping of resources we support to their basic
// operations. This makes it easy to implement new resource types.
var resourceMap *resource.Map

func init() {
	resourceMap = &resource.Map{
		Mapping: map[string]resource.Resource{
			"dummy_record": resource.Resource{
				ConfigValidator: resource_dummy_record_validation(),
				Create:          resource_dummy_record_create,
				Destroy:         resource_dummy_record_destroy,
				Diff:            resource_dummy_record_diff,
				Update:          resource_dummy_record_update,
				Refresh:         resource_dummy_record_refresh,
			},
		},
	}
}

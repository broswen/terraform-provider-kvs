package kvs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KVS_HOST", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kvs_pair": resourcePair(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"kvs_pair": dataSourcepair(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type ProviderConfig struct {
	Host string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var host *string

	hostVal, ok := d.GetOk("host")
	if ok {
		tempHost := hostVal.(string)
		host = &tempHost
	}

	var diags diag.Diagnostics

	return ProviderConfig{
		Host: *host,
	}, diags
}

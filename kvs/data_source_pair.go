package kvs

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcepair() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePairRead,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConfig := m.(ProviderConfig)
	var diags diag.Diagnostics

	key := d.Get("key").(string)
	resp, err := http.Get(fmt.Sprintf("%s/%s", providerConfig.Host, key))
	if err != nil {
		return diag.FromErr(err)
	}
	if resp.StatusCode == http.StatusNotFound {
		d.SetId("")
	}

	defer resp.Body.Close()
	value, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("value", string(value)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("key", key); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(key)

	return diags
}

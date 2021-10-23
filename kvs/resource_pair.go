package kvs

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePairCreate,
		ReadContext:   resourcePairRead,
		UpdateContext: resourcePairUpdate,
		DeleteContext: resourcePairDelete,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePairCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConfig := m.(ProviderConfig)
	var diags diag.Diagnostics

	key := d.Get("key").(string)
	value := d.Get("value").(string)
	_, err := http.Post(fmt.Sprintf("%s/%s", providerConfig.Host, key), "text/plain", strings.NewReader(value))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(key)

	resourcePairRead(ctx, d, m)

	return diags
}

func resourcePairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConfig := m.(ProviderConfig)
	var diags diag.Diagnostics

	key := d.Get("key").(string)
	resp, err := http.Get(fmt.Sprintf("%s/%s", providerConfig.Host, key))
	if err != nil {
		return diag.FromErr(err)
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

	return diags
}

func resourcePairUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConfig := m.(ProviderConfig)
	var diags diag.Diagnostics

	key := d.Get("key").(string)
	value := d.Get("value").(string)
	_, err := http.Post(fmt.Sprintf("%s/%s", providerConfig.Host, key), "text/plain", strings.NewReader(value))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(key)

	resourcePairRead(ctx, d, m)

	return diags
}

func resourcePairDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	providerConfig := m.(ProviderConfig)
	var diags diag.Diagnostics

	key := d.Get("key").(string)
	_, err := http.Post(fmt.Sprintf("%s/%s", providerConfig.Host, key), "text/plain", nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	resourcePairRead(ctx, d, m)

	return diags
}

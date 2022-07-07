package flintlock

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	flintlockv1 "github.com/weaveworks-liquidmetal/flintlock/api/services/microvm/v1alpha1"
)

func resourceMicrovm() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMicrovmCreate,
		ReadContext:   resourceMicrovmRead,
		//UpdateContext: nil,
		DeleteContext: resourceMicrovmDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "default",
				Optional: true,
				ForceNew: true,
			},
			"vcpu": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  2,
				Optional: true,
				ForceNew: true,
			},
			"memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  2048,
				Optional: true,
				ForceNew: true,
			},
			"kernel_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kernel_filename": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "vmlinux",
				Optional: true,
				ForceNew: true,
			},
			"root_volume_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMicrovmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*fc.)
	c := m.(flintlockv1.MicroVMClient)

	fmt.Println(c)

	//TODO convert and call create

	var diags diag.Diagnostics

	return diags
}

func resourceMicrovmRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*fc.)
	c := m.(flintlockv1.MicroVMClient)

	var diags diag.Diagnostics

	vmID := d.Id()

	resp, err := c.GetMicroVM(ctx, &flintlockv1.GetMicroVMRequest{
		Uid: vmID,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if resp.Microvm != nil {
		// TODO set the properties of d
		//d.Set()
	}

	return diags
}

func resourceMicrovmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*fc.)
	c := m.(flintlockv1.MicroVMClient)

	var diags diag.Diagnostics

	vmID := d.Id()

	_, err := c.DeleteMicroVM(ctx, &flintlockv1.DeleteMicroVMRequest{
		Uid: vmID,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

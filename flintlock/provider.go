package flintlock

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	flintlockv1 "github.com/weaveworks-liquidmetal/flintlock/api/services/microvm/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hosts": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FLINTLOCT_HOSTS", nil),
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				// ValidateFunc: , //TODO: add validation to ensure its a host:port
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"flintlock_microvm": resourceMicrovm(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}

}

func createFlintlockClient(address string) (flintlockv1.MicroVMClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("creating grpc connection: %w", err)
	}

	flClient := flintlockv1.NewMicroVMClient(conn)

	return flClient, nil
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	fmt.Printf("in provider config")

	//TODO: create the grpc client and return this
	return nil, diags
}

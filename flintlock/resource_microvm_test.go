package flintlock

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	flintlockv1 "github.com/weaveworks-liquidmetal/flintlock/api/services/microvm/v1alpha1"
)

func TestCreateMicrovmBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMicrovmDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMicrovmConfigBasic("vm1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVMExists(""),
				),
			},
		},
	})
}

func testAccCheckMicrovmDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(flintlockv1.MicroVMClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "hashicups_order" {
			continue
		}

		vmID := rs.Primary.ID

		_, err := c.DeleteMicroVM(context.TODO(), &flintlockv1.DeleteMicroVMRequest{
			Uid: vmID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccMicrovmConfigBasic(name string) string {
	return fmt.Sprintf(`
	resource "flintlock_microvm" "new" {
		name = %s
		kernel_image = "myimagehere"
    	root_volume_image = "myrootimage"
	}
	`, name)
}

func testAccCheckVMExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No OrderID set")
		}

		return nil
	}
}

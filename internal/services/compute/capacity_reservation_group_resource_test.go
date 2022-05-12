package compute_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/compute/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type CapacityReservationGroupResource struct{}

func TestAccCapacityReservationGroup_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_capacity_reservation_group", "test")
	r := CapacityReservationGroupResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccCapacityReservationGroup_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_capacity_reservation_group", "test")
	r := CapacityReservationGroupResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r CapacityReservationGroupResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.CapacityReservationGroupID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.Compute.CapacityReservationGroupsClient.Get(ctx, id.ResourceGroup, id.Name, "")
	if err != nil {
		return nil, fmt.Errorf("retrieving Compute Capacity Reservation Group %q", id)
	}

	return utils.Bool(resp.ID != nil), nil
}

func (CapacityReservationGroupResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-compute-%[1]d"
  location = "%s"
}

resource "azurerm_capacity_reservation_group" "test" {
  name                = "acctestCRG-compute-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}
`, data.RandomInteger, data.Locations.Primary)
}

func (CapacityReservationGroupResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-compute-%[1]d"
  location = "%s"
}

resource "azurerm_capacity_reservation_group" "test" {
  name                = "acctestCRG-compute-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  zones               = ["1", "2", "3"]
  tags = {
    key = "value1"
  }
}
`, data.RandomInteger, data.Locations.Primary)
}

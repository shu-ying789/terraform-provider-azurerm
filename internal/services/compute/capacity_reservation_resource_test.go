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

type CapacityReservationResource struct{}

func TestAccCapacityReservation_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_capacity_reservation", "test")
	r := CapacityReservationResource{}

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

func TestAccCapacityReservationResource_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_capacity_reservation", "test")
	r := CapacityReservationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccCapacityReservationResource_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_capacity_reservation", "test")
	r := CapacityReservationResource{}

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

func (r CapacityReservationResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := parse.CapacityReservationID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.Compute.CapacityReservationsClient.Get(ctx, id.ResourceGroup, id.CapacityReservationGroupName, id.Name, "")
	if err != nil {
		return nil, fmt.Errorf("retrieving %q: %+v", id, err)
	}

	return utils.Bool(resp.ID != nil), nil
}

func (r CapacityReservationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_capacity_reservation" "test" {
  name                = "acctestCR-compute-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}
`, r.template(data), data.RandomInteger)
}

func (r CapacityReservationResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_capacity_reservation" "import" {
  resource_group_name = azurerm_capacity_reservation.test.resource_group_name
  name                = azurerm_capacity_reservation.test.name
  location            = azurerm_capacity_reservation.test.location
}
`, r.basic(data))
}

func (r CapacityReservationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_capacity_reservation" "test" {
  name                = "acctestCR-compute-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  //zones               = ["1"]
  //sku = {
  //    capacity = 1
  //    name     = "Standard_D2s_v3"
  //  }
}
`, r.template(data), data.RandomInteger)
}

func (r CapacityReservationResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-compute-%[1]d"
  location = "%[2]s"
}
`, data.RandomInteger, data.Locations.Primary)
}

package compute

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-11-01/compute"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/compute/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/compute/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceCapacityReservation() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceCapacityReservationCreateUpdate,
		Read:   resourceCapacityReservationRead,
		Update: resourceCapacityReservationCreateUpdate,
		Delete: resourceCapacityReservationDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.CapacityReservationID(id)
			return err
		}),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"capacity_reservation_group_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.CapacityReservationGroupID,
			},

			"location": azure.SchemaLocation(),

			"sku": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"name": {
							Type:     pluginsdk.TypeString,
							Required: true,
						},

						"tier": {
							Type:     pluginsdk.TypeString,
							Optional: true,
							Computed: true,
						},

						"capacity": {
							Type:     pluginsdk.TypeInt,
							Required: true,
						},
					},
				},
			},

			"zones": azure.SchemaZones(),

			"tags": tags.Schema(),
		},
	}
}

func resourceCapacityReservationCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.CapacityReservationsClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	capacityId, err := parse.CapacityReservationGroupID(d.Get("capacity_reservation_group_id").(string))
	if err != nil {
		return err
	}
	id := parse.NewCapacityReservationID(subscriptionId, capacityId.ResourceGroup, d.Get("name").(string), "")

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.CapacityReservationGroupName, id.Name, "")
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of %s: %+v", id, err)
			}
		}
		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_capacity_reservation", id.ID())
		}
	}

	parameters := compute.CapacityReservation{
		Location: utils.String(azure.NormalizeLocation(d.Get("location").(string))),
		Sku:      expandCapacityReservationSku(d.Get("sku").([]interface{})),
		Tags:     tags.Expand(d.Get("tags").(map[string]interface{})),
	}

	if v, ok := d.GetOk("zones"); ok {
		parameters.Zones = utils.ExpandStringSlice(v.([]interface{}))
	}

	if _, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.CapacityReservationGroupName, id.Name, parameters); err != nil {
		return fmt.Errorf("creating %s: %+v", id, err)
	}

	d.SetId(id.ID())

	return resourceCapacityReservationRead(d, meta)
}

func resourceCapacityReservationRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.CapacityReservationsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.CapacityReservationID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.CapacityReservationGroupName, id.Name, "")
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO]  Capacity Reservation %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("reading %q: %+v", id, err)
	}

	d.Set("name", id.Name)
	d.Set("capacity_reservation_group_id", parse.NewCapacityReservationGroupID(id.SubscriptionId, id.ResourceGroup, id.Name).ID())
	d.Set("location", location.NormalizeNilable(resp.Location))

	d.Set("zones", resp.Zones)

	if err := d.Set("sku", flattenCapacityReservationSku(resp.Sku)); err != nil {
		return fmt.Errorf("setting `sku`: %+v", err)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceCapacityReservationDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.CapacityReservationsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.CapacityReservationID(d.Id())
	if err != nil {
		return err
	}

	if _, err := client.Delete(ctx, id.ResourceGroup, id.CapacityReservationGroupName, id.Name); err != nil {
		return fmt.Errorf("deleting %q: %+v", id, err)
	}

	return nil
}

func expandCapacityReservationSku(input []interface{}) *compute.Sku {
	if len(input) == 0 {
		return nil
	}

	raw := input[0].(map[string]interface{})
	return &compute.Sku{
		Tier:     utils.String(raw["tier"].(string)),
		Name:     utils.String(raw["name"].(string)),
		Capacity: utils.Int64(int64(raw["capacity"].(int))),
	}
}

func flattenCapacityReservationSku(input *compute.Sku) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	capacity := 0
	if input.Capacity != nil {
		capacity = int(*input.Capacity)
	}

	return []interface{}{
		map[string]interface{}{
			"capacity": capacity,
			"name":     input.Name,
			"tier":     input.Tier,
		},
	}
}

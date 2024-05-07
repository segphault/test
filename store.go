// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petstorefix

import (
	"context"
	"net/http"

	"github.com/segphault/test/internal/apijson"
	"github.com/segphault/test/internal/requestconfig"
	"github.com/segphault/test/option"
	"github.com/segphault/test/shared"
)

// StoreService contains methods and other services that help with interacting with
// the petstore API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewStoreService] method instead.
type StoreService struct {
	Options []option.RequestOption
	Order   *StoreOrderService
}

// NewStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoreService(opts ...option.RequestOption) (r *StoreService) {
	r = &StoreService{}
	r.Options = opts
	r.Order = NewStoreOrderService(opts...)
	return
}

// Place a new order in the store
func (r *StoreService) NewOrder(ctx context.Context, body StoreNewOrderParams, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "store/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a map of status codes to quantities
func (r *StoreService) Inventory(ctx context.Context, opts ...option.RequestOption) (res *StoreInventoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "store/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StoreInventoryResponse map[string]int64

type StoreNewOrderParams struct {
	Order shared.OrderParam `json:"order,required"`
}

func (r StoreNewOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Order)
}

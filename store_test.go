// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petstorefix_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/segphault/test"
	"github.com/segphault/test/internal/testutil"
	"github.com/segphault/test/option"
	"github.com/segphault/test/shared"
)

func TestStoreNewOrderWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := petstorefix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.NewOrder(context.TODO(), petstorefix.StoreNewOrderParams{
		Order: shared.OrderParam{
			ID:       petstorefix.F(int64(10)),
			PetID:    petstorefix.F(int64(198772)),
			Quantity: petstorefix.F(int64(7)),
			ShipDate: petstorefix.F(time.Now()),
			Status:   petstorefix.F(shared.OrderStatusApproved),
			Complete: petstorefix.F(true),
		},
	})
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreInventory(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := petstorefix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.Inventory(context.TODO())
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

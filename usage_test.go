// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petstorefix_test

import (
	"context"
	"os"
	"testing"

	"github.com/segphault/test"
	"github.com/segphault/test/internal/testutil"
	"github.com/segphault/test/option"
	"github.com/segphault/test/shared"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Store.NewOrder(context.TODO(), petstorefix.StoreNewOrderParams{
		Order: shared.OrderParam{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", order.ID)
}

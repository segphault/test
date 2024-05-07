// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petstorefix_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/segphault/test"
	"github.com/segphault/test/internal/testutil"
	"github.com/segphault/test/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), petstorefix.PetNewParams{
		Pet: petstorefix.PetParam{
			ID:   petstorefix.F(int64(10)),
			Name: petstorefix.F("doggie"),
			Category: petstorefix.F(petstorefix.PetCategoryParam{
				ID:   petstorefix.F(int64(1)),
				Name: petstorefix.F("Dogs"),
			}),
			PhotoURLs: petstorefix.F([]string{"string", "string", "string"}),
			Tags: petstorefix.F([]petstorefix.PetTagParam{{
				ID:   petstorefix.F(int64(0)),
				Name: petstorefix.F("string"),
			}, {
				ID:   petstorefix.F(int64(0)),
				Name: petstorefix.F("string"),
			}, {
				ID:   petstorefix.F(int64(0)),
				Name: petstorefix.F("string"),
			}}),
			Status: petstorefix.F(petstorefix.PetStatusAvailable),
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), petstorefix.PetUpdateParams{
		Pet: petstorefix.PetParam{
			ID:   petstorefix.F(int64(10)),
			Name: petstorefix.F("doggie"),
			Category: petstorefix.F(petstorefix.PetCategoryParam{
				ID:   petstorefix.F(int64(1)),
				Name: petstorefix.F("Dogs"),
			}),
			PhotoURLs: petstorefix.F([]string{"string", "string", "string"}),
			Tags: petstorefix.F([]petstorefix.PetTagParam{{
				ID:   petstorefix.F(int64(0)),
				Name: petstorefix.F("string"),
			}, {
				ID:   petstorefix.F(int64(0)),
				Name: petstorefix.F("string"),
			}, {
				ID:   petstorefix.F(int64(0)),
				Name: petstorefix.F("string"),
			}}),
			Status: petstorefix.F(petstorefix.PetStatusAvailable),
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), petstorefix.PetFindByStatusParams{
		Status: petstorefix.F(petstorefix.PetFindByStatusParamsStatusAvailable),
	})
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), petstorefix.PetFindByTagsParams{
		Tags: petstorefix.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		int64(0),
		petstorefix.PetUpdateByIDParams{
			Name:   petstorefix.F("string"),
			Status: petstorefix.F("string"),
		},
	)
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		int64(0),
		petstorefix.PetUploadImageParams{
			AdditionalMetadata: petstorefix.F("string"),
		},
	)
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

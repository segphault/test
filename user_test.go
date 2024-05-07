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

func TestUserNewWithOptionalParams(t *testing.T) {
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
	err := client.User.New(context.TODO(), petstorefix.UserNewParams{
		User: petstorefix.UserParam{
			ID:         petstorefix.F(int64(10)),
			Username:   petstorefix.F("theUser"),
			FirstName:  petstorefix.F("John"),
			LastName:   petstorefix.F("James"),
			Email:      petstorefix.F("john@email.com"),
			Password:   petstorefix.F("12345"),
			Phone:      petstorefix.F("12345"),
			UserStatus: petstorefix.F(int64(1)),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "string")
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.User.Update(
		context.TODO(),
		"string",
		petstorefix.UserUpdateParams{
			User: petstorefix.UserParam{
				ID:         petstorefix.F(int64(10)),
				Username:   petstorefix.F("theUser"),
				FirstName:  petstorefix.F("John"),
				LastName:   petstorefix.F("James"),
				Email:      petstorefix.F("john@email.com"),
				Password:   petstorefix.F("12345"),
				Phone:      petstorefix.F("12345"),
				UserStatus: petstorefix.F(int64(1)),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.User.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithList(t *testing.T) {
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
	_, err := client.User.NewWithList(context.TODO(), petstorefix.UserNewWithListParams{
		Items: []petstorefix.UserParam{{
			ID:         petstorefix.F(int64(10)),
			Username:   petstorefix.F("theUser"),
			FirstName:  petstorefix.F("John"),
			LastName:   petstorefix.F("James"),
			Email:      petstorefix.F("john@email.com"),
			Password:   petstorefix.F("12345"),
			Phone:      petstorefix.F("12345"),
			UserStatus: petstorefix.F(int64(1)),
		}, {
			ID:         petstorefix.F(int64(10)),
			Username:   petstorefix.F("theUser"),
			FirstName:  petstorefix.F("John"),
			LastName:   petstorefix.F("James"),
			Email:      petstorefix.F("john@email.com"),
			Password:   petstorefix.F("12345"),
			Phone:      petstorefix.F("12345"),
			UserStatus: petstorefix.F(int64(1)),
		}, {
			ID:         petstorefix.F(int64(10)),
			Username:   petstorefix.F("theUser"),
			FirstName:  petstorefix.F("John"),
			LastName:   petstorefix.F("James"),
			Email:      petstorefix.F("john@email.com"),
			Password:   petstorefix.F("12345"),
			Phone:      petstorefix.F("12345"),
			UserStatus: petstorefix.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Login(context.TODO(), petstorefix.UserLoginParams{
		Password: petstorefix.F("string"),
		Username: petstorefix.F("string"),
	})
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *petstorefix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petstorefix

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/segphault/test/internal/apijson"
	"github.com/segphault/test/internal/apiquery"
	"github.com/segphault/test/internal/param"
	"github.com/segphault/test/internal/requestconfig"
	"github.com/segphault/test/option"
)

// PetService contains methods and other services that help with interacting with
// the petstore API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPetService] method instead.
type PetService struct {
	Options []option.RequestOption
}

// NewPetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPetService(opts ...option.RequestOption) (r *PetService) {
	r = &PetService{}
	r.Options = opts
	return
}

// Add a new pet to the store
func (r *PetService) New(ctx context.Context, body PetNewParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a single pet
func (r *PetService) Get(ctx context.Context, petID int64, opts ...option.RequestOption) (res *Pet, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing pet by Id
func (r *PetService) Update(ctx context.Context, body PetUpdateParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// delete a pet
func (r *PetService) Delete(ctx context.Context, petID int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Multiple status values can be provided with comma separated strings
func (r *PetService) FindByStatus(ctx context.Context, query PetFindByStatusParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet/findByStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3
// for testing.
func (r *PetService) FindByTags(ctx context.Context, query PetFindByTagsParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = append(r.Options[:], opts...)
	path := "pet/findByTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates a pet in the store with form data
func (r *PetService) UpdateByID(ctx context.Context, petID int64, body PetUpdateByIDParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// uploads an image
func (r *PetService) UploadImage(ctx context.Context, petID int64, body PetUploadImageParams, opts ...option.RequestOption) (res *APIResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("pet/%v/uploadImage", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponse struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Type    string          `json:"type"`
	JSON    apiResponseJSON `json:"-"`
}

// apiResponseJSON contains the JSON metadata for the struct [APIResponse]
type apiResponseJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseJSON) RawJSON() string {
	return r.raw
}

type Pet struct {
	Name      string      `json:"name,required"`
	PhotoURLs []string    `json:"photoUrls,required"`
	ID        int64       `json:"id"`
	Category  PetCategory `json:"category"`
	// pet status in the store
	Status PetStatus `json:"status"`
	Tags   []PetTag  `json:"tags"`
	JSON   petJSON   `json:"-"`
}

// petJSON contains the JSON metadata for the struct [Pet]
type petJSON struct {
	Name        apijson.Field
	PhotoURLs   apijson.Field
	ID          apijson.Field
	Category    apijson.Field
	Status      apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petJSON) RawJSON() string {
	return r.raw
}

type PetCategory struct {
	ID   int64           `json:"id"`
	Name string          `json:"name"`
	JSON petCategoryJSON `json:"-"`
}

// petCategoryJSON contains the JSON metadata for the struct [PetCategory]
type petCategoryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petCategoryJSON) RawJSON() string {
	return r.raw
}

// pet status in the store
type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

func (r PetStatus) IsKnown() bool {
	switch r {
	case PetStatusAvailable, PetStatusPending, PetStatusSold:
		return true
	}
	return false
}

type PetTag struct {
	ID   int64      `json:"id"`
	Name string     `json:"name"`
	JSON petTagJSON `json:"-"`
}

// petTagJSON contains the JSON metadata for the struct [PetTag]
type petTagJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PetTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r petTagJSON) RawJSON() string {
	return r.raw
}

type PetParam struct {
	Name      param.Field[string]           `json:"name,required"`
	PhotoURLs param.Field[[]string]         `json:"photoUrls,required"`
	ID        param.Field[int64]            `json:"id"`
	Category  param.Field[PetCategoryParam] `json:"category"`
	// pet status in the store
	Status param.Field[PetStatus]     `json:"status"`
	Tags   param.Field[[]PetTagParam] `json:"tags"`
}

func (r PetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetCategoryParam struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetTagParam struct {
	ID   param.Field[int64]  `json:"id"`
	Name param.Field[string] `json:"name"`
}

func (r PetTagParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PetNewParams struct {
	Pet PetParam `json:"pet,required"`
}

func (r PetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Pet)
}

type PetUpdateParams struct {
	Pet PetParam `json:"pet,required"`
}

func (r PetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Pet)
}

type PetFindByStatusParams struct {
	// Status values that need to be considered for filter
	Status param.Field[PetFindByStatusParamsStatus] `query:"status"`
}

// URLQuery serializes [PetFindByStatusParams]'s query parameters as `url.Values`.
func (r PetFindByStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status values that need to be considered for filter
type PetFindByStatusParamsStatus string

const (
	PetFindByStatusParamsStatusAvailable PetFindByStatusParamsStatus = "available"
	PetFindByStatusParamsStatusPending   PetFindByStatusParamsStatus = "pending"
	PetFindByStatusParamsStatusSold      PetFindByStatusParamsStatus = "sold"
)

func (r PetFindByStatusParamsStatus) IsKnown() bool {
	switch r {
	case PetFindByStatusParamsStatusAvailable, PetFindByStatusParamsStatusPending, PetFindByStatusParamsStatusSold:
		return true
	}
	return false
}

type PetFindByTagsParams struct {
	// Tags to filter by
	Tags param.Field[[]string] `query:"tags"`
}

// URLQuery serializes [PetFindByTagsParams]'s query parameters as `url.Values`.
func (r PetFindByTagsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUpdateByIDParams struct {
	// Name of pet that needs to be updated
	Name param.Field[string] `query:"name"`
	// Status of pet that needs to be updated
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [PetUpdateByIDParams]'s query parameters as `url.Values`.
func (r PetUpdateByIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUploadImageParams struct {
	// Additional Metadata
	AdditionalMetadata param.Field[string] `query:"additionalMetadata"`
}

// URLQuery serializes [PetUploadImageParams]'s query parameters as `url.Values`.
func (r PetUploadImageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

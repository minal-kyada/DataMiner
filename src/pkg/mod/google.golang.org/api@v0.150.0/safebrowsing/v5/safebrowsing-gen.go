// Copyright 2023 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package safebrowsing provides access to the Safe Browsing API.
//
// For product documentation, see: https://developers.google.com/safe-browsing/
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/safebrowsing/v5"
//	...
//	ctx := context.Background()
//	safebrowsingService, err := safebrowsing.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	safebrowsingService, err := safebrowsing.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	safebrowsingService, err := safebrowsing.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package safebrowsing // import "google.golang.org/api/safebrowsing/v5"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "safebrowsing:v5"
const apiName = "safebrowsing"
const apiVersion = "v5"
const basePath = "https://safebrowsing.googleapis.com/"
const mtlsBasePath = "https://safebrowsing.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Hashes = NewHashesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Hashes *HashesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewHashesService(s *Service) *HashesService {
	rs := &HashesService{s: s}
	return rs
}

type HashesService struct {
	s *Service
}

// GoogleSecuritySafebrowsingV5FullHash: The full hash identified with
// one or more matches.
type GoogleSecuritySafebrowsingV5FullHash struct {
	// FullHash: The matching full hash. This is the SHA256 hash. The length
	// will be exactly 32 bytes.
	FullHash string `json:"fullHash,omitempty"`

	// FullHashDetails: Unordered list. A repeated field identifying the
	// details relevant to this full hash.
	FullHashDetails []*GoogleSecuritySafebrowsingV5FullHashFullHashDetail `json:"fullHashDetails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FullHash") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FullHash") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleSecuritySafebrowsingV5FullHash) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleSecuritySafebrowsingV5FullHash
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleSecuritySafebrowsingV5FullHashFullHashDetail: Details about a
// matching full hash. An important note about forward compatibility:
// new threat types and threat attributes may be added by the server at
// any time; those additions are considered minor version changes. It is
// Google's policy not to expose minor version numbers in APIs (see
// https://cloud.google.com/apis/design/versioning), so clients MUST be
// prepared to receive FullHashDetail messages containing ThreatType
// enum values or ThreatAttribute enum values that are considered
// invalid by the client. Therefore, it is the client's responsibility
// to check for the validity of all ThreatType and ThreatAttribute enum
// values; if any value is considered invalid, the client MUST disregard
// the entire FullHashDetail message.
type GoogleSecuritySafebrowsingV5FullHashFullHashDetail struct {
	// Attributes: Unordered list. Additional attributes about those full
	// hashes. This may be empty.
	//
	// Possible values:
	//   "THREAT_ATTRIBUTE_UNSPECIFIED" - Unknown.
	//   "CANARY" - Indicates that the threat_type should not be used for
	// enforcement.
	//   "FRAME_ONLY" - Indicates that the threat_type should only be used
	// for enforcement on frames.
	Attributes []string `json:"attributes,omitempty"`

	// ThreatType: The type of threat. This field will never be empty.
	//
	// Possible values:
	//   "THREAT_TYPE_UNSPECIFIED" - Unknown.
	//   "MALWARE" - Malware threat type.
	//   "SOCIAL_ENGINEERING" - Social engineering threat type.
	//   "UNWANTED_SOFTWARE" - Unwanted software threat type.
	//   "POTENTIALLY_HARMFUL_APPLICATION" - Potentially harmful application
	// threat type.
	ThreatType string `json:"threatType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleSecuritySafebrowsingV5FullHashFullHashDetail) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleSecuritySafebrowsingV5FullHashFullHashDetail
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleSecuritySafebrowsingV5SearchHashesResponse: The response
// returned after searching threat hashes. Note that if nothing is
// found, the server will return an OK status (HTTP status code 200)
// with the `full_hashes` field empty, rather than returning a NOT_FOUND
// status (HTTP status code 404).
type GoogleSecuritySafebrowsingV5SearchHashesResponse struct {
	// CacheDuration: The client-side cache duration. The client shall add
	// this duration to the current time to determine the expiration time.
	// The expiration time then applies to every hash prefix queried by the
	// client in the request, regardless of how many full hashes are
	// returned in the response. Even if the server returns no full hashes
	// for a particular hash prefix, this fact should also be cached by the
	// client. Important: the client must not assume that the server will
	// return the same cache duration for all responses. The server may
	// choose different cache durations for different responses depending on
	// the situation.
	CacheDuration string `json:"cacheDuration,omitempty"`

	// FullHashes: Unordered list. The unordered list of full hashes found.
	FullHashes []*GoogleSecuritySafebrowsingV5FullHash `json:"fullHashes,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CacheDuration") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CacheDuration") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleSecuritySafebrowsingV5SearchHashesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleSecuritySafebrowsingV5SearchHashesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "safebrowsing.hashes.search":

type HashesSearchCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Search: Search for full hashes matching the specified prefixes. This
// is a custom method as described by guidance at
// https://google.aip.dev/136
func (r *HashesService) Search() *HashesSearchCall {
	c := &HashesSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// HashPrefixes sets the optional parameter "hashPrefixes": Required.
// The hash prefixes to be looked up.
func (c *HashesSearchCall) HashPrefixes(hashPrefixes ...string) *HashesSearchCall {
	c.urlParams_.SetMulti("hashPrefixes", append([]string{}, hashPrefixes...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HashesSearchCall) Fields(s ...googleapi.Field) *HashesSearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *HashesSearchCall) IfNoneMatch(entityTag string) *HashesSearchCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *HashesSearchCall) Context(ctx context.Context) *HashesSearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *HashesSearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *HashesSearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/"+internal.Version)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v5/hashes:search")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "safebrowsing.hashes.search" call.
// Exactly one of *GoogleSecuritySafebrowsingV5SearchHashesResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleSecuritySafebrowsingV5SearchHashesResponse.ServerResponse.Heade
// r or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *HashesSearchCall) Do(opts ...googleapi.CallOption) (*GoogleSecuritySafebrowsingV5SearchHashesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &GoogleSecuritySafebrowsingV5SearchHashesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Search for full hashes matching the specified prefixes. This is a custom method as described by guidance at https://google.aip.dev/136",
	//   "flatPath": "v5/hashes:search",
	//   "httpMethod": "GET",
	//   "id": "safebrowsing.hashes.search",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "hashPrefixes": {
	//       "description": "Required. The hash prefixes to be looked up.",
	//       "format": "byte",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v5/hashes:search",
	//   "response": {
	//     "$ref": "GoogleSecuritySafebrowsingV5SearchHashesResponse"
	//   }
	// }

}

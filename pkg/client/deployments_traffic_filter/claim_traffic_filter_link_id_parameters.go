// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewClaimTrafficFilterLinkIDParams creates a new ClaimTrafficFilterLinkIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClaimTrafficFilterLinkIDParams() *ClaimTrafficFilterLinkIDParams {
	return &ClaimTrafficFilterLinkIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClaimTrafficFilterLinkIDParamsWithTimeout creates a new ClaimTrafficFilterLinkIDParams object
// with the ability to set a timeout on a request.
func NewClaimTrafficFilterLinkIDParamsWithTimeout(timeout time.Duration) *ClaimTrafficFilterLinkIDParams {
	return &ClaimTrafficFilterLinkIDParams{
		timeout: timeout,
	}
}

// NewClaimTrafficFilterLinkIDParamsWithContext creates a new ClaimTrafficFilterLinkIDParams object
// with the ability to set a context for a request.
func NewClaimTrafficFilterLinkIDParamsWithContext(ctx context.Context) *ClaimTrafficFilterLinkIDParams {
	return &ClaimTrafficFilterLinkIDParams{
		Context: ctx,
	}
}

// NewClaimTrafficFilterLinkIDParamsWithHTTPClient creates a new ClaimTrafficFilterLinkIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewClaimTrafficFilterLinkIDParamsWithHTTPClient(client *http.Client) *ClaimTrafficFilterLinkIDParams {
	return &ClaimTrafficFilterLinkIDParams{
		HTTPClient: client,
	}
}

/*
ClaimTrafficFilterLinkIDParams contains all the parameters to send to the API endpoint

	for the claim traffic filter link id operation.

	Typically these are written to a http.Request.
*/
type ClaimTrafficFilterLinkIDParams struct {

	/* Body.

	   The specification for traffic filter claimed link id.
	*/
	Body *models.TrafficFilterClaimedLinkIDRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the claim traffic filter link id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClaimTrafficFilterLinkIDParams) WithDefaults() *ClaimTrafficFilterLinkIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the claim traffic filter link id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClaimTrafficFilterLinkIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) WithTimeout(timeout time.Duration) *ClaimTrafficFilterLinkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) WithContext(ctx context.Context) *ClaimTrafficFilterLinkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) WithHTTPClient(client *http.Client) *ClaimTrafficFilterLinkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) WithBody(body *models.TrafficFilterClaimedLinkIDRequest) *ClaimTrafficFilterLinkIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the claim traffic filter link id params
func (o *ClaimTrafficFilterLinkIDParams) SetBody(body *models.TrafficFilterClaimedLinkIDRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ClaimTrafficFilterLinkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CancelDeploymentResourcePendingPlanReader is a Reader for the CancelDeploymentResourcePendingPlan structure.
type CancelDeploymentResourcePendingPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelDeploymentResourcePendingPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelDeploymentResourcePendingPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelDeploymentResourcePendingPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelDeploymentResourcePendingPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelDeploymentResourcePendingPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelDeploymentResourcePendingPlanOK creates a CancelDeploymentResourcePendingPlanOK with default headers values
func NewCancelDeploymentResourcePendingPlanOK() *CancelDeploymentResourcePendingPlanOK {
	return &CancelDeploymentResourcePendingPlanOK{}
}

/*
CancelDeploymentResourcePendingPlanOK describes a response with status code 200, with default header values.

Standard Deployment Resource Crud Response
*/
type CancelDeploymentResourcePendingPlanOK struct {
	Payload *models.DeploymentResourceCrudResponse
}

// IsSuccess returns true when this cancel deployment resource pending plan o k response has a 2xx status code
func (o *CancelDeploymentResourcePendingPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel deployment resource pending plan o k response has a 3xx status code
func (o *CancelDeploymentResourcePendingPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel deployment resource pending plan o k response has a 4xx status code
func (o *CancelDeploymentResourcePendingPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel deployment resource pending plan o k response has a 5xx status code
func (o *CancelDeploymentResourcePendingPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel deployment resource pending plan o k response a status code equal to that given
func (o *CancelDeploymentResourcePendingPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cancel deployment resource pending plan o k response
func (o *CancelDeploymentResourcePendingPlanOK) Code() int {
	return 200
}

func (o *CancelDeploymentResourcePendingPlanOK) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanOK  %+v", 200, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanOK) String() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanOK  %+v", 200, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanOK) GetPayload() *models.DeploymentResourceCrudResponse {
	return o.Payload
}

func (o *CancelDeploymentResourcePendingPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDeploymentResourcePendingPlanBadRequest creates a CancelDeploymentResourcePendingPlanBadRequest with default headers values
func NewCancelDeploymentResourcePendingPlanBadRequest() *CancelDeploymentResourcePendingPlanBadRequest {
	return &CancelDeploymentResourcePendingPlanBadRequest{}
}

/*
CancelDeploymentResourcePendingPlanBadRequest describes a response with status code 400, with default header values.

The Resource does not have a pending plan. (code: `deployments.resource_does_not_have_a_pending_plan`)
*/
type CancelDeploymentResourcePendingPlanBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this cancel deployment resource pending plan bad request response has a 2xx status code
func (o *CancelDeploymentResourcePendingPlanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel deployment resource pending plan bad request response has a 3xx status code
func (o *CancelDeploymentResourcePendingPlanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel deployment resource pending plan bad request response has a 4xx status code
func (o *CancelDeploymentResourcePendingPlanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel deployment resource pending plan bad request response has a 5xx status code
func (o *CancelDeploymentResourcePendingPlanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel deployment resource pending plan bad request response a status code equal to that given
func (o *CancelDeploymentResourcePendingPlanBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cancel deployment resource pending plan bad request response
func (o *CancelDeploymentResourcePendingPlanBadRequest) Code() int {
	return 400
}

func (o *CancelDeploymentResourcePendingPlanBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanBadRequest) String() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanBadRequest  %+v", 400, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CancelDeploymentResourcePendingPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDeploymentResourcePendingPlanNotFound creates a CancelDeploymentResourcePendingPlanNotFound with default headers values
func NewCancelDeploymentResourcePendingPlanNotFound() *CancelDeploymentResourcePendingPlanNotFound {
	return &CancelDeploymentResourcePendingPlanNotFound{}
}

/*
CancelDeploymentResourcePendingPlanNotFound describes a response with status code 404, with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type CancelDeploymentResourcePendingPlanNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this cancel deployment resource pending plan not found response has a 2xx status code
func (o *CancelDeploymentResourcePendingPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel deployment resource pending plan not found response has a 3xx status code
func (o *CancelDeploymentResourcePendingPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel deployment resource pending plan not found response has a 4xx status code
func (o *CancelDeploymentResourcePendingPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel deployment resource pending plan not found response has a 5xx status code
func (o *CancelDeploymentResourcePendingPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel deployment resource pending plan not found response a status code equal to that given
func (o *CancelDeploymentResourcePendingPlanNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cancel deployment resource pending plan not found response
func (o *CancelDeploymentResourcePendingPlanNotFound) Code() int {
	return 404
}

func (o *CancelDeploymentResourcePendingPlanNotFound) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanNotFound  %+v", 404, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanNotFound) String() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanNotFound  %+v", 404, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CancelDeploymentResourcePendingPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDeploymentResourcePendingPlanInternalServerError creates a CancelDeploymentResourcePendingPlanInternalServerError with default headers values
func NewCancelDeploymentResourcePendingPlanInternalServerError() *CancelDeploymentResourcePendingPlanInternalServerError {
	return &CancelDeploymentResourcePendingPlanInternalServerError{}
}

/*
CancelDeploymentResourcePendingPlanInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type CancelDeploymentResourcePendingPlanInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this cancel deployment resource pending plan internal server error response has a 2xx status code
func (o *CancelDeploymentResourcePendingPlanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel deployment resource pending plan internal server error response has a 3xx status code
func (o *CancelDeploymentResourcePendingPlanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel deployment resource pending plan internal server error response has a 4xx status code
func (o *CancelDeploymentResourcePendingPlanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel deployment resource pending plan internal server error response has a 5xx status code
func (o *CancelDeploymentResourcePendingPlanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel deployment resource pending plan internal server error response a status code equal to that given
func (o *CancelDeploymentResourcePendingPlanInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cancel deployment resource pending plan internal server error response
func (o *CancelDeploymentResourcePendingPlanInternalServerError) Code() int {
	return 500
}

func (o *CancelDeploymentResourcePendingPlanInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{resource_kind}/{ref_id}/plan/pending][%d] cancelDeploymentResourcePendingPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelDeploymentResourcePendingPlanInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CancelDeploymentResourcePendingPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

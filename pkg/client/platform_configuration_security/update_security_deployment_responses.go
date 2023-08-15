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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateSecurityDeploymentReader is a Reader for the UpdateSecurityDeployment structure.
type UpdateSecurityDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecurityDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSecurityDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSecurityDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSecurityDeploymentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSecurityDeploymentOK creates a UpdateSecurityDeploymentOK with default headers values
func NewUpdateSecurityDeploymentOK() *UpdateSecurityDeploymentOK {
	return &UpdateSecurityDeploymentOK{}
}

/*
UpdateSecurityDeploymentOK describes a response with status code 200, with default header values.

The security deployment was successfully updated
*/
type UpdateSecurityDeploymentOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.IDResponse
}

// IsSuccess returns true when this update security deployment o k response has a 2xx status code
func (o *UpdateSecurityDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update security deployment o k response has a 3xx status code
func (o *UpdateSecurityDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update security deployment o k response has a 4xx status code
func (o *UpdateSecurityDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update security deployment o k response has a 5xx status code
func (o *UpdateSecurityDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update security deployment o k response a status code equal to that given
func (o *UpdateSecurityDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update security deployment o k response
func (o *UpdateSecurityDeploymentOK) Code() int {
	return 200
}

func (o *UpdateSecurityDeploymentOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityDeploymentOK) String() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityDeploymentOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *UpdateSecurityDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-resource-created
	hdrXCloudResourceCreated := response.GetHeader("x-cloud-resource-created")

	if hdrXCloudResourceCreated != "" {
		o.XCloudResourceCreated = hdrXCloudResourceCreated
	}

	// hydrates response header x-cloud-resource-last-modified
	hdrXCloudResourceLastModified := response.GetHeader("x-cloud-resource-last-modified")

	if hdrXCloudResourceLastModified != "" {
		o.XCloudResourceLastModified = hdrXCloudResourceLastModified
	}

	// hydrates response header x-cloud-resource-version
	hdrXCloudResourceVersion := response.GetHeader("x-cloud-resource-version")

	if hdrXCloudResourceVersion != "" {
		o.XCloudResourceVersion = hdrXCloudResourceVersion
	}

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityDeploymentNotFound creates a UpdateSecurityDeploymentNotFound with default headers values
func NewUpdateSecurityDeploymentNotFound() *UpdateSecurityDeploymentNotFound {
	return &UpdateSecurityDeploymentNotFound{}
}

/*
UpdateSecurityDeploymentNotFound describes a response with status code 404, with default header values.

The security deployment was not found. (code: `security_deployment.not_found`)
*/
type UpdateSecurityDeploymentNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this update security deployment not found response has a 2xx status code
func (o *UpdateSecurityDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update security deployment not found response has a 3xx status code
func (o *UpdateSecurityDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update security deployment not found response has a 4xx status code
func (o *UpdateSecurityDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update security deployment not found response has a 5xx status code
func (o *UpdateSecurityDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update security deployment not found response a status code equal to that given
func (o *UpdateSecurityDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update security deployment not found response
func (o *UpdateSecurityDeploymentNotFound) Code() int {
	return 404
}

func (o *UpdateSecurityDeploymentNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSecurityDeploymentNotFound) String() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSecurityDeploymentNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSecurityDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSecurityDeploymentConflict creates a UpdateSecurityDeploymentConflict with default headers values
func NewUpdateSecurityDeploymentConflict() *UpdateSecurityDeploymentConflict {
	return &UpdateSecurityDeploymentConflict{}
}

/*
	UpdateSecurityDeploymentConflict describes a response with status code 409, with default header values.

	* There is a version conflict. (code: `security_deployment.version_conflict`)

* There is a version conflict. (code: `security_deployment.already_exists`)
*/
type UpdateSecurityDeploymentConflict struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this update security deployment conflict response has a 2xx status code
func (o *UpdateSecurityDeploymentConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update security deployment conflict response has a 3xx status code
func (o *UpdateSecurityDeploymentConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update security deployment conflict response has a 4xx status code
func (o *UpdateSecurityDeploymentConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update security deployment conflict response has a 5xx status code
func (o *UpdateSecurityDeploymentConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update security deployment conflict response a status code equal to that given
func (o *UpdateSecurityDeploymentConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update security deployment conflict response
func (o *UpdateSecurityDeploymentConflict) Code() int {
	return 409
}

func (o *UpdateSecurityDeploymentConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentConflict  %+v", 409, o.Payload)
}

func (o *UpdateSecurityDeploymentConflict) String() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentConflict  %+v", 409, o.Payload)
}

func (o *UpdateSecurityDeploymentConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSecurityDeploymentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

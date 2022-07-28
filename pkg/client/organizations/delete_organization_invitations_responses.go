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

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteOrganizationInvitationsReader is a Reader for the DeleteOrganizationInvitations structure.
type DeleteOrganizationInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrganizationInvitationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationInvitationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationInvitationsOK creates a DeleteOrganizationInvitationsOK with default headers values
func NewDeleteOrganizationInvitationsOK() *DeleteOrganizationInvitationsOK {
	return &DeleteOrganizationInvitationsOK{}
}

/* DeleteOrganizationInvitationsOK describes a response with status code 200, with default header values.

Organization invitations deleted successfully
*/
type DeleteOrganizationInvitationsOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteOrganizationInvitationsOK) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/invitations/{invitation_tokens}][%d] deleteOrganizationInvitationsOK  %+v", 200, o.Payload)
}
func (o *DeleteOrganizationInvitationsOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteOrganizationInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationInvitationsBadRequest creates a DeleteOrganizationInvitationsBadRequest with default headers values
func NewDeleteOrganizationInvitationsBadRequest() *DeleteOrganizationInvitationsBadRequest {
	return &DeleteOrganizationInvitationsBadRequest{}
}

/* DeleteOrganizationInvitationsBadRequest describes a response with status code 400, with default header values.

No valid invitation token was supplied. (code: `root.invalid_data`)
*/
type DeleteOrganizationInvitationsBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteOrganizationInvitationsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/invitations/{invitation_tokens}][%d] deleteOrganizationInvitationsBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteOrganizationInvitationsBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteOrganizationInvitationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrganizationInvitationsNotFound creates a DeleteOrganizationInvitationsNotFound with default headers values
func NewDeleteOrganizationInvitationsNotFound() *DeleteOrganizationInvitationsNotFound {
	return &DeleteOrganizationInvitationsNotFound{}
}

/* DeleteOrganizationInvitationsNotFound describes a response with status code 404, with default header values.

 * Organization not found. (code: `organization.not_found`)
* Invitation not found. (code: `organization.invitation_not_found`)
*/
type DeleteOrganizationInvitationsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteOrganizationInvitationsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization_id}/invitations/{invitation_tokens}][%d] deleteOrganizationInvitationsNotFound  %+v", 404, o.Payload)
}
func (o *DeleteOrganizationInvitationsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteOrganizationInvitationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
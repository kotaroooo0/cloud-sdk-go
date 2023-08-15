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

// SetTLSCertificateReader is a Reader for the SetTLSCertificate structure.
type SetTLSCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetTLSCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSetTLSCertificateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetTLSCertificateAccepted creates a SetTLSCertificateAccepted with default headers values
func NewSetTLSCertificateAccepted() *SetTLSCertificateAccepted {
	return &SetTLSCertificateAccepted{}
}

/*
SetTLSCertificateAccepted describes a response with status code 202, with default header values.

The TLS update has been accepted for the given service and will take effect throughout the system
*/
type SetTLSCertificateAccepted struct {
	Payload *models.UpdatedTLSChain
}

// IsSuccess returns true when this set Tls certificate accepted response has a 2xx status code
func (o *SetTLSCertificateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set Tls certificate accepted response has a 3xx status code
func (o *SetTLSCertificateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set Tls certificate accepted response has a 4xx status code
func (o *SetTLSCertificateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this set Tls certificate accepted response has a 5xx status code
func (o *SetTLSCertificateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this set Tls certificate accepted response a status code equal to that given
func (o *SetTLSCertificateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the set Tls certificate accepted response
func (o *SetTLSCertificateAccepted) Code() int {
	return 202
}

func (o *SetTLSCertificateAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/tls/{service_name}][%d] setTlsCertificateAccepted  %+v", 202, o.Payload)
}

func (o *SetTLSCertificateAccepted) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/tls/{service_name}][%d] setTlsCertificateAccepted  %+v", 202, o.Payload)
}

func (o *SetTLSCertificateAccepted) GetPayload() *models.UpdatedTLSChain {
	return o.Payload
}

func (o *SetTLSCertificateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatedTLSChain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

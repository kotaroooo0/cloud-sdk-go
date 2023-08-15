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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterInstanceMemoryInfo Information about the specific instances memory capacity and its usage
//
// swagger:model ClusterInstanceMemoryInfo
type ClusterInstanceMemoryInfo struct {

	// The memory capacity in MB of the instance
	// Required: true
	InstanceCapacity *int32 `json:"instance_capacity"`

	// The planned memory capacity in MB of the instance (only shown when an override is present)
	InstanceCapacityPlanned int32 `json:"instance_capacity_planned,omitempty"`

	// The % memory pressure of Elasticsearch JVM heap space if available (60-75% consider increasing capacity, >75% can incur significant performance and stability issues)
	MemoryPressure int32 `json:"memory_pressure,omitempty"`

	// The % memory pressure of the instance Docker container (if available)
	NativeMemoryPressure int32 `json:"native_memory_pressure,omitempty"`
}

// Validate validates this cluster instance memory info
func (m *ClusterInstanceMemoryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceCapacity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterInstanceMemoryInfo) validateInstanceCapacity(formats strfmt.Registry) error {

	if err := validate.Required("instance_capacity", "body", m.InstanceCapacity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster instance memory info based on context it is used
func (m *ClusterInstanceMemoryInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterInstanceMemoryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterInstanceMemoryInfo) UnmarshalBinary(b []byte) error {
	var res ClusterInstanceMemoryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

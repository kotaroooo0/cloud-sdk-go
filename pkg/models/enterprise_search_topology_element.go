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
)

// EnterpriseSearchTopologyElement Defines the topology of the Enterprise Search nodes (e.g. number/capacity of nodes, and where they can be allocated)
//
// swagger:model EnterpriseSearchTopologyElement
type EnterpriseSearchTopologyElement struct {

	// allocator filter
	AllocatorFilter interface{} `json:"allocator_filter,omitempty"`

	// enterprise search
	EnterpriseSearch *EnterpriseSearchConfiguration `json:"enterprise_search,omitempty"`

	// Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the id of an existing instance configuration.
	InstanceConfigurationID string `json:"instance_configuration_id,omitempty"`

	// The version of the Instance Configuration Id. If it is unset, the meaning depends on read vs writes. For deployment reads, it is equivalent to version 0 (or the IC is unversioned); for deployment creates and deployment template use, it is equivalent to 'the latest version'; and for deployment updates, it is equivalent to 'retain the current version'.
	InstanceConfigurationVersion int32 `json:"instance_configuration_version,omitempty"`

	// memory per node
	MemoryPerNode interface{} `json:"memory_per_node,omitempty"`

	// node configuration
	NodeConfiguration string `json:"node_configuration,omitempty"`

	// node count per zone
	NodeCountPerZone interface{} `json:"node_count_per_zone,omitempty"`

	// Defines the EnterpriseSearch node type
	NodeType *EnterpriseSearchNodeTypes `json:"node_type,omitempty"`

	// size
	Size *TopologySize `json:"size,omitempty"`

	// number of zones in which nodes will be placed
	ZoneCount int32 `json:"zone_count,omitempty"`
}

// Validate validates this enterprise search topology element
func (m *EnterpriseSearchTopologyElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnterpriseSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchTopologyElement) validateEnterpriseSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.EnterpriseSearch) { // not required
		return nil
	}

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchTopologyElement) validateNodeType(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeType) { // not required
		return nil
	}

	if m.NodeType != nil {
		if err := m.NodeType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_type")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchTopologyElement) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this enterprise search topology element based on the context it is used
func (m *EnterpriseSearchTopologyElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnterpriseSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchTopologyElement) contextValidateEnterpriseSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchTopologyElement) contextValidateNodeType(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeType != nil {
		if err := m.NodeType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_type")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchTopologyElement) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnterpriseSearchTopologyElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnterpriseSearchTopologyElement) UnmarshalBinary(b []byte) error {
	var res EnterpriseSearchTopologyElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

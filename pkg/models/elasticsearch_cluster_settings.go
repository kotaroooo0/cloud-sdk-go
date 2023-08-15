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

// ElasticsearchClusterSettings The settings for an Elasticsearch cluster.
//
// swagger:model ElasticsearchClusterSettings
type ElasticsearchClusterSettings struct {

	// The curation settings for this deployment. When provided, curation settings are changed as specified. A `null` value reverts the field to the default value. Otherwise, all curation settings remain as they were set previously.
	Curation *ClusterCurationSettings `json:"curation,omitempty"`

	// Threshold starting from which the number of instances in the cluster results in the introduction of dedicated masters. If the cluster is downscaled to a number of nodes below this one, dedicated masters will be removed. Limit is inclusive. When provided the threshold setting is updated. A `null` value removes the field. Otherwise, the setting remains as it was set previously.
	DedicatedMastersThreshold int32 `json:"dedicated_masters_threshold,omitempty"`

	// The contents of the Elasticsearch keystore. It's a write only field.
	KeystoreContents *KeystoreContents `json:"keystore_contents,omitempty"`

	// metadata
	Metadata *ClusterMetadataSettings `json:"metadata,omitempty"`

	// The monitoring settings for this deployment. When provided, monitoring settings are changed as specified. A `null` value reverts the field to the default value. Otherwise, all monitoring settings remain as they were set previously.
	Monitoring *ManagedMonitoringSettings `json:"monitoring,omitempty"`

	// The snapshot settings for this deployment. When provided, snapshot settings are changed as specified. A `null` value reverts the field to the default value. Otherwise, all snapshot settings remain as they were set previously.
	Snapshot *ClusterSnapshotSettings `json:"snapshot,omitempty"`

	// The rulesets to apply to all resources in this cluster. When specified, the set of rulesets is updated and the same rulesets will be applied to Kibana and APM clusters as well. If not specified, the rulesets remain as they were set previously.
	TrafficFilter *TrafficFilterSettings `json:"traffic_filter,omitempty"`

	// Configuration of trust with other clusters. When provided, trust settings are changed as specified. A `null` value reverts the field to the default value. Otherwise, all trust settings remain as they were set previously.
	Trust *ElasticsearchClusterTrustSettings `json:"trust,omitempty"`
}

// Validate validates this elasticsearch cluster settings
func (m *ElasticsearchClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeystoreContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoring(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrust(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterSettings) validateCuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Curation) { // not required
		return nil
	}

	if m.Curation != nil {
		if err := m.Curation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("curation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("curation")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateKeystoreContents(formats strfmt.Registry) error {
	if swag.IsZero(m.KeystoreContents) { // not required
		return nil
	}

	if m.KeystoreContents != nil {
		if err := m.KeystoreContents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keystore_contents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keystore_contents")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateMonitoring(formats strfmt.Registry) error {
	if swag.IsZero(m.Monitoring) { // not required
		return nil
	}

	if m.Monitoring != nil {
		if err := m.Monitoring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoring")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitoring")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateTrafficFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.TrafficFilter) { // not required
		return nil
	}

	if m.TrafficFilter != nil {
		if err := m.TrafficFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traffic_filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("traffic_filter")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateTrust(formats strfmt.Registry) error {
	if swag.IsZero(m.Trust) { // not required
		return nil
	}

	if m.Trust != nil {
		if err := m.Trust.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trust")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this elasticsearch cluster settings based on the context it is used
func (m *ElasticsearchClusterSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeystoreContents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitoring(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrafficFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrust(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateCuration(ctx context.Context, formats strfmt.Registry) error {

	if m.Curation != nil {
		if err := m.Curation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("curation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("curation")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateKeystoreContents(ctx context.Context, formats strfmt.Registry) error {

	if m.KeystoreContents != nil {
		if err := m.KeystoreContents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keystore_contents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keystore_contents")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateMonitoring(ctx context.Context, formats strfmt.Registry) error {

	if m.Monitoring != nil {
		if err := m.Monitoring.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoring")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitoring")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.Snapshot != nil {
		if err := m.Snapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateTrafficFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.TrafficFilter != nil {
		if err := m.TrafficFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traffic_filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("traffic_filter")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) contextValidateTrust(ctx context.Context, formats strfmt.Registry) error {

	if m.Trust != nil {
		if err := m.Trust.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trust")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterSettings) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

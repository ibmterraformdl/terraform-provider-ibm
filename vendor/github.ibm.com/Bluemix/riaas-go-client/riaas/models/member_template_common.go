// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MemberTemplateCommon MemberTemplateCommon
// swagger:model MemberTemplateCommon
type MemberTemplateCommon struct {

	// The port number of the application running in the server member.
	// Maximum: 65535
	// Minimum: 1
	Port int64 `json:"port,omitempty"`

	// target
	Target *MemberTemplateCommonTarget `json:"target,omitempty"`

	// Weight of the server member. This takes effect only when the load balancing algorithm of its belonging pool is `weighted_round_robin`.
	// Maximum: 100
	// Minimum: 1
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this member template common
func (m *MemberTemplateCommon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberTemplateCommon) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *MemberTemplateCommon) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *MemberTemplateCommon) validateWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", int64(m.Weight), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", int64(m.Weight), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberTemplateCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberTemplateCommon) UnmarshalBinary(b []byte) error {
	var res MemberTemplateCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
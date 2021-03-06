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

// PatchDocument A JSONPatch document as defined by RFC 6902
//
// swagger:model patchDocument
type PatchDocument struct {

	// A JSON-pointer
	From string `json:"from,omitempty"`

	// The operation to be performed
	// Example: \"replace\
	// Required: true
	Op *string `json:"op"`

	// A JSON-pointer
	// Example: \"/name\
	// Required: true
	Path *string `json:"path"`

	// The value to be used within the operations
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this patch document
func (m *PatchDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchDocument) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	return nil
}

func (m *PatchDocument) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch document based on context it is used
func (m *PatchDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchDocument) UnmarshalBinary(b []byte) error {
	var res PatchDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRuntimeEnvCredentialRequest openpitrix modify runtime env credential request
// swagger:model openpitrixModifyRuntimeEnvCredentialRequest
type OpenpitrixModifyRuntimeEnvCredentialRequest struct {

	// content
	Content map[string]string `json:"content,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// runtime env credential id
	RuntimeEnvCredentialID string `json:"runtime_env_credential_id,omitempty"`
}

// Validate validates this openpitrix modify runtime env credential request
func (m *OpenpitrixModifyRuntimeEnvCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRuntimeEnvCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRuntimeEnvCredentialRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRuntimeEnvCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
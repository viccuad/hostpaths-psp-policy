// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// SecretKeySelector SecretKeySelector selects a key of a Secret.
//
// swagger:model SecretKeySelector
type SecretKeySelector struct {

	// The key of the secret to select from.  Must be a valid secret key.
	// Required: true
	Key *string `json:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Specify whether the Secret or its key must be defined
	Optional bool `json:"optional,omitempty"`
}

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Repo repo

swagger:model repo
*/
type Repo struct {

	/* URL

	Required: true
	Min Length: 1
	*/
	URL *string `json:"URL"`

	/* name

	Required: true
	Min Length: 1
	*/
	Name *string `json:"name"`

	/* source
	 */
	Source string `json:"source,omitempty"`
}

// Validate validates this repo
func (m *Repo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Repo) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("URL", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("URL", "body", string(*m.URL), 1); err != nil {
		return err
	}

	return nil
}

func (m *Repo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

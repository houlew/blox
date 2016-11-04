// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the License). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the license file accompanying this file. This file is distributed
// on an AS IS BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TaskContainerOverrideModel task container override model
// swagger:model TaskContainerOverrideModel
type TaskContainerOverrideModel struct {

	// command
	Command []string `json:"command"`

	// environment
	Environment []*TaskEnvironmentModel `json:"environment"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this task container override model
func (m *TaskContainerOverrideModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
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

func (m *TaskContainerOverrideModel) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	return nil
}

func (m *TaskContainerOverrideModel) validateEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	for i := 0; i < len(m.Environment); i++ {

		if swag.IsZero(m.Environment[i]) { // not required
			continue
		}

		if m.Environment[i] != nil {

			if err := m.Environment[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *TaskContainerOverrideModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

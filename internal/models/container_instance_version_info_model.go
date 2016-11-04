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

	"github.com/go-openapi/errors"
)

/*ContainerInstanceVersionInfoModel container instance version info model

swagger:model ContainerInstanceVersionInfoModel
*/
type ContainerInstanceVersionInfoModel struct {

	/* agent hash
	 */
	AgentHash string `json:"agentHash,omitempty"`

	/* agent version
	 */
	AgentVersion string `json:"agentVersion,omitempty"`

	/* docker version
	 */
	DockerVersion string `json:"dockerVersion,omitempty"`
}

// Validate validates this container instance version info model
func (m *ContainerInstanceVersionInfoModel) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

package e2einstancesteps

import (
	"github.com/aws/amazon-ecs-event-stream-handler/internal/models"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/pkg/errors"
)

func ValidateInstancesMatch(ecsInstance ecs.ContainerInstance, eshInstance models.ContainerInstanceModel) error {
	if *ecsInstance.ContainerInstanceArn != *eshInstance.ContainerInstanceARN ||
		*ecsInstance.Status != *eshInstance.Status {
		return errors.New("Container instances don't match")
	}
	return nil
}

func ValidateListContainsInstance(ecsInstance ecs.ContainerInstance, eshInstanceList []models.ContainerInstanceModel) error {
	instanceARN := *ecsInstance.ContainerInstanceArn
	var eshInstance models.ContainerInstanceModel
	for _, i := range eshInstanceList {
		if *i.ContainerInstanceARN == instanceARN {
			eshInstance = i
			break
		}
	}
	if eshInstance.ContainerInstanceARN == nil {
		return errors.Errorf("Instance with ARN '%s' not found in response", instanceARN)
	}
	return ValidateInstancesMatch(ecsInstance, eshInstance)
}

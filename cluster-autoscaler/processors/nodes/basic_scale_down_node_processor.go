/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodes

import (
	"sort"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
)

type BasicScaleDownNodeProcessor struct{}

func (p *BasicScaleDownNodeProcessor) GetPodDestinationCandidates(_ *context.AutoscalingContext, candidates []*apiv1.Node) ([]*apiv1.Node, errors.AutoscalerError) {
	return candidates, nil
}

// GetScaleDownCandidates sort scale down candidates by name in order to have stable order each CA loop.
// Having stable order is needed for scale down simulation.
func (p *BasicScaleDownNodeProcessor) GetScaleDownCandidates(_ *context.AutoscalingContext, candidates []*apiv1.Node) ([]*apiv1.Node, errors.AutoscalerError) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Name < candidates[j].Name
	})
	return candidates, nil
}

func (p *BasicScaleDownNodeProcessor) CleanUp() {}

/*
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
package pbr

import (
	"github.com/iter8-tools/iter8-controller/pkg/analytics/algorithm"
	"github.com/iter8-tools/iter8-controller/pkg/analytics/api"
	iter8v1alpha1 "github.com/iter8-tools/iter8-controller/pkg/apis/iter8/v1alpha1"
)

const (
	Strategy string = "posterior_bayesian_routing"

	SCKeyMinMax     string = "min_max"
	TCKeyConfidence string = "confidence"
)

var _ algorithm.Interface = Impl{}

type Impl struct {
}

type MinMax struct {
	// Minimum value of the metric
	Min float64 `json:"min"`

	// Maximum value of the metric
	Max float64 `json:"max"`
}

func (i Impl) SupplementSuccessCriteria(specSC iter8v1alpha1.SuccessCriterion, sc api.SuccessCriterion) (api.SuccessCriterion, error) {
	if specSC.MinMax != nil {
		sc[SCKeyMinMax] = MinMax{
			Min: specSC.MinMax.Min,
			Max: specSC.MinMax.Max,
		}
	}
	return sc, nil
}

func (i Impl) SupplementTrafficControl(instance *iter8v1alpha1.Experiment, tc api.TrafficControl) api.TrafficControl {
	tc[TCKeyConfidence] = instance.Spec.TrafficControl.GetConfidence()
	return tc
}

func (i Impl) GetPath() string {
	return api.AnalyticsAPIPath + Strategy
}

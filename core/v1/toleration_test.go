/*
Copyright 2017 The Kubernetes Authors.

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

package v1

import (
	"testing"
)

func TestTolerationToleratesTaint(t *testing.T) {

	testCases := []struct {
		description     string
		toleration      Toleration
		taint           Taint
		expectTolerated bool
	}{
		{
			description: "toleration and taint have the same key and effect, and operator is Exists, and taint has no value, expect tolerated",
			toleration: Toleration{
				Key:      ref("foo"),
				Operator: ref(TolerationOpExists),
				Effect:   ref(TaintEffectNoSchedule),
			},
			taint: Taint{
				Key:    ref("foo"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectTolerated: true,
		},
		{
			description: "toleration and taint have the same key and effect, and operator is Exists, and taint has some value, expect tolerated",
			toleration: Toleration{
				Key:      ref("foo"),
				Operator: ref(TolerationOpExists),
				Effect:   ref(TaintEffectNoSchedule),
			},
			taint: Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectTolerated: true,
		},
		{
			description: "toleration and taint have the same effect, toleration has empty key and operator is Exists, means match all taints, expect tolerated",
			toleration: Toleration{
				Key:      ref(""),
				Operator: ref(TolerationOpExists),
				Effect:   ref(TaintEffectNoSchedule),
			},
			taint: Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectTolerated: true,
		},
		{
			description: "toleration and taint have the same key, effect and value, and operator is Equal, expect tolerated",
			toleration: Toleration{
				Key:      ref("foo"),
				Operator: ref(TolerationOpEqual),
				Value:    ref("bar"),
				Effect:   ref(TaintEffectNoSchedule),
			},
			taint: Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectTolerated: true,
		},
		{
			description: "toleration and taint have the same key and effect, but different values, and operator is Equal, expect not tolerated",
			toleration: Toleration{
				Key:      ref("foo"),
				Operator: ref(TolerationOpEqual),
				Value:    ref("value1"),
				Effect:   ref(TaintEffectNoSchedule),
			},
			taint: Taint{
				Key:    ref("foo"),
				Value:  ref("value2"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectTolerated: false,
		},
		{
			description: "toleration and taint have the same key and value, but different effects, and operator is Equal, expect not tolerated",
			toleration: Toleration{
				Key:      ref("foo"),
				Operator: ref(TolerationOpEqual),
				Value:    ref("bar"),
				Effect:   ref(TaintEffectNoSchedule),
			},
			taint: Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoExecute),
			},
			expectTolerated: false,
		},
	}
	for _, tc := range testCases {
		if tolerated := tc.toleration.ToleratesTaint(&tc.taint); tc.expectTolerated != tolerated {
			t.Errorf("[%s] expect %v, got %v: toleration %+v, taint %s", tc.description, tc.expectTolerated, tolerated, tc.toleration, tc.taint.ToString())
		}
	}
}

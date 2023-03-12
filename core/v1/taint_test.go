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

func TestTaintToString(t *testing.T) {
	testCases := []struct {
		taint          *Taint
		expectedString string
	}{
		{
			taint: &Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectedString: "foo=bar:NoSchedule",
		},
		{
			taint: &Taint{
				Key:    ref("foo"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectedString: "foo:NoSchedule",
		},
		{
			taint: &Taint{
				Key: ref("foo"),
			},
			expectedString: "foo",
		},
		{
			taint: &Taint{
				Key:   ref("foo"),
				Value: ref("bar"),
			},
			expectedString: "foo=bar:",
		},
	}

	for i, tc := range testCases {
		if tc.expectedString != tc.taint.ToString() {
			t.Errorf("[%v] expected taint %v converted to %s, got %s", i, tc.taint, tc.expectedString, tc.taint.ToString())
		}
	}
}

func TestMatchTaint(t *testing.T) {
	testCases := []struct {
		description  string
		taint        *Taint
		taintToMatch Taint
		expectMatch  bool
	}{
		{
			description: "two taints with the same key,value,effect should match",
			taint: &Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			taintToMatch: Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectMatch: true,
		},
		{
			description: "two taints with the same key,effect but different value should match",
			taint: &Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			taintToMatch: Taint{
				Key:    ref("foo"),
				Value:  ref("different-value"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectMatch: true,
		},
		{
			description: "two taints with the different key cannot match",
			taint: &Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			taintToMatch: Taint{
				Key:    ref("different-key"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			expectMatch: false,
		},
		{
			description: "two taints with the different effect cannot match",
			taint: &Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectNoSchedule),
			},
			taintToMatch: Taint{
				Key:    ref("foo"),
				Value:  ref("bar"),
				Effect: ref(TaintEffectPreferNoSchedule),
			},
			expectMatch: false,
		},
	}

	for _, tc := range testCases {
		if tc.expectMatch != tc.taint.MatchTaint(&tc.taintToMatch) {
			t.Errorf("[%s] expect taint %s match taint %s", tc.description, tc.taint.ToString(), tc.taintToMatch.ToString())
		}
	}
}

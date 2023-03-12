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
	"fmt"
)

// MatchTaint checks if the taint matches taintToMatch. Taints are unique by key:effect,
// if the two taints have same key:effect, regard as they match.
func (t *Taint) MatchTaint(taintToMatch *Taint) bool {
	return value(t.Key) == value(taintToMatch.Key) && value(t.Effect) == value(taintToMatch.Effect)
}

// taint.ToString() converts taint struct to string in format '<key>=<value>:<effect>', '<key>=<value>:', '<key>:<effect>', or '<key>'.
func (t *Taint) ToString() string {
	if len(value(t.Effect)) == 0 {
		if len(value(t.Value)) == 0 {
			return fmt.Sprintf("%v", value(t.Key))
		}
		return fmt.Sprintf("%v=%v:", value(t.Key), value(t.Value))
	}
	if len(value(t.Value)) == 0 {
		return fmt.Sprintf("%v:%v", value(t.Key), value(t.Effect))
	}
	return fmt.Sprintf("%v=%v:%v", value(t.Key), value(t.Value), value(t.Effect))
}

func value[T any](v *T) (z T) {
	if v != nil {
		return *v
	}
	return
}

func ref[T any](v T) *T {
	return &v
}

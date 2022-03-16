// Copyright 2022 Neil Alexander
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package atomic

import "sync/atomic"

type Value[T any] struct {
	v atomic.Value
}

func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.v.CompareAndSwap(old, new)
}

func (v *Value[T]) Load() (val T) {
	return v.v.Load().(T)
}

func (v *Value[T]) Store(val T) {
	v.v.Store(val)
}

func (v *Value[T]) Swap(new T) (old T) {
	return v.v.Swap(new).(T)
}

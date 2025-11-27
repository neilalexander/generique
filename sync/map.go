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

package sync

import stdsync "sync"

type Map[TK, TV any] struct {
	m stdsync.Map
}

func (m *Map[TK, TV]) Delete(k TK) {
	m.m.Delete(k)
}

func (m *Map[TK, TV]) Load(key TK) (TV, bool) {
	value, ok := m.m.Load(key)
	if !ok {
		var empty TV
		return empty, false
	}
	return value.(TV), true
}

func (m *Map[TK, TV]) LoadAndDelete(key TK) (TV, bool) {
	value, ok := m.m.LoadAndDelete(key)
	if !ok {
		var empty TV
		return empty, false
	}
	return value.(TV), true
}

func (m *Map[TK, TV]) LoadOrStore(key TK, value TV) (TV, bool) {
	actual, ok := m.m.LoadOrStore(key, value)
	return actual.(TV), ok
}

func (m *Map[TK, TV]) Range(f func(TK, TV) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(TK), value.(TV))
	})
}

func (m *Map[TK, TV]) Store(key TK, value TV) {
	m.m.Store(key, value)
}

func (m *Map[TK, TV]) Swap(key TK, value TV) (TV, bool) {
	prev, ok := m.m.Swap(key, value)
	if !ok {
		var empty TV
		return empty, false
	}
	return prev.(TV), true
}

func (m *Map[TK, TV]) CompareAndSwap(key TK, old, new TV) bool {
	return m.m.CompareAndSwap(key, old, new)
}

func (m *Map[TK, TV]) CompareAndDelete(key TK, old TV) bool {
	return m.m.CompareAndDelete(key, old)
}

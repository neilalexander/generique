package sync

import stdsync "sync"

type Map[TK, TV any] struct {
	m stdsync.Map
}

func NewMap[TK, TV any]() *Map[TK, TV] {
	return &Map[TK, TV]{}
}

func (m *Map[TK, TV]) Delete(k TK) {
	m.m.Delete(k)
}

func (m *Map[TK, TV]) Load(key TK) (TV, bool) {
	value, ok := m.m.Load(key)
	return value.(TV), ok
}

func (m *Map[TK, TV]) LoadAndDelete(key TK) (TV, bool) {
	value, ok := m.m.LoadAndDelete(key)
	return value.(TV), ok
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

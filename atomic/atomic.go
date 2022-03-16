package atomic

import "sync/atomic"

type Value[T any] struct {
	v atomic.Value
}

func NewValue[T any]() *Value[T] {
	return &Value[T]{}
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

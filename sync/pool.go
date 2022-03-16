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

type Pool[T any] struct {
	pool stdsync.Pool
}

func NewPool[T any](newfunc func() any) *Pool[T] {
	return &Pool[T]{
		pool: stdsync.Pool{
			New: newfunc,
		},
	}
}

func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

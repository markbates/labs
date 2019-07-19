package valid

import (
	"encoding/json"
	"sync"
)

type Errors struct {
	moot   *sync.RWMutex
	errors map[string][]string
}

func (e Errors) Error() string {
	var msg string

	e.rwlock(func() {
		b, err := json.Marshal(e.errors)
		if err != nil {
			msg = err.Error()
			return
		}
		msg = string(b)
	})

	return msg
}

func (e *Errors) Add(key string, value string) {
	e.write(func() {
		e.errors[key] = append(e.errors[key], value)
	})
}

func (e *Errors) Set(key string, values []string) {
	e.write(func() {
		e.errors[key] = values
	})
}

func (e *Errors) Range(f func(key string, values []string) bool) {
	e.rwlock(func() {
		for k, v := range e.errors {
			if !f(k, v) {
				break
			}
		}
	})
}

func (e *Errors) Get(key string) ([]string, bool) {
	var values []string
	var ok bool
	e.read(func() {
		var x []string
		x, ok = e.errors[key]
		values = make([]string, len(x), len(x))
		copy(values, x)
	})
	return values, ok
}

func (e *Errors) rwlock(fn func()) {
	if e.moot == nil {
		e.moot = &sync.RWMutex{}
	}
	e.moot.RLock()
	defer e.moot.RUnlock()
	fn()
}

func (e *Errors) lock(fn func()) {
	if e.moot == nil {
		e.moot = &sync.RWMutex{}
	}
	e.moot.Lock()
	defer e.moot.Unlock()
	fn()
}

func (e *Errors) read(fn func()) {
	e.lock(func() {
		if e.errors == nil {
			e.errors = map[string][]string{}
		}
	})
	e.rwlock(fn)
}

func (e *Errors) write(fn func()) {
	e.lock(func() {
		if e.errors == nil {
			e.errors = map[string][]string{}
		}
	})
	e.lock(fn)
}

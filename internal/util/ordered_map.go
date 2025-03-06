package util

import (
	"encoding/json"
	"slices"
)

type OrderedMap[T any] struct {
	names  []string
	values map[string]T
}

func NewOrderedMap[T any]() *OrderedMap[T] {
	return &OrderedMap[T]{
		names:  []string{},
		values: map[string]T{},
	}
}

func (m *OrderedMap[T]) Add(name string, value T) {
	if _, ok := m.values[name]; ok {
		return //The item is already added - error?
	}

	m.values[name] = value
	m.names = append(m.names, name)
}

func (m *OrderedMap[T]) Get(name string) (T, bool) {
	value, ok := m.values[name]
	return value, ok
}

func (m *OrderedMap[T]) Iterate(f func(string, T) error) error {
	sorted_names := make([]string, len(m.names))
	copy(sorted_names, m.names)
	slices.Sort(sorted_names)

	for _, name := range sorted_names {
		value := m.values[name]

		if err := f(name, value); err != nil {
			return err
		}
	}

	return nil
}

func (m *OrderedMap[T]) MarshalJSON() ([]byte, error) {
	data := []map[string]T{}

	for _, name := range m.names {
		value := m.values[name]

		data = append(data, map[string]T{name: value})
	}

	return json.Marshal(data)
}

func (m *OrderedMap[T]) UnmarshalJSON(b []byte) error {
	m.values = map[string]T{}

	data := []map[string]T{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	for _, pair := range data {
		for name, value := range pair {
			m.Add(name, value)
		}
	}

	return nil
}

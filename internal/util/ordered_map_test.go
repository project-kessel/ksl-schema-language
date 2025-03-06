package util

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMapIteratesInInsertionOrder(t *testing.T) {
	m := NewOrderedMap[int]()

	m.Add("2", 2)
	m.Add("1", 1)
	m.Add("4", 4)
	m.Add("5", 5)
	m.Add("3", 3)

	actual := []int{}
	err := m.Iterate(func(s string, i int) error {
		actual = append(actual, i)
		return nil
	})
	assert.NoError(t, err)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, actual)
}

func TestOrderedMapGetAndAdd(t *testing.T) {
	m := NewOrderedMap[int]()

	m.Add("one", 1)

	v, ok := m.Get("one")
	assert.Equal(t, 1, v)
	assert.Equal(t, true, ok)

	_, ok = m.Get("two")
	assert.Equal(t, false, ok)
}

func TestOrderedMapReturnsErrorFromIterate(t *testing.T) {
	m := NewOrderedMap[int]()
	m.Add("one", 1)

	err := m.Iterate(func(s string, i int) error {
		return io.ErrUnexpectedEOF
	})

	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestOrderedMapSerializationRoundTrip(t *testing.T) {
	m := NewOrderedMap[int]()
	m.Add("1", 1)
	m.Add("3", 3)
	m.Add("2", 2)

	holder := mapHolder{OM: m}
	data, err := json.Marshal(holder)
	assert.NoError(t, err)

	holder2 := mapHolder{}
	err = json.Unmarshal(data, &holder2)
	assert.NoError(t, err)

	values := []int{}
	err = holder2.OM.Iterate(func(s string, i int) error {
		values = append(values, i)
		return nil
	})
	assert.NoError(t, err)

	assert.Equal(t, []int{1, 2, 3}, values)
}

type mapHolder struct {
	OM *OrderedMap[int] `json:"map"`
}

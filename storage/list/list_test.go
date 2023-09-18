package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()
	assert.NotNil(t, l)
	assert.Equal(t, int64(0), l.length, "New List should have 0 length.")
	assert.Nil(t, l.firstNode, "New List should not have firstNode.")
}

func TestList_Len(t *testing.T) {
	l := NewList()
	actual := l.Len()
	assert.Equal(t, int64(0), actual)
}

func TestList_Add(t *testing.T) {
	l := NewList()
	l.Add(42)
	assert.Equal(t, int64(1), l.Len())
	assert.Equal(t, int64(1), l.length)
	assert.NotNil(t, l.firstNode)
	assert.Equal(t, int64(42), l.firstNode.value)
	assert.Nil(t, l.firstNode.nextNode)

	l.Add(13)
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, int64(2), l.length)
	assert.NotNil(t, l.firstNode)
	assert.Equal(t, int64(42), l.firstNode.value)
	assert.NotNil(t, l.firstNode.nextNode)
	assert.Equal(t, int64(13), l.firstNode.nextNode.value)
	assert.Nil(t, l.firstNode.nextNode.nextNode)
}

func TestList_GetAll(t *testing.T) {
	l := NewList()
	actualValue, actualStatus := l.GetAll()
	assert.Nil(t, actualValue)
	assert.Equal(t, false, actualStatus)

	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	actualValue, actualStatus = l.GetAll()
	expectedValue := []int64{13, 42, 112, 256}
	assert.Equal(t, expectedValue, actualValue)
	assert.Equal(t, true, actualStatus)
}

func TestList_RemoveByIndex(t *testing.T) {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{256}, actualValue)
	assert.Equal(t, true, actualStatus)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	var expectedValue []int64
	assert.Equal(t, expectedValue, actualValue)
	assert.Equal(t, false, actualStatus)

	assert.Equal(t, false, l.RemoveByIndex(0))
}

func TestList_RemoveByIndex2(t *testing.T) {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(4), l.Len())

	assert.Equal(t, true, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(3), l.Len())

	assert.Equal(t, true, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(2), l.Len())

	assert.Equal(t, true, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(1), l.Len())

	assert.Equal(t, false, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(1), l.Len())
}

func TestList_RemoveByIndex3(t *testing.T) {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(4), l.Len())

	assert.Equal(t, true, l.RemoveByIndex(2))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(3), l.Len())

	assert.Equal(t, true, l.RemoveByIndex(2))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(2), l.Len())

	assert.Equal(t, false, l.RemoveByIndex(2))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(2), l.Len())
}

func TestList_RemoveByValue(t *testing.T) {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 256, 13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(8), l.Len())

	assert.Equal(t, true, l.RemoveByValue(13))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(7), l.Len())

	assert.Equal(t, true, l.RemoveByValue(13))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())

	assert.Equal(t, false, l.RemoveByValue(13))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())
}

func TestList_RemoveByValue2(t *testing.T) {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 256, 13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(8), l.Len())

	assert.Equal(t, true, l.RemoveByValue(256))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(7), l.Len())

	assert.Equal(t, true, l.RemoveByValue(256))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 13, 42, 112}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())

	assert.Equal(t, false, l.RemoveByValue(256))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 13, 42, 112}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())
}

func TestList_RemoveByValue3(t *testing.T) {
	l := NewList()
	l.Add(42)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{42}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(1), l.Len())

	l.RemoveByValue(42)
	actualValue, actualStatus = l.GetAll()
	assert.Nil(t, actualValue)
	assert.Equal(t, false, actualStatus)
}

package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()
	assert.NotNil(t, l)
	assert.Equal(t, int64(0), l.length, "New List should have 0 length.")
	assert.Equal(t, int64(0), l.idCounter)
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
	assert.Equal(t, int64(1), l.idCounter)
	assert.Equal(t, int64(1), l.firstNode.id)
	assert.Nil(t, l.firstNode.nextNode)
	assert.Equal(t, l.firstNode, l.lastNode)

	l.Add(13)
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, int64(2), l.length)
	assert.NotNil(t, l.firstNode)
	assert.Equal(t, int64(42), l.firstNode.value)
	assert.NotNil(t, l.firstNode.nextNode)
	assert.Equal(t, int64(13), l.firstNode.nextNode.value)
	assert.Equal(t, int64(2), l.idCounter)
	assert.Equal(t, int64(1), l.firstNode.id)
	assert.Equal(t, int64(2), l.firstNode.nextNode.id)
	assert.Nil(t, l.firstNode.nextNode.nextNode)
	assert.Equal(t, l.firstNode.nextNode, l.lastNode)
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
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(0))
	actualValue, actualStatus = l.GetAll()
	var expectedValue []int64
	assert.Equal(t, expectedValue, actualValue)
	assert.Equal(t, false, actualStatus)
	assert.Nil(t, l.lastNode)

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
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(3), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(1), l.Len())
	assert.Equal(t, int64(13), l.lastNode.value)

	assert.Equal(t, false, l.RemoveByIndex(1))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(1), l.Len())
	assert.Equal(t, int64(13), l.lastNode.value)
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
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(2))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(3), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByIndex(2))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, int64(42), l.lastNode.value)

	assert.Equal(t, false, l.RemoveByIndex(2))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(2), l.Len())
	assert.Equal(t, int64(42), l.lastNode.value)
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
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByValue(13))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(7), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByValue(13))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, false, l.RemoveByValue(13))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)
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
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByValue(256))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 13, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(7), l.Len())
	assert.Equal(t, int64(256), l.lastNode.value)

	assert.Equal(t, true, l.RemoveByValue(256))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 13, 42, 112}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())
	assert.Equal(t, int64(112), l.lastNode.value)

	assert.Equal(t, false, l.RemoveByValue(256))
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{13, 42, 112, 13, 42, 112}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(6), l.Len())
	assert.Equal(t, int64(112), l.lastNode.value)
}

func TestList_RemoveByValue3(t *testing.T) {
	l := NewList()
	l.Add(42)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{42}, actualValue)
	assert.Equal(t, true, actualStatus)
	assert.Equal(t, int64(1), l.Len())
	assert.Equal(t, int64(42), l.lastNode.value)

	l.RemoveByValue(42)
	actualValue, actualStatus = l.GetAll()
	assert.Nil(t, actualValue)
	assert.Equal(t, false, actualStatus)
	assert.Nil(t, l.lastNode)
}

func TestList_RemoveByID(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4
	assert.Equal(t, int64(4), l.Len())
	assert.Equal(t, int64(1), l.firstNode.id)
	assert.Equal(t, int64(2), l.firstNode.nextNode.id)
	assert.Equal(t, int64(3), l.firstNode.nextNode.nextNode.id)
	assert.Equal(t, int64(4), l.firstNode.nextNode.nextNode.nextNode.id)
	assert.Equal(t, int64(4), l.lastNode.id)

	assert.Equal(t, true, l.RemoveByID(3))
	assert.Equal(t, int64(3), l.Len())
	assert.Equal(t, int64(1), l.firstNode.id)
	assert.Equal(t, int64(2), l.firstNode.nextNode.id)
	assert.Equal(t, int64(4), l.firstNode.nextNode.nextNode.id)
	assert.Equal(t, int64(4), l.lastNode.id)

	l.Add(1007)
	assert.Equal(t, int64(4), l.Len())
	assert.Equal(t, int64(1), l.firstNode.id)
	assert.Equal(t, int64(2), l.firstNode.nextNode.id)
	assert.Equal(t, int64(4), l.firstNode.nextNode.nextNode.id)
	assert.Equal(t, int64(5), l.firstNode.nextNode.nextNode.nextNode.id)
	assert.Equal(t, int64(5), l.lastNode.id)

	assert.Equal(t, false, l.RemoveByID(10))
	assert.Equal(t, int64(4), l.Len())

	assert.Equal(t, true, l.RemoveByID(1))
	assert.Equal(t, true, l.RemoveByID(2))
	assert.Equal(t, false, l.RemoveByID(3))
	assert.Equal(t, true, l.RemoveByID(4))
	assert.Equal(t, true, l.RemoveByID(5))

	assert.Equal(t, int64(0), l.Len())
	assert.Equal(t, int64(5), l.idCounter)

	assert.Equal(t, false, l.RemoveByID(2))
}

func TestList_RemoveByID2(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4

	assert.Equal(t, true, l.RemoveByID(4))
	assert.Equal(t, int64(3), l.Len())
	assert.Equal(t, int64(1), l.firstNode.id)
	assert.Equal(t, int64(2), l.firstNode.nextNode.id)
	assert.Equal(t, int64(3), l.firstNode.nextNode.nextNode.id)
	assert.Equal(t, int64(3), l.lastNode.id)

}

func TestList_RemoveAllByValue(t *testing.T) {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)
	l.Add(13)
	l.Add(42)
	l.Add(112)
	l.Add(256)

	l.RemoveAllByValue(13)
	actualValue, actualStatus := l.GetAll()
	assert.Equal(t, []int64{42, 112, 256, 42, 112, 256}, actualValue)
	assert.Equal(t, true, actualStatus)

	l.RemoveAllByValue(256)
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{42, 112, 42, 112}, actualValue)
	assert.Equal(t, true, actualStatus)

	l.RemoveAllByValue(42)
	actualValue, actualStatus = l.GetAll()
	assert.Equal(t, []int64{112, 112}, actualValue)
	assert.Equal(t, true, actualStatus)

	l.RemoveAllByValue(112)
	actualValue, actualStatus = l.GetAll()
	assert.Nil(t, actualValue)
	assert.Equal(t, false, actualStatus)
}

func TestList_GetValueByIndex(t *testing.T) {
	l := NewList()
	l.Add(13)  // index: 0
	l.Add(42)  // index: 1
	l.Add(112) // index: 2
	l.Add(256) // index: 3

	actualValue, actualStatus := l.GetValueByIndex(0)
	assert.Equal(t, int64(13), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByIndex(1)
	assert.Equal(t, int64(42), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByIndex(2)
	assert.Equal(t, int64(112), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByIndex(3)
	assert.Equal(t, int64(256), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByIndex(4)
	assert.Equal(t, int64(0), actualValue)
	assert.Equal(t, false, actualStatus)

}

func TestList_GetValueByID(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4

	actualValue, actualStatus := l.GetValueByID(1)
	assert.Equal(t, int64(13), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByID(2)
	assert.Equal(t, int64(42), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByID(3)
	assert.Equal(t, int64(112), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByID(4)
	assert.Equal(t, int64(256), actualValue)
	assert.Equal(t, true, actualStatus)

	actualValue, actualStatus = l.GetValueByID(5)
	assert.Equal(t, int64(0), actualValue)
	assert.Equal(t, false, actualStatus)

	actualValue, actualStatus = l.GetValueByID(0)
	assert.Equal(t, int64(0), actualValue)
	assert.Equal(t, false, actualStatus)

	actualValue, actualStatus = l.GetValueByID(-1)
	assert.Equal(t, int64(0), actualValue)
	assert.Equal(t, false, actualStatus)
}

func TestList_GetValueByID2(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4
	l.RemoveByID(3)

	actualValue, actualStatus := l.GetValueByID(3)
	assert.Equal(t, int64(0), actualValue)
	assert.Equal(t, false, actualStatus)
}

func TestList_GetIndexByValue(t *testing.T) {
	l := NewList()
	l.Add(13)  // index: 0
	l.Add(42)  // index: 1
	l.Add(112) // index: 2
	l.Add(256) // index: 3

	actualIndex, actualStatus := l.GetIndexByValue(13)
	assert.Equal(t, int64(0), actualIndex)
	assert.Equal(t, true, actualStatus)

	actualIndex, actualStatus = l.GetIndexByValue(42)
	assert.Equal(t, int64(1), actualIndex)
	assert.Equal(t, true, actualStatus)

	actualIndex, actualStatus = l.GetIndexByValue(112)
	assert.Equal(t, int64(2), actualIndex)
	assert.Equal(t, true, actualStatus)

	actualIndex, actualStatus = l.GetIndexByValue(256)
	assert.Equal(t, int64(3), actualIndex)
	assert.Equal(t, true, actualStatus)

	actualIndex, actualStatus = l.GetIndexByValue(100000)
	assert.Equal(t, l.Len(), actualIndex)
	assert.Equal(t, false, actualStatus)
}

func TestList_GetIDByValue(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4

	actualID, actualStatus := l.GetIDByValue(13)
	assert.Equal(t, int64(1), actualID)
	assert.Equal(t, true, actualStatus)

	actualID, actualStatus = l.GetIDByValue(42)
	assert.Equal(t, int64(2), actualID)
	assert.Equal(t, true, actualStatus)

	actualID, actualStatus = l.GetIDByValue(112)
	assert.Equal(t, int64(3), actualID)
	assert.Equal(t, true, actualStatus)

	actualID, actualStatus = l.GetIDByValue(256)
	assert.Equal(t, int64(4), actualID)
	assert.Equal(t, true, actualStatus)

	actualID, actualStatus = l.GetIDByValue(1000000)
	assert.Equal(t, int64(0), actualID)
	assert.Equal(t, false, actualStatus)
}

func TestList_GetAllIDsByValue(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4
	l.Add(13)  // id: 5
	l.Add(42)  // id: 6
	l.Add(112) // id: 7
	l.Add(256) // id: 8

	actualIDs, actualStatus := l.GetAllIDsByValue(13)
	assert.Equal(t, []int64{1, 5}, actualIDs)
	assert.Equal(t, true, actualStatus)

	actualIDs, actualStatus = l.GetAllIDsByValue(256)
	assert.Equal(t, []int64{4, 8}, actualIDs)
	assert.Equal(t, true, actualStatus)

	actualIDs, actualStatus = l.GetAllIDsByValue(112)
	assert.Equal(t, []int64{3, 7}, actualIDs)
	assert.Equal(t, true, actualStatus)

	actualIDs, actualStatus = l.GetAllIDsByValue(100000)
	assert.Nil(t, actualIDs)
	assert.Equal(t, false, actualStatus)

}

func TestList_Clear(t *testing.T) {
	l := NewList()
	l.Add(13)  // id: 1
	l.Add(42)  // id: 2
	l.Add(112) // id: 3
	l.Add(256) // id: 4
	l.Add(13)  // id: 5
	l.Add(42)  // id: 6
	l.Add(112) // id: 7
	l.Add(256) // id: 8

	l.Clear()
	assert.Nil(t, l.firstNode)
	assert.Nil(t, l.lastNode)
	assert.Equal(t, int64(0), l.Len())

	l.Add(1007)
	assert.Equal(t, int64(1007), l.firstNode.value)
	assert.Equal(t, int64(1007), l.lastNode.value)
	assert.Equal(t, int64(9), l.firstNode.id)
	assert.Equal(t, int64(9), l.lastNode.id)
}

func ExampleList_Print() {
	l := NewList()
	l.Add(13)
	l.Add(42)
	l.Add(112)

	l.Print()
	// Output: [13, 42, 112]
}

func ExampleList_Print2() {
	l := NewList()

	l.Print()
	// Output: []
}

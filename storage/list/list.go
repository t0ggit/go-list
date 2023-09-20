package list

import (
	"fmt"
	"strconv"
)

type List struct {
	length    int64 // Текущая длина списка (количество узлов)
	firstNode *node // Указатель на первый узел
	lastNode  *node // Указатель на последний узел (для ускорения вставки элемента в конец)

	/*
		Счетчик идентификаторов,
		растет по мере добавления, но не уменьшается при удалении элементов,
		начинается с 1, т.е. первый вставленный узел будет иметь идентификатор 1.
	*/
	idCounter int64
}

// NewList создает новый пустой односвязный список
func NewList() (l *List) {
	return &List{length: 0, firstNode: nil, lastNode: nil, idCounter: 0}
}

// Len возвращает количество элементов в списке
func (l *List) Len() (len int64) {
	len = l.length
	return len
}

// Add добавляет элемент в конец списка, возвращает идентификатор добавленного элемента
func (l *List) Add(data int64) (id int64) {
	l.idCounter++
	newNode := &node{value: data, id: l.idCounter}
	l.length++
	nextNode := l.firstNode
	if nextNode == nil {
		l.firstNode = newNode
		l.lastNode = l.firstNode
		return l.idCounter
	}
	l.lastNode.nextNode = newNode
	l.lastNode = l.lastNode.nextNode
	return l.idCounter
}

// RemoveByIndex удаляет элемент по индексу (текущему порядковому номеру)
func (l *List) RemoveByIndex(index int64) (ok bool) {
	if l.length <= index {
		return false
	}

	var previousNode *node = nil
	currentNode := l.firstNode
	for currentIndex := int64(0); currentIndex < index; currentIndex++ {
		previousNode = currentNode
		currentNode = currentNode.nextNode
	}
	if previousNode == nil {
		l.firstNode = currentNode.nextNode
		if l.firstNode == nil {
			l.lastNode = nil
		}
	} else if currentNode.nextNode != nil {
		previousNode.nextNode = currentNode.nextNode
	} else {
		previousNode.nextNode = nil
		l.lastNode = previousNode
	}
	l.length--
	return true
}

// RemoveByID удаляет элемент по уникальному идентификатору
func (l *List) RemoveByID(id int64) (ok bool) {
	if l.firstNode == nil {
		return false
	}

	if l.firstNode.id == id {
		if l.firstNode.nextNode == l.lastNode {
			l.lastNode = l.firstNode
		}
		l.firstNode = l.firstNode.nextNode
		if l.firstNode == nil {
			l.lastNode = nil
		}
		l.length--
		return true
	}

	currentNode := l.firstNode
	for ; currentNode.nextNode != nil && currentNode.nextNode.id != id; currentNode = currentNode.nextNode {
	}
	if currentNode.nextNode == nil {
		return false // not found
	}
	if currentNode.nextNode == l.lastNode {
		l.lastNode = currentNode
	}
	currentNode.nextNode = currentNode.nextNode.nextNode
	l.length--
	return true
}

// RemoveByValue удаляет первый встретившийся элемент с данным значением
func (l *List) RemoveByValue(value int64) (ok bool) {
	if l.firstNode == nil {
		return false
	}

	if l.firstNode.value == value {
		if l.firstNode.nextNode == l.lastNode {
			l.lastNode = l.firstNode
		}
		l.firstNode = l.firstNode.nextNode
		if l.firstNode == nil {
			l.lastNode = nil
		}
		l.length--
		return true
	}

	currentNode := l.firstNode
	for ; currentNode.nextNode != nil && currentNode.nextNode.value != value; currentNode = currentNode.nextNode {
	}
	if currentNode.nextNode == nil {
		return false // not found
	}
	if currentNode.nextNode == l.lastNode {
		l.lastNode = currentNode
	}
	currentNode.nextNode = currentNode.nextNode.nextNode
	l.length--
	return true
}

// RemoveAllByValue удаляет все элементы с данным значением
func (l *List) RemoveAllByValue(value int64) {
	for l.RemoveByValue(value) {
	}
}

// GetValueByIndex возвращает значение элемента с данным индексом
func (l *List) GetValueByIndex(index int64) (value int64, ok bool) {
	if index >= l.Len() || index < 0 {
		return 0, false
	}

	currentNode := l.firstNode
	for currentIndex := int64(0); currentIndex < index; currentIndex++ {
		currentNode = currentNode.nextNode
	}
	return currentNode.value, true
}

// GetValueByID возвращает значение элемента с данным идентификатором
func (l *List) GetValueByID(id int64) (value int64, ok bool) {
	if id > l.idCounter || id <= 0 {
		return 0, false
	}
	currentNode := l.firstNode
	for ; currentNode.id != id; currentNode = currentNode.nextNode {
		if currentNode.nextNode == nil {
			return 0, false
		}
	}
	return currentNode.value, true
}

// GetIndexByValue возвращает индекс первого по порядку элемента с данным значением
// Если не находит элемента с данным значением, возвращает len и false
func (l *List) GetIndexByValue(value int64) (index int64, ok bool) {
	var (
		currentIndex  int64 = 0
		currentValue  int64
		currentStatus bool
	)
	for ; ; currentIndex++ {
		currentValue, currentStatus = l.GetValueByIndex(currentIndex)
		if currentValue == value {
			return currentIndex, true
		}
		if !currentStatus {
			return l.length, false
		}
	}
}

// GetIDByValue возвращает идентификатор первого по порядку элемента с данным значением
// Если не находит элемента с данным значением, возвращает 0 и false
func (l *List) GetIDByValue(value int64) (id int64, ok bool) {
	for currentNode := l.firstNode; currentNode != nil; currentNode = currentNode.nextNode {
		if currentNode.value == value {
			return currentNode.id, true
		}
	}
	return 0, false
}

// GetAllIDsByValue возвращает индексы всех элементов с данным значением
// Если элементы с данным значением не найдены, возвращает nil и false
func (l *List) GetAllIDsByValue(value int64) (ids []int64, ok bool) {
	for currentNode := l.firstNode; currentNode != nil; currentNode = currentNode.nextNode {
		if currentNode.value == value {
			ids = append(ids, currentNode.id)
		}
	}
	if len(ids) == 0 {
		return nil, false
	}
	return ids, true
}

// GetAll возвращает значения всех элементов
func (l *List) GetAll() (values []int64, ok bool) {
	if l.length == 0 || l.firstNode == nil {
		return nil, false
	}

	for currentNode := l.firstNode; currentNode != nil; currentNode = currentNode.nextNode {
		values = append(values, currentNode.value)
	}
	return values, true
}

// Clear удаляет все элементы из списка
func (l *List) Clear() {
	l.firstNode = nil
	l.lastNode = nil
	l.length = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	if l.Len() == int64(0) {
		fmt.Printf("[]\n")
		return
	}
	fmt.Printf("[")
	currentNode := l.firstNode
	for ; currentNode.nextNode != nil; currentNode = currentNode.nextNode {
		fmt.Printf(strconv.FormatInt(currentNode.value, 10) + ", ")
	}
	fmt.Printf(strconv.FormatInt(currentNode.value, 10) + "]\n")
	return
}

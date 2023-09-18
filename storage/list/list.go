package list

import "fmt"

type List struct {
	length    int64
	firstNode *node
}

// NewList создает новый пустой односвязный список
func NewList() (l *List) {
	return &List{length: 0, firstNode: nil}
}

// Len возвращает количество элементов в списке
func (l *List) Len() (len int64) {
	len = l.length
	return len
}

// Add добавляет элемент в конец списка, возвращает индекс добавленного элемента
func (l *List) Add(data int64) (id int64) {
	newNode := &node{value: data}
	l.length++
	var currentID int64 = 0
	nextNode := l.firstNode
	if nextNode == nil {
		l.firstNode = newNode
		return currentID
	}
	for ; nextNode.nextNode != nil; nextNode = nextNode.nextNode {
		currentID++
	}
	nextNode.nextNode = newNode
	return currentID
}

// RemoveByIndex удаляет элемент по индексу
func (l *List) RemoveByIndex(id int64) (ok bool) {
	if l.length == 0 || l.firstNode == nil {
		return false
	}
	if l.length <= id {
		return false
	}

	var previousNode *node = nil
	currentNode := l.firstNode
	for currentIndex := int64(0); currentIndex < id; currentIndex++ {
		previousNode = currentNode
		currentNode = currentNode.nextNode
	}
	if previousNode == nil {
		l.firstNode = currentNode.nextNode
	} else if currentNode.nextNode != nil {
		previousNode.nextNode = currentNode.nextNode
	} else {
		previousNode.nextNode = nil
	}
	l.length--
	return true
}

// RemoveByValue удаляет первый встретившийся элемент с данным значением
func (l *List) RemoveByValue(value int64) (ok bool) {
	if l.firstNode == nil {
		return false
	}

	if l.firstNode.value == value {
		l.firstNode = l.firstNode.nextNode
		l.length--
		return true
	}

	currentNode := l.firstNode
	for ; currentNode.nextNode != nil && currentNode.nextNode.value != value; currentNode = currentNode.nextNode {
	}
	if currentNode.nextNode == nil {
		return false
	}
	currentNode.nextNode = currentNode.nextNode.nextNode
	l.length--
	return true
}

// RemoveAllByValue удаляет все элементы с данным значением
func (l *List) RemoveAllByValue(value int64) {
	return
}

// GetValueByIndex возвращает значение элемента с данным индексом
func (l *List) GetValueByIndex(id int64) (value int64, ok bool) {
	return
}

// GetIndexByValue возвращает индекс первого встретившегося элемента с данным значением
func (l *List) GetIndexByValue(value int64) (id int64, ok bool) {
	return
}

// GetAllIndicesByValue возвращает индексы всех элементов с данным значением
func (l *List) GetAllIndicesByValue(value int64) (ids []int64, ok bool) {
	return
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

}

// Print выводит список в консоль
func (l *List) Print() {
	if l.firstNode == nil {
		fmt.Println("[]")
		return
	}
	for n := l.firstNode; n != nil; n = n.nextNode {
		fmt.Println(n.value)
	}
	return
}

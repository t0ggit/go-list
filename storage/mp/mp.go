package mp

import "fmt"

type Map struct {
	mp        map[int64]int64
	idCounter int64
}

// NewMap возвращает новую таблицу
func NewMap() (m *Map) {
	return &Map{idCounter: int64(1), mp: make(map[int64]int64)}
}

// Len возвращает количество элементов в таблице
func (m *Map) Len() int64 {
	return int64(len(m.mp))
}

// Add добавляет значение в таблицу и возвращает его идентификатор
func (m *Map) Add(value int64) int64 {
	m.mp[m.idCounter] = value
	m.idCounter++
	return m.idCounter - 1
}

// RemoveByID удаляет элемент из таблицы по идентификатору
func (m *Map) RemoveByID(id int64) {
	delete(m.mp, id)
}

// RemoveByValue удаляет один элемент из таблицы по значению
func (m *Map) RemoveByValue(value int64) {
	for key, val := range m.mp {
		if val == value {
			delete(m.mp, key)
			return
		}
	}
}

// GetByID возвращает значение элемента по идентификатору.
//
// Если элемента с таким идентификатором нет, то возвращается 0 и false.
func (m *Map) GetByID(id int64) (value int64, ok bool) {
	value, ok = m.mp[id]
	return value, ok
}

// GetByValue возвращает идентификатор первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (m *Map) GetByValue(value int64) (id int64, ok bool) {
	for key, val := range m.mp {
		if val == value {
			return key, true
		}
	}
	return 0, false
}

// Clear очищает список
func (m *Map) Clear() {
	m.mp = make(map[int64]int64)
}

// Print выводит список в консоль
func (m *Map) Print() {
	for key, val := range m.mp {
		fmt.Printf("%d: %d\n", key, val)
	}
}

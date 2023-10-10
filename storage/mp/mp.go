package mp

import (
	"fmt"
	"list/storage"
	"reflect"
	"sync"
)

type Map struct {
	mp map[int64]any
	// Тип id (ключа) всегда int64
	idCounter int64
	mu        sync.RWMutex
	// V - тип значения, фиксируется при добавлении первого элемента, сбрасывается при удалении последнего элемента
	V reflect.Type
}

// NewMap возвращает новую таблицу
func NewMap() (m *Map) {
	return &Map{idCounter: int64(0), mp: make(map[int64]any)}
}

// Len возвращает количество элементов в таблице
func (m *Map) Len() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return int64(len(m.mp))
}

// Add добавляет значение в таблицу и возвращает его идентификатор
func (m *Map) Add(value any) (int64, error) {
	if m.V != nil {
		m.V = reflect.TypeOf(value)
	} else if m.V != reflect.TypeOf(value) {
		return 0, storage.ErrMismatchType
	}
	m.mu.Lock()
	//defer fmt.Println("unlocked") // потом это
	defer m.mu.Unlock() // defer - отложенное выполнение функции, то есть это будет сделано после return
	//defer fmt.Println("done")     // сначала это
	m.mp[m.idCounter] = value
	m.idCounter++
	// defer защищает от паники, которая может не выполнить Unlock(), экстренно выйдя из функции,
	// оставив мьютекс залоченным навсегда
	return m.idCounter - 1, nil
}

// RemoveByID удаляет элемент из таблицы по идентификатору
func (m *Map) RemoveByID(id int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.mp, id)
	if m.Len() == 0 {
		m.V = nil
	}
}

// RemoveByValue удаляет один элемент из таблицы по значению
func (m *Map) RemoveByValue(value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, val := range m.mp {
		if val == value {
			delete(m.mp, key)
			if m.Len() == 0 {
				m.V = nil
			}
			return
		}
	}

}

// RemoveAllByValue удаляет все элементы из таблицы по значению
func (m *Map) RemoveAllByValue(value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, val := range m.mp {
		if val == value {
			delete(m.mp, key)
		}
	}
	if m.Len() == 0 {
		m.V = nil
	}
}

// GetByID возвращает значение элемента по идентификатору.
//
// Если элемента с таким идентификатором нет, то возвращается 0 и false.
func (m *Map) GetByID(id int64) (value any, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok = m.mp[id]
	return value, ok
}

// GetByValue возвращает идентификатор первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (m *Map) GetByValue(value any) (id int64, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.V != reflect.TypeOf(value) {
		return 0, false
	}
	for key, val := range m.mp {
		if val == value {
			return key, true
		}
	}
	return 0, false
}

// Clear очищает список
func (m *Map) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp = make(map[int64]any)
	m.V = nil
	m.idCounter = int64(0)
}

// Print выводит список в консоль
func (m *Map) Print() {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for key, val := range m.mp {
		fmt.Println("Key:", key, "Value:", val)
	}
}

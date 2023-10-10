package main

import (
	"fmt"
	"reflect"
)

// global variables
var (
	id int64
	st string
)

// constants
const (
	idCounter int64 = 0
)

//func main() {
//	r := Sum(1, 2)
//	fmt.Println(r)
//
//}

func Sum(a, b int64) int64 {
	s := func(a, b int64) int64 {
		return a + b
	}
	return s(a, b)
}

// следующее дз по интерфейсам:
// лист и мапу доработать, чтобы они хранили любой тип данных, обязательно должен быть контроль типа данных,
// Add должна уметь возвращать ошибку, если пытаемся в один лист или мап вставить другой тип данных,
// то есть флаг для типа данных значений, при очищении он сбрасывается и можно опять пихать туда любой другой тип,
// для проверки типа данных можно использовать golang reflect, понадобится reflect.TypeOf() и reflect.ValueOf()
// плюс еще многопоточность прикрутить везде

//var s []string
//
//func apd(a any) (err error) {
//	switch a.(type) {
//	case string:
//		s = append(s, a.(string))
//		return nil
//	default:
//		err = errors.New("type not supported")
//		return err
//	}
//}

// многопоточность

//func main() {
//
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	t1 := time.Now()
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 10000; i++ {
//			fmt.Println("msg from goroutine 1 number", i)
//			time.Sleep(10 * time.Millisecond)
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 10000; i++ {
//			fmt.Println("msg from goroutine 2 number", i)
//			time.Sleep(12 * time.Millisecond)
//		}
//	}()
//
//	// если main завершается, то остальные горутины завершаются автоматически
//	//time.Sleep(20 * time.Second)
//	fmt.Println("runtime:", time.Since(t1))
//	return
//	// нужно реализовывать контексты (чтобы во время завершения данные как-то сохранялись)
//
//	// консоль это такой же файл, поток ввода-вывода, то есть во время вывода происходит блокировка
//}

func main() {
	var V reflect.Type
	var abc any
	V = reflect.TypeOf(abc)
	fmt.Println(V == reflect.TypeOf(abc))

}

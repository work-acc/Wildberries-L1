Сколько существует способов задать переменную типа slice или map?

Ответ:
```
Используя литералы:
Для slice можно использовать литералы с указанием элементов: 
mySlice := []int{1, 2, 3}.
Для map можно использовать литералы с указанием пар ключ-значение: 
myMap := map[string]int{"a": 1, "b": 2}.

Используя функцию make:
Для slice: mySlice := make([]int, 0, 10).
Для map: myMap := make(map[string]int).

Инициализация пустым значением:
Для slice: var mySlice []int.
Для map: var myMap map[string]int.

Используя копирование существующих значений:
Для slice: newSlice := []int{1, 2, 3}; mySlice := newSlice[:].
Для map: newMap := map[string]int{"a": 1, "b": 2}; myMap := newMap.
```

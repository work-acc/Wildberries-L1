Что такое интерфейсы, как они применяются в Go?

Ответ:
```
В языке программирования Go интерфейсы представляют собой коллекцию 
методов без их реализации. Интерфейсы позволяют определить набор методов, 
которые тип должен реализовать, чтобы считаться реализацией этого интерфейса.

type share interface {
Area() float64, Perimetr() float64 }

В языке Go есть «пустой» интерфейс interface{}. Он не содержит в себе никаких 
методов, описывающих поведение программной сущности.
Это значит, что любое значение может быть приведено к типу interface{}.
Пустой интерфейс используется, когда заранее неизвестен тип данных.
```

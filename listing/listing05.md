Какой размер у структуры struct{}{}?

Ответ:
```
В Go пустая структура struct{}{} не занимает памяти. 
Она не имеет полей и, следовательно, не занимает места в памяти. 
Однако, поскольку каждая переменная должна иметь уникальный тип, 
компилятор Go выделяет нулевой байт памяти для экземпляров пустой 
структуры, чтобы гарантировать, что каждая переменная имеет 
уникальный тип.
```

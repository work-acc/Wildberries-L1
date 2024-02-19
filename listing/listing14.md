Что выведет данная программа и почему?

```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}

```

Ответ:
```
Output: [b b a][a a]

Передается slice в анонимную функцию, затем выполняется 
добавление элемента с помощью append и в памяти создается 
новая ссылка на массив, и уже у нее меняются значения. 
При этом изначальный slice никак не меняется.
```

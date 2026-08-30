Oto 40 najpopularniejszych funkcji biblioteki `go-funk` wraz z przykładami kodu dla każdej z nich:

### 1. **Map**
   - Przekształca każdy element listy na podstawie podanej funkcji.
   ```go
   result := funk.Map([]int{1, 2, 3}, func(x int) int {
       return x * 2
   })
   ```

### 2. **Filter**
   - Filtruje elementy listy, pozostawiając tylko te, które spełniają określony warunek.
   ```go
   result := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 3. **Reduce**
   - Redukuje listę do jednej wartości, stosując funkcję akumulacyjną.
   ```go
   result := funk.Reduce([]int{1, 2, 3, 4}, func(acc int, x int) int {
       return acc + x
   }, 0)
   ```

### 4. **Contains**
   - Sprawdza, czy lista zawiera dany element.
   ```go
   result := funk.Contains([]string{"a", "b", "c"}, "b")
   ```

### 5. **Uniq**
   - Usuwa zduplikowane wartości z listy.
   ```go
   result := funk.Uniq([]int{1, 2, 2, 3, 4, 4, 5})
   ```

### 6. **Flatten**
   - Spłaszcza listę zagnieżdżonych list do jednej listy.
   ```go
   result := funk.Flatten([][]int{{1, 2}, {3, 4}})
   ```

### 7. **Reverse**
   - Odwraca kolejność elementów w liście.
   ```go
   result := funk.Reverse([]int{1, 2, 3})
   ```

### 8. **Every**
   - Sprawdza, czy każdy element w liście spełnia podany warunek.
   ```go
   result := funk.Every([]int{2, 4, 6}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 9. **Some**
   - Sprawdza, czy przynajmniej jeden element w liście spełnia podany warunek.
   ```go
   result := funk.Some([]int{1, 2, 3}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 10. **All**
   - Alias dla `Every`. Sprawdza, czy wszystkie elementy spełniają warunek.
   ```go
   result := funk.All([]int{2, 4, 6}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 11. **None**
   - Sprawdza, czy żaden element w liście nie spełnia podanego warunku.
   ```go
   result := funk.None([]int{1, 3, 5}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 12. **Find**
   - Zwraca pierwszy element, który spełnia podany warunek.
   ```go
   result := funk.Find([]int{1, 2, 3}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 13. **Chunk**
   - Dzieli listę na mniejsze listy o określonym rozmiarze.
   ```go
   result := funk.Chunk([]int{1, 2, 3, 4}, 2)
   ```

### 14. **Difference**
   - Zwraca elementy, które znajdują się w jednej liście, ale nie w drugiej.
   ```go
   result := funk.Difference([]int{1, 2, 3}, []int{2, 3, 4})
   ```

### 15. **Intersection**
   - Zwraca elementy, które znajdują się w obu listach.
   ```go
   result := funk.Intersection([]int{1, 2, 3}, []int{2, 3, 4})
   ```

### 16. **Union**
   - Łączy dwie listy, usuwając zduplikowane elementy.
   ```go
   result := funk.Union([]int{1, 2}, []int{2, 3})
   ```

### 17. **Zip**
   - Łączy elementy z dwóch list w pary.
   ```go
   result := funk.Zip([]string{"a", "b"}, []int{1, 2})
   ```

### 18. **Unzip**
   - Rozdziela listę par na dwie listy.
   ```go
   result := funk.Unzip([][]interface{}{{"a", 1}, {"b", 2}})
   ```

### 19. **IndexOf**
   - Zwraca indeks pierwszego wystąpienia elementu w liście.
   ```go
   result := funk.IndexOf([]string{"a", "b", "c"}, "b")
   ```

### 20. **LastIndexOf**
   - Zwraca indeks ostatniego wystąpienia elementu w liście.
   ```go
   result := funk.LastIndexOf([]string{"a", "b", "c", "b"}, "b")
   ```

### 21. **GroupBy**
   - Grupuje elementy listy według klucza zwróconego przez funkcję.
   ```go
   result := funk.GroupBy([]string{"apple", "banana", "apricot"}, func(x string) string {
       return x[:1]
   })
   ```

### 22. **Partition**
   - Dzieli listę na dwie na podstawie podanego warunku.
   ```go
   result := funk.Partition([]int{1, 2, 3, 4}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 23. **Shuffle**
   - Losowo przetasowuje elementy listy.
   ```go
   result := funk.Shuffle([]int{1, 2, 3, 4})
   ```

### 24. **Drop**
   - Usuwa pierwsze N elementów z listy.
   ```go
   result := funk.Drop([]int{1, 2, 3, 4}, 2)
   ```

### 25. **Take**
   - Zwraca pierwsze N elementów z listy.
   ```go
   result := funk.Take([]int{1, 2, 3, 4}, 2)
   ```

### 26. **Keys**
   - Zwraca klucze z mapy.
   ```go
   result := funk.Keys(map[string]int{"a": 1, "b": 2})
   ```

### 27. **Values**
   - Zwraca wartości z mapy.
   ```go
   result := funk.Values(map[string]int{"a": 1, "b": 2})
   ```

### 28. **ZipWith**
   - Łączy dwie listy na podstawie funkcji.
   ```go
   result := funk.ZipWith([]int{1, 2}, []int{3, 4}, func(a, b int) int {
       return a + b
   })
   ```

### 29. **Max**
   - Zwraca maksymalną wartość w liście.
   ```go
   result := funk.Max([]int{1, 2, 3})
   ```

### 30. **Min**
   - Zwraca minimalną wartość w liście.
   ```go
   result := funk.Min([]int{1, 2, 3})
   ```

### 31. **Compact**
   - Usuwa puste wartości (np. `nil`, `0`, `""`) z listy.
   ```go
   result := funk.Compact([]interface{}{0, 1, "", "hello", nil, "world"})
   ```

### 32. **Without**
   - Usuwa określone wartości z listy.
   ```go
   result := funk.Without([]int{1, 2, 3, 4}, 2, 3)
   ```

### 33. **Pull**
   - Usuwa określone elementy z listy.
   ```go
   result := funk.Pull([]int{1, 2, 3, 4}, 2)
   ```

### 34. **Reject**
   - Odrzuca elementy, które spełniają podany warunek.
   ```go
   result := funk.Reject([]int{1, 2, 3, 4}, func(x int) bool {
       return x%2 == 0
   })
   ```

### 35. **Merge**
   - Łączy dwie mapy w jedną.
   ```go
   result := funk.Merge(map[string]int{"a": 1}, map[string]int{"b": 

2})
   ```

### 36. **Omit**
   - Usuwa określone klucze z mapy.
   ```go
   result := funk.Omit(map[string]int{"a": 1, "b": 2}, "a")
   ```

### 37. **Pluck**
   - Wyciąga wartości z listy map na podstawie klucza.
   ```go
   result := funk.Pluck([]map[string]interface{}{
       {"name": "John", "age": 30},
       {"name": "Jane", "age": 25},
   }, "name")
   ```

### 38. **Invoke**
   - Wywołuje funkcję na każdym elemencie listy.
   ```go
   result := funk.Invoke([]string{"hello", "world"}, "ToUpper")
   ```

### 39. **Sort**
   - Sortuje listę elementów.
   ```go
   result := funk.Sort([]int{3, 1, 4, 2})
   ```

### 40. **IntersectionBy**
   - Zwraca elementy wspólne w dwóch listach na podstawie wartości zwróconej przez funkcję.
   ```go
   result := funk.IntersectionBy([]map[string]int{{"id": 1}, {"id": 2}}, []map[string]int{{"id": 2}}, "id")
   ```


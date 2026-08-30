#### Oto zestawienie 40 najpopularniejszych metod, które można wykonać na `DataFrame` w bibliotece Gota. Każda z tych metod oferuje różne funkcjonalności, które są użyteczne w analizie i przetwarzaniu danych:

1. **ReadCSV**: Wczytuje dane z pliku CSV do `DataFrame`.
   ```go
   df := dataframe.ReadCSV(file)
   ```

2. **WriteCSV**: Zapisuje `DataFrame` do pliku CSV.
   ```go
   df.WriteCSV(file)
   ```

3. **Head**: Zwraca pierwsze N wierszy `DataFrame`.
   ```go
   df.Head(10)
   ```

4. **Tail**: Zwraca ostatnie N wierszy `DataFrame`.
   ```go
   df.Tail(10)
   ```

5. **Filter**: Filtruje wiersze `DataFrame` na podstawie określonych warunków.
   ```go
   df.Filter(dataframe.F{Colname: "column", Comparator: series.Greater, Comparando: 100})
   ```

6. **GroupBy**: Grupuje wiersze `DataFrame` na podstawie wartości w jednej lub więcej kolumnach.
   ```go
   df.GroupBy("column1", "column2")
   ```

7. **Aggregation**: Agreguje dane w `DataFrame` po wcześniejszym grupowaniu.
   ```go
   df.GroupBy("column").Aggregation(
       []dataframe.AggregationType{dataframe.Aggregation_SUM},
       []string{"column_to_sum"},
   )
   ```

8. **Select**: Wybiera określone kolumny z `DataFrame`.
   ```go
   df.Select([]string{"column1", "column2"})
   ```

9. **Drop**: Usuwa określone kolumny z `DataFrame`.
   ```go
   df.Drop("column")
   ```

10. **Mutate**: Dodaje lub modyfikuje kolumny w `DataFrame`.
    ```go
    df.Mutate(dataframe.NewSeries("new_column", nil, []int{1, 2, 3}))
    ```

11. **Arrange**: Sortuje wiersze `DataFrame` na podstawie wartości w określonej kolumnie.
    ```go
    df.Arrange(dataframe.Sort("column"))
    ```

12. **Join**: Łączy dwa `DataFrame` na podstawie wartości w określonych kolumnach.
    ```go
    df.Join(otherDF, dataframe.InnerJoin, "key_column")
    ```

13. **Describe**: Zwraca statystyki opisowe dla kolumn liczbowych w `DataFrame`.
    ```go
    df.Describe()
    ```

14. **Summary**: Zwraca podsumowanie zawartości `DataFrame`, w tym liczność, średnią, odchylenie standardowe itp.
    ```go
    df.Summary()
    ```

15. **Crosstab**: Tworzy tabelę krzyżową (pivot table) dla danych w `DataFrame`.
    ```go
    dataframe.Crosstab(df, "column1", "column2")
    ```

16. **Apply**: Zastosowuje funkcję do każdej kolumny lub wiersza w `DataFrame`.
    ```go
    df.Apply(dataframe.ApplyOptions{
        SrcColname: "column",
        DstColname: "new_column",
        Fn: func(s series.Series) series.Series {
            return s.Add(10)
        },
    })
    ```

17. **Map**: Modyfikuje każdy wiersz `DataFrame` przy użyciu funkcji.
    ```go
    df.Map(func(row dataframe.Row) dataframe.Row {
        row.Set("column", row.Get("column").(int)+1)
        return row
    })
    ```

18. **InnerJoin**: Wykonuje złączenie wewnętrzne dwóch `DataFrame`.
    ```go
    df.InnerJoin(otherDF, "key_column")
    ```

19. **OuterJoin**: Wykonuje złączenie zewnętrzne dwóch `DataFrame`.
    ```go
    df.OuterJoin(otherDF, "key_column")
    ```

20. **Concat**: Łączy dwa lub więcej `DataFrame` pionowo (dodając wiersze).
    ```go
    dataframe.Concat([]dataframe.DataFrame{df1, df2})
    ```

21. **Unique**: Zwraca unikalne wartości w określonej kolumnie.
    ```go
    df.Col("column").Unique()
    ```

22. **IsNaN**: Sprawdza, czy kolumna zawiera wartości `NaN`.
    ```go
    df.Col("column").IsNaN()
    ```

23. **FillNaN**: Zastępuje wartości `NaN` w kolumnie określoną wartością.
    ```go
    df.Col("column").FillNaN(0)
    ```

24. **Max**: Zwraca maksymalną wartość w kolumnie liczbowej.
    ```go
    df.Col("column").Max()
    ```

25. **Min**: Zwraca minimalną wartość w kolumnie liczbowej.
    ```go
    df.Col("column").Min()
    ```

26. **Mean**: Oblicza średnią wartość w kolumnie liczbowej.
    ```go
    df.Col("column").Mean()
    ```

27. **Median**: Oblicza medianę w kolumnie liczbowej.
    ```go
    df.Col("column").Median()
    ```

28. **StdDev**: Oblicza odchylenie standardowe w kolumnie liczbowej.
    ```go
    df.Col("column").StdDev()
    ```

29. **Var**: Oblicza wariancję w kolumnie liczbowej.
    ```go
    df.Col("column").Var()
    ```

30. **Quantile**: Zwraca określony kwantyl dla kolumny liczbowej.
    ```go
    df.Col("column").Quantile(0.75)
    ```

31. **RollingMean**: Oblicza średnią kroczącą w kolumnie liczbowej.
    ```go
    df.Col("column").RollingMean(3)
    ```

32. **RollingSum**: Oblicza sumę kroczącą w kolumnie liczbowej.
    ```go
    df.Col("column").RollingSum(3)
    ```

33. **Shift**: Przesuwa wartości w kolumnie o określoną liczbę miejsc.
    ```go
    df.Col("column").Shift(1)
    ```

34. **DropNA**: Usuwa wiersze zawierające `NaN` w określonych kolumnach.
    ```go
    df.DropNA("column")
    ```

35. **ResetIndex**: Resetuje indeks `DataFrame` do domyślnych wartości (0, 1, 2, ...).
    ```go
    df.ResetIndex()
    ```

36. **SetIndex**: Ustawia kolumnę jako nowy indeks `DataFrame`.
    ```go
    df.SetIndex("column")
    ```

37. **Transpose**: Transponuje `DataFrame`, zamieniając wiersze z kolumnami.
    ```go
    df.Transpose()
    ```

38. **Slice**: Zwraca podzbiór `DataFrame` na podstawie indeksu wierszy.
    ```go
    df.Slice(0, 10)
    ```

39. **Subset**: Zwraca podzbiór `DataFrame` na podstawie indeksów kolumn.
    ```go
    df.Subset([]string{"column1", "column2"})
    ```

40. **Merge**: Łączy dwa `DataFrame` na podstawie wartości wspólnych kolumn.
    ```go
    df.Merge(otherDF, "key_column")
    ```

Te metody pozwalają na szerokie spektrum operacji na danych, od prostych operacji jak sortowanie i filtrowanie, po bardziej złożone, jak grupowanie, agregacja, czy też transpozycja. Dzięki tym funkcjom, `DataFrame` w Gota staje się potężnym narzędziem do analizy danych, które może zaspokoić potrzeby wielu różnych przypadków użycia w Go.
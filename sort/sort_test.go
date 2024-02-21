package sort

import (
	"math/rand"
	"testing"
)

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}
	return ar
}

func TestSort_SelectionSort(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{input: []int{5, 2, 3, 1, 9, 12}},
		{input: []int{3, 1, 7, 9}},
		{input: []int{-5, 17, -9, 11, 34, 48}},
	}

	ascendingOrder := func(in []int) bool {
		// Устанавливаем prev равным первому элементу слайса
		prev := in[0]

		// Проходимся по всем элементам, начиная со второго
		for _, val := range in[1:] {
			// Если текущий элемент меньше предыдущего,
			// значит слайс не отсортирован по возрастанию
			if val < prev {
				return false
			}

			// Обновляем prev перед проверкой следующего элемента
			prev = val
		}

		// Если цикл закончился, значит все элементы отсортированы по возрастанию
		return true

	}

	// Проходимся по всем тестовым случаям, определенным в начале функции.
	// Для каждого тестового случая:
	for _, tt := range testCases {
		// Вызываем SelectionSort, передавая ему входной слайс.
		SelectionSort(tt.input)

		// Проверяем, отсортирован ли слайс по возрастанию с помощью функции ascendingOrder.
		if !ascendingOrder(tt.input) {
			// Если слайс не отсортирован по возрастанию,
			// выводим сообщение об ошибке в лог теста.
			t.Errorf("SelectionSort(%v) = %v", tt.input, tt.input)
		}
	}

	// Проверяем работу SelectionSort на случайно сгенерированном слайсе.
	// Для этого создаем слайс заданной длины с помощью функции generateSlice.
	t.Run("check random slice", func(t *testing.T) {
		// Создаем слайс из 100 целых чисел в диапазоне от -100 до 100.
		randomSlice := generateSlice(100, 100)

		// Вызываем SelectionSort, передавая ему входной слайс.
		SelectionSort(randomSlice)

		// Проверяем, отсортирован ли слайс по возрастанию с помощью функции ascendingOrder.
		if !ascendingOrder(randomSlice) {
			// Если слайс не отсортирован по возрастанию,
			// выводим сообщение об ошибке в лог теста.
			t.Errorf("failed to sort a randomly generated slice")
		}
	})

}

// BenchmarkSelectionSort проводит бенчмаркинг алгоритма сортировки SelectionSort.
func BenchmarkSelectionSort(b *testing.B) {
	// Включаем отчет о выделении памяти во время бенчмарка.
	b.ReportAllocs()
	b.StopTimer() // Стоп таймера, чтобы он не мешал нашей функции.

	// Проводим бенчмаркинг для массивов различных размеров.
	// Для каждого размера:
	b.Run("small arrays", func(b *testing.B) {

		// Запускаем таймер, который будет измерять время выполнения цикла сортировки.
		b.StopTimer()

		// Выполняем цикл сортировки N раз, где N - число повторений, определенное тестом.
		for i := 0; i < b.N; i++ {

			// Создаем слайс из 10 целых чисел в диапазоне от 0 до 100 с помощью функции generateSlice.
			ar := generateSlice(10, 100)

			// Запускаем таймер, который будет измерять время выполнения сортировки.
			b.StartTimer()

			// Вызываем SelectionSort, передавая ему входной слайс.
			SelectionSort(ar)

			// Останавливаем таймер, чтобы измерить время выполнения сортировки.
			b.StopTimer()
		}
	})

	// Добавляем здесь другие бенчмарки для других размеров массивов.

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 10000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
		}
	})
}

func TestSort_BubbleSort(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{input: []int{5, 2, 3, 1, 9, 12}},
		{input: []int{3, 1, 7, 9}},
		{input: []int{-5, 17, -9, 11, 34, 48}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}
		return true
	}

	for _, tt := range testCases {
		BubbleSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Errorf("BubbleSort(%v) = %v", tt.input, tt.input)
		}
	}

	t.Run("check random slice", func(t *testing.T) {
		randomSlice := generateSlice(100, 100)
		BubbleSort(randomSlice)
		if !ascendingOrder(randomSlice) {
			t.Errorf("failed to sort a randomly generated slice")
		}
	})

}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			BubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			BubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 10000)
			b.StartTimer()
			BubbleSort(ar)
			b.StopTimer()
		}
	})
}

func TestSort_insertionSort(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{input: []int{5, 2, 3, 1, 9, 12}},
		{input: []int{3, 1, 7, 9}},
		{input: []int{-5, 17, -9, 11, 34, 48}},
	}
	ascendingOerder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}
		return true
	}
	for _, tt := range testCases {
		InsertionSort(tt.input)
		if !ascendingOerder(tt.input) {
			t.Errorf("SelectionSort(%v) = %v", tt.input, tt.input)
		}
	}
	t.Run("check panic on empty slice", func(t *testing.T) {
		randomSlice := generateSlice(100, 100)
		InsertionSort(randomSlice)
		if !ascendingOerder(randomSlice) {
			t.Errorf("failed to sort a randomle generated slice")
		}
	})
}

// Benchmarking functions
func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("medium arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 10000)
			b.StartTimer()
			InsertionSort(ar)
			b.StopTimer()
		}
	})

}

func TestSort_QuickSort(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{input: []int{5, 2, 3, 1, 9, 12}},
		{input: []int{3, 1, 7, 9}},
		{input: []int{-5, 17, -9, 11, 34, 48}},
	}
	ascendingOerder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}
		return true
	}
	for _, tt := range testCases {
		QuickSort(tt.input)
		if !ascendingOerder(tt.input) {
			t.Errorf("QuickSort(%v) = %v", tt.input, tt.input)
		}
	}
	t.Run("check panic on empty slice", func(t *testing.T) {
		randomSlice := generateSlice(100, 100)
		QuickSort(randomSlice)
		if !ascendingOerder(randomSlice) {
			t.Errorf("failed to sort a randomle generated slice")
		}
	})
}

// Benchmarking functions
func BenchmarkQuickSort(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("medium arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 10000)
			b.StartTimer()
			QuickSort(ar)
			b.StopTimer()
		}
	})

}

func TestSort_MergeSort(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{input: []int{5, 2, 3, 1, 9, 12}},
		{input: []int{3, 1, 7, 9}},
		{input: []int{-5, 17, -9, 11, 34, 48}},
	}
	ascendingOerder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}
		return true
	}
	for _, tt := range testCases {
		MergeSort(tt.input)
		if !ascendingOerder(tt.input) {
			t.Errorf("MergeSort(%v) = %v", tt.input, tt.input)
		}
	}
	t.Run("check panic on empty slice", func(t *testing.T) {
		randomSlice := generateSlice(100, 100)
		MergeSort(randomSlice)
		if !ascendingOerder(randomSlice) {
			t.Errorf("failed to sort a randomle generated slice")
		}
	})
}

// Benchmarking functions
func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("medium arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(1000, 10000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

}

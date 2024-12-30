package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

const (
	testString = "example"
	iterations = 100000
)

// Подготовка тестовых данных
func generateTestData() []string {
	data := make([]string, iterations)
	for i := 0; i < iterations; i++ {
		data[i] = testString
	}
	return data
}

// Конкатенация с помощью оператора +
func BenchmarkPlusOperator(b *testing.B) {
	data := generateTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result string
		for _, s := range data {
			result += s
		}
	}
}

// Конкатенация с помощью strings.Builder
func BenchmarkStringsBuilder(b *testing.B) {
	data := generateTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.Grow(len(testString) * iterations) // Предварительное выделение памяти
		for _, s := range data {
			builder.WriteString(s)
		}
		_ = builder.String()
	}
}

// Конкатенация с помощью bytes.Buffer
func BenchmarkBytesBuffer(b *testing.B) {
	data := generateTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.Grow(len(testString) * iterations) // Предварительное выделение памяти
		for _, s := range data {
			buffer.WriteString(s)
		}
		_ = buffer.String()
	}
}

// Конкатенация с помощью fmt.Sprintf
func BenchmarkSprintf(b *testing.B) {
	data := generateTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result string
		for _, s := range data {
			result = fmt.Sprintf("%s%s", result, s)
		}
	}
}

// Добавим функцию для вывода результатов
func BenchmarkCompareAll(b *testing.B) {
	data := generateTestData()
	b.ResetTimer()

	results := make(map[string]time.Duration)

	// Тестируем оператор +
	start := time.Now()
	var result string
	for _, s := range data {
		result += s
	}
	results["Оператор +"] = time.Since(start)

	// Тестируем strings.Builder
	start = time.Now()
	var builder strings.Builder
	builder.Grow(len(testString) * iterations)
	for _, s := range data {
		builder.WriteString(s)
	}
	_ = builder.String()
	results["strings.Builder"] = time.Since(start)

	// Тестируем bytes.Buffer
	start = time.Now()
	var buffer bytes.Buffer
	buffer.Grow(len(testString) * iterations)
	for _, s := range data {
		buffer.WriteString(s)
	}
	_ = buffer.String()
	results["bytes.Buffer"] = time.Since(start)

	// Тестируем fmt.Sprintf
	start = time.Now()
	result = ""
	for _, s := range data {
		result = fmt.Sprintf("%s%s", result, s)
	}
	results["fmt.Sprintf"] = time.Since(start)

	// Выводим результаты в отформатированном виде
	fmt.Printf("\nРезультаты сравнения методов конкатенации строк (%d итераций):\n", iterations)
	fmt.Println("============================================================")
	fmt.Printf("%-20s | %-20s | %-15s\n", "Метод", "Время выполнения", "Относительно Builder")
	fmt.Println("------------------------------------------------------------")

	builderTime := results["strings.Builder"]
	for method, duration := range results {
		ratio := float64(duration) / float64(builderTime)
		fmt.Printf("%-20s | %-20s | %.2fx медленнее\n",
			method,
			duration.Round(time.Microsecond),
			ratio)
	}
	fmt.Println("============================================================")
}

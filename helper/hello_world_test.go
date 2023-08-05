package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Slamet",
			request: "Slamet",
		},
		{
			name:    "Mahendra",
			request: "Mahendra",
		},
		{
			name:    "Putra",
			request: "Mahendra Putra",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			HelloWorld(benchmark.request)
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Slamet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Slamet")
		}
	})
	b.Run("Mahendra", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Mahendra")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Slamet")
	}
}

func BenchmarkHelloWorldMahendra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Slamet")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Slamet",
			request:  "Slamet",
			expected: "Hello Slamet",
		},
		{
			name:     "Mahendra",
			request:  "Mahendra",
			expected: "Hello Mahendra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Slamet", func(t *testing.T) {
		result := HelloWorld("Slamet")
		require.Equal(t, "Hello Slamet", result, "Result must be 'Hello Slamet'")
	})
	t.Run("Mahendra", func(t *testing.T) {
		result := HelloWorld("Mahendra")
		require.Equal(t, "Hello Mahendra", result, "Result must be 'Hello Mahendra'")
	})
}

func TestMain(m *testing.M) {
	//Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//After
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Windows")
	}
	// not run
	result := HelloWorld("Slamet")
	require.Equal(t, "Hello Slamet", result, "Result must be 'Hello Slamet'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Slamet")
	require.Equal(t, "Hello Slamet", result, "Result must be 'Hello Slamet'")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Slamet")
	assert.Equal(t, "Hello Slamet", result, "Result must be 'Hello Slamet'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Slamet")

	if result != "Hello Slamet" {
		// unit test failed
		t.Error("Result must be 'Hello Slamet'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldMahendra(t *testing.T) {
	result := HelloWorld("Mahendra")

	if result != "Hello Mahendra" {
		// unit test failed
		t.Fatal("Result must be 'Hello Mahendra'")
	}
	fmt.Println("TestHelloWorldMahendra Done")
}

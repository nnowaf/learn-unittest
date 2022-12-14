package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nowaf")
	if result != "Hello Nowaf" {
		t.Errorf("HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Nowaf")
	}
}

func TestHelloWorldWleo(t *testing.T) {
	result := HelloWorld("Wleo")
	if result != "Hello Wleo" {
		t.Errorf("HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Wleo")
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Nowaf")
	assert.Equal(t, "Hello Nowaf", result, "HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Nowaf")
	fmt.Println("TestHelloWorldAssert passed")
}

func TestHelloWorldRequered(t *testing.T) {
	result := HelloWorld("Nowaf")
	require.Equal(t, "Hello Nowaf", result, "HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Nowaf")
	fmt.Println("TestHelloWorldRequered passed")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run this test on Windows")
	}
	result := HelloWorld("Nowaf")
	assert.Equal(t, "Hello Nowaf", result, "HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Nowaf")
	fmt.Println("TestHelloWorldAssert passed")
}

// this function is like Init in LifeCycle Flutter, so it will run before all tests
// and remember this TestMain always running in the same package,
// so if you want to run it in another package, you should move it to another package, or created TestMain
func TestMain(m *testing.M) {
	fmt.Println("Starting tests")
	m.Run()
	fmt.Println("Tests finished")
}

func TestSubTest(t *testing.T) {
	t.Run("Nowaf", func(t *testing.T) {
		result := HelloWorld("Nowaf")
		assert.Equal(t, "Hello Nowaf", result, "HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Nowaf")
		fmt.Println("TestHelloWorldAssert passed")
	})
	t.Run("Wleo", func(t *testing.T) {
		result := HelloWorld("Wleo")
		assert.Equal(t, "Hi Wleo", result, "HelloWorld was incorrect, got: %s, want: %s.", result, "Hello Wleo")
		fmt.Println("TestHelloWorldAssert passed")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Nowaf",
			request:  "Nowaf",
			expected: "Hello Nowaf",
		},
		{
			name:     "Wleo",
			request:  "Wleo",
			expected: "Hello Wleo",
		},
		{
			name:     "Jokowi",
			request:  "Jokowi",
			expected: "Yah Jokowi gak ada di sini",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "HelloWorld was incorrect, got: %s, want: %s.", result, test.expected)
		})
	}
}

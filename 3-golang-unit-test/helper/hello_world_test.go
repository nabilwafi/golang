package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Testing Start");

	m.Run()

	fmt.Println("Testing Done");
}

func BenchmarkTableBenchmark(b *testing.B) {
	bench := []struct {
		Name string
		Request string
	}{
		{
			Name: "HelloWorld(Nadin)",
			Request: "Nadin",
		},
		{
			Name: "HelloWorld(Fira)",
			Request: "Fira",
		},
	}

	for _, val := range bench {
		b.Run(val.Name, func (b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(val.Request)
			}
		})
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Nabil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nabil")
		}
	})

	b.Run("NabilWafi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("NabilWafi")
		}
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		Name string
		Request string
		Expected string
		ErrorMessage string
	} {
		{
			Name: "HelloWorld(Nadin)",
			Request: "Nadin",
			Expected: "Hello, Nadin",
			ErrorMessage: "return must be 'hello, nadin'",
		},
		{
			Name: "HelloWorld(Fira)",
			Request: "Fira",
			Expected: "Hello, Fira",
			ErrorMessage: "return must be 'hello, Fira'",
		},
		{
			Name: "HelloWorld(Jaka)",
			Request: "Jaka",
			Expected: "Hello, Jaka",
			ErrorMessage: "return must be 'hello, Jaka'",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := HelloWorld(test.Request)
			assert.Equal(t, test.Expected, res, test.ErrorMessage)
		})
	}

}

// SUB TEST
func TestSubTest(t *testing.T) {
	t.Run("Wafi", func (t *testing.T) {
		res := HelloWorld("Wafi")
		assert.Equal(t, res, "Hello, Wafi", res, "return mus be 'Hello, Wafi'")
	})

	t.Run("Wafi", func (t *testing.T) {
		res := HelloWorld("Restu")
		assert.Equal(t, res, "Hello, Restu", res, "return mus be 'Hello, Restu'")
	})
}

func TestHelloWorld(t *testing.T) {
	res := HelloWorld("Nabil")
	assert.Equal(t, "Hello, Nabil", res, "return must be 'Hello, Nabil'")
}

func TestHelloWorldFaris(t *testing.T) {
	res := HelloWorld("Nabil")
	assert.Equal(t, "Hello, Faris", res, "return must be 'Hello, Faris'")
}
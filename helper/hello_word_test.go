package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Irfan", func(t *testing.T) {
		result:= HelloWorld("Irfan")
		require.Equal(t, "Hello Irfan", result, "Result must be 'Hello Irfan'")
        
    })
	t.Run("Zidni", func(t *testing.T) {
		result:= HelloWorld("Zidni")
		require.Equal(t, "Hi Zidni", result, "Result must be 'Hello Zidni'")
        
    })
}

func TestTableHelloWorld(t *testing.T) {
    tests := []struct {
        name     string
        request  string
        expected string
    }{
        {
            name:     "Irfan",
            request:  "Irfan",
            expected: "Hello Irfan",
        },
        {
            name:     "Zidni",
            request:  "Zidni",
            expected: "Hello Zidni",
        },
        {
            name:     "Muhammad",
            request:  "Muhammad",
            expected: "Hello Muhammad",
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := HelloWorld(test.request)
            require.Equal(t, test.expected, result, "Result must be 'Hello "+test.name+"'")
        })
    }
}
		



func TestMain(m *testing.M) {
	// before
	m.Run()
	fmt.Println("BEFORE UNIT TESTING")

	//after
	fmt.Println("AFTER UNIT TESTING")
}

func TestSkip(t *testing.T){
	if runtime.GOOS=="darwin" {
		t.Skip("Can not run Linux OS")

	}
	result:= HelloWorld("Irfan")
	require.Equal(t, "Hello Irfan", result, "Result must be 'Hello Irfan'")
}
func TestHelloWorldRequire(t *testing.T){
	result:= HelloWorld("Irfan")
	require.Equal(t, "Hello Irfan", result, "Result must be 'Hello Irfan'")
	fmt.Println("Test Hello World with Assert Done")
}





func TestHelloWorldAssert(t *testing.T){
	result:= HelloWorld("Irfan")
	assert.Equal(t, "Hello Irfan", result, "Result must be 'Hello Irfan'")
	fmt.Println("Test Hello World with Assert Done")
}

func TestHelloWorldIrfan(t *testing.T) {
	result:= HelloWorld("Irfan")

	if result != "Hello Irfan"{
		// error
		t.Error("Result must be 'Hello Irfan'")
	}
	fmt.Println("TestHelloWordIrfan Done")
}

func TestHelloWorldZidni(t *testing.T) {
	result:= HelloWorld("Zidni")

	if result != "Hello Zidni"{
		// error // ketika menemui FailNow, otomatis program teehrnti testnya
		t.Fatal("Result must be 'Hello Zidni'")
	}
	fmt.Println("TestHelloWordZidni Done")
}
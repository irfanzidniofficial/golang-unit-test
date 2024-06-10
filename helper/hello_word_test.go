package helper

import (
	"fmt"
	"testing"
)

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
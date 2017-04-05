package randHashGenerator

import (
	"testing"
)

var hashGen *HashGenerator
var randStr string

func Test_Get_1(t *testing.T) {
	length := 8
	hashGen = &HashGenerator{
		make(chan string),
		length,
	}
	hashGen.init()
	randStr = hashGen.Get()
	if len(randStr) == length {
		t.Log("第1个测试:", randStr)
	} else {
		t.Error("第1个测试没通过")
	}
}

func Test_Get_2(t *testing.T) {
	length := 10
	hashGen = &HashGenerator{
		make(chan string),
		length,
	}
	hashGen.init()
	randStr = hashGen.Get()
	if len(randStr) == length {
		t.Log("第2个测试:", randStr)
	} else {
		t.Error("第2个测试没通过")
	}

}

func Test_Get_3(t *testing.T) {
	length := 16
	hashGen = &HashGenerator{
		make(chan string),
		length,
	}
	hashGen.init()
	randStr = hashGen.Get()
	if len(randStr) == length {
		t.Log("第3个测试:", randStr)
	} else {
		t.Error("第3个测试没通过")
	}
}

func Test_Get_4(t *testing.T) {
	length := 7
	hashGen := NewHashGen(length)
	for i := 0; i < 100; i++ {
		randStr = hashGen.Get()
		t.Log(i, randStr)
	}
}

func NewHashGen(length int) *HashGenerator {
	hashGen = &HashGenerator{
		make(chan string),
		length,
	}
	hashGen.init()
	return hashGen
}

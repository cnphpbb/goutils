package randHashGenerator

import (
	"testing"
)

var hashGen *HashGenerator
var randStr string

func init() {
	length := 9
	hashGen = NewHashGen(length)
}

func Test_Get_0(t *testing.T) {
	t.Log("第0个测试开始")
	for i := 0; i < 100; i++ {
		randStr = hashGen.Get()
		t.Log(i, randStr)
	}
	t.Log("第0个测试结束")
}

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
	t.Log("第4个测试开始")
	for i := 0; i < 100; i++ {
		randStr = hashGen.Get()
		t.Log(i, randStr)
	}
	t.Log("第4个测试结束")
}

func Test_Get_5(t *testing.T) {
	length := 8
	hashGen = new(HashGenerator).New(length)
	randStr = hashGen.Get()
	if len(randStr) == length {
		t.Log("第5个测试:", randStr)
	} else {
		t.Error("第5个测试没通过")
	}
}

func Test_Get_6(t *testing.T) {
	length := 8
	hashGen = NewHashGen(length)
	randStr = hashGen.Get()
	if len(randStr) == length {
		t.Log("第6个测试:", randStr)
	} else {
		t.Error("第6个测试没通过")
	}
}

package goutils

import (
	"testing"
)

//基准测试
func Benchmark_Get_1(b *testing.B) {
	length := 8
	hashGen := NewHashGen(length)
	for i := 0; i < b.N; i++ {
		hashGen.Get()
	}
}

// 测试并发效率
func Benchmark_Get_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		length := 8
		hashGen := NewHashGen(length)
		for pb.Next() {
			hashGen.Get()
		}
	})
}

//功能测试
func Test_Get_1(t *testing.T) {
	length := 8
	hashGen := NewHashGen(length)
	randStr := hashGen.Get()
	if len(randStr) == length {
		t.Log("功能测试:", randStr)
	} else {
		t.Error("功能测试没通过")
	}
}

func Test_New_1(t *testing.T) {
	length := 8
	hashGen := new(HashGenerator).New(length)
	t.Log("New function 测试:", hashGen.length)
}

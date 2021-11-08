package goutils

import "testing"

//功能测试
func Test_ConversionInPath_1(t *testing.T) {
	inPath := "~/test.log"
	rs_inPath := ConversionInPath(inPath)
	t.Log("New function 测试:", rs_inPath)
}

func Test_ConversionInPath_2(t *testing.T) {
	inPath := "$HOME/test/test.log"
	rs_inPath := ConversionInPath(inPath)
	t.Log("New function 测试:", rs_inPath)
}

package goutils

import "testing"

//功能测试
func Test_TitleCasedName_1(t *testing.T) {
	str := "ppos-vue-reports"
	newstr := TitleCasedName(str)
	t.Log("TitleCasedName function 测试:", newstr)
	if newstr != "PposVueReports" {
		t.Error("功能测试没通过")
	}
}
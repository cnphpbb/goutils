# go_utils
整理出来的一些比较不错的go的pkg包

### func:: HashGenerator

生成随机Hash值  
Rand Hash Generator 

* 采用了`crypto/rand`来生产0～9，A～Z，a~z的随机字符串。可以用于生成密码，盐值，命名文件等。
* 采用了`golang`的特性`goroutine`及`channel`。

例子：
```golang
hashGen := &HashGenerator{
	hashGetter: make(chan string),
	length: 8,
}

hashGen.init()

hashGen.Get()
```   
把以上代码加载上就可以获得随机字符串。

详细的使用方法可以参考[单元测试文件](./hashGenerator_test.go)。

### func:: TitleCasedName

例子：
```golang
_str := "is-check-hash"
str := TitleCasedName(_str)
fmt.Println(str) // IsCheckHash
```   


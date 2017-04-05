Rand Hash Generator
========

采用了`crypto/rand`来生产0～9，A～Z，a~z的随机字符串。可以用于生成密码，盐值，命名文件等。

使用方法：
```golang
hashGen := &HashGenerator{
	hashGetter: make(chan string),
	length: 8,
}

hashGen.init()

hashGen.Get()
```   
把以上代码加载上 就可以获得随机字符串。
package pack

/*
	工作去结构
	$GOPATH/
		src/   存储源代码，每个包应该在这个文件夹的子文件夹下
		bin/   保存便宜后的可执行文件
		pkg/   用于保存编译后的包的目标文件

	下载包
	go get 路径     使用...安装路径下全部包
	eg； go get github.com/golang/lint/golint
	-v 版本控制
	-u 获取最新版本

	构建包
	go build 默认情况下 go build 命令会构建指定的包和它以来的包,然后丢弃除了可执行文件之外的中间便编译结果,到项目
	过大的时候会导致每次构建都会消耗较长时间。
	go install 作用和go build相似，但是它会保留编译的中间结果到($GOPAHT/pkg/...) 可执行文件会被保存在$GOPATH/bin
	下。
	go build -i 安装每个包所依赖的包。

	工具

	go doc 包名   获取包文档
	go doc 包名.对象名 获取这个对象的文档
	go list 路径  查询可用包，打印它的导入信息
	go list ...xxx... 查询和xxx相关的所有包
	go list -json xxx  查询xxx包的所有信息
*/

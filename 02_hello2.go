protoc -I . --go_out=plugins=micro:. proto/example/example.proto
micro new --type "web" iot/webiot
micro new --type "srv" iot/PostLogin
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
#二进制编译
$ CGO_ENABLED=0 GOOS=linux /usr/local/go/bin/go build -a -installsuffix cgo -ldflags '-w' -i -o webiot-web ./main.go
#编译需要在root账户下进行
#指明cgo工具是否可用的标识在这里表示禁用
CGO_ENABLED=0
#目标平台（编译后的目标平台）的操作系统（darwin、freebsd、linux、windows）
GOOS=linux
#由于没有在root下安装go所以我们需要使用go的绝对路径进行使用
/usr/local/go/bin/go build
#强制重新编译所有涉及的go语言代码包
-a
#为了使当前的输出目录与默认的编译输出目录分离，可以使用这个标记。此标记的值会作为结果文件的父目录名
称的后缀。
-installsuffix
cgo
# 给 cgo指定命令
-ldflags
#关闭所有警告信息
'-w'
#标志安装目标的依赖包。
-i
#命名
-o ihomeweb
#编译的main.go地址
./main.go



FROM alpine:3.2
#拷贝文件
ADD conf /conf
#拷贝文件
ADD html /html
#拷贝二进制
ADD ihomeweb /ihomeweb
WORKDIR /
ENTRYPOINT [ "/ihomeweb" ]
EXPOSE 8999



FROM alpine:3.2
ADD conf /conf
ADD getarea-srv /getarea-srv
ENTRYPOINT [ "/getarea-srv" ]


consul:
#覆盖启动后的执行命令
command: agent -server -bootstrap-expect=1 -node=node1 -client 0.0.0.0 -ui -
bind=0.0.0.0 -join 127.0.0.2
#command: agent -server -bootstrap -rejoin -ui
#镜像：镜像名称:版本号
image: consul:latest
#主机名
hostname: "registry"
#暴露端口
ports:
- "8300:8300"
- "8400:8400"
- "8500:8500"
- "8600:53/udp"
#web主页
web:
#覆盖启动后的执行命令
command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10
web
#镜像构建的dockerfile文件地址
build: ./ihomeweb
links:
- consul
ports:


- "8999:8999"
#获取地区
getarea:
#覆盖启动后的执行命令
command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10
web
#镜像构建的dockerfile文件地址
build: ./getarea
links:
- consul
#注册三部曲
getimagecd:
#覆盖启动后的执行命令
command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10
web
#镜像构建的dockerfile文件地址
build: ./getimagecd
links:
- consul
getsmscd:
#覆盖启动后的执行命令
command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10
web
#镜像构建的dockerfile文件地址
build: ./getsmscd
links:
- consul
postret:
#覆盖启动后的执行命令
command: --registry_address=registry:8500 --register_interval=5 --register_ttl=10
web
#镜像构建的dockerfile文件地址
build: ./postret
links:
- consul

golang 自动下载所有依赖包
如何自动下载所有依赖包？
大部分情况下大家下载 Go 项目都是使用go get命令，它除了会下载指定的项目代码，还会去下载这个项目所依赖的所有项目。

但是有的时候我们的项目由于各种原因并不是通过go get下载的，是通过git clone下载的，这样代码下下来就没有依赖包了，没办法编译通过的。这样的话怎么办呢？

go get -d -v ./...
-d标志只下载代码包，不执行安装命令；
-v打印详细日志和调试日志。这里加上这个标志会把每个下载的包都打印出来；
./...这个表示路径，代表当前目录下所有的文件。



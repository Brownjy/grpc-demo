# grpc golang准备工作
+ 安装protoc编译器
```text
https://github.com/protocolbuffers/protobuf/releases?q=&expanded=true
```
+ 安装protoc-gen-go 配合protoc编译器,生成go文件
```text
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
+ 安装grpc
```text
go get google.golang.org/grpc
```
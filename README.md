# mn-hosted
mn-hosted基于[micro](https://github.com/micro/micro)，编写主节点托管服务使用micro框架范例。
# 1.环境安装：
## 1.golang
go1.12.2,请自行搜索安装
## 2.protobuf
[下载](https://github.com/protocolbuffers/protobuf/archive/v3.6.0.1.zip)或wget https://github.com/protocolbuffers/protobuf/archive/v3.6.0.1.zip
./autogen.sh && ./configure && make && make check  
sudo make install    
sudo ldconfig  
## 3.protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
## 4.protoc-gen-micro
go get github.com/micro/protoc-gen-micro
## 5.依赖环境启动(仅方便测试)
依赖的consul（注册中心）、nsq（消息队列），如无docker软件请自行安装
### 启动
docker-compose up &
# 2.测试：
## 1.运行rpcsvr
go run rpcsvr/user/main.go
## 2.运行gateway
micro --registry=consul --registry_address=127.0.0.1:8500 --server_advertise=192.168.0.118:8080 api --handler=rpc --address=0.0.0.0:8080 --namespace=go.mnhosted.srv
## 3.测试
curl -H 'Content-Type: application/json' -d '{"name": "john", "passwd":"123456"}' "http://localhost:8080/user/User/SignIn"
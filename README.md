# mn-hosted
mn-hosted基于[micro](https://github.com/micro/micro)，编写主节点托管服务使用micro框架范例。
# 1.环境安装：
## 1.golang
go1.12.2,请自行搜索安装
## 2.protobuf
[下载](https://github.com/protocolbuffers/protobuf/archive/v3.6.0.1.zip)或wget https://github.com/protocolbuffers/protobuf/archive/v3.6.0.1.zip
./autogen.sh && ./configure && make  
sudo make install    
sudo ldconfig  
## 3.protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
## 4.protoc-gen-micro
go get github.com/micro/protoc-gen-micro
## 5.依赖环境启动(仅方便测试)
依赖的consul（注册中心）、nsq（消息队列），如无docker软件请自行安装
### 编绎及启动
cd bin
./install.sh
cd ..
docker-compose up &
cd bin
./start.sh
# 2.测试：
## 1.运行rpcsvr
### go run rpcsvr/user/main.go
### go run apisvr/user/main.go(如采用参数形式需要启动)
## 2.运行gateway
### micro --registry=consul --registry_address=127.0.0.1:8500 --server_advertise=192.168.0.118:8080 api --handler=rpc --address=0.0.0.0:8080 --namespace=go.mnhosted.srv
### 或micro --registry=consul --registry_address=127.0.0.1:8500 --server_advertise=192.168.0.118:8080 api --handler=api --address=0.0.0.0:8080 --namespace=go.mnhosted.api(如采用参数形式需要启动)
## 3.测试
### curl -H 'Content-Type: application/json' -d '{"account": "john", "passwd":"123456"}' "http://localhost:8080/user/User/SignUp"
### 或curl "http://localhost:8080/user/User/SignUp?account=lixu&passwd=123456"
### curl -H 'Content-Type: application/json' -d '{"account": "john", "passwd":"123456"}' "http://localhost:8080/user/User/SignIn"
### curl -H 'Content-Type: application/json' -d '{"userId": 1, "mnAddress":"dddd", "outputIndex":1, "vps":"", "alias":"","txId":"", "mnConf":""}' "http://localhost:8080/node/Masternode/New"


# TODO
## 1.gateway鉴权，可参考https://blog.csdn.net/linux_Allen/article/details/89914647，后面处理
## 2.忘记密码处理

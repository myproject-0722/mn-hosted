go build ../rpcsvr/wallet/rpc-wallet-srv.go
go build ../rpcsvr/order/rpc-order-srv.go
go build ../rpcsvr/user/rpc-user-srv.go
go build -gcflags "-N -l" ../rpcsvr/node/rpc-node-srv.go
go build ../apisvr/user/api-user-srv.go
go build ../apisvr/node/api-node-srv.go
go build ../apisvr/order/api-order-srv.go
#go build ../micro/
go build -o time-srv ../timesrv/
go build -gcflags "-N -l" ../notifysrv/alipay/alipay-notify.go 

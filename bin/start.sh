dashd -testnet -daemon
#./micro --registry=consul --registry_address=127.0.0.1:8500 --server_advertise=127.0.0.1:8000 api --handler=api --address=0.0.0.0:8000 --namespace=go.mnhosted.api &
./micro --registry=consul --registry_address=161.189.42.122:8500 --server_advertise=127.0.0.1:8000 api --handler=api --address=0.0.0.0:8000 --namespace=go.mnhosted.api &
./api-node-srv &
./api-user-srv &
./rpc-node-srv &
./rpc-order-srv &
./rpc-user-srv &
./rpc-wallet-srv &
./time-srv &

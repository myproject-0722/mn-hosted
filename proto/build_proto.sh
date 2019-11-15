protoc --proto_path=./user --micro_out=./user --go_out=./user ./user/*.proto
protoc --proto_path=./node --micro_out=./node --go_out=./node ./node/*.proto
protoc --proto_path=./wallet --micro_out=./wallet --go_out=./wallet ./wallet/*.proto
protoc --proto_path=./order --micro_out=./order --go_out=./order ./order/*.proto
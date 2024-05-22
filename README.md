# pitaya_study
study https://github.com/topfreegames/pitaya 

protoc version: 3.20.3

pitaya库中的proto生成go：
    进入protos目录下执行：
    protoc.exe -I ../pitaya/pitaya-protos/ ../pitaya/pitaya-protos/*.proto --go_out=../pitaya/protos --go-grpc_out=../pitaya/protos

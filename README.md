## 编写链码
1. 功能列表
    - 初始化链码，设置账户余额
    - 调用链码进行转账
    - 查询某个账户余额
    - 删除某个账户

## 安装并运行链码

- 将编写的链码放到
/home/GoFile/src/github.com/hyperledger/fabric/scripts/fabric-samples/chaincode/developChaincode文件夹下
- docker exec -it chaincode sh 进入chaincode容器
- ls即可查看到自己的链码
- 编译自己的链码再启动
   以官方的例子为例
        go build
        
        CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=testcc:0 ./developChaincode
        
- 安装链码
   新启动终端<br>
   docker exec -it cli bash <br>
   peer chaincode install -p chaincodedev/chaincode/developChaincode -n testcc -v 0
   `注意这里的-p 对应的值会默认加上/opt/gopath/src/`,所以指定的时候需要注意
   
- 实例化代码
    
    peer chaincode instantiate -n testcc -C myc -v 0 -c '{"Args":["init"]}'

- 调用链码
    
    调用转账方法<br>
    peer chaincode invoke -n testcc -c '{"Args":["invoke","a","b","300"]}' -C myc
    <br>调用查询方法<br>
    peer chaincode invoke -n testcc -c '{"Args":["query","a"]}' -C myc
    <br>调用删除方法<br>
    peer chaincode invoke -n testcc -c '{"Args":["delete","b"]}' -C myc
    
     
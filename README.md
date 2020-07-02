# fabricStudy
1. Go链码开发与编写 https://www.cnblogs.com/zongmin/p/11874792.html
 

## develop chaincode

- two interface's func

    - init(xx ChaincodeStudInterface)
    
        when instantiate/upgrade chaincode,it works
    - invoke(xx ChaincodeStudInterface)
    
        when invoke/query chaincode,it works

- two packages

    - shim
    
        this package is used to query/modify data status, context and call other chaincode's API,
        chaincode use funcs which shim.ChaincodeStub provides to read/write the status of ledger.
        
    - peer
    
        this package contains some API to get the response of the call of chaincode.
        peer.Response contains the response info.
        
   
## 链码相关的API

### shim

shim包提供了如下几种类型的接口，ChaincodeStubInterface 

- 参数解析api 
- 账本状态数据操作api 
- 交易信息获取api
- 对私密数据操作的api （private data)
- 其他api（包括事件设置，调用其他链码操作）


## 如何在开发模式测试链码

https://www.jianshu.com/p/21231847fe81

1. fabric提供了基础的测试环境

    进入fabric-sample/chaincode-docker-devmode文件夹下,执行以下命令

        docker-compose -f docker-compose-simple.yaml up

    如果报: ERROR: An HTTP request took too long to complete,注释掉`docker-compose
-simple.yaml`的tty:true即可。

2. 安装并运行链码

- 将编写的链码放到
/home/GoFile/src/github.com/hyperledger/fabric/scripts/fabric-samples/chaincode/developChaincode文件夹下
- docker exec -it chaincode sh 进入chaincode容器
- ls即可查看到自己的链码
- 编译自己的链码再启动
   以官方的例子为例

        cd chaincode_example02/go
        go build -o chaincode_example02
        CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./chaincode_example02
        
- 安装链码
   新启动终端<br>
   docker exec -it cli bash <br>
   peer chaincode install -p chaincodedev/chaincode/developChaincode -n testcc -v 0
   `注意这里的-p 对应的值会默认加上/opt/gopath/src/`,所以指定的时候需要注意
   
- 实例化代码
    
    peer chaincode instantiate -n testcc -C myc -v 0 -c '{"Args":["init"]}'

- 调用链码

     
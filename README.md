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

- 将编写的链码放到fabric-sample/chaincode/文件夹下
- ls即可查看到自己的链码

#### 下载并安装
```
go get -u github.com/gin-gonic/gin
```

#### 在代码中导入它
```
import "github.com/gin-gonic/gin"
```

#### 运行
```
cd $GOPATH/src/micro/ch01/
go build -o gin
./gin
```

#### 校验
```
//postman、curl
GET: curl http://127.0.0.1:8080/
GET: curl http://127.0.0.1:8080/person/66
POST: curl -d "first_name=will&last_name=wu" http://127.0.0.1:8080/person
```

#### 参考资料
```
https://github.com/BroQiang/gin/blob/master/README_ZH.md
```

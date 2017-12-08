[![image](http://b2oks-cover.b0.upaiyun.com/default/cyberlife-logo.jpg)](http://cyber-life.cn)
# go-api-demo

### 安装golang
    # yum install go

### 编辑环境变量
    $ vi ~/.bashrc
    export GOROOT=/usr/lib/golang
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

### 安装beego
    $ go get github.com/astaxie/beego
    $ go get github.com/beego/bee

### 安装依赖包
    $ go get -u github.com/go-sql-driver/mysql
    $ go get github.com/satori/go.uuid
    $ go get gopkg.in/mgo.v2
    $ go get github.com/bradfitz/gomemcache/memcache
    $ go get github.com/casbin/casbin

### 创建工程
    $ bee api go-api-demo
    $ cd go-api-demo/
    $ bee run -gendoc=true -downdoc=true

### 启动
    $ cd rpc
    $ go run rpc_server.go
    $ bee run -gendoc=true -downdoc=true


## multiple http fetches in parallel (for mysql)
    ./http_load -p 100 -s 10 url.txt
    $ vi url.txt
    http://go.domicake.com/api/test
    
    // test for mysql
    1709 fetches, 100 max parallel, 1.98757e+06 bytes, in 10 seconds
    1163 mean bytes/connection
    170.9 fetches/sec, 198757 bytes/sec
    msecs/connect: 0.0987104 mean, 3.829 max, 0.042 min
    msecs/first-response: 553.364 mean, 3483.48 max, 103.149 min
    HTTP response codes:
      code 200 -- 1709


## multiple http fetches in parallel (for mongodb)
    ./http_load -p 10 -s 10 url.txt
    $ vi url.txt
    http://go.domicake.com/api/person
    
    6269 fetches, 100 max parallel, 677052 bytes, in 10 seconds
    108 mean bytes/connection
    626.9 fetches/sec, 67705.2 bytes/sec
    msecs/connect: 0.111617 mean, 3.937 max, 0.028 min
    msecs/first-response: 124.162 mean, 8285.98 max, 33.903 min
    HTTP response codes:
      code 200 -- 6269


## multiple http fetches in parallel (for memcache)
    ./http_load -p 100 -s 10 url.txt
    $ vi url.txt
    http://go.domicake.com/api/memcache
    
    15866 fetches, 100 max parallel, 2.91934e+06 bytes, in 10 seconds
    184 mean bytes/connection
    1586.6 fetches/sec, 291934 bytes/sec
    msecs/connect: 0.0603796 mean, 5.536 max, 0.022 min
    msecs/first-response: 62.6842 mean, 139.218 max, 1.171 min
    HTTP response codes:
      code 200 -- 15866


## multiple http fetches in parallel (for rpc)
    2649 fetches, 100 max parallel, 447681 bytes, in 10 seconds
    169 mean bytes/connection
    264.9 fetches/sec, 44768.1 bytes/sec
    msecs/connect: 0.117413 mean, 8.438 max, 0.03 min
    msecs/first-response: 345.759 mean, 3774.71 max, 66.315 min
    HTTP response codes:
      code 200 -- 2649

# go-api-demo

# golang-mysql

''# yum install go

$ vi ~/.bashrc
export GOROOT=/usr/lib/golang
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


$ go get github.com/astaxie/beego
$ go get github.com/beego/bee


$ bee api go-api-demo
$ cd go-api-demo/
$ bee run -gendoc=true -downdoc=true


$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/satori/go.uuid


## multiple http fetches in parallel
./http_load -p 100 -s 10 url.txt
$ vi url.txt
http://go.domicake.com/api/test

''# test for mysql
1709 fetches, 100 max parallel, 1.98757e+06 bytes, in 10 seconds
1163 mean bytes/connection
170.9 fetches/sec, 198757 bytes/sec
msecs/connect: 0.0987104 mean, 3.829 max, 0.042 min
msecs/first-response: 553.364 mean, 3483.48 max, 103.149 min
HTTP response codes:
  code 200 -- 1709


# golang-mongodb

$ go get gopkg.in/mgo.v2

''# test for mongodb
2571 fetches, 10 max parallel, 277668 bytes, in 10 seconds
108 mean bytes/connection
257.1 fetches/sec, 27766.8 bytes/sec
msecs/connect: 0.0540397 mean, 1.56 max, 0.018 min
msecs/first-response: 38.7706 mean, 325.844 max, 33.872 min
HTTP response codes:
  code 200 -- 2571


# golang-memcache
$ go get github.com/bradfitz/gomemcache/memcache

8686 fetches, 10 max parallel, 3.43822e+06 bytes, in 10 seconds
184 mean bytes/connection
1868.6 fetches/sec, 343822 bytes/sec
msecs/connect: 0.0410613 mean, 3.079 max, 0.018 min
msecs/first-response: 5.30374 mean, 19.324 max, 0.903 min
HTTP response codes:
  code 200 -- 18686

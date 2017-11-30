# go-api-demo

# yum install go

$ vi ~/.bashrc
export GOROOT=/usr/lib/golang
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


$ go get github.com/astaxie/beego
$ go get github.com/beego/bee


$ bee api go-api-demo
$ cd go-api-demo/
$ bee run -gendoc=true -downdoc=true


$ go get code.google.com/p/go-mysql-driver/mysql
$ go get github.com/satori/go.uuid


# multiple http fetches in parallel
./http_load -p 100 -s 10 url.txt
$ vi url.txt
http://go.domicake.com/api/test

1709 fetches, 100 max parallel, 1.98757e+06 bytes, in 10 seconds
1163 mean bytes/connection
170.9 fetches/sec, 198757 bytes/sec
msecs/connect: 0.0987104 mean, 3.829 max, 0.042 min
msecs/first-response: 553.364 mean, 3483.48 max, 103.149 min
HTTP response codes:
  code 200 -- 1709

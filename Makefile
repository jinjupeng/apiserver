SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir = "apiserver/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"


all: gotool
	# -w 为去掉调试信息（无法使用 gdb 调试），这样可以使编译后的二进制文件更小
	@go build -v -ldflags ${ldflags} .
clean:
	# 删除二进制文件
	rm -f apiserver
	# 删除vim swp文件
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	# 格式化代码
	gofmt -w . | grep -v vendor;true
	# 源代码静态检查
	go tool vet . | grep -v vendor;true
ca:
	# 生成证书
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	# 打印help信息
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help



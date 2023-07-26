# @command：命令前加`@`表示执行时不显示当前命令，仅显示结果
# 使用 := 声明变量, 使用$()或${}或"$()"打印变量
# `-`此条命令出错,make也会继续执行后续的命令, 如"-rm -f main.o"
# .PHONY:clean 伪目标声明, 将不会检查该目标是否存在或者该目标是否需要更新，直接执行, PHONY是一个伪目标，可以有效防止在Makefile文件中定义的可执行命令的目标规则和工作目录下的实际文件出现名称冲突，并提高了执行Makefile的性能
# 目录下有与make 命令 同名的文件时 执行make 命令就会出现错误，如出现"clean"文件，使用`.PHONY: clean`伪目标即可解决, 而`clean.go`不会冲突,如二进制可执行文件名可鞥会冲突




# 变量声明
Message := Hello World
# 函数：匹配所有的go文件
SRC := $(wildcard ./*.go)
#OBJ = $(patsubst %.go, $(SRC))

RELEASE := $(shell date +"%Y-%m-%d")

# make run
run:
	@echo "============开始编译============"
	@ls
	@du -sh ./*
	@echo "run"
	@echo "============编译结束============"

# make deploy
deploy:
	@echo "============开始编译============"
	@pwd
	@echo "deploy"
	@echo "============编译结束============"

dev:
	# 变量使用
	@echo $(Message)


.PHONY: clean
clean:
	@echo "============开始编译============"
	@echo "clean"
	@echo "============编译结束============"

build: $(SRC)
	@echo "============开始编译============"
	@echo "build"
	@echo $(SRC)
	@echo "============编译结束============"

release:
	@echo $(RELEASE)
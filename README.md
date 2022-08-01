# goworkspace overview
Belong to my GO language learning space

## Use Git submodule
```
 git submodule update --init --recursive
```
准备源代码

1. `git@github.com:mao888/goworkspace.git`
2. `cd goworkspace`
3. `git submodule init`(注册所有模块，会检出所有模块，包括核心仓库和非核心仓库)(下一条命令，二选一)
4. `git submodule init argo bluebell go_web GoWeb_ToDoList uber_go_guide_cn`(注册列出的模块，会检出对应模块)
5. `git submodule sync`
6. `git submodule update`


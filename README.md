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

附：[Git - git-submodule Documentation](https://git-scm.com/docs/git-submodule/en)

常见问题：
    - [No submodule mapping found in .gitmodules for path ......](https://blog.csdn.net/SaberJYang/article/details/124431029)

## submodule repositories

- [ ] [**submodule** "**go_web/GoWeb_ToDoList**"]
  - path = go_web/GoWeb_ToDoList 
  - url = [git@github.com:mao888/bluebell.git](git@github.com:mao888/bluebell.git)
- [ ] [**submodule** "**bluebell**"]
  - path = bluebell 
  - url = [git@github.com:mao888/bluebell.git](git@github.com:mao888/bluebell.git)
- [ ] [**submodule** "**argo**"]
  - path = argo 
  - url = [ git@github.com:mao888/argo.git]( git@github.com:mao888/argo.git)

- [ ] [**submodule** "**uber_go_guide_cn**"]
  - path = uber_go_guide_cn 
  - url = https://github.com/mao888/uber_go_guide_cn.git
- [ ] [**submodule** "**GoWeb_ToDoList**"]
  - path = GoWeb_ToDoList 
  - url = [git@github.com:mao888/GoWeb_ToDoList.git](git@github.com:mao888/GoWeb_ToDoList.git)
- [ ] [**submodule** **"Microservice"**]
  - path = Microservice 
  - url = [git@github.com:mao888/Microservice.git](git@github.com:mao888/Microservice.git)

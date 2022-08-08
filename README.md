# goworkspace overview
Belong to my GO language learning space

## [Use Git submodule](https://zhuanlan.zhihu.com/p/87053283)
### 一、 创建 submodule
  1.` git submodule add <submodule_url> 命令可以在项目中创建一个子模块。`

### 二、获取 submodule

准备源代码

1. `git@github.com:mao888/goworkspace.git`
2. `cd goworkspace`
3. `git submodule init`(注册所有模块，会检出所有模块，包括核心仓库和非核心仓库)(下一条命令，三选一)
4. `git submodule init argo bluebell go_web GoWeb_ToDoList uber_go_guide_cn`(注册列出的模块，会检出对应模块)
5. ` git submodule update --init --recursive`  子模块代码也获取到
6. `git submodule sync`
7. `git submodule update`

### 三、子模块内容的更新
对于子模块而言，并不需要知道引用自己的主项目的存在。对于自身来讲，子模块就是一个完整的 Git 仓库，按照正常的 Git 代码管理规范操作即可。

对于主项目而言，子模块的内容发生变动时，通常有三种情况：

1）当前项目下子模块文件夹内的内容发生了未跟踪的内容变动；

2）当前项目下子模块文件夹内的内容发生了版本变化；

3）当前项目下子模块文件夹内的内容没变，远程有更新；


> 情况汇总
终上所述，可知在不同场景下子模块的更新方式如下：

对于子模块，只需要管理好自己的版本，并推送到远程分支即可；
对于父模块，若子模块版本信息未提交，需要更新子模块目录下的代码，并执行 commit 操作提交子模块版本信息；
对于父模块，若子模块版本信息已提交，需要使用 git submodule update ，Git 会自动根据子模块版本信息更新所有子模块目录的相关代码。
### 四、删除子模块
根据官方文档的说明，应该使用 git submodule deinit 命令卸载一个子模块。这个命令如果添加上参数 --force，则子模块工作区内即使有本地的修改，也会被移除。

1. `git submodule deinit <sub-module> `
2. `git rm <sub-module>`
3. `git commit -m "delete submodule project-sub-1"`


附：[Git - git-submodule Documentation](https://git-scm.com/docs/git-submodule/en)

常见问题：
    - [No submodule mapping found in .gitmodules for path ......](https://blog.csdn.net/SaberJYang/article/details/124431029)

## submodule repositories

- [x] [submodule "golang-guide"]
  - path = golang-guide 
  - ssh = [git@github.com:mao888/golang-guide.git](git@github.com:mao888/golang-guide.git)
  - url = [https://github.com/mao888/golang-guide.git](https://github.com/mao888/golang-guide.git)
  - description = 「Golang学习+面试指南」一份涵盖大部分 Golang程序员所需要掌握的核心知识。准备 Golang面试，首选 GolangGuide！

- [x] [**submodule** "**go_web/GoWeb_ToDoList**"]
  - path = go_web/GoWeb_ToDoList 
  - ssh = [git@github.com:mao888/bluebell.git](git@github.com:mao888/bluebell.git)
  - url = [https://github.com/mao888/GoWeb_ToDoList.git](https://github.com/mao888/GoWeb_ToDoList.git)
  - description: go web course source code.
- [x] [**submodule** "**bluebell**"]
  - path = bluebell 
  - ssh = [git@github.com:mao888/bluebell.git](git@github.com:mao888/bluebell.git)
  - url = [https://github.com/mao888/bluebell.git](https://github.com/mao888/bluebell.git)
  - description: 基于vue+go+gin+mysql+redis的博客论坛web项目
- [x] [**submodule** "**argo**"]
  - path = argo 
  - ssh = [ git@github.com:mao888/argo.git]( git@github.com:mao888/argo.git)
  - url = [ https://github.com/mao888/argo.git]( https://github.com/mao888/argo.git)
  - description: 使用go语言实现数据结构与算法，涵盖字符串、数组、链表、队列、栈、树、图等数据结构。在实现算法的基础上，进行go语言实战。此外也包含经典算法在go实战项目中的应用，以及开源项目算法方面源码分析。
- [x] [**submodule** "**uber_go_guide_cn**"]
  - path = uber_go_guide_cn 
  - ssh = [git@github.com:mao888/uber_go_guide_cn.git](git@github.com:mao888/uber_go_guide_cn.git)
  - url = [https://github.com/mao888/uber_go_guide_cn.git](https://github.com/mao888/uber_go_guide_cn.git)
  - description: Uber Go 语言编码规范中文版. The Uber Go Style Guide .
- [x] [**submodule** "**GoWeb_ToDoList**"]
  - path = GoWeb_ToDoList 
  - ssh = [git@github.com:mao888/GoWeb_ToDoList.git](git@github.com:mao888/GoWeb_ToDoList.git)
  - url = [https://github.com/mao888/GoWeb_ToDoList.git](https://github.com/mao888/GoWeb_ToDoList.git)
  - description: A ToDoList Demo based on Vue+Gin+Gorm+mysql
- [x] [**submodule** **"Microservice"**]
  - path = Microservice 
  - ssh = [git@github.com:mao888/Microservice.git](git@github.com:mao888/Microservice.git)
  - url = [https://github.com/mao888/Microservice.git](https://github.com/mao888/Microservice.git)
  - description: go-zero 是一个集成了各种工程实践的 web 和 rpc 框架
- [x] [submodule "go-utils"]
  - path = go-utils 
  - ssh = [git@github.com:mao888/go-utils.git](git@github.com:mao888/go-utils.git)
  - url = [https://github.com/mao888/go-utils.git](https://github.com/mao888/go-utils.git)
  - description: 基于日常学习及项目需要所积累的go语言常用工具库
- [x] [submodule "mao888"]
  - path = mao888 
  - ssh = [git@github.com:mao888/mao888.git](git@github.com:mao888/mao888.git)
  - url = [https://github.com/mao888/mao888](https://github.com/mao888/mao888)
  - description: my github public profile
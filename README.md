# gopl-exercises

[《Go 程序设计语言》](https://book.douban.com/subject/27044219/)练习题参考答案。

## 搭建练习环境

推荐使用 VS Code 简单几步搭建练习环境。

1. 安装 [VS Code go 拓展](https://marketplace.visualstudio.com/items?itemName=golang.go)。
2. 打开 VS Code 的 Command Palette，执行 `Go: Install/Update Tools`。
   > 在一些互联网访问受限的地区，这一步可能会执行失败，一个解决方案是设置代理。例如在中国大陆地区，终端运行 `go env -w GOPROXY=https://goproxy.cn,direct`，然后重启 VS Code 并重新执行 `Go: Install/Update Tools`。
3. 打开 VS Code settings.json 文件，添加如下配置将 gofmt 设置为 go 代码默认的格式化工具：
   ```json
   "[go]": {
     "editor.defaultFormatter": "golang.go"
   }
   ```
4. 完成，开始愉快地编写 go 程序吧！

## gopl 学习方法

以下是我使用的学习方法，可供参考：

1. 首先通篇逐字阅读《Go 程序设计语言》（有条件的建议阅读英文原版），对 go 语言建立一个初步的系统性认知。
2. 认真独立地完成每一章每一小节的练习题，在编辑器中编写代码并编译运行，得到正确的结果，如果练习过程中遇到困难，借此机会回顾书本的内容，试着从书本中寻找解决办法，实在无法解决再试着参考他人的解答。
3. 练习练习再练习！

## 参考资料

- [gopl](https://github.com/linehk/gopl)：一个包含了[《Go 程序设计语言》](https://book.douban.com/subject/27044219/)中全部示例代码和全部练习题参考答案的项目。

# Gee - Go Web Framework (学习实现)

本仓库是我跟随 [极客兔兔的 7 天实现 Gee 框架教程](https://geektutu.com/post/gee.html) 的学习与实践记录。  
Gee 是一个用 Go 语言实现的简易 Web 框架，参考了 Gin 的设计思路，代码简洁、小而美，适合入门者学习 Web 框架的核心原理。

## 📚 教程目录
- **Day1**：HTTP 基础 - `http.Handler` 接口实现  
- **Day2**：上下文 - 统一封装 `Context`，便于扩展  
- **Day3**：前缀树路由 - 使用 Trie 实现动态路由匹配  
- **Day4**：分组控制 - 路由分组 & 控制前缀  
- **Day5**：中间件 - 支持类似 Gin 的中间件机制  
- **Day6**：模板 - 集成 HTML 模板渲染  
- **Day7**：错误恢复 - `panic` 恢复机制  

每一个从day2开始，都是在前一个day的基础上做内容的增加或优化

如果熟悉postman、curl工具进行接口测试，那其中day3的router_test.go可以忽略，因为目前还没写入使用测试样例进行测试的细节说明。

## 🚀 项目目标

- 理解 Web 框架的核心设计（路由、中间件、模板、错误恢复）。  
- 掌握 Go 语言在 Web 开发中的实践。  
- 将学习过程与实现代码逐步记录在 GitHub。  

## 🎯 补充说明

- 在Day7错误恢复中，对于Recovery中间件的作用 正确理解应该是“保证所有的请求都能正常响应，否则panic后其他请求就没有回应了” 

## 🙌 致谢

本项目基于极客兔兔的教程实现，原教程地址：[https://geektutu.com/post/gee.html](https://geektutu.com/post/gee.html)  

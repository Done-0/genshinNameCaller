# README

## 介绍

欢迎来到 `genshinNameCaller` 项目，这是一个基于 `wails` 和 `vue.js` 构建的应用。

## 快速开始

您可以通过以下命令启动项目：

```bash
wails dev
```

## 键位

双击 Gui 切换全屏

## 关于

您可以通过编辑 `wails.json` 来配置项目。有关项目设置的更多信息，请访问：https://wails.io/docs/reference/project-config

## 实时开发

要在实时开发模式下运行，请在项目目录中运行 `wails dev`。这将启动一个 `Vite` 开发服务器，为您的前端更改提供非常快速的热重载。  
如果您想在浏览器中进行开发并访问您的 Go 方法，还有一个开发服务器在 [http://localhost:34115](http://localhost:34115) 上运行。  
请在浏览器中连接到此地址，您可以通过开发者工具调用您的 Go 代码。

## Building

要构建一个可再分发的生产模式包，请使用 `wails build`。

执行下面的命令将构建一个 Windows/Linux 的可执行软件：
```bash
wails build
```

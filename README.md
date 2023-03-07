<p align="center">
  <a href="https://github.com/songquanpeng/go-file"><img src="https://user-images.githubusercontent.com/39998050/108494937-1a573e80-72e3-11eb-81c3-5545d7c2ed6e.jpg" width="200" height="200" alt="go-file"></a>
</p>

<div align="center">

# Go File

_✨ 文件分享工具，仅单个可执行文件，开箱即用，可用于局域网内分享文件和文件夹，直接跑满本地带宽 ✨_  

</div>

<p align="center">
  <a href="https://raw.githubusercontent.com/songquanpeng/go-file/master/LICENSE">
    <img src="https://img.shields.io/github/license/songquanpeng/go-file?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/songquanpeng/go-file/releases/latest">
    <img src="https://img.shields.io/github/v/release/songquanpeng/go-file?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://github.com/songquanpeng/go-file/releases/latest">
    <img src="https://img.shields.io/github/downloads/songquanpeng/go-file/total?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://hub.docker.com/repository/docker/justsong/go-file">
    <img src="https://img.shields.io/docker/pulls/justsong/go-file?color=brightgreen" alt="docker pull">
  </a>
  <a href="https://goreportcard.com/report/github.com/songquanpeng/go-file">
  <img src="https://goreportcard.com/badge/github.com/songquanpeng/go-file" alt="GoReportCard">
  </a>
</p>

<p align="center">
  <a href="https://github.com/songquanpeng/go-file/projects/1">开发规划</a>
  ·
  <a href="https://github.com/songquanpeng/go-file/releases">程序下载</a>
  ·
  <a href="https://github.com/songquanpeng/gofile-launcher/releases/latest">启动器下载</a>
  ·
  <a href="https://github.com/songquanpeng/gofile-cli/releases/latest">CLI 下载</a>
  ·
  <a href="https://iamazing.cn/page/LAN-SHARE-使用教程">使用教程</a>
  ·
  <a href="#演示">截图展示</a>
</p>


<details>
<summary><strong><i>English</i></strong></summary>
<div>

Warning: The English version is outdated.

## Description
File sharing tool, can be used to share files in a LAN.

## Features
1. No need to configure environment and there is only a single executable file.
2. Automatically open browser to make you share file more quickly.
3. Generate QR codes for your mobile phone to scan.
4. Easily share all the content of a local dir.

## Usage
*For v0.3.4 and below.*

Just double-click to use with default port `3000` and default token (used to verify identity when user try to delete files) `token`.

If you want to change the port and token, run it like this:`./go-file.exe --port 80 --token private`.

Your can also public a local path by providing a `path` like this : `./go-file.exe --path ./this/is/a/path` 

```
Usage of go-file.exe:
  -host string
        the server's ip address or domain (default "localhost")
  -path string
        specify a local path to public
  -port int
        specify the server listening port. (default 3000)
  -token string
        specify the private token. (default "token")
  -video string
        specify a video folder to public
```

## Demo
Please visit https://go-file.herokuapp.com/ to have a try yourself.

![index page](https://user-images.githubusercontent.com/39998050/130427067-80bf3cc5-5fee-488a-bea5-e323b9458064.png)
![explorer page](https://user-images.githubusercontent.com/39998050/177032568-8af95d7e-87ab-4e60-804b-5e49addfb6ab.png)
![image page](https://user-images.githubusercontent.com/39998050/177032659-c8c68186-09f4-4142-9f57-70bcb4a4cda1.png)
![video page](https://user-images.githubusercontent.com/39998050/177032588-8946abde-a8da-45a2-a389-c16dba9cea34.png)


## Others
[Node.js version is here.](https://github.com/songquanpeng/lan-share)
</div>
</details>

## 规划
基于 `up-master`分支上即最初原作者的运营版本进行二次开发。

## 项目启动
参数示例`go run *.go --port 8001 --video /d/system/video --path /d/Document --no-browser true`
- port：端口
- video：视频路径
- path：共享文件夹目录
- no-browser：不自动打开页面


## 更新记录
1. 废弃标准库os中利用`GetEnv`来读取或设置参数配置，变更为 `yml` 文件形式。
2. 使用`gorm.io/gorm`官方包来做数据库驱动、引擎。



# Neuvillette-Helper
Neuvillette automatic circle program. 

那维莱特自动滋水程序



## 使用方法

### 编译可执行文件

克隆代码后在本地编译即可

```bash
git clone https://github.com/xiabee/Neuvillette-Helper.git
cd ./Neuvillette-Helper
go mod download

GOOS=windows GOARCH=amd64 go build -o neuvillette.exe .
```



不会编译的话直接去 release 页面下载最新版本可执行文件即可



### 模拟滋水

* 按住鼠标中键开始转圈滋水，同时每隔三秒开始放技能；

* 松开鼠标中键停止滋水



## 注意

首次使用 gohook / robotgo 可能需要先安装相关依赖

### robotgo 依赖项

1. **GCC 或 MinGW-w64**：用于编译 C 语言部分。Windows 上推荐使用 MinGW-w64。
2. **Xcode（macOS）**：如果您在 macOS 上编译，需要安装 Xcode，因为它包含了 macOS 开发所需的工具链和库。
3. **X11 库（Linux）**：Linux 系统上需要 X11 库，用于屏幕捕获和模拟输入等功能。通常需要安装 `libX11`、`libXtst`、`libXScrnSaver`、`libpng` 等。
4. **libpng**：用于图像处理。
5. **其他依赖**：根据您的需求，可能还需要其他依赖库，如 `zlib` 等。

# 视频转 mp4 格式

程序可编译在 windows 下运行，只测试过 win11 支持。

基于 ffmpeg 实现的将当前目录下所有非 mp4 视频文件转码为 mp4 文件。

in.avi 为测试视频，运行程序后将会生成 in.mp4 文件。生成文件不会覆盖，注意及时移走不需要转码的非 mp4 文件。

build/dependencies/ffmpeg-win 是 windows 下的 ffmpeg 可执行程序，版本为 5.0.1。

开发基于 https://github.com/u2takey/ffmpeg-go 库。

## 下载

https://github.com/zhan3333/converter/releases

## 运行

```shell
make run
```

## 编译

```shell
make build
```

编译完成后将生成 build/converter.exe 文件，build 目录可以独立运行程序，release 发包是打包了 build 目录。

## 发包

```shell
make release
```

执行完毕后会生成 converter-win.zip 文件
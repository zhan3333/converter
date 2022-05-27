# 视频下载和转换工具

程序可编译在 Windows、Mac 平台运行。

基于 ffmpeg 实现的将当前目录下所有非 mp4 视频文件转码为 mp4 文件。

基于 lux 项目实现的网络视频下载功能。

in.avi 为测试视频，运行程序后将会生成 in.mp4 文件。生成文件不会覆盖，注意及时移走不需要转码的非 mp4 文件。

build/dependencies/ffmpeg-win 是 windows 下的 ffmpeg 可执行程序，版本为 5.0.1。

使用 fyne 制作了跨平台 GUI。

## 功能



### 下载网络视频

GUI

![GUI](./data/gui.png)

![下载视频](./data/download-video.png)

支持下载的站点: 

| Site       | URL                          | 🎬 Videos | 🌁 Images | 🔊 Audio | 📚 Playlist | 🍪 VIP adaptation | Build Status                                                                                                                                                                |
| ---------- | ---------------------------- | --------- | --------- | -------- | ----------- | ----------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 抖音       | <https://www.douyin.com>     | ✓         | ✓         |          |             |                   | [![douyin](https://github.com/iawia002/lux/actions/workflows/stream_douyin.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_douyin.yml)             |
| 哔哩哔哩   | <https://www.bilibili.com>   | ✓         |           |          | ✓           | ✓                 | [![bilibili](https://github.com/iawia002/lux/actions/workflows/stream_bilibili.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_bilibili.yml)       |
| 半次元     | <https://bcy.net>            |           | ✓         |          |             |                   | [![bcy](https://github.com/iawia002/lux/actions/workflows/stream_bcy.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_bcy.yml)                      |
| pixivision | <https://www.pixivision.net> |           | ✓         |          |             |                   | [![pixivision](https://github.com/iawia002/lux/actions/workflows/stream_pixivision.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_pixivision.yml) |
| 优酷       | <https://www.youku.com>      | ✓         |           |          |             | ✓                 | [![youku](https://github.com/iawia002/lux/actions/workflows/stream_youku.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_youku.yml)                |
| YouTube    | <https://www.youtube.com>    | ✓         |           |          | ✓           |                   | [![youtube](https://github.com/iawia002/lux/actions/workflows/stream_youtube.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_youtube.yml)          |
| 西瓜视频（头条）    | <https://m.toutiao.com>, <https://v.ixigua.com>, <https://www.ixigua.com>    | ✓         |           |          |             |                   | [![ixigua](https://github.com/iawia002/lux/actions/workflows/stream_ixigua.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_ixigua.yml)   |
| 爱奇艺     | <https://www.iqiyi.com>      | ✓         |           |          |             |                   | [![iqiyi](https://github.com/iawia002/lux/actions/workflows/stream_iqiyi.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_iqiyi.yml)                |
| 新片场    | <https://www.xinpianchang.com>       | ✓         |           |          |             |                   | [![xinpianchang](https://github.com/iawia002/lux/actions/workflows/stream_xinpianchang.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_xinpianchang.yml)                   |
| 芒果 TV    | <https://www.mgtv.com>       | ✓         |           |          |             |                   | [![mgtv](https://github.com/iawia002/lux/actions/workflows/stream_mgtv.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_mgtv.yml)                   |
| 糖豆广场舞 | <https://www.tangdou.com>    | ✓         |           |          |             |                   | [![tangdou](https://github.com/iawia002/lux/actions/workflows/stream_tangdou.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_tangdou.yml)          |
| Tumblr     | <https://www.tumblr.com>     | ✓         | ✓         |          |             |                   | [![tumblr](https://github.com/iawia002/lux/actions/workflows/stream_tumblr.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_tumblr.yml)             |
| Vimeo      | <https://vimeo.com>          | ✓         |           |          |             |                   | [![vimeo](https://github.com/iawia002/lux/actions/workflows/stream_vimeo.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_vimeo.yml)                |
| Facebook   | <https://facebook.com>       | ✓         |           |          |             |                   | [![facebook](https://github.com/iawia002/lux/actions/workflows/stream_facebook.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_facebook.yml)       |
| 斗鱼视频   | <https://v.douyu.com>        | ✓         |           |          |             |                   | [![douyu](https://github.com/iawia002/lux/actions/workflows/stream_douyu.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_douyu.yml)                |
| 秒拍       | <https://www.miaopai.com>    | ✓         |           |          |             |                   | [![miaopai](https://github.com/iawia002/lux/actions/workflows/stream_miaopai.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_miaopai.yml)          |
| 微博       | <https://weibo.com>          | ✓         |           |          |             |                   | [![weibo](https://github.com/iawia002/lux/actions/workflows/stream_weibo.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_weibo.yml)                |
| Instagram  | <https://www.instagram.com>  | ✓         | ✓         |          |             |                   | [![instagram](https://github.com/iawia002/lux/actions/workflows/stream_instagram.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_instagram.yml)    |
| Twitter    | <https://twitter.com>        | ✓         |           |          |             |                   | [![twitter](https://github.com/iawia002/lux/actions/workflows/stream_twitter.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_twitter.yml)          |
| 腾讯视频   | <https://v.qq.com>           | ✓         |           |          |             |                   | [![qq](https://github.com/iawia002/lux/actions/workflows/stream_qq.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_qq.yml)                         |
| 网易云音乐 | <https://music.163.com>      | ✓         |           |          |             |                   | [![netease](https://github.com/iawia002/lux/actions/workflows/stream_netease.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_netease.yml)          |
| 音悦台     | <https://yinyuetai.com>      | ✓         |           |          |             |                   | [![yinyuetai](https://github.com/iawia002/lux/actions/workflows/stream_yinyuetai.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_yinyuetai.yml)    |
| 极客时间   | <https://time.geekbang.org>  | ✓         |           |          |             |                   | [![geekbang](https://github.com/iawia002/lux/actions/workflows/stream_geekbang.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_geekbang.yml)       |
| Pornhub    | <https://pornhub.com>        | ✓         |           |          |             |                   | [![pornhub](https://github.com/iawia002/lux/actions/workflows/stream_pornhub.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_pornhub.yml)          |
| XVIDEOS    | <https://xvideos.com>        | ✓         |           |          |             |                   | [![xvideos](https://github.com/iawia002/lux/actions/workflows/stream_xvideos.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_xvideos.yml)          |
| 聯合新聞網 | <https://udn.com>            | ✓         |           |          |             |                   | [![udn](https://github.com/iawia002/lux/actions/workflows/stream_udn.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_udn.yml)                      |
| TikTok     | <https://www.tiktok.com>     | ✓         |           |          |             |                   | [![tiktok](https://github.com/iawia002/lux/actions/workflows/stream_tiktok.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_tiktok.yml)             |
| 好看视频   | <https://haokan.baidu.com>   | ✓         |           |          |             |                   | [![haokan](https://github.com/iawia002/lux/actions/workflows/stream_haokan.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_haokan.yml)             |
| AcFun      | <https://www.acfun.cn>       | ✓         |           |          | ✓           |                   | [![acfun](https://github.com/iawia002/lux/actions/workflows/stream_acfun.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_acfun.yml)                |
| Eporner    | <https://eporner.com>        | ✓         |           |          |             |                   | [![eporner](https://github.com/iawia002/lux/actions/workflows/stream_eporner.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_eporner.yml)          |
| StreamTape | <https://streamtape.com>     | ✓         |           |          |             |                   | [![streamtape](https://github.com/iawia002/lux/actions/workflows/stream_streamtape.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_streamtape.yml) |
| 虎扑       | <https://hupu.com>           | ✓         |           |          |             |                   | [![hupu](https://github.com/iawia002/lux/actions/workflows/stream_hupu.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_hupu.yml)                   |
| 虎牙视频   | <https://v.huya.com>         | ✓         |           |          |             |                   | [![huya](https://github.com/iawia002/lux/actions/workflows/stream_huya.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_huya.yml)                   |
| 喜马拉雅   | <https://www.ximalaya.com>   |           |           | ✓        |             |                   | [![ximalaya](https://github.com/iawia002/lux/actions/workflows/stream_ximalaya.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_ximalaya.yml)       |
| 快手       | <https://www.kuaishou.com>   | ✓         |           |          |             |                   | [![kuaishou](https://github.com/iawia002/lux/actions/workflows/stream_kuaishou.yml/badge.svg)](https://github.com/iawia002/lux/actions/workflows/stream_kuaishou.yml)       |

### 视频转为 mp4 格式

![视频转为 mp4 格式](./data/converter-to-mp4.png)

## 下载

https://github.com/zhan3333/converter/releases

下载 converter.zip 文件，解压后 windows 平台双击运行 converter.exe, mac 平台双击或命令行运行 converter。

## 开发

### 运行

```shell
make run
```

### 编译

```shell
make build
```

编译完成后将生成 build/converter.exe 文件，build 目录可以独立运行程序，release 发包是打包了 build 目录。

### 发包

```shell
make release
```

执行完毕后会生成 converter.zip 文件

## 其他

开发基于 https://github.com/u2takey/ffmpeg-go

视频下载基于 https://github.com/iawia002/lux
# regbrand
Scrawer  built by golang

下载Redis Server

```
https://github.com/tporadowski/redis/releases/download/v5.0.10/Redis-x64-5.0.10.msi
```

安装

注意添加到PATH，如下图打勾选中

![QQ截图20201204142101](C:\Users\dengyun\Pictures\QQ截图20201204142101.png)

修改配置 C:\Program Files\Redis\redis.windows.conf 设置密码

```
requirepass 123456
```

运行Redis Server

```
redis-server.exe redis.windows.conf
```



运行regbrand

命令行如下

```
regbrand.exe -p 1721 -c 25 -t 361
```




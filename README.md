# gotomongo
本工具用于mongodb性能压测

## 执行本程序-windows下

```
双击打开
```

## 执行本程序-linux

```
chmod 777 gotomongo
./gotomongo
```
具体命令：
```
COMMANDS:
   insert, i  输入：[命令] [总数] [并发数] [地址] 如：i 20000 10 mongodb://127.0.0.1:27017
   query, q   输入：[命令] [总数] [并发数] [地址] 如：q 20000 10 mongodb://127.0.0.1:27017
   help, h    Shows a list of commands or help for one command

```
执行效果如下：
```
请输入命令或者回车Enter提示命令，或输入exit退出！
i 10000 10 mongodb://192.168.1.117:27017
当前命令：i 10000 10 mongodb://192.168.1.117:27017
exec 0
exec 1000
exec 2000
exec 3000
exec 4000
exec 5000
exec 6000
exec 7000
exec 8000
exec 9000
exec cost=29.48秒
exec TPS=339.21
exec NPS=3392.09

```

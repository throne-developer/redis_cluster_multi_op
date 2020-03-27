# redis_cluster_multi_op
execute redis commands on redis cluster

# 中文说明
对redis集群执行批量操作，解决redis集群不能使用pipeline的问题。

使用方法：
（1）修改config.json，设置redis集群的ip列表和密码，以及要执行的命令文件路径
```javascript
{
    "redis_addrs":[
        "10.11.3.104:6000",
        "10.11.3.125:6000",
        "10.11.3.104:6000"
    ],
    "redis_pwd":"pass",
    "command_file":"commands.csv"
}
```
（2）准备命令文件，每一行为一条命令，用逗号分隔，比如 "set key1 abc" 这条命令在文件为 "set,key1,abc"，如下所示：
```javascript
set,key1,11
set,key2,22
set,key3,11
```
（3）go build得到exe程序，也可以直接使用编译好的redis_cluster_multiop程序（linux版本），将config.json和exe放在同一目录下，然后执行exe，输出如下：
```javascript
2020-03-27 18:06:24 [info]  loadFile commands.csv 3
reply  0 OK
reply  1 OK
reply  2 OK
done count  3
```

# Usage
exec multi redis command on redis cluster, resolve problem that redis cluster cann't use pipeline

（1）modify config.json， set redis cluster ip addr and password，and command file path
```javascript
{
    "redis_addrs":[
        "10.11.3.104:6000",
        "10.11.3.125:6000",
        "10.11.3.104:6000"
    ],
    "redis_pwd":"pass",
    "command_file":"commands.csv"
}
```
（2）prepare command file，one command per line，seperate by comma，such as "set key1 abc" command is "set,key1,abc" in file，for example：
```javascript
set,key1,11
set,key2,22
set,key3,11
```
（3）go build and get execute program，put config.json and exe in same folder，then execute exe，output is：
```javascript
2020-03-27 18:06:24 [info]  loadFile commands.csv 3
reply  0 OK
reply  1 OK
reply  2 OK
done count  3
```

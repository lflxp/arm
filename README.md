# arm
fast and stable control all pc
* 快速扫描局域网所有服务器
* 控制服务器 进行服务注册
* 批量环境安装

# 实现
## 客户端 - 保护udp服务发现和tcp rpc远程过程调用
* UDPClient 利用广播的方式上线寻找服务器
* 客户端发送信息参照dhcp 第一个字段buf[0] 填写Discover信息
* 等待接收Request信息，包含keepalive信息
* 等待下一次keepalive
* 等待Release、Decline主动停止
## 服务端
* 服务器接受到后返回一个连接地址
* 获取真实服务器IP 注册到本地缓存、数据库、etcd等
* 反馈request信息，包含keepalive信息
* 客户端控制终端
## 题目
1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

## 回答
### 题目1
查询了网上的资料：
**粘包和半包**
粘包问题是指当发送两条消息时，例如“消息1：ABC，消息2：DEF”，接收端本应收到“结果1：ABC，结果2：DEF”，但却收到了“结果1：ABCDEF”，像这种一次性读取了多条数据的情况叫做粘包。
半包问题是指对于消息：ABC，接收端收到的变成了“消息1：AB，消息2：C”。

**粘包和半包的产生原因**
TCP是面向连接的传输协议，传输的数据是以流的形式，而数据流没有明确的开始与结尾的边界。
在发送和接收的过程中，为了减少上下文切换带来的系统开销，数据流会先放入缓冲区中，等缓冲区满了以后再统一进行发送或读取。
* 当发送方发送的数据<socket缓冲区大小时，一个缓冲区内会包含不止一条数据，此时就会产生**粘包**；
* 反之，当发送方发送的数据>socket缓冲区大小时，一条消息被截断分开发送，此时就会产生**半包**；
* 接收方在读取缓冲区内的数据时，若读取不够及时，将导致缓冲区内数据堆积了多条消息，此时也会产生**粘包**；
* 当发送的数据大于协议的MTU（Maximum Transmission Unit，最大传输单元），此时必须拆包，导致**半包**。
  
**粘包的解包方式**

**1.fix length**
无论一次接受多少数据，都按照预先设置的固定长度进行解码。如果是半包消息，则会缓存半包消息并等待下个包到达后进行拼包，直到读取到一个完整的包。
**2.delimiter based**
发送的时候在数据包中添加一个符号作为分隔符，用来标记数据包边界。
**3.length field based frame decoder**
发送数据的时候，在前边附加上数据包的长度

**应用**：Netty中的FixedLengthFrameDecoder和DelimiterBasedFrameDecoder

### 题目2
以毛老师视频课程中的goim协议为例，client发送数据，并对数据编码；server接收数据，解码并输出。
见该目录下代码。

**测试效果**

**客户端**
```
PS E:\go-project\geektime\homework\9th_week\client> go run .\main.go
Please input data...
sadlkjjhfakldsfhjjkahdf
server response:Receive Success!
hi i am zhy
server response:Receive Success!
```

**服务端**
```
PS E:\go-project\geektime\homework\9th_week\server> go run .\main.go
Listening....

packet length:41
header length:16
version:2
operation:3
sequence:9
body:sadlkjjhfakldsfhjjkahdf
-----------------------------------------------
packet length:29
header length:16
version:2
operation:3
sequence:9
body:hi i am zhy
-----------------------------------------------
```
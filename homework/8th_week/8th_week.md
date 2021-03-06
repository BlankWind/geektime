## 题目
1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

## 结果
### 题目1
**命令**
```
# ./redis-benchmark -h 127.0.0.1 -p 6379 -t get,set -d 10/20/50/100/200/1024/5120
```
**set结果**
|value size| Duration/100k requests | RPS |
|-|-|-|
| 10 |1.66s|60350.03|
| 20 |1.76s|56753.69|
| 50 |1.66s|60204.70|
| 100 |1.73s|57971.02|
| 200 |1.75s|57240.98|
| 1k |1.62s|61728.39|
| 5k |2.08s|48169.56|

**get结果**
|value size| Duration/100k requests | RPS |
|-|-|-|
| 10 |1.67s|59737.16|
| 20 |1.62s|61728.39|
| 50 |1.69s|59171.59|
| 100 |1.86s|53879.31|
| 200 |1.69s|59101.65|
| 1k |1.64s|61012.81|
| 5k |2.23s|44863.16|

**分析**
整体来看，value size越大，性能越差，这也符合预期；
但局部数据并非线性，尤其是1k时性能较好。
可能跟redis本身的优化有关，也可能是因为自己的机器性能不稳定。
### 题目2
**写入数据**
分别写入10000、50000、100000条数据，其中value大小分别为10、100、1000字节

**结果**
|数据量|value size|写入前|写入后|diff|平均每个key的占用空间（取整）
|-|-|-|-|-|-|
|1w|10|used_memory:872320|used_memory:1803312|930992|83|
|5w|10|used_memory:872472|used_memory:5396680|4524208|80|
|10w|10|used_memory:872624|used_memory:9921200|9048576|80|
|1w|100|used_memory:872136|used_memory:2763208|1891072|89|
|5w|100|used_memory:872288|used_memory:10196576|9324288|86|
|10w|100|used_memory:872440|used_memory:19521016|18648576|86|
|1w|1000|used_memory:872592|used_memory:11883584|11010992|101|
|5w|1000|used_memory:872744|used_memory:55796952|54924208|98|
|10w|1000|used_memory:872896|used_memory:110721472|109848576|98|

**分析**
根据最后一列数据可看出，**在数据量相同的情况下，value越大，每个key的占用空间越大**
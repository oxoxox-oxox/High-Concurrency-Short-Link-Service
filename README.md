# High-Concurrency Short Link Service

基于 Go 语言构建的高并发短链转换服务。采用经典的 **Cache-Aside 旁路缓存架构**，能够轻松应对秒杀、营销推广等极高频次的短链点击和重定向场景。

## 核心设计与算法

本项目摒弃了容易产生冲突的哈希（如 MD5）算法，采用业内标准的 **数据库自增 ID + 62 进制转换算法**：

1. **长链存入**：每次接收长链，存入 MySQL，获取唯一的自增 ID。
2. **生成短码**：将 10 进制的 ID 转换为 62 进制（包含 `0-9`, `a-z`, `A-Z`），生成极短且绝对唯一的短码（如 `2Bi`）。
3. **高效重定向**：访问短链时，利用 HTTP `302 Found` 状态码实现瞬间跳转。

## 技术栈

* **Web 框架**: [Gin](https://github.com/gin-gonic/gin) (极致轻量，路由高性能)
* **持久化存储**: MySQL 8.0 + [GORM](https://gorm.io/) (处理核心数据落地)
* **高性能缓存**: Redis 5.0 (应对海量并发读请求)
* **压测工具**: [hey](https://github.com/rakyll/hey)

## 性能压测表现 (高并发抗压能力)

通过引入 Redis 缓存层，本系统将绝大部分的高频读请求挡在内存层面，极大地保护了 MySQL。

**压测条件：**

* 操作系统：Windows (AMD Ryzen 9 8945HX, 16C32T)
* 压测工具：`hey` (`-n 50000 -c 300 -disable-redirects`)
* 并发数：300
* 总请求数：50,000 次

**极限压测成绩：**

* **QPS (Requests/sec)**: **10191.92** 
* **Average Latency**: **0.028 secs** 
* **TP50 (50% requests in)**: **0.028 secs**
* **TP99 (99% requests in)**: **0.084 secs**

## 如何在本地运行

### 1. 环境依赖

请确保你的电脑已安装并启动以下环境：

* Go 1.20+
* MySQL 8.0+
* Redis 5.0+

### 2. 数据库配置

修改 `cmd/server/main.go` 中的数据库连接 DSN：

```go
// 替换为你的本地用户名、密码和数据库名
dsn := "root:123456@tcp(127.0.0.1:3306)/shortlink_db?charset=utf8mb4&parseTime=True&loc=Local"
```

### 3. 启动服务

在项目根目录下执行：

```bash
go run cmd/server/main.go
```

服务默认运行在 `:8080` 端口。

### 4. API 测试

**生成短链：**

```bash
curl -X POST http://localhost:8080/api/v1/shorten -H "Content-Type: application/json" -d "{\"url\":\"https://github.com/gin-gonic/gin\"}"
```

**测试重定向：**
在浏览器中直接访问返回的 `short_url` (如 `http://localhost:8080/5`)，即可体验瞬间跳转。

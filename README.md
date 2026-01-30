# ai-alert
自建AI大模型代理后端服务，用AI助手生成前端页面，端到端闭环演示AI协同分析告警事件流程。


## Quick Start

### 1. 构建前端

由claude+minimax根据提示词实现, 购买Minimax的Code Plan套餐，[建议月套餐](https://platform.minimaxi.com/subscribe/coding-plan?code=LqPLjgdFQQ\&source=link)。

[claude code和minimax环境搭建](https://juejin.cn/post/7597709339981479976)

替换`internal/models/ai.go`文件的`AI_API_KEY`为`<your-minimax-api-key>`.


```prompt
1. 用react+typescript技术栈实现一个前端应用
2. 代理api请求路径http://localhost:8090/api/v1/chat, form表单参数content、rule_id、rule_name、deep和search_ql,
3. api代理携带bear token，模拟即可,请求头添加content-type:application/x-www-form-urlencoded
4. 页面左右分栏比例为1:2，左边填写表单，deep字段为select类型，枚举值为true或是false，search_ql填充默认值*，右边展示结果用Markdown格式
5. 右侧返回取response的data字段进行markdown ui展示
```

```
cd web && npm ci && npm run build
```

### 2. 构建后端

```windows
go build -o ai-alert.exe cmd/ai-model/main.go
```

```Linux
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ai-alert ./...
```

### 3. 嵌入前端页面

***将构建后的前端文件复制到internal/static目录***

```
rm -rf internal/static/dist && mkdir -p internal/static/dist
	cp -rf web/build/* internal/static/dist
```

### 4. 运行后端服务

```Windows
./ai-alert.exe
```

```Linux
./ai-alert
```

### 5. 访问前端页面

```
浏览器地址输入http://localhost:8090
```

***表单填充***

| 字段 | 样例值 | 描述 |
| --- | --- | --- |
| Content | 磁盘使用率达到90% | 告警事件 |
| Rule ID | 1 | 规则ID |
| Rule Name | 磁盘空间不足 | 规则名称 |
| Deep| 1 | 进行深度思索 |
| SearchQL| * | 搜索条件 |

## TODO List

### 前端

- []前端添加SSE方式接收数据功能
- []前端Form表单样式美化
- []前端添加Markdown样式美化

### 后端

- []后端添加SSE方式推送数据功能
- []后端添加Markdown样式美化

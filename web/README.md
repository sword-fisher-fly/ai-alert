# Web Frontend Application

这是一个基于 React + TypeScript 的前端应用，提供表单输入和 Markdown 显示功能。

## 功能特性

- 表单输入字段：content、rule_id、rule_name、deep、search_ql
- 集成 Markdown UI 展示返回内容
- 响应式设计，适配不同屏幕尺寸
- 错误处理和加载状态显示

## 技术栈

- React 18
- TypeScript
- React Markdown (用于 Markdown 渲染)
- Remark GFM (GitHub Flavored Markdown)
- React Scripts

## 安装和运行

### 安装依赖

```bash
cd web
npm install
```

### 启动开发服务器

```bash
npm start
```

应用将在 http://localhost:3000 启动

### 构建生产版本

```bash
npm run build
```

### 运行测试

```bash
npm test
```

## 项目结构

```
web/
├── public/
│   └── index.html          # HTML 模板
├── src/
│   ├── components/
│   │   ├── FormComponent.tsx    # 表单组件
│   │   └── MarkdownDisplay.tsx  # Markdown 显示组件
│   ├── App.tsx              # 主应用组件
│   ├── App.css              # 应用样式
│   ├── index.tsx            # 入口文件
│   └── index.css            # 全局样式
├── package.json
├── tsconfig.json
└── README.md
```

## API 集成

应用通过 POST 请求调用 `http://localhost:8090/api/v1/ai/chat` API

请求体格式：
```json
{
  "content": "内容",
  "rule_id": "规则ID",
  "rule_name": "规则名称",
  "deep": "深度值",
  "search_ql": "查询语言"
}
```

## 使用说明

1. 在表单中填写所有必填字段
2. 点击 "Submit" 按钮提交表单
3. 等待 API 响应，结果将以 Markdown 格式显示
4. 如果出现错误，错误信息将显示在表单下方

## 样式说明

- 使用现代的 Material Design 风格
- 表单具有清晰的标签和占位符
- 响应式布局，适配移动端和桌面端
- Markdown 内容具有专门的样式，包括代码块、表格等

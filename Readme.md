# Ai-Visual Core
- Ai-Visual Core 是 Ai-Visual 项目的核心服务模块，基于 Go 语言开发。

- 该模块负责接收来自检测引擎的数据并进行统一处理，遵循 GA/T1400 公安行业标准。

## 核心职责

- 消费 RabbitMQ 报警消息
- 将报警信息转换为 GA/T 1400 标准格式
- 提供 RESTful 查询接口
- 接收 WVP-Pro 设备事件（Webhook）

## 技术栈

- Go 语言
- RabbitMQ 消息队列
- RESTful API
- GA/T 1400 公安行业标准

## 项目结构

```
aivisual-core/
├── cmd/              # 应用程序入口
├── configs/          # 配置文件
├── internal/         # 内部模块
│   ├── app/          # 应用层
│   ├── config/       # 配置模块
│   ├── domain/       # 领域模型
│   ├── infra/        # 基础设施层
│   └── service/      # 服务层
└── pkg/              # 公共包
```

## 核心标准

本项目严格遵循 GA/T 1400 公安行业标准，确保与公安系统平台的兼容性。

## Git工作流规范

本项目采用标准的Git工作流进行版本控制：

### 分支策略

1. **main分支** - 生产环境稳定版本
2. **develop分支** - 开发环境最新版本
3. **feature分支** - 功能开发分支，命名规范：`feature/功能名称`
4. **hotfix分支** - 紧急修复分支，命名规范：`hotfix/问题描述`
5. **release分支** - 发布准备分支，命名规范：`release/版本号`

### 提交规范

提交信息遵循以下格式：
```
<type>(<scope>): <subject>

<body>

<footer>
```

常用type类型：
- feat: 新功能
- fix: 修复bug
- docs: 文档更新
- style: 代码格式调整
- refactor: 代码重构
- test: 测试相关
- chore: 构建过程或辅助工具的变动

示例：
```
feat(api): 添加离岗检测功能

实现离岗检测核心算法，支持ROI区域设置和时间阈值配置

Closes #123
```

### 工作流程

1. 从develop分支创建feature分支
```bash
git checkout develop
git pull origin develop
git checkout -b feature/新功能名称
```

2. 开发并提交代码
```bash
git add .
git commit -m "feat: 实现新功能"
```

3. 推送分支并创建Pull Request
```bash
git push origin feature/新功能名称
```

4. 代码审查通过后合并到develop分支
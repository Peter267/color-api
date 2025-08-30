# 🎨 图片主色调 API

一个基于 Go 语言构建的高性能 Serverless API，它可以根据你提供的图片 URL，快速分析并返回该图片的主色调。项目被设计为在 Vercel 上进行一键部署。

## ✨ 主要功能

-   **快速颜色提取**: 输入任意图片的 URL，即可获得其主色调的十六进制 (HEX) 颜色码。
-   **高性能**: 使用 Go 语言编写，拥有极快的执行速度和非常低的冷启动延迟。
-   **Serverless 架构**: 部署在 Vercel 平台，无需管理服务器，具备高可用性和自动扩缩容能力。
-   **广泛的图片格式支持**: 支持常见的图片格式，包括 JPEG, PNG, GIF, 以及 WebP。
-   **性能优化**: 通过先将图片缩放为缩略图再进行分析，极大地减少了网络下载和颜色计算的耗时。


## 🚀 API 使用指南

### 请求地址

```
GET /api/color
```

### 请求参数

| 参数 | 类型   | 描述               | 是否必须 |
| :--- | :----- | :----------------- | :------- |
| `url`  | string | 待分析图片的公开 URL | **是**   |

### 请求示例 (使用 curl)

将 `[YOUR_DEPLOYED_URL]` 替换为你部署后 Vercel 提供的 URL。

```bash
curl "[YOUR_DEPLOYED_URL]/api/color?url=https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
```

### 响应格式

#### ✅ 成功响应

```json
{
  "RGB": "#4285f4"
}
```

#### ❌ 失败响应

```json
{
  "error": "Missing image url"
}
```


## 🛠️ 技术栈

-   **后端语言**: [Go](https://go.dev/)
-   **部署平台**: [Vercel](https://vercel.com/)
-   **核心 Go 库**:
    -   `net/http`: 处理 HTTP 请求与响应
    -   `image`: 标准库，用于解码图片
    -   `github.com/nfnt/resize`: 用于高质量的图片缩放
    -   `golang.org/x/image/webp`: 提供 WebP 格式的解码支持


## 本地开发与部署

### 准备工作

1.  安装 [Go](https://go.dev/dl/) (版本 1.18 或更高)。
2.  安装 [Node.js](https://nodejs.org/) (用于安装 Vercel CLI)。
3.  安装 [Vercel CLI](https://vercel.com/docs/cli):
    ```bash
    npm install -g vercel
    ```

### 步骤

1.  **克隆仓库**
    ```bash
    git clone https://github.com/Peter267/color-api.git
    cd color-api
    ```

2.  **初始化 Go 模块** (如果 `go.mod` 文件不存在)
    ```bash
    go mod init <你的模块名>
    go mod tidy
    ```

3.  **本地运行**
    使用 Vercel CLI 可以在本地模拟 Serverless 环境。
    ```bash
    vercel dev
    ```
    服务启动后，你可以在 `http://localhost:3000` 访问你的 API。

4.  **部署到 Vercel**
    在项目根目录运行以下命令，并根据提示操作即可完成部署。
    ```bash
    vercel
    ```
    或者直接点击按钮：
    [![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FPeter267%2Fcolor-api)


## 📁 项目结构

```
.
├── api/
│   └── index.go       # Serverless 函数核心逻辑
├── go.mod             # Go 模块依赖文件
├── go.sum             # 依赖项校验和
├── README.md          # 项目说明文件
└── vercel.json        # Vercel 部署配置文件
```

## 📄 开源许可

本项目采用 [MIT License](LICENSE) 开源许可。

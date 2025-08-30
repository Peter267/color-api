# 🎨 图片主色调 API

一个基于 Python 构建的 Serverless API，它可以根据你提供的图片 URL，快速分析并返回该图片的主色调。项目被设计为在 Vercel 上进行一键部署，非常适合轻量级应用和个人项目。

## ✨ 主要功能

-   **快速颜色提取**: 输入任意图片的 URL，即可获得其主色调的十六进制 (HEX) 颜色码。
-   **简洁高效**: 使用 Python 和成熟的图像处理库，代码简洁且易于维护。
-   **Serverless 架构**: 部署在 Vercel 平台，无需管理服务器，具备高可用性和自动扩缩容能力。
-   **广泛的图片格式支持**: 依赖 Pillow 库，支持包括 JPEG, PNG, GIF, WebP 在内的多种主流图片格式。
-   **性能优化**: 通过先将图片缩放为缩略图再进行分析，极大地减少了网络下载和颜色计算的耗时。


## 🚀 API 使用指南

### 请求地址

```
GET /api/color
```

### 请求参数

| 参数  | 类型   | 描述                 | 是否必须 |
| :---- | :----- | :------------------- | :------- |
| `url` | string | 待分析图片的公开 URL | **是**   |

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

-   **后端语言**: [Python](https://www.python.org/)
-   **部署平台**: [Vercel](https://vercel.com/)
-   **核心 Python 库**:
    -   `requests`: 用于从 URL 下载图片。
    -   `Pillow`: 强大的图像处理库，用于打开、缩放和处理图片。
    -   `colorgram.py`: 用于从图片中轻松提取颜色。


## 本地开发与部署

### 准备工作

1.  安装 [Python](https://www.python.org/downloads/) (版本 3.8 或更高)。
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

2.  **安装 Python 依赖**
    建议在虚拟环境中进行此操作。
    ```bash
    # (可选) 创建并激活虚拟环境
    python -m venv venv
    source venv/bin/activate  # on Windows, use `venv\Scripts\activate`
    
    # 安装依赖
    pip install -r requirements.txt
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
    或者点击按钮部署 / [Fork](https://github.com/Peter267/color-api/fork) 仓库自行部署

    [![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FPeter267%2Fcolor-api)


## 📁 项目结构

```
.
├── api/
│   └── index.py         # Serverless 函数核心逻辑
├── README.md            # 项目说明文件
├── requirements.txt     # Python 依赖列表
└── vercel.json          # Vercel 部署配置文件
```

### `requirements.txt` 示例内容

```
requests
Pillow
colorgram.py
```

## ⚠️ 使用限制
  - 图片大小建议不要过大

  - 图片URL必须可公开访问

  - API调用频率请参考Vercel的无服务器函数限制

## 🤝 贡献
欢迎提交 Issue 和 Pull Request 来改进这个项目。

## 📄 开源许可

本项目采用 [MIT License](LICENSE) 开源许可。
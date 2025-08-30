import requests
import colorgram
from http.server import BaseHTTPRequestHandler
from urllib.parse import urlparse, parse_qs
import json
from io import BytesIO

class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        # 解析请求 URL 和查询参数
        parsed_path = urlparse(self.path)
        query_params = parse_qs(parsed_path.query)
        
        # 从查询参数中获取图片 URL
        image_url = query_params.get('url', [None])[0]

        if not image_url:
            self.send_response(400)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps({"error": "Missing image url"}).encode())
            return

        try:
            # 下载图片
            response = requests.get(image_url, stream=True)
            response.raise_for_status()
            
            # 从图片中提取颜色
            # colorgram.extract 需要一个文件对象，所以我们用 BytesIO 来包装下载的内容
            colors = colorgram.extract(BytesIO(response.content), 1)
            
            if not colors:
                self.send_response(400)
                self.send_header('Content-type', 'application/json')
                self.end_headers()
                self.wfile.write(json.dumps({"error": "Could not extract color from image"}).encode())
                return
            
            # 获取主色调
            dominant_color = colors[0]
            rgb = dominant_color.rgb
            
            # 将 RGB 转换为十六进制格式
            hex_color = '#{:02x}{:02x}{:02x}'.format(rgb.r, rgb.g, rgb.b)

            # 发送响应
            self.send_response(200)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps({"RGB": hex_color}).encode())

        except requests.exceptions.RequestException as e:
            self.send_response(400)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps({"error": f"Failed to download image: {e}"}).encode())
        except Exception as e:
            self.send_response(500)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps({"error": f"An error occurred: {e}"}).encode())
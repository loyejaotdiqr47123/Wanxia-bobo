from flask import Flask, redirect, request, send_file
import random

app = Flask(__name__)

# 访问根目录时显示API文档
@app.route('/')
def index():
    return send_file('index.html')

# 处理/pc.php请求，随机返回pc.txt中的图片链接
@app.route('/pc.php')
def random_pc_image():
    # 从pc.txt中读取图片链接列表
    with open('pc.txt', 'r') as f:
        images = f.readlines()

    # 随机选择一张图片链接
    random_image = random.choice(images).strip()

    # 通过301方式重定向到随机图片链接
    return redirect(random_image, code=301)

# 处理/pe.php请求，随机返回pe.txt中的图片链接
@app.route('/pe.php')
def random_pe_image():
    # 从pe.txt中读取图片链接列表
    with open('pe.txt', 'r') as f:
        images = f.readlines()

    # 随机选择一张图片链接
    random_image = random.choice(images).strip()

    # 通过301方式重定向到随机图片链接
    return redirect(random_image, code=301)

# 处理/api.php请求，根据请求头中的User-Agent进行重定向
@app.route('/api.php')
def redirect_based_on_ua():
    user_agent = request.headers.get('User-Agent')

    # 判断User-Agent是否为手机，如果是则重定向到/pe.php
    if 'Mobile' in user_agent:
        return redirect('/pe.php', code=301)
    # 如果不是手机，则重定向到/pc.php
    else:
        return redirect('/pc.php', code=301)

if __name__ == '__main__':
    # 启动Flask应用，并设置为多线程处理请求
    app.run(threaded=True)

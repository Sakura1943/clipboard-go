<h1 align="center">Online Clipboard</h1>

## Demo
[https://clipboard.sakura1943.top](https://clipboard.sakura1943.top)

## 描述
使用Golang为后端，Vue.js为前端编写的前后端分离的Web在线剪切板

> 准备条件<br>

Node >= 18.0.0<br>
Go >= 1.20<br>
Python >= 3.6<br>


## 🤖 构建
### 手动构建
#### 前端
```shell
# yarn
yarn
yarn run build
# npm
npm install
npm run build
# pnpm
pnpm install
pnpm build
```
#### 后端
```shell
cd backend
go mod tidy
go build
```

#### 编写`config.toml`配置文件
```toml
[base]
allowed_origins = [
    "http://127.0.0.1:80",
    "https://127.0.0.1:443",
    "http://127.0.0.1:8080",
    "http://localhost:80",
    "https://localhost:443",
    "http://localhost:8080",
    "http://localhost:5173",
    "http://127.0.0.1:5173" 
] # 这里是用到的允许跨域的全端页面地址，有域名的得把域名加上(必填)
server_port = 8000 # 服务端运行地址(必填)
server_host = "127.0.0.1" # 服务端运行地址(可空)
gin_mode = "release" # Go gin后端运行模式(分为debug和release, debug会打印更多信息，上线后推荐release)(必填)
```

### 使用安装脚本构建
#### 初始化
```shell
## 安装依赖
pip install -r requirements.txt
## 设置执行权限
chmod +x ./build
./build --init
```
#### 编辑`config.toml`配置文件
存放在`backend`文件夹内
```toml
[base]
allowed_origins = [
    "http://127.0.0.1:80",
    "https://127.0.0.1:443",
    "http://127.0.0.1:8080",
    "http://localhost:80",
    "https://localhost:443",
    "http://localhost:8080",
    "http://localhost:5173",
    "http://127.0.0.1:5173" 
] # 这里是用到的允许跨域的全端页面地址，有域名的得把域名加上(必填)
server_port = 8000 # 服务端运行地址(必填)
server_host = "127.0.0.1" # 服务端运行地址(可空)
gin_mode = "release" # Go gin后端运行模式(分为debug和release, debug会打印更多信息，上线后推荐release)(必填)
```
#### 开始构建
```shell
./build --build
```

### 放置在同一文件夹
```shell
# 回到项目根目录
mkdir -p bin/{backend,frontend}
cp -rf ./dist/* bin/frontend
cp -f ./backend/backend bin/backend
cp -f ./backend/config.toml bin/backend
```
## ⚙️运行
### 运行后端服务
```shell
# 返回项目根目录
cd bin
./backend/backend
## 一定要进入backend文件夹，因为程序需要读取`config.toml`文件
```
### 运行前端服务
```shell
# 使用能够解析前端页面的服务端程序， 比如nginx，apache, caddy等，这里我使用字自己编写的一个小工具进行解析
# 返回项目根目录
cd bin/frontend
simple_server -p 80
```

## 📖 关于在服务器进行后端服务进程挂起的问题
可以使用`tmux`或者`screen`的方式挂起<br>
`Linux`用户也可以编写`systemd`的`service`文件使用`systemd`进行进程守护，但需要指定`WorkingDirectory`变量到程序所在目录，下面是`service`
文件的编写实例

### 编写service文件
```shell
vim clipcoard-go.service
```
### 文件内容
```ini
[Unit]
Description=clipboard-go # 程序描述
After=network.target # 在网络单元启动后允许
 
[Service]
Type=simple
WorkingDirectory=/home/xxx/code/projects/clipboard/bin/backend # 程序工作目录
ExecStart=/home/xxx/code/projects/clipboard/bin/backend/backend # 嫌太长可以把backend所在目录添加到PATH变量中，然后按下面一行的命令执行
## mv /home/xxx/code/projects/clipboard/bin/backend/backend /home/xxx/code/projects/clipboard/bin/backend/clipboard-go-server
## echo "export $PATH:/home/xxx/code/projects/clipboard/bin/backend" >> ~/.bashrc
## 然后将ExecStart改成下面这样
#ExecStart=clipboard-go-server
Restart=on-failure
 
[Install]
WantedBy=multi-user.target
```

加载该`service`并启动
```shell
# 使用`root`权限，`sudo`也可以， `doas`也可以
cp clipboard-go.service /usr/lib/systemd/system/clipboard-go.service
systemctl daemon-reload
systemctl start clipboard-go.service
```

## 📖 使用
### 文件上传
访问后端接口 `/api/login`， 输入表单数据`name`(用户名)， `password`(用户密码)登录，获取数据`JSON`数据中的`extra.token`字段，保存token。

然后访问后端接口`/api/document/upload`， headers带`token`字段，form表单`file`参数为文件，获取返回的`JSON`数据的`extra.file-path`字段，然后访问前端界面即可获取内容`http(s)://前端地址/{extra.path}`。

## 📖 后端公共接口

| 接口                   | 接口类型 | 描述                    | 使用方法                                         |
| ---------------------- | -------- | ----------------------- | ------------------------------------------------ |
| `/api/document/upload` | `POST`   | 文件上传                | 传入文件表单参数`file`                           |
| `/api/login`      | `POST`   | 用户登录，以获取`token` | 传入form表单数据`name`(用户名), `password`(密码) |

## License
The MIT License ([MIT](https://opensource.org/licenses/MIT))

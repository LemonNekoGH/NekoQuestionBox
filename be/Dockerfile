from golang:alpine

# 创建 app 文件夹
run mkdir /app
# 指定 app 文件夹为工作目录
workdir /app
# 将当前文件夹中所有内容复制到镜像中
copy . /app/build
# 执行构建
run cd /app/build && go build -o qboxb main/main.go && mv qboxb ../qboxb && ls ..
# 删除源文件
run rm -rf /app/build
# 设定文件映射
volume ["/etc/qboxb/config"]
# 设定环境变量
env QBOX_CONFIG_PATH="/etc/qboxb/config/config.yaml"
# 设定入口点
entrypoint ["/app/qboxb"]
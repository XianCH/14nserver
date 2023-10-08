# 14nserver
介绍： http文件下载服务

**文件传输**:
通过post请求表单来处理文件流
返回json格式的响应到前段

**登录逻辑**

1. 通过post请求提交json格式的登录表单，
2. 向mysql查询 返回结构，
3. 如果查询到则生成token 并将token 和用户数据存储到redis中


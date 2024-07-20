## GO-WEB-SAMPLE
a sample golang web service for product API service

#api: 
there is two API for query product:/products/:id and /products,  refer to API doc

#database: 
选择可扩展的关系型数据库 postgre server hosted in Azure, init.sql .. password

#web backend server:（介绍以下使用到的常用框架）
gin:https://gin-gonic.com/docs/quickstart/
gorm: https://gorm.io/docs/index.html

#test:
unit test: in test folder,go test -v
E2E test: 不需要本地做任何配置 在云上github codebase就可以运行后端server 查看接口

other things:

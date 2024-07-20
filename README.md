# GO-WEB-SAMPLE
a sample golang web service for product read API service

## api: 
there is two API for query product:/products/:id and /products,  refer to [API doc](https://github.com/marsforwill/GO-WEB-SAMPLE/blob/main/API.md)

## database: 
选择可扩展的关系型数据库 postgre server hosted in Azure, init.sql .. password

## web backend frame:
### Gin 
Gin is a high-performance HTTP web framework written in Golang. It provides a robust set of features for building web applications and APIs. Key features of Gin include:

- **Fast**: Gin uses httprouter, which provides a lightweight and fast HTTP request router.
- **Middleware**: Gin supports middleware, allowing you to easily extend and customize your application.
- **Crash-Free**: Gin includes a built-in recovery mechanism that prevents the server from crashing.
- **Routing**: Gin provides easy-to-use and powerful routing capabilities.

For more information, visit the [Gin Documentation](https://gin-gonic.com/docs/quickstart/).

### GORM

GORM is the fantastic ORM library for Golang. It simplifies database operations and includes features like ORM, hooks, preload, transactions, and more. Key features of GORM include:

- **Full-Featured ORM**: Supports CRUD operations, associations (including has one, has many, belongs to, and many to many), hooks, eager loading, and more.
- **Migrations**: Provides schema migrations that can be integrated with your codebase.
- **SQL Builder**: Allows you to construct complex SQL queries.
- **Extendable**: GORM can be extended with plugins for additional functionality.

For more information, visit the [GORM Documentation](https://gorm.io/docs/index.html).


## test:
###  E2E test: 
highly recommended to run server through the github codebase virtual environment, which means that you do not need to do any environment configuration locally, just through the browser, you can see the server and API response
1) fork this repo and start/enter the corresponding [codeSpaces](https://docs.github.com/en/codespaces/overview)
<img width="1085" alt="image" src="https://github.com/user-attachments/assets/a7f47e54-a6e3-4318-980e-e3b25821cd9e">
2) in codeSpace, no need to set up any dependency, you can run server in root directory by command:
<img width="1000" alt="image" src="https://github.com/user-attachments/assets/8795c0ba-98be-4c57-a6e6-d80b546c45e1">
3) after server start, you can easily check the api and server status by your own query
<img width="1102" alt="image" src="https://github.com/user-attachments/assets/9042d2ba-4d86-49be-942a-e748116dbd9b">



### unit test: 
in test folder GO-WEB-SAMPLE/src/test, run command: go test -v, and you can see unit test output and result like this, it will ingest some test data into the same db instance, and verify the controller function output 
![image](https://github.com/user-attachments/assets/3d8acc66-7034-4a53-84b8-eedd5cb9738d)


## other things:

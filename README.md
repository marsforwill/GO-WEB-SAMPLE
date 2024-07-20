# GO-WEB-SAMPLE
a sample golang web service for product read API, using azure postgreSQL as database, using common Golang web frame Gin and Gorm, using github codeSpace as dev/run/test env.

## api: 
there is two API for query product:/products/:id and /products,  refer to [API doc](https://github.com/marsforwill/GO-WEB-SAMPLE/blob/main/API.md)

## database: PostGreSQL
1. **Scalable Relational Database**:
   It use PostgreSQL, a powerful, open-source object-relational database system to store data commonly. Hosting the server on a personal private Azure cloud instance allows for easy access and scalability. This setup eliminates the need for local database configuration and ensures that the backend server can be easily accessed and runnable from anywhere.
   ![image](https://github.com/user-attachments/assets/d96e88f9-a0ad-4101-9beb-0487f22f6000)


3. **Database Initialization**:
   The corresponding table/indices and some data injection can be viewed in the [`init.sql`](https://github.com/marsforwill/GO-WEB-SAMPLE/blob/main/init.sql).

4. **Database Connection Parameters**:
   The database connection parameters can be found in the `.env` file. While storing sensitive information in plain text is not a best practice, this is done here for ease of use and to facilitate running the project. The credentials and database instance information will be removed/updated soon for personal security.

5. **Database visualization**:
   You can use the connection parameters provided in the `.env` file with any PostgreSQL visualization tool to connect to the cloud-hosted database. This allows you to view and modify the data directly.
   <img width="1117" alt="image" src="https://github.com/user-attachments/assets/a459ee41-6ef6-4b6d-9bd2-d1cab47110b0">


## web backend frame:
### Gin 
Gin is a high-performance HTTP web framework written in Golang. It provides a robust set of features for building web applications and APIs. Key features of Gin include:

For more information, visit the [Gin Documentation](https://gin-gonic.com/docs/quickstart/).

### GORM

GORM is the fantastic ORM library for Golang. It simplifies database operations and includes features like ORM, hooks, preload, transactions, and more. Key features of GORM include:

For more information, visit the [GORM Documentation](https://gorm.io/docs/index.html).


## run&test:
###  run&test in cloud: 
highly recommended to run server through the github codebase virtual environment, which means that you **do not need to do any environment configuration locally**, just through the browser, you can see the server and API response
1) fork this repo and start/enter the corresponding [codeSpaces](https://docs.github.com/en/codespaces/overview)
<img width="1085" alt="image" src="https://github.com/user-attachments/assets/a7f47e54-a6e3-4318-980e-e3b25821cd9e">
2) in codeSpace, no need to set up any dependency, you can run server in root directory by command:
<img width="1000" alt="image" src="https://github.com/user-attachments/assets/8795c0ba-98be-4c57-a6e6-d80b546c45e1">
3) after server start, you can easily check the api and server status by your own query
<img width="872" alt="image" src="https://github.com/user-attachments/assets/ee76aeb6-27fc-4c12-ba2b-a2028d47bca9">


You can still run it locally as long as the environment is fully prepared (mainly golang sdk and package dependencies)

### unit test: 
in test folder GO-WEB-SAMPLE/src/test, run command: go test -v, and you can see unit test output and result like this, it will ingest some test data into the same db instance, and verify the controller function output 
![image](https://github.com/user-attachments/assets/3d8acc66-7034-4a53-84b8-eedd5cb9738d)


## other things:
we can scale this service if bottleneck in:
- database: scale the PostgreSQL database Vertical/Horizontal/Partitioning/Sharding
- web service: multi web instance
- lantency: Caching 

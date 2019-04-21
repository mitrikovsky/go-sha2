# go sha2 hasher

The service is for multiple sha2 hashing of the string. It starts API on the web server (host and port from config) and listens for requests for adding and getting hashing jobs.

* To start the app run `main.go`
* It starts HTTP API that located in `api` package
* API writes data to DB using `db` package
* When new hashing job is added to DB - API starts hash calculating process in background
* Hash calculating methods located in `core` package
* Configuration and installation files located in `config` package

## Configuration and installation
* Configuration for DB (PostgreSQL) and API is located in `config/config.go`
* Replace values with your own for correct work
* Before starting the app execute sql query from `config/db.sql`

## API endpoints
### POST /job
Adds hashing job for calculating, writes to DB, returns `id` of the job.
* Request
```
{
    "payload": <string>, // string for hashing
    "hash_rounds_cnt": <int> // rounds count for hashing
}
```
* Response
```
{
    "id": <int> // id of added job
}
```
### GET /job/\<id\>
Gets the hashing job info by provided \<id\>, returns all fields.
```
{
    "id": <int>, // id of the job
    "payload": <string>, // string for hashing
    "hash_rounds_cnt": <int>, // rounds count for hashing
    "status": <string>, // status of the job (in progress / finished)
    "hash": <string> // the final hash result
}
```

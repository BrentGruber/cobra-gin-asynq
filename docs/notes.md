* Tools:
    * database
        * queries: [sqlc](https://sqlc.dev/)
        * migrations: [migrate](https://github.com/golang-migrate/migrate)
    * HTTP: [gin](https://gin-gonic.com/)
    * Distributed task queue: [asynq](https://github.com/hibiken/asynq)
    * event bus: [redis streams](https://redis.io/docs/data-types/streams/)
* Thoughts:
    * would be interested in exploring Atlas for DB migrations
* Should be able to run in monolithic or microservice mode
* All services should be able to start from same docker image
* Not sure if workers running in go routines is worth it

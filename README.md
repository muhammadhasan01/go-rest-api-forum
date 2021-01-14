# Onboarding Hasan

A repository for Golang RESTful API  that provides the backend for a discussion forum with Users, Threads and Posts

## Prerequisites

Before using this project, you will need to install some software:
1. The programming language [Golang](https://golang.org/)
2. Storing the data with a database such as [PostgreSQL](https://www.postgresql.org/)
3. Testing out the endpoints with [Postman](https://www.postman.com/downloads/) (Optional)

## How to Run the project

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

1. Make a new `.env` file by changing the contents from [.env.example](.env.example)
2. Make a new database, and put the name in the `DB_NAME` assignment on the `.env` file
3. Run the project by typing `go run main.go` on the terminal
4. Your server will be hosted in `localhost:8888`

**Notes:**
- You don't need to worry about creating any tables, the migrations will do it for you
- If you want to seed the tables, you can use the [seeder](./migrations/seeder.go) and call it from the main function, for an example, use `migrations.CreateAccounts()` to seed the user table in [main.go](main.go)

## API Specification

You can see the API Specification on the [docs](.docs) folder, mainly on the [index.html](./docs/index.html) (generated using swagger), you can also load the [oas.json](oas.json) file to https://editor.swagger.io/

## Running the tests

There is some testing made in the [test](./test) folder, and you can run it by typing `go test` in the terminal, the test runs simultaneously so it is **recommended** that you change it manually and testing it one by one.

## Authors

* **Muhammad Hasan** - [muhammadhasan01](https://github.com/muhammadhasan01)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Golang Tutorials
* The Hydra Squad
* Pinhome CTO, Mr. Ahmed Aljunied :)
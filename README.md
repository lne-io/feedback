# Feedback

A simple service to get websites' users feedback.

## Description

- When you have a website or websites and you want to get visitors feedback, you can have a small button at the right bottom of any page and when a visitor clicks on it, a modal or a form shows up. This form can contain any kind of information you want your website's visitors to send you. Like email, subject and a message. In this application I have chosen to only get the subject and a description of the feedback, but before the user submits the form some javascript code will get the operating system and also the browser name and its version and also set the website uid. After clicking on submit button a request is sent to feedback endpoint. Request's payload then gets saved and a thank you message is shown to the user.
- I had some time and I wanted to build a simple service application. I only built the backend using Golang.
- I didn't write any unit nor integration tests.

## Getting Started

### Dependencies

To run this application you need to install these three tools

* Go 1.16
* SQLite
* Redis

### Installing

* Install all application dependencies inside go.mod file
* Update .env file `REDIS_ADDR` with your redis address 

### Executing program

There are two steps to run the application. First open project directory inside a terminal then:
* Run Task Queue server
``` bash
go run main.go -wroker
```
* Run gofiber web server
``` bash
go run main.go
```

## Version History

* 0.1
    * Initial Release

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgments

Libraries and tools used to build this application
* [Redis](https://github.com/redis/redis)
* [Go](https://github.com/golang/go)
* [SQLite](https://github.com/sqlite/sqlite)
* [Gofiber](https://github.com/gofiber/fiber)
* [Gorm](https://github.com/go-gorm/gorm)
* [Asynq](https://github.com/hibiken/asynq)
* [GoDotEnv](github.com/joho/godotenv)
* [ozzo-validation](github.com/go-ozzo/ozzo-validation)
* [uuid](github.com/google/uuid)
* [README template](https://gist.github.com/DomPizzie/7a5ff55ffa9081f2de27c315f5018afc)


# toggle-game
This is a rest Api that servers a the backend for games like Poker or Blackjack.

## Requirements
To run this app you should have the following installed:

- docker - [Get it here](https://www.docker.com/get-started)
- make - [Get it here](https://www.gnu.org/software/make/)

## To run in development
```
$ make run
```
## To run tests
```
$ make test_app
```

## To run tests while developing the app
```
$ make test_continuous
```
## Available Endpoints

| Endpoint                    | Method | Query Param | Datatype                                |
| /deck                       | POST   | shuffle     | optional boolean                        |
|                             |        | cards       | optional string of card codes separated |
|                             |        |             | by commas. Ex. AS,KD,AC,2C,KH           |
|                             |        |             |                                         |
| /deck/{uuid}                | GET    |             |                                         |
| /deck/{uuid}/draw?count=int | GET    | count       |                                         |

# TODO
- Add online hosted CI testing configuration such as CircleCI or TravisCi
- Add unit tests

[![Go Reference](https://pkg.go.dev/badge/github.com/agoblet/chesscompubapi.svg)](https://pkg.go.dev/github.com/agoblet/chesscompubapi)
[![codecov](https://codecov.io/gh/agoblet/chesscompubapi/branch/main/graph/badge.svg?token=U93I8XCM3X)](https://codecov.io/gh/agoblet/chesscompubapi)

# ♟️ chesscompubapi 

Go client for the chess.com PubAPI.

## Installation

```
$ go get -u github.com/agoblet/chesscompubapi
```

## Quick Start

```go
c := chesscompubapi.NewClient()
profile, err := c.GetPlayerProfile("hikaru")
```

## Endpoint Implementation Status

- ✅ /club/{url-ID}
- ✅ /country/{iso}
- ✅ /country/{iso}/clubs
- ✅ /country/{iso}/players
- ✅ /player/{username}
- ✅ /player/{username}/clubs
- ✅ /player/{username}/games/archives
- ✅ /player/{username}/games/{YYYY}/{MM}
- ✅ /player/{username}/games/{YYYY}/{MM}/pgn
- ✅ /player/{username}/stats
- ✅ /streamers
- ✅ /titled/{title-abbrev}
- ❌ /club/{url-ID}/members
- ❌ /club/{url-ID}/matches
- ❌ /leaderboards
- ❌ /match/{ID}
- ❌ /match/{ID}/{board}
- ❌ /match/live/{ID}
- ❌ /match/live/{ID}/{board}
- ❌ /player/{username}/games
- ❌ /player/{username}/games/to-move
- ❌ /player/{username}/is-online
- ❌ /player/{username}/matches
- ❌ /player/{username}/tournaments
- ❌ /puzzle
- ❌ /puzzle/random
- ❌ /tournament/{url-ID}
- ❌ /tournament/{url-ID}/{round}
- ❌ /tournament/{url-ID}/{round}/{group}

## Contributing

### Dev dependencies

- make
- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)

### Testing your changes locally

```
$ make
```

## Links 

- [chess.com PubAPI Documentation](https://www.chess.com/news/view/published-data-api)

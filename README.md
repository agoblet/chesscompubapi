[![Go Reference](https://pkg.go.dev/badge/github.com/agoblet/chesscompubapi.svg)](https://pkg.go.dev/github.com/agoblet/chesscompubapi)
[![codecov](https://codecov.io/gh/agoblet/chesscompubapi/branch/main/graph/badge.svg?token=U93I8XCM3X)](https://codecov.io/gh/agoblet/chesscompubapi)

# ♟️ chesscompubapi 

Go client for the chess.com Published-Data API (PubAPI).
This package depends on the standard library only.
It uses
- [net/http.Client](https://pkg.go.dev/net/http#Client) to send HTTP requests to chess.com
- [encoding/json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) to deserialize the response

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
- ✅ /club/{url-ID}/members
- ✅ /country/{iso}
- ✅ /country/{iso}/clubs
- ✅ /country/{iso}/players
- ✅ /player/{username}
- ✅ /player/{username}/clubs
- ✅ /player/{username}/games/archives
- ✅ /player/{username}/games/to-move
- ✅ /player/{username}/games/{YYYY}/{MM}
- ✅ /player/{username}/games/{YYYY}/{MM}/pgn
- ✅ /player/{username}/stats
- ✅ /puzzle
- ✅ /puzzle/random
- ✅ /streamers
- ✅ /titled/{title-abbrev}
- ❌ /club/{url-ID}/matches
- ❌ /leaderboards
- ❌ /match/{ID}
- ❌ /match/{ID}/{board}
- ❌ /match/live/{ID}
- ❌ /match/live/{ID}/{board}
- ❌ /player/{username}/games
- ❌ /player/{username}/matches
- ❌ /player/{username}/tournaments
- ❌ /tournament/{url-ID}
- ❌ /tournament/{url-ID}/{round}
- ❌ /tournament/{url-ID}/{round}/{group}

## Contributing

### Dev dependencies

- Go
- make
- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)

### Testing your changes locally

```
$ make
```

## Links 

- [chess.com PubAPI Documentation](https://www.chess.com/news/view/published-data-api)

# ♟️ chesscompubapi [![GoDoc][doc-img]][doc]

Go client for the [chess.com PubAPI](https://www.chess.com/news/view/published-data-api#pubapi-endpoint-games-archive).

## Installation

`go get -u github.com/agoblet/chesscompubapi`

## Quick Start

```go
c := chesscompubapi.NewClient()
profile, err := c.GetPlayerProfile("hikaru")
```

## Endpoint Implementation Status

- ✅ /country/{iso}
- ✅ /player/{username}
- ✅ /player/{username}/clubs
- ✅ /player/{username}/games/archives
- ✅ /player/{username}/games/{YYYY}/{MM}
- ❌ /club/{url-ID}
- ❌ /club/{url-ID}/members
- ❌ /club/{url-ID}/matches
- ❌ /country/{iso}/players
- ❌ /country/{iso}/clubs
- ❌ /leaderboards
- ❌ /match/{ID}
- ❌ /match/{ID}/{board}
- ❌ /match/live/{ID}
- ❌ /match/live/{ID}/{board}
- ❌ /player/{username}/games
- ❌ /player/{username}/games/to-move
- ❌ /player/{username}/games/{YYYY}/{MM}/pgn
- ❌ /player/{username}/is-online
- ❌ /player/{username}/matches
- ❌ /player/{username}/stats
- ❌ /player/{username}/tournaments
- ❌ /puzzle
- ❌ /puzzle/random
- ❌ /streamers
- ❌ /titled/{title-abbrev}
- ❌ /tournament/{url-ID}
- ❌ /tournament/{url-ID}/{round}
- ❌ /tournament/{url-ID}/{round}/{group}

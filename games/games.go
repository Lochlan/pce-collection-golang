package games

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

type Game struct {
    Id *int `jsonapi:"primary,games"`
    Name *string `jsonapi:"attr,name"`
    Slug *string `jsonapi:"attr,slug"`
    Developer *string `jsonapi:"attr,developer"`
}

func (g *Game) ToString() string {
    return *g.Name
}

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil { panic(err) }
    if db == nil { panic("db nil") }
    return db
}

func CreateTable(db *sql.DB) {
    sql_table := `
    CREATE TABLE IF NOT EXISTS games(
        Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Name TEXT,
        Slug TEXT UNIQUE,
        Developer TEXT,
        InsertedDatetime DATETIME
    );
    `

    _, err := db.Exec(sql_table)
    if err != nil { panic(err) }
}

func NewGame(db *sql.DB, game *Game) {
    sql_addgame := `
    INSERT INTO games(
        Name,
        Slug,
        Developer,
        InsertedDatetime
    ) values(?, ?, ?, CURRENT_TIMESTAMP)
    `

    stmt, err := db.Prepare(sql_addgame)
    if err != nil { panic(err) }
    defer stmt.Close()

    _, err2 := stmt.Exec(game.Name, game.Slug, game.Developer)
    if err2 != nil { panic(err2) }
}

func StoreGame(db *sql.DB, games []Game) {
    sql_addgame := `
    INSERT OR REPLACE INTO games(
        Id,
        Name,
        Slug,
        Developer,
        InsertedDatetime
    ) values(?, ?, ?, ?, CURRENT_TIMESTAMP)
    `

    stmt, err := db.Prepare(sql_addgame)
    if err != nil { panic(err) }
    defer stmt.Close()

    for _, game := range games {
        _, err2 := stmt.Exec(game.Id, game.Name, game.Slug, game.Developer)
        if err2 != nil { panic(err2) }
    }
}

func ReadGame(db *sql.DB) []Game {
    sql_readall := `
    SELECT Id, Name, Slug FROM games
    ORDER BY datetime(InsertedDatetime) DESC
    `

    rows, err := db.Query(sql_readall)
    if err != nil { panic(err) }
    defer rows.Close()

    var result []Game
    for rows.Next() {
        game := Game{}
        err2 := rows.Scan(&game.Id, &game.Name, &game.Slug, &game.Developer)
        if err2 != nil { panic(err2) }
        result = append(result, game)
    }
    return result
}

func ReadGameBySlug(db *sql.DB, slug string) *Game {
    sql_readbyslug := `
    SELECT Id, Name, Slug, Developer FROM games
    WHERE Slug=?
    `
    game := &Game{}
    err := db.QueryRow(sql_readbyslug, slug).Scan(&game.Id, &game.Name, &game.Slug, &game.Developer)

    if err == sql.ErrNoRows { return nil }
    if err != nil { log.Fatal(err) }

    return game
}

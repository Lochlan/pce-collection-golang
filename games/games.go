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

func (g *Game) LoadFromSlug(db *sql.DB) bool {
    sqlReadBySlug := `
    SELECT Id, Name, Developer FROM games
    WHERE Slug=?
    `

    err := db.QueryRow(sqlReadBySlug, *g.Slug).Scan(&g.Id, &g.Name, &g.Developer)
    if err == sql.ErrNoRows { return false }
    if err != nil { log.Fatal(err) }

    return true
}

func (g *Game) Save(db *sql.DB) {
    if g.Id == nil {
        createGame(db, g)
        return
    }
    updateGame(db, g)
}

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil { panic(err) }
    if db == nil { panic("db nil") }
    return db
}

func CreateTable(db *sql.DB) {
    sqlTable := `
    CREATE TABLE IF NOT EXISTS games(
        Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Name TEXT,
        Slug TEXT UNIQUE,
        Developer TEXT,
        InsertedDatetime DATETIME
    );
    `

    _, err := db.Exec(sqlTable)
    if err != nil { panic(err) }
}

func createGame(db *sql.DB, game *Game) {
    sqlCreateGame := `
    INSERT INTO games(
        Name,
        Slug,
        Developer,
        InsertedDatetime
    ) values(?, ?, ?, CURRENT_TIMESTAMP)
    `

    stmt, err := db.Prepare(sqlCreateGame)
    if err != nil { panic(err) }
    defer stmt.Close()

    _, err2 := stmt.Exec(game.Name, game.Slug, game.Developer)
    if err2 != nil { panic(err2) }
}

func updateGame(db *sql.DB, game *Game) {
    sqlUpdateGame := `
    INSERT OR REPLACE INTO games(
        Id,
        Name,
        Slug,
        Developer,
        InsertedDatetime
    ) values(?, ?, ?, ?, CURRENT_TIMESTAMP)
    `

    stmt, err := db.Prepare(sqlUpdateGame)
    if err != nil { panic(err) }
    defer stmt.Close()

    _, err2 := stmt.Exec(game.Id, game.Name, game.Slug, game.Developer)
    if err2 != nil { panic(err2) }
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

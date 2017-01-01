package main

import (
    "bufio"
    "encoding/csv"
    "github.com/lochlan/pce-collection-golang/games"
    "os"
)

var pwd, _ = os.Getwd()

func getCsvData(filepath string) [][]string {
    f, _ := os.Open(filepath)
    r := csv.NewReader(bufio.NewReader(f))

    records, err := r.ReadAll()
    if err != nil { panic(err) }

    return records
}

func main() {
    db := games.InitDB(pwd + "/test.db")
    defer db.Close()
    games.CreateTable(db)

    records := getCsvData(pwd + "/seed/game-list.csv")
    for _, record := range records {
        name := record[0]
        slug := record[1]
        // region := record[2]
        // format := record[3]
        // year := record[4]
        developer := record[5]

        new_game := &games.Game{
            Name: name,
            Slug: slug,
            Developer: developer,
        }
        games.NewGame(db, new_game)
    }
}

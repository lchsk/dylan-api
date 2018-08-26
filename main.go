package main

import (
    "bufio"
    "encoding/csv"
    "strconv"
    "fmt"
    "io"
    "log"
    "os"
)

type Song struct {
    Album string
    AlbumYear  int
    Name   string
    Text   string
}

func main() {
    csvFile, _ := os.Open("data/all_songs_with_metadata.csv")

    reader := csv.NewReader(bufio.NewReader(csvFile))

    var songs []Song

    for {
        line, error := reader.Read()

        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }

        year, err := strconv.Atoi(line[1])

        if err != nil {
            year = 0
        }

        songs = append(songs, Song{
            Album: line[0],
            AlbumYear:  year,
            Name:  line[2],
            Text: line[3],
        })
    }
    fmt.Println(songs[0].Text)
    fmt.Println(len(songs))
}

package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"rohmat.co.id/model"
	"rohmat.co.id/serverconfig"
	"strings"
	//"io"
	_ "github.com/lib/pq"
)

func connectDB() (db *sql.DB) {
	const (
		//10.10.5.167
		host     = "10.10.5.165"
		port     = 5432
		user     = "nexchief"
		password = "nexChief"
		dbname   = "nexSOFT"
		schema   = "nexchief"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schema)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return
}

type WriteCounter struct {
	Total uint64
}

func (wc WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %d complete", wc.Total)
}
func printProgressBar(total, counter int) {
	counter++
	intervalPrint := total / 100
	if counter%intervalPrint == 0 || total == counter {
		fmt.Println("progress ... ", rune((float64(counter)/float64(total))*100), "%")
	}
}

// Entry represents each stream. If the stream fails, an error will be present.
type Entry struct {
	Error error
	Data  []byte
}

// Stream helps transmit each streams withing a channel.
type Stream struct {
	stream chan Entry
}

// NewJSONStream returns a new `Stream` type.
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch watches JSON streams. Each stream entry will either have an error or a
// User object. Client code does not need to explicitly exit after catching an
// error as the `Start` method will close the channel automatically.
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts streaming JSON file line by line. If an error occurs, the channel
// will be closed.
func (s Stream) Start(path string) {
	// Stop streaming channel as soon as nothing left to read in the file.
	defer close(s.stream)

	// Open file to read.
	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("open file: %w", err)}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening delimiter. `[` or `{`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
		return
	}

	// Read file content as long as there is something.
	i := 1
	for decoder.More() {
		var temp interface{}
		if err := decoder.Decode(&temp); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode line %d: %w", i, err)}
			return
		}
		tempData, _ := json.Marshal(temp)
		//fmt.Println(tempData)
		s.stream <- Entry{Data: tempData}

		i++
	}

	// Read closing delimiter. `]` or `}`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		return
	}
}

func StartReadFile(path string, start func(db *sql.DB, data []byte) model.ErrorModel, partData string) {
	db := serverconfig.ServerAttribute.DBConnection
	//errString := "dial tcp 10.10.5.167:5432: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond."
	//errString2 := "read tcp 10.10.3.194:56851->10.10.5.167:5432: wsarecv: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond."

	stream := NewJSONStream()
	i := 0
	var err  model.ErrorModel
	go func() {
		for data := range stream.Watch() {
			i++
			if data.Error != nil {
				log.Println(data.Error)
				continue
			}
			//if i == 156 {
			//	return
			//}
			err = start(db, data.Data)
			if err.Error != nil && err.Code == 500 {
				log.Println(err.Error)
				//_, _ = fmt.Fprintln(serverconfig.ServerAttribute.Write, fmt.Sprintf(partData+" - %d error : %s", i, err.Error))
				//os.Exit(3)
				continue
			} else if err.Error != nil && err.Code != 500 {
				continue
			}

			fmt.Println(partData, "-", i)
			_, _ = fmt.Fprintln(serverconfig.ServerAttribute.Write, fmt.Sprintf(partData+" - %d", i))
		}
	}()
	stream.Start(path)
	fmt.Println(partData, "finish")
}

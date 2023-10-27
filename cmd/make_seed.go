package main

import (
	"flag"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//conf := config.All()
	//dbConfig := coreDB.NewConfiguration(&conf.Database)
	//db := coreDB.InitDB(dbConfig)

	flag.Parse()
	//args := flag.Args()

	//seeder := seeder.New(db)
	//seeder.Execute(args[0:]...)
	os.Exit(0)
}

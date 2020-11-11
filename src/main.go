package main

import (
	"context"
	"time"

	"github.com/sky0621/try-sqlboiler/src/boiled"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	dsn := "host=localhost port=15432 dbname=localdb user=postgres password=yuckyjuice sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Err(err).Send()
		return
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Err(err).Send()
		return
	}
	boil.SetLocation(loc)

	log.Debug().Msg("setupLocalRDB___END")

	ctx := context.Background()
	item, err := boiled.FindItem(ctx, db, 1)
	if err != nil {
		log.Err(err).Send()
		return
	}

	log.Info().Msgf("%#+v", item)
}

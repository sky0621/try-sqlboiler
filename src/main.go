package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sky0621/try-sqlboiler/src/boiled"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

	orderKey := "name"
	orderDirection := "ASC"
	table := boiled.TableNames.Customer
	baseCondition := "1=1"
	compareSymbol := ">"
	decodedCursor := 5
	limit := 5

	q := `
		SELECT * FROM (
			SELECT ROW_NUMBER() OVER (ORDER BY %s %s) AS rownum, *
			FROM %s
		) AS tmp
		WHERE %s
		AND rownum %s %d
		LIMIT %d
	`
	results, err := boiled.Customers(qm.SQL(fmt.Sprintf(q,
		orderKey, orderDirection,
		table,
		baseCondition, compareSymbol, decodedCursor,
		limit,
	))).All(ctx, db)
	if err != nil {
		log.Err(err).Send()
		return
	}

	for _, result := range results {
		log.Info().Msgf("%#+v", result)
	}
}

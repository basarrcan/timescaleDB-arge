package inits

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// connect to database using a single connection
func InitTable() {
	/***********************************************/
	/* Single Connection to TimescaleDB/ PostresQL */
	/***********************************************/
	ctx := context.Background()
	connStr := "postgres://postgres:password@localhost:5432/example"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	/********************************************/
	/* Create relational table                      */
	/********************************************/

	//Create relational table called sensors
	queryCreateTable := `CREATE TABLE sensors (id SERIAL PRIMARY KEY, type VARCHAR(50), location VARCHAR(50));`
	_, err = conn.Exec(ctx, queryCreateTable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create SENSORS table: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created relational table SENSORS")

}

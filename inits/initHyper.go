package inits

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// connect to database using a single connection
func InitHyperTable() {
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
	/* Create Hypertable                        */
	/********************************************/
	// Create hypertable of time-series data called sensor_data

	//formulate statement
	queryCreateHypertable := `CREATE TABLE sensor_data (
		time TIMESTAMPTZ NOT NULL,
		sensor_id INTEGER,
		temperature DOUBLE PRECISION,
		cpu DOUBLE PRECISION,
		FOREIGN KEY (sensor_id) REFERENCES sensors (id)
		);
		SELECT create_hypertable('sensor_data', 'time');
		`

	// execute statement
	_, err = conn.Exec(ctx, queryCreateHypertable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create SENSOR_DATA hypertable: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created hypertable SENSOR_DATA")

}

package utils_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/4epyx/authrpc/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

func getAllTables(ctx context.Context, db *pgxpool.Pool) ([]string, error) {
	res, err := db.Query(ctx, "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema'")
	if err != nil {
		return nil, err
	}

	tables := make([]string, 0, 5)
	for res.Next() {
		var table string
		res.Scan(&table)
		tables = append(tables, table)
	}

	return tables, nil
}

func TestConnectToDB(t *testing.T) {
	type args struct {
		ctx           context.Context
		connectionURL string
	}

	url, ok := os.LookupEnv("TEST_DB_URL")
	if !ok {
		t.Fatal("have not ENV")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				ctx:           ctx,
				connectionURL: url,
			},
			wantErr: false,
		},
		{
			name: "incorrect url",
			args: args{
				ctx:           ctx,
				connectionURL: "pg://fakeuser:incorrectpassword@incorrecthost:1111/db",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := utils.ConnectToDB(tt.args.ctx, tt.args.connectionURL)
			defer db.Close()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMigrateTable(t *testing.T) {
	type args struct {
		ctx  context.Context
		conn *pgxpool.Pool
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := utils.ConnectToDB(ctx, os.Getenv("TEST_DB_URL"))
	if err != nil {
		t.Error("Failed to connect to database")
	}
	defer db.Close()

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				ctx:  ctx,
				conn: db,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := utils.MigrateTable(tt.args.ctx, tt.args.conn); (err != nil) != tt.wantErr {
				t.Fatalf("MigrateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
			defer db.Exec(ctx, "DROP TABLE users")

			tables, err := getAllTables(ctx, db)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			found := false
			for _, t := range tables {
				if t == "users" {
					found = true
					break
				}
			}
			if !found {
				t.Error("table not created")
			}

		})
	}
}

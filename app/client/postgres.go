package client

import (
	"database/sql"
	"fmt"
)

type PostgresClient struct {
	Db *sql.DB
}

func (c *PostgresClient) GetSchemas() (Result, error) {
	var result Result

	data, err := executeQuery(c.Db, "SELECT * FROM information_schema.schemata")
	if err != nil {
		return result, err
	}
	result.Data = data

	info, err := executeQuery(c.Db, "SELECT * FROM information_schema.columns WHERE table_schema = 'information_schema' AND table_name = 'schemata'")
	if err != nil {
		return result, err
	}
	result.Info = info

	return result, nil
}

func (c *PostgresClient) GetTables(schema string) (Result, error) {
	var result Result

	data, err := executeQuery(c.Db, "SELECT * FROM information_schema.tables WHERE table_schema = $1", schema)
	if err != nil {
		return result, err
	}
	result.Data = data

	info, err := executeQuery(c.Db, "SELECT * FROM information_schema.columns WHERE table_schema = 'information_schema' AND table_name = 'tables'")
	if err != nil {
		return result, err
	}
	result.Info = info

	return result, nil
}

func (c *PostgresClient) GetTable(schema string, table string) (Result, error) {
	var result Result

	data, err := executeQuery(c.Db, fmt.Sprintf("SELECT * FROM %s.%s", schema, table))
	if err != nil {
		return result, err
	}
	result.Data = data

	info, err := executeQuery(c.Db, "SELECT * FROM information_schema.columns WHERE table_schema = $1 AND table_name = $2", schema, table)
	if err != nil {
		return result, err
	}
	result.Info = info

	return result, nil
}

func (c *PostgresClient) ExecuteQuery(query string, args ...interface{}) (QueryResult, error) {
	return executeQuery(c.Db, query, args...)
}

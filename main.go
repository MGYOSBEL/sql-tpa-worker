package main

import (
	"fmt"
	"os"

	"github.com/MGYOSBEL/querybuilder"
)

func main() {
	config, err := querybuilder.LoadConfig("test.json")
	if err != nil {
		fmt.Println("Error loading the config file")
		os.Exit(1)
	}

	for _, table := range config.Tables {
		query := querybuilder.SelectColumnsFromTableWhereCondition(table.Columns, table.Name, table.WhereCondition, table.Options)
		fmt.Println("SELECT query: ")
		fmt.Println(query)

		query = querybuilder.UpdateTableSetColumnWherePrimaryKeyInIds(table.Name, table.UpdateStatement, table.PrimaryKey[0], "1, 2, 3")
		fmt.Println("UPDATE query: ")
		fmt.Println(query)
	}

}

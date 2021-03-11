package main

import (
	"encoding/json"
	"fmt"
	"github.com/MGYOSBEL/querybuilder"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error reading the file")
		os.Exit(1)
	}

	var config querybuilder.Config
	err = json.Unmarshal(data, &config)

	fmt.Printf("%s", data)




	for _, table := range config.Tables  {
		query := querybuilder.SelectColumnsFromTableWhereCondition(table.Columns, table.Name, table.WhereCondition, table.Options )
		fmt.Println("SELECT query: ")
		fmt.Println(query)

		query = querybuilder.UpdateTableSetColumnWherePrimaryKeyInIds(table.Name, table.UpdateStatement, table.PrimaryKey[0], "1, 2, 3")
		fmt.Println("UPDATE query: ")
		fmt.Println(query)
	}

}

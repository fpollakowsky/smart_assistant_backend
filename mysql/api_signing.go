package mysql

import "fmt"

func GetApiKeys() [][]string {
	db := connect()
	var keys [][]string
	var row1 []string
	var row2 []string

	results, err := db.Query("SELECT * FROM ehome.api_keys")
	if err != nil {
		fmt.Println(err)
	}

	for results.Next() {
		var key, company string
		err = results.Scan(&key, &company)
		if err != nil {
			panic(err.Error())
		}
		row1 = append(row1, key)
		row2 = append(row2, company)
	}

	keys = append(keys, row1)
	keys = append(keys, row2)

	defer db.Close()
	return keys
}

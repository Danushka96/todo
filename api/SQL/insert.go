package SQL

import (
	"fmt"
	"github.com/danushka96/gowork/SQL/Connection"
	"github.com/danushka96/gowork/SQL/Models"
	"log"
	"strconv"
)

func InsertData(t Models.Table, m map[string]string) bool {
	db := Connection.GetInstance()
	tx, err := db.Begin()
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	fieldString := "("
	valuesString := "("
	for k, v := range m {
		fieldString += k + ","
		valuesString += strconv.Quote(v) + ","
	}
	fieldString = string([]rune(fieldString)[0:len(fieldString)-1]) + ")"
	valuesString = string([]rune(valuesString)[0:len(valuesString)-1]) + ")"
	//fieldString += ")"
	//valuesString += ")"
	queryString := fmt.Sprintf("INSERT INTO %s%s VALUES %s", t.Name, fieldString, valuesString)
	stmt, err := tx.Prepare(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	_, err = stmt.Exec()
	if err!=nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return true
}

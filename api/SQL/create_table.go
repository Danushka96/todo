package SQL

import (
	"fmt"
	"github.com/danushka96/gowork/SQL/Connection"
	"github.com/danushka96/gowork/SQL/Models"
	"log"
)

func CreateTable(t Models.Table) bool {
	db := Connection.GetInstance()
	queryString := "CREATE TABLE " + t.Name + " ( "
	for i := 0; i < len(t.Fields); i++ {
		var singleType = fmt.Sprintf("%s %s ", t.Fields[i].Name, t.Fields[i].DataType)
		queryString += singleType
		if i != len(t.Fields)-1 {
			queryString += ","
		}
	}
	queryString += " )"
	_, err := db.Exec(queryString)
	if err != nil {
		log.Printf("%q: %s\n", err, queryString)
		return false
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return true
}

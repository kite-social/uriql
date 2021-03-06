package builder

import (
	"fmt"
	"github.com/restra-social/uriql/models"
	"strings"
)

type cypherQueryBuilder struct {
}

// GetCypherBuilder : Get Cypher Builder Object
func GetCypherBuilder() *cypherQueryBuilder {
	return &cypherQueryBuilder{}
}

func (n *cypherQueryBuilder) Build(allparam []models.QueryParam) string {

	var str string
	//total := len(allparam)
	for _, model := range allparam {

		var conNVal string
		//#todo#fix token condition need to be fixed
		switch model.Condition {
		case "like":
			conNVal = fmt.Sprintf("%s '%%%s%%'", model.Condition, model.Value.Value)
		case "=":
			conNVal = fmt.Sprintf("%s '%s'", model.Condition, model.Value.Value)
		default:
			conNVal = fmt.Sprintf("%s '%s'", model.Condition, model.Value.Value)
		}

		fmt.Println(conNVal)

		switch model.FHIRType {
		case "relation":

			str += fmt.Sprintf("MATCH (u:%s{id:'%s'})-[r:%s]-(f)", model.RequestInfo.Type, model.RequestInfo.UserID, strings.ToUpper(model.Resource))

			str += fmt.Sprintf(" WHERE ")
			str += fmt.Sprintf("r.%s %s", model.FieldsInfo.ObjectPath, conNVal)

			str += fmt.Sprintf(" RETURN f")

		case "node":

			str += fmt.Sprintf("MATCH (u:%s{id:'%s'})-[:%s]-(f) RETURN f", model.RequestInfo.Type, model.RequestInfo.UserID, strings.ToUpper(model.Resource))
		}

	}

	return str
}

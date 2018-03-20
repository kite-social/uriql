package dictionary

import "github.com/restra-social/uriql/models"

func N1QLDictionary() map[string]map[string]models.SearchParam {

	dict := map[string]map[string]models.SearchParam{

		"product": map[string]models.SearchParam{

			"_id": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"id"},
			},
			"category": models.SearchParam{
				Type:      "join",
				FieldType: "string",
				Path:      []string{"name"},
				Join:      []string{"category", "variances"},
			},
		},

		"order": map[string]models.SearchParam{

			"_id": models.SearchParam{ // will be used by admin
				Type:      "string",
				FieldType: "string",
				Path:      []string{"id"},
			},
			"store_id": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Select:    []string{"id", "type", "status", "[]stores.id"},
				Path:      []string{"[]stores.id"}, // #todo must be case insensitive
			},
			"product_id": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"[]stores.[]orders.id"}, // #todo must be case insensitive
			},
		},

		"profile": map[string]models.SearchParam{

			"_id": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"id"},
			},
			"gender": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"gender"},
			},
			"name": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"name.prefix", "name.first_name", "name.last_name"},
			},
			"address": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"address.city.name", "address.state.name", "address.street"},
			},
			"address-street": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"address.street"},
			},
			"hobbies": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"[]hobbies"},
			},
			"postal": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"address.postal"},
			},
		},

		"restaurant": map[string]models.SearchParam{

			"_id": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"id"},
			},
			"title": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"title"},
			},
			"status": models.SearchParam{
				Type:      "string",
				FieldType: "string",
				Path:      []string{"status"},
			},
		},
	}

	return dict
}

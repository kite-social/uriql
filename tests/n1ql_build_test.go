package tests

import (
	decoder "github.com/restra-social/uriql"
	search "github.com/restra-social/uriql/builder"
	"github.com/restra-social/uriql/dictionary"
	"github.com/restra-social/uriql/models"
	"testing"
)

func printResult(t *testing.T, p string, qp interface{}, q string) {
	t.Logf("Decoding : %s", p)
	//t.Logf("Decoded to : %+v", qp)
	t.Logf("Query %s", q)
}

func TestN1QLBuild(t *testing.T) {

	dict := &models.Dictionary{Model: dictionary.N1QLDictionary(), Bucket: "restra"}

	decode := decoder.GetQueryDecoder(dict)
	builder := search.GetN1QLQueryBuilder(dict.Bucket, "type")

	t.Log("Testing Restaurant Parameter : ")

	p := "profile?name=mr&hobbies=sports"
	qp := decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "profile", Query: p})
	q := builder.Build(qp)
	printResult(t, p, qp, q)

	p = "profile?hobbies=sports"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "profile", Query: p, Filter: models.Filter{Limit: 10, Page: 1}})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "profile?name=mr"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "profile", Query: p, Filter: models.Filter{Limit: 10, Page: 3}})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	/*p = "Patient?name:contains=Mr."
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "Patient", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "Patient?organization=Medical/10234"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "Patient", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "Patient?identifier=http://acme.org/patient|2345"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "Patient", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "Patient?general-practitioner=Practitioner/2345"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "Patient", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "Patient?active=true"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "Patient", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "Patient?address-use=Dhaka"
	qp = decode.DecodeQueryString(models.RequestInfo{UserID: "1234567890", Type: "Patient", Query: p})
	q = builder.Build(qp)
	printResult(t, p, qp, q)*/

	///Observation?subject=Patient/23

	/*decode = search.GetQueryDecoder(dictionary.RestaurantItemsDictionary())
	builder = builder.GetN1QLQueryBuilder()

	t.Log("Testing Restaurant Items Parameter : ")


	p = "restaurant-items?foods=burger"
	qp = decode.DecodeQueryString(p)
	q = builder.Build(qp)
	printResult(t, p, qp, q)

	p = "restaurant-items?foods-price=45"
	qp = decode.DecodeQueryString(p)
	q = builder.Build(qp)
	printResult(t, p, qp, q)*/

}

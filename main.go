package main

import (
	"io/ioutil"
	"kktestgo/controllers"
	"log"
	// "github.com/krazybee/gofpdf"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jung-kurt/gofpdf"
)

func pdf() {
	var pdf *gofpdf.Fpdf

	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "OK")
	if pdf.Error() != nil {
		log.Printf("not expecting error when rendering text")
	}

	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.Cell(40, 10, "Not OK") // Font not set
	if pdf.Error() == nil {
		log.Printf("expecting error when rendering text without having set font")
	}
}

// Handler is executed by AWS Lambda in the main function. Once the request
// is processed, it returns an Amazon API Gateway response object to AWS Lambda
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	index, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	controllers.Test()
	pdf()

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(index),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}

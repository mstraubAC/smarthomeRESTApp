package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func getLocation(locationId string) {
	// start tracer
	tracer := otel.Tracer("testapi")

	// start a span
	ctxTrace, span := tracer.Start(context.Background(), "GetLocation-Call") //, trace.WithStackTrace(true), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	if span.IsRecording() {
		println("Span is recording")
		span.SetAttributes(
			attribute.Int64("userId", 508),
			attribute.String("userName", "testuer508"),
		)
	}

	// create http client
	// create an instrumented HTTP client
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	url := "http://localhost:7777/v1/locations/" + locationId

	req, err := http.NewRequestWithContext(ctxTrace, "GET", url, nil)
	if err != nil {
		// error handling
		println(" --> Error while creating request: " + err.Error())
		return
	}

	// print(tracer.

	resp, err := client.Do(req)
	if err != nil {
		println("  --> Execution failed at Do: " + err.Error())
		log.Fatal(err)
	}

	bodyBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		println("  --> Execution failed while reading response: " + err.Error())
		log.Fatal(readErr)
	}
	bodyString := string(bodyBytes)
	println("  --> Status Code: " + strconv.Itoa(resp.StatusCode))
	println(bodyString)
}

func main() {
	println("Call should be successful")
	getLocation("0")

	println("")
	println("-------------------------------------------")
	println("Should be failing call")
	getLocation("h")
}

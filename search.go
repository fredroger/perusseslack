package main

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

const developerKey = "AIzaSyAXu54yqnQKQTQdVMvQgF-HMJwZJTXY7Zc"

func searchListByKeyword(part string, maxResults int64, q string, typeArgument string) *youtube.SearchListResponse {

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	call := service.Search.List(part)
	if maxResults != 0 {
		call = call.MaxResults(maxResults)
	}
	if q != "" {
		call = call.Q(q)
	}
	if typeArgument != "" {
		call = call.Type(typeArgument)
	}

	response, err := call.Do()
	if err != nil {
		log.Printf("Error making search API call: %v", err)
	}

	return response
}

package Dbservice

import (
	"context"
	"log"

	"google.golang.org/api/option"

	"cloud.google.com/go/bigquery"
)
const (
   
	datasetID = "<DATASET ID>"
   projectID = "<project id>"
)
func GetBQClient()*bigquery.Client {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile("<creds>"))
	if err != nil {
	  log.Fatalf("Failed to create client: %v", err)
	}
	
   return client
}
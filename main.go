package main

import (
	"changefeed/Dbservice"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)


const (
   PORT = ":8080"
   //datasetID = "staging"
   datasetID = "production"
   projectID = "iserveuprod"
)

func main() {
    fmt.Println("In Main")
	http.HandleFunc("/",changedDataCapture)
	http.ListenAndServe(PORT, nil)
}
type mark map[string]interface{}

func changedDataCapture(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Coming From Changefeed")
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Println("Data coming from Changefeed ",string(reqBody))
    // Decoding the interface data
    digit := json.NewDecoder(strings.NewReader(string(reqBody)))
    for digit.More() {
        var result mark
        err := digit.Decode(&result)
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
        fmt.Println("final_data ", result )
        //check for interface is nil or not
        if result != nil && len(result) > 0 {
        Balance := result.Data("after").Balance("balance")
        fmt.Println("Balance:: ", Balance)
        Id := result.Data("after").Id("id")
        fmt.Println("id:: ", Id)
        //bigquery client
        ctx := context.Background()
        client := Dbservice.GetBQClient()
        defer client.Close()
        Date:=time.Now().Format("2006-01-02 15:04:05")
        table_date:= time.Now().Format("20060102")
	    tableID := "wallet3_cdc_"+table_date 
        fmt.Println("Current TableId of BigQuery:",tableID)
        //checking tableid form bigquery
        md,err := client.Dataset(datasetID).Table(tableID).Metadata(ctx)
        if err != nil {
            fmt.Println("Failed to detect any table metadata::",err)
            return
        }
        fmt.Printf("Number of rows present in the tableid:: %s is:: %d\n",tableID,md.NumRows)
        //Big query Table Details
        tabledetails := projectID+"."+datasetID+"."+tableID
        //Data to be insertrd in bigquery table
        amount:=fmt.Sprintf("%v",Balance) 
        id:=fmt.Sprintf("%v",Id)
        fmt.Println(id,amount,Date)
        //query to be executed for bigquery
        q:= client.Query(`insert `+"`"+tabledetails+"`"+` (id,balance,date) values(`+id+`,`+amount+`,`+"'"+Date+"'"+`)`)
        job, err := q.Run(ctx)
        if err != nil {
        fmt.Println("Error in insert BigQuery",err)
        return
    }
    //after inertion returning jobid
    fmt.Println("Inserted Successfully in bigquery with the ID :",job.ID())
    } else {
        fmt.Println("Interface is nil")
    }
 }
}

func (d mark) Data(s string) mark {
    return d[s].(map[string]interface{})
 }
 
 func (d mark) Balance(s string) float64 {
    return d[s].(float64)
 }

 func (d mark) Id(s string) float64 {
    return d[s].(float64)
 }

package main

//CDC OPTIMISED USING ndjson.
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/olivere/ndjson"
)


var (
PORT = ":8080"
)

func main() {
    fmt.Println("In Main")
	http.HandleFunc("/",changedDataCapture)
	http.ListenAndServe(PORT, nil)
}


type Data struct {
	AfterData After `json:"after"`
	Key       []int `json:"key"`
  }
  
type After struct {
	Id	int64 `json:"id"`
	Balance float64 `json:"balance"`
	HoldBalance float64 `json:"hold_balance"`
	Status int64 `json:"status"`
	MinimumBalance float64 `json:"minimum_balance"`
	MaximumBalance float64 `json:"maximum_balance"`
}

func changedDataCapture(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Coming From Changefeed")
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Println("Data coming from Changefeed ",string(reqBody))
	
	digit := ndjson.NewReader(strings.NewReader(string(reqBody)))
    for digit.Next() {
      var data Data
      if err := digit.Decode(&data); err != nil {
        fmt.Println("Decode failed", err)
        return
      }
      fmt.Println("CDC data from DB",data.AfterData)
      fmt.Println("ID",data.AfterData.Id)
      fmt.Println("status",data.AfterData.Status)
      fmt.Println("Balance",data.AfterData.Balance)
      

	
} 
}

 




 
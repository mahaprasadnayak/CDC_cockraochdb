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
// topic      *pubsub.Topic
// my_project = "iserveustaging"
// my_topic   = "staging_topic"
// ctx = context.Background()
// Client *pubsub.Client
)

func main() {
    fmt.Println("In Main")
	http.HandleFunc("/",changedDataCapture)
	http.ListenAndServe(PORT, nil)
}

// type message struct {
// 	Key string            `json:"key"`
// 	Msg After `json:"message"`
// }
type Data struct {
	AfterData After `json:"after"`
	Key       []int `json:"key"`
  }
  
type After struct {
	// Id               int64 `json:"id"`
	// PreviousBalance  float64 `json:"previousbalance"`
	// Amount           float64 `json:"amount"`
	// CurrentBalance   float64 `json:"currentbalance"`
	// WalletId         int64 `json:"walletid"`
	// Status           string  `json:"status"`
	// Type             int64 `json:"type"`
	// ReversalId       int64 `json:"reversalid"`
	// StatusCode      int64 `json:"status_code"`
	// Txnid            int64 `json:"txnid"`
	// CreatedDate      string `json:"createddate"`
	// UpdatedDate      string `json:"updateddate"`
	// TransactionType string `json:"transaction_type"`
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
	// createclient_pubsub()
    // defer Client.Close()
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
      

	// creating pubsub client
        
	// m := message{
	// 	Key: "wallet1",
	// 	Msg: data.AfterData,
	// }
       
	// jsonData, err := json.Marshal(m)
	// if err != nil {
	// fmt.Println(err)
	// }
	// topic = Client.Topic(my_topic)
	// defer topic.Stop()
	// topic.PublishSettings.NumGoroutines = 2
	// result := topic.Publish(ctx, &pubsub.Message{Data: jsonData})
	// id, err := result.Get(ctx)
	// if err != nil {
	// 	fmt.Println("Error in Publish: ", err)
	// }
    // fmt.Println("Published a message with msg ID:", id,"txnid::",data.AfterData.Txnid)
} 
}

 

// func createclient_pubsub() {
// 	var err error
// 	Client, err = pubsub.NewClient(ctx, my_project)

// 	if err != nil {
// 		fmt.Println("Errror in Pub/Sub client creation globally :", err)
// 	} else {
// 		topic = Client.Topic(my_topic)
// 		fmt.Println("Pubsub client created successfullly and Client!,  and topic :", topic)
// 	}

// }


 
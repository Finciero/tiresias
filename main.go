package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/rodrwan/cat-grpc/categoryapi"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

var (
	url = flag.String("url", "localhost:10000", "categorizer url")
)

func categorizeHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("w")

	conn, err := grpc.Dial(*url, grpc.WithInsecure())
	c := pb.NewCategoryAPIClient(conn)

	response, err := c.Categorize(context.Background(), &pb.Transaction{
		Description: word,
	})
	if err != nil {
		grpclog.Fatalf("failed to categorize: %v", err)
	}

	fmt.Fprintf(w, "categorize: %s\ncategory: %s", word, response.CategoryName)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/categorize", categorizeHandler)
	fmt.Println("Listening on 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

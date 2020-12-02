package main

import (
	"context"
	"github.com/anishj0shi/cf-kyma-k8s/pkg"
	"github.com/anishj0shi/cf-kyma-k8s/pkg/commerce_client"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
	"net/http"
)

var cartURL = "https://accstorefront.crz8nbwca-internalc6-d29-public.model-t.cc.commerce.ondemand.com/?site=electronics"

func main() {
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	h, err := cloudevents.NewHTTPReceiveHandler(ctx, p, receiveFn)
	if err != nil {
		log.Fatalf("failed to create handler: %s", err.Error())
	}

	router := http.NewServeMux()
	router.HandleFunc("/createCart", handleCart)
	router.Handle("/", h)

	log.Printf("will listen on :8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("unable to start http server, %s", err)
	}

	//client, err := cloudevent.NewDefaultClient()
	//if err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//if err := client.StartReceiver(context.Background(), receiveFn); err != nil {
	//	log.Fatalf("%+v", err)
}

func receiveFn(event cloudevents.Event) {
	log.Printf("event : %v", event)
	eventData := &pkg.OrderCreatedEvent{}
	err := event.DataAs(eventData)
	if err != nil {
		log.Printf("Error received: %v", err)
	}
	orderClient := commerce_client.NewOrderClient(eventData)
	orderClient.SendRecommendation()
}

func handleCart(w http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query()["user"]
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	productIds := req.URL.Query()["products"]
	if len(productIds) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	log.Printf("products..%v", productIds)
	log.Printf("userid..%v", userId)

	cartClient := commerce_client.NewShoppingCartClient(userId[0])
	cartClient.CreateShoppingCart()
	cartClient.AddProductsToCart(productIds...)

	http.Redirect(w, req, cartURL, http.StatusSeeOther)
}

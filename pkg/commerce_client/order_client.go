package commerce_client

import (
	"encoding/json"
	"fmt"
	"github.com/anishj0shi/cf-kyma-k8s/pkg"
	"github.com/anishj0shi/cf-kyma-k8s/pkg/slack_client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type orderClient struct {
	gatewayUrl string
	orderEvent *pkg.OrderCreatedEvent
}

type OrderClient interface {
	GetOrders() interface{}
	SendRecommendation()
	CreateOrder() error
}

func NewOrderClient(event *pkg.OrderCreatedEvent) OrderClient {
	return &orderClient{
		gatewayUrl: os.Getenv("GATEWAY_URL"),
		orderEvent: event,
	}
}

func (o *orderClient) GetOrders() interface{} {
	url := fmt.Sprintf("%s/electronics/orders/%s", o.gatewayUrl, o.orderEvent.OrderCode)
	res, err := http.Get(url)
	if err != nil {
		log.Error("Unable to form Order Request")
	}
	var responseMap map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&responseMap)
	if err != nil {
		log.Warningf("Error Decoding Json, %v", err)
	}
	log.Printf("Order Response: %+v", responseMap)
	return nil
}

func (o *orderClient) SendRecommendation() {
	str := `Based on your order, we recommend you the following products:
	 - Monopod 100 - Floor Standing Monopod
	 - ACK-E5 AC Adapter Kit
	 - BG-E5 Battery Grip
	 To Buy these items click <https://demo.cf.test-4.cf-kyma-k8.shoot.canary.k8s-hana.ondemand.com/createCart?user=anish.joshi@sap.com&products=1099285&products=1422222&products=1422224|here>.`
	client := slack_client.NewSlackClient()
 	client.SendMessage(str)
}

func (o *orderClient) CreateOrder() error {
	return nil
}

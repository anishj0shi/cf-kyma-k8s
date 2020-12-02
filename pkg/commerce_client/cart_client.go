package commerce_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anishj0shi/cf-kyma-k8s/pkg"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

type CartClient interface {
	CreateShoppingCart()
	AddProductsToCart(productIds ...string)
}

type cartClient struct {
	userid     string
	cartId     *string
	gatewayUrl string
}

func NewShoppingCartClient(userid string) CartClient {
	return &cartClient{
		userid:     userid,
		cartId:     nil,
		gatewayUrl: os.Getenv("GATEWAY_URL"),
	}
}

func (c *cartClient) CreateShoppingCart() {
	if c.cartId == nil {
		url := fmt.Sprintf("%s/electronics/users/%s/carts", c.gatewayUrl, c.userid)
		res, err := http.Post(url, "application/json", nil)
		if err != nil {
			log.Warnf("Error in Creating CArt, err: %+v", err)
		}
		var responseMap map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&responseMap)
		if err != nil {
			log.Warningf("error received: %+v", err)
		}
		log.Printf("response received: %v", responseMap)

		str := responseMap["code"].(string)
		c.cartId = &str
		log.Printf("Cart ID : %s", *c.cartId)
	}
}

func (c *cartClient) AddProductsToCart(productIds ...string) {
	url := fmt.Sprintf("%s/electronics/users/%s/carts/%s/entries", c.gatewayUrl, c.userid, *c.cartId)
	log.Printf("CArt Url: %s", url)
	for id := range Recommendation {
		cartEntry := pkg.CartEntry{
			Quantity: 1,
			Product: map[string]string{
				"code": id,
			},
		}
		req, err := json.Marshal(cartEntry)
		if err != nil {
			log.Warningf("Error in json marshalling of cart request. err: %+v", err)
		}
		res, err := http.Post(url,"applicaion/json", bytes.NewReader(req))
		if err != nil {
			log.Warningf("Unable to add product to the Cart. err: %+v", err)
		}
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		responseString := string(responseData)
		log.Printf("Cart Response: %s", responseString)

	}
}

package items

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/dg185200/qrappstore/internal/handler"
	"github.com/dg185200/qrappstore/internal/httperror"
)

const nepOrganization = "nep-organization"

var (
	menu = map[string][]*item{
		"Our Specials": {},
		"Appetizers":   {},
		"Salads":       {},
		"Entrees":      {},
		"Desserts":     {},
	}
	categories = []string{
		"Our Specials",
		"Appetizers",
		"Salads",
		"Entrees",
		"Desserts",
	}
	getRandomKey = func(keys []string) string {
		i := rand.Intn(len(keys))
		return keys[i]
	}
)

type itemsHandler struct {
	username string
	password string
	c        *http.Client
}

func NewHandler(username, password string) handler.RequestHandler {
	return &itemsHandler{
		username: username,
		password: password,
		c: &http.Client{
			Timeout:   30 * time.Second,
			Transport: http.DefaultTransport,
		}}
}

func (h *itemsHandler) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	organization := r.Header.Get(nepOrganization)
	if organization == "" {
		return httperror.StatusError{Code: 400, Err: errors.New("items: nep-organization header is required")}
	}
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	url := fmt.Sprintf("https://gateway-staging.ncrcloud.com/catalog/item-details?pageSize=10")
	r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return httperror.StatusError{Code: 500, Err: err}
	}
	r.SetBasicAuth(h.username, h.password)
	r.Header.Add("nep-organization", "red-robin-qa")
	r.Header.Add("nep-enterprise-unit", "7ec5fed5d44f4c91a1c885f3968b755c")
	r.Header.Add("nep-service-version", "2:2")

	// b, _ := httputil.DumpRequestOut(r, true)
	// log.Printf("%s\n", b)

	resp, err := h.c.Do(r)
	if err != nil {
		return httperror.StatusError{Code: 502, Err: err}
	}
	defer resp.Body.Close()

	// b, _ = httputil.DumpResponse(resp, true)
	// log.Printf("%s\n", b)

	var respPayload struct {
		PageContent []*catalogItem `json:"pageContent,omitempty"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respPayload); err != nil {
		return httperror.StatusError{Code: 500, Err: err}
	}
	finalItems := []*item{}
	for _, catalogItem := range respPayload.PageContent {
		name := ""
		price := 0.0
		for _, values := range catalogItem.ShortDescription.Values {
			name = values.Value
		}
		for _, itemPrice := range catalogItem.ItemPrices {
			price = itemPrice.Price
		}
		finalItems = append(finalItems, &item{name, price})
	}
	itemsResponse := &items{Items: finalItems}
	finalResponse := []struct {
		Category string
		Items    []*item
	}{}
	for k, v := range buildItemCategories(itemsResponse) {
		finalResponse = append(finalResponse, struct {
			Category string
			Items    []*item
		}{
			Category: k,
			Items:    v,
		})
	}
	if err := json.NewEncoder(w).Encode(&finalResponse); err != nil {
		return httperror.StatusError{Code: 500, Err: err}
	}
	return nil
}

func buildItemCategories(items *items) map[string][]*item {
	for _, item := range items.Items {
		c := getRandomKey(categories)
		menu[c] = append(menu[c], item)
	}
	return menu
}

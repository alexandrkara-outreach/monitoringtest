package http

import (
	"encoding/json"
	"net/http"

	"github.com/alexandrkara-outreach/monitoringtest/internal/service"

	"log"
)

type Controler struct {
	service *service.Heavy
}

func NewController(service *service.Heavy) *Controler {
	return &Controler{
		service: service,
	}
}

func (c *Controler) SuperEndpoint(w http.ResponseWriter, r *http.Request) {

	n := c.service.Compute()

	json.NewEncoder(w).Encode(map[string]service.Result{"computed": n})
	log.Println("request served")
}

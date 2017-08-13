package farmer

import (
	"encoding/json"
	"net/http"
)

type Farmer struct {
	FarmerCode   string
	FarmerName   string
	GLNNumber    string
	SerialNumber string
	PlotNumber   string
}

type FarmerAPI struct {
	db *db.DB
}

func (api *FarmerAPI) All(w http.ResponseWriter, r *http.Request) {
	var farmers []Farmer
	api.db.All("Farmer", &farmers)
	json.NewDecoder(w).Decode(farmers)
}

func (api *FarmerAPI) Register(mux *http.ServeMux) {
	mux.HandleFunc("/list", api.All)
}

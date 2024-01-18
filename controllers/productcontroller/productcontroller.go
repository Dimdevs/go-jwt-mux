package productcontroller

import (
	"net/http"

	"github.com/jeypac/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "TAS MAHAL",
			"stock":        1000,
		},
		{
			"id":           2,
			"nama_product": "TAS MURAH",
			"stock":        100,
		},
		{
			"id":           3,
			"nama_product": "CELANA",
			"stock":        50,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}

package photocontroller

import (
	"net/http"

	"task-5-vix-btpns/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"Title": 		"Fotoku",
			"PhotoUrl":     "https://www.google.com/",
			"caption":      "fotoku",
		},
		{
			"id":           2,
			"Title": 		"Danang Wahyu Hermawan",
			"PhotoUrl":     "https://www.google.com/",
			"caption":      "danang",
		},
	}

	helpers.ResponseJSON(w, http.StatusOK, data)

}

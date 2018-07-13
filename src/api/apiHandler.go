package api

import (
	"datecalculate"
	"encoding/json"
	"net/http"
)

type Duration struct {
	Days string `json:"resultDay"`
}

func ApiCalculateDate(responseWriter http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	startDay := queryString.Get("startDate")
	endDay := queryString.Get("endDate")

	startDateWithoutSlash := datecalculate.RemoveSlashFromStringDate(startDay)
	endDateWithoutSlash := datecalculate.RemoveSlashFromStringDate(endDay)

	durationResponse, _ := json.Marshal(datecalculate.MakeJson(startDateWithoutSlash, endDateWithoutSlash))
	responseWriter.Write(durationResponse)
}

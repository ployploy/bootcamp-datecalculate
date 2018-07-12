package api

import (
	"datecalculate"
	"encoding/json"
	"net/http"
	"strconv"
)

type Duration struct {
	Days string `json:"days"`
}

func ApiCalculateDate(responseWriter http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	startDay, _ := strconv.Atoi(queryString.Get("startDay"))
	startMonth, _ := strconv.Atoi(queryString.Get("startMonth"))
	startYear, _ := strconv.Atoi(queryString.Get("startYear"))
	endDay, _ := strconv.Atoi(queryString.Get("endDay"))
	endMonth, _ := strconv.Atoi(queryString.Get("endMonth"))
	endYear, _ := strconv.Atoi(queryString.Get("endYear"))
	startDate := datecalculate.NewDate(startDay, startMonth, startYear)
	endDate := datecalculate.NewDate(endDay, endMonth, endYear)

	durationResponse, _ := json.Marshal(datecalculate.MakeJson(startDate, endDate))
	responseWriter.Write(durationResponse)
}

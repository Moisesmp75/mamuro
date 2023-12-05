package services

import (
	"bytes"
	"encoding/json"
	"io"
	"mamuro_mail-api/api/dto"
	"mamuro_mail-api/api/models"
	"net/http"
)

var (
	url      = "http://localhost:4080/es/enron_mail/_search"
	username = "admin"
	password = "Complexpass#123"
)

func errors(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	response := dto.ErrorResponse(message)
	w.Write([]byte(dto.JsonResponse[*string](response)))
}

func validateRequest(w http.ResponseWriter, request *dto.Request, body []byte) bool {
	err := json.Unmarshal(body, &request)
	if err != nil {
		errors(w, 400, err.Error())
		return false
	}
	if err := dto.ValidateRequest(*request); err != "" {
		errors(w, 400, err)
		return false
	}
	return true
}

func SearchData(w http.ResponseWriter, r *http.Request) {
	bodyRequest := make([]byte, r.ContentLength)
	r.Body.Read(bodyRequest)

	var request dto.Request
	if !validateRequest(w, &request, bodyRequest) {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(dto.RequestToString(request))))
	if err != nil {
		errors(w, 500, err.Error())
		return
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errors(w, 500, err.Error())
		return
	}
	defer resp.Body.Close()

	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		errors(w, 400, err.Error())
		return
	}

	var response models.SearchDataResponse
	if err := models.ConvertToResponse(bodyResponse, &response); err != "" {
		errors(w, 400, err)
		return
	}
	data_result := response.Hits.Hits
	mails := models.MapResource(data_result, models.HitsToSource)
	response_mails := dto.NewResponse[[]models.Source](mails)
	w.Write([]byte(dto.JsonResponse[[]byte](response_mails)))
	// data_result := dto.NewResponse[models.SearchDataResponse](response)
	// w.Write([]byte(dto.JsonResponse[[]byte](data_result)))
}

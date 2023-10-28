package opaservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mhshahin/cool-service-go/config"
	"github.com/mhshahin/cool-service-go/model"
)

type OpaService struct {
	httpClient *http.Client
	cfg        *config.AppConfig
}

func NewOpaService(cfg *config.AppConfig) *OpaService {
	client := &http.Client{
		Timeout: cfg.OPA.Timeout,
	}

	return &OpaService{
		httpClient: client,
		cfg:        cfg,
	}
}

func (ops OpaService) Validate(payload model.OpaRequest) (bool, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	requestBody := bytes.NewBuffer(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, ops.cfg.OPA.URL, requestBody)
	if err != nil {
		return false, err
	}

	resp, err := ops.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("non ok status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	newResponse := model.OpaResponse{}
	err = json.Unmarshal(responseBody, &newResponse)
	if err != nil {
		return false, err
	}

	if !newResponse.Allowed {
		return false, nil
	}

	return true, nil
}

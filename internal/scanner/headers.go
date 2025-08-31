package scanner

import (
	"fmt"
	"net/http"
	"time"
	"websecscanner/internal/models"
)

// ScanHeaders faz a requisição e retorna os resultados dos checks
func ScanHeaders(url string) ([]models.Result, error) {
	resp, err := getResponse(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return scanHeadersFromResponse(resp), nil
}

// getResponse faz a requisição HTTP com timeout
func getResponse(url string) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao requisitar %s: %w", url, err)
	}

	return resp, nil
}

// scanHeadersFromResponse chama cada check e agrega os resultados
func scanHeadersFromResponse(resp *http.Response) []models.Result {
	return []models.Result{
		checkServerHeader(resp),
		checkXPoweredBy(resp),
		checkCSP(resp),
		checkXFrameOptions(resp),
		checkXContentTypeOptions(resp),
		checkReferrerPolicy(resp),
	}
}

// ------------------- CHECKS -------------------

func checkServerHeader(resp *http.Response) models.Result {
	server := resp.Header.Get("Server")
	if server != "" {
		return models.Result{
			Check:   "Server Header",
			Details: fmt.Sprintf("Exposto: %s", server),
			Status:  "warning",
		}
	}
	return models.Result{
		Check:   "Server Header",
		Details: "Não exposto",
		Status:  "OK",
	}
}

func checkXPoweredBy(resp *http.Response) models.Result {
	powered := resp.Header.Get("X-Powered-By")
	if powered != "" {
		return models.Result{
			Check:   "X-Powered-By Header",
			Details: fmt.Sprintf("Exposto: %s", powered),
			Status:  "warning",
		}
	}
	return models.Result{
		Check:   "X-Powered-By Header",
		Details: "Não exposto",
		Status:  "OK",
	}
}

func checkCSP(resp *http.Response) models.Result {
	csp := resp.Header.Get("Content-Security-Policy")
	if csp != "" {
		return models.Result{
			Check:   "CSP Header",
			Details: "Definido",
			Status:  "OK",
		}
	}
	return models.Result{
		Check:   "CSP Header",
		Details: "Não definido",
		Status:  "warning",
	}
}

func checkXFrameOptions(resp *http.Response) models.Result {
	xfo := resp.Header.Get("X-Frame-Options")
	if xfo != "" {
		return models.Result{
			Check:   "X-Frame-Options Header",
			Details: fmt.Sprintf("Definido: %s", xfo),
			Status:  "OK",
		}
	}
	return models.Result{
		Check:   "X-Frame-Options Header",
		Details: "Não definido",
		Status:  "warning",
	}
}

func checkXContentTypeOptions(resp *http.Response) models.Result {
	xcto := resp.Header.Get("X-Content-Type-Options")
	if xcto != "" {
		return models.Result{
			Check:   "X-Content-Type-Options Header",
			Details: fmt.Sprintf("Definido: %s", xcto),
			Status:  "OK",
		}
	}
	return models.Result{
		Check:   "X-Content-Type-Options Header",
		Details: "Não definido",
		Status:  "warning",
	}
}

func checkReferrerPolicy(resp *http.Response) models.Result {
	rp := resp.Header.Get("Referrer-Policy")
	if rp != "" {
		return models.Result{
			Check:   "Referrer-Policy Header",
			Details: fmt.Sprintf("Definido: %s", rp),
			Status:  "OK",
		}
	}
	return models.Result{
		Check:   "Referrer-Policy Header",
		Details: "Não definido",
		Status:  "warning",
	}
}

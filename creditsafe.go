package creditsafe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	m "github.com/sapient/creditsafe/models"
)

type CreditSafe struct {
	baseUrl  string
	username string
	password string
	token    string
	client   http.Client
	timeout  time.Duration
}

type Option func(*CreditSafe)

func New(options ...Option) *CreditSafe {
	cs := &CreditSafe{
		baseUrl:  "https://api.creditsafe.com/v1",
		username: "",
		password: "",
		timeout:  time.Second * 30,
	}

	for _, option := range options {
		option(cs)
	}

	cs.client = http.Client{
		Timeout: cs.timeout,
	}

	return cs
}

func SetSandbox(sandbox bool) Option {
	return func(cs *CreditSafe) {
		if sandbox {
			cs.baseUrl = "https://connect.sandbox.creditsafe.com/v1"
		} else {
			cs.baseUrl = "https://connect.creditsafe.com/v1"
		}
	}
}

func SetUsername(username string) Option {
	return func(cs *CreditSafe) {
		cs.username = username
	}
}

func SetPassword(password string) Option {
	return func(cs *CreditSafe) {
		cs.password = password
	}
}

func (cs *CreditSafe) Login() error {
	login := m.LoginRequest{
		Username: cs.username,
		Password: cs.password,
	}

	loginJson, _ := json.Marshal(login)
	req, err := http.NewRequest("POST", cs.baseUrl+"/authenticate", bytes.NewBuffer(loginJson))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	httpResponse, err := cs.client.Do(req)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()
	body, _ := io.ReadAll(httpResponse.Body)

	loginRes := m.LoginResponse{}
	json.Unmarshal(body, &loginRes)

	cs.token = loginRes.Token
	return nil
}

func (cs *CreditSafe) GetCompanyReport(companyId string) (m.CompanyReportResponse, error) {
	req, err := http.NewRequest("GET", cs.baseUrl+"/companies/"+companyId, nil)
	if err != nil {
		return m.CompanyReportResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+cs.token)

	res, err := cs.client.Do(req)
	if err != nil {
		return m.CompanyReportResponse{}, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	report := m.CompanyReportResponse{}
	json.Unmarshal(body, &report)

	return report, nil
}

package rest

import (
	"context"
	"gis-crawler/pkg/logging"
	"io"
	"net/http"
	"net/url"
	"time"
)

type GisClient struct {
	ctx     context.Context
	baseUrl *url.URL
	params  map[string]string

	req    *http.Request
	client *http.Client
}

var Host string

func NewGisClient(ctx context.Context, host string) *GisClient {
	Host = host
	return &GisClient{
		ctx:    ctx,
		params: map[string]string{},
	}
}

func (c *GisClient) CreateRequest(method string, path string) {
	var err error
	c.baseUrl, err = url.Parse(Host)
	if err != nil {
		logging.Log().Panicln(err)
	}

	c.baseUrl = c.baseUrl.JoinPath(path)

	//params
	params := c.baseUrl.Query()
	for k, v := range c.params {
		params.Set(k, v)
	}
	c.baseUrl.RawQuery = params.Encode() // Кодируем параметры обратно в строку запроса

	logging.Log().Debugln("parse from url: ", c.baseUrl.String())
	ctx2, _ := context.WithTimeout(c.ctx, time.Second*5)
	c.req, err = http.NewRequestWithContext(ctx2, method, c.baseUrl.String(), nil)
	if err != nil {
		logging.Get().Errorln(err)
	}
	c.req.Header.Add("Accept", "application/json")
	c.client = &http.Client{}
}

func (c *GisClient) AddParam(key string, value string) {
	c.params[key] = value
}

func (c *GisClient) Send() []byte {

	resp, err := c.client.Do(c.req)
	if err != nil {
		logging.Get().Errorln("error gis request", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logging.Get().Warnf("http code error: %d \n", resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Get().Errorln("ошибка при чтении тела ответа:", err)
		return nil
	}

	return body
}

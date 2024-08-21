package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gis-crawler/pkg/logging"
	"net/http"
	"time"
)

type Message struct {
	Message string `json:"message"`
}

func SendTgMessage(ctx context.Context, URL string, token string, message string) error {
	ctx2, _ := context.WithTimeout(ctx, time.Second*5)

	m := Message{
		Message: message,
	}

	jsonStr, err := json.Marshal(m)
	if err != nil {
		logging.Get().Errorln(err)
	}
	req, err := http.NewRequestWithContext(ctx2, "POST", URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		logging.Log().Errorln(err)
	}

	bearer := "Bearer " + token

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("error response is %d \n", resp.StatusCode))
	}

	return nil
}

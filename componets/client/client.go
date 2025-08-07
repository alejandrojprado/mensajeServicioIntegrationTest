package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mensajeServiceIntegrationTests/componets/models"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: time.Second * 30},
	}
}

func (c *Client) doRequest(method, path string, body interface{}, headers map[string]string) (*http.Response, error) {
	var req *http.Request
	var err error

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(jsonBody))
	} else {
		req, err = http.NewRequest(method, c.baseURL+path, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.httpClient.Do(req)
}

func (c *Client) CreateMessage(userID string, content string) (*models.Message, error) {
	request := models.RequestCreateMessage{Content: content}
	headers := map[string]string{"X-User-ID": userID}

	resp, err := c.doRequest("POST", "/messages", request, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d", resp.StatusCode)
	}

	var message models.Message
	err = json.NewDecoder(resp.Body).Decode(&message)
	return &message, err
}

func (c *Client) GetUserMessages(userID string) ([]models.Message, error) {
	headers := map[string]string{"X-User-ID": userID}

	resp, err := c.doRequest("GET", "/messages", nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d", resp.StatusCode)
	}

	var messages []models.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)
	return messages, err
}

func (c *Client) FollowUser(followerID string, followingID string) error {
	request := models.RequestFollow{FollowingID: followingID}
	headers := map[string]string{"X-User-ID": followerID}

	resp, err := c.doRequest("POST", "/follows", request, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("status: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) GetUserTimeline(userID string) ([]models.Message, error) {
	headers := map[string]string{"X-User-ID": userID}

	resp, err := c.doRequest("GET", "/timeline", nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d", resp.StatusCode)
	}

	var timeline []models.Message
	err = json.NewDecoder(resp.Body).Decode(&timeline)
	return timeline, err
}

package repository

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/manochatt/line-noti/domain"
)

func (lnr *lineNotifyRepository) SendNotify(c context.Context, payload *bytes.Buffer) error {
	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", domain.LineNotifyURL, payload)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer r5AGBIGYqxNudoJunkfLVZVQzGiZLeqnlgbKzHsHAHThZCDyQvfT/wsYWZxdDrddN8uD5Q1HkvKQknSbzONIXQAKRgGxJXCnomOq5Yvc3OE85L7xg4pgPvdrmZcFAi7yhmAU2MdklOzSfiJxYKEa5gdB04t89/1O/w1cDnyilFU=")

	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	log.Println("API request successful")
	return nil

}

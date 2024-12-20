package pkg

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/sling"
)

const AOC_INPUT_URL = "https://adventofcode.com/2024/day/%d/input"

type AocService struct {
	session string
}

func NewAocService(session string) *AocService {
	return &AocService{session: session}
}

func (s *AocService) GetInput(day int) string {
	req, err := sling.New().Get(fmt.Sprintf(AOC_INPUT_URL, day)).Request()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", s.session))

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)[:len(body)-1]
}

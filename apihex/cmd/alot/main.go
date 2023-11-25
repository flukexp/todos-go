package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const max = 20

var wg sync.WaitGroup

func main() {
	wg.Add(max)
	for id := 0; id < max; id++ {
		go request(id)
	}
	wg.Wait()
}

func request(id int) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		os.Getenv("TRANSFER_URL")+"/"+strconv.Itoa(id),
		nil,
	)
	if err != nil {
		slog.Error(err.Error(), slog.Int("id", id))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err.Error(), slog.Int("id", id))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error(err.Error(), slog.Int("id", id))
			return
		}
		slog.Error(string(b), slog.Int("id", id))
		return
	}

	type responseMessate struct {
		Message string `json:"message"`
	}

	var response responseMessate
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		slog.Error(err.Error(), slog.Int("id", id))
		return
	}

	slog.Info(response.Message, slog.Int("id", id))
}

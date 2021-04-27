package scripts

import (
	"context"
	"net/http"
	"time"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/helpers"
)

func CheckConnection() {
	i := 10
	d := time.Duration(i)
	ticker := time.NewTicker(d * time.Second)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			ok1 := testMetadataConn()
			ok2 := testCommandConn()
			ok3 := testDataConn()

			if ok1 && ok2 && ok3 {
				ticker.Stop()
				return
			}
		}
	}
}

func testMetadataConn() (ok bool) {
	ok = false

	// try get device from core metadata
	path := "/api/v1/device"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", helpers.CoreMetadataURL+path, nil)
	if err != nil {
		return ok
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ok
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		ok = true
	}

	return ok
}

func testCommandConn() (ok bool) {
	ok = false

	// try get device from core command
	path := "/api/v1/device"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", helpers.CoreCommandURL+path, nil)
	if err != nil {
		return ok
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ok
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		ok = true
	}

	return ok
}

func testDataConn() (ok bool) {
	ok = false

	// try get readings from core data
	path := "/api/v1/reading/device/Random-Temperature-Generator01/10"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", helpers.CoreDataURL+path, nil)
	if err != nil {
		return ok
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ok
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		ok = true
	}

	return ok
}

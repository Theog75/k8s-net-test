package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Res struct {
	Respo string `json:"Respo"`
	Stat  string `json:"Stat"`
}

func main() {
	var dest string
	var delayvar string
	var fork string
	var forkval int
	var delay time.Duration
	dest = os.Getenv("DEST")
	delayvar = os.Getenv("DELAYVAR")
	fork = os.Getenv("FORK")
	delay, _ = time.ParseDuration(delayvar)
	forkval, _ = strconv.Atoi(fork)
	print("DELAYVAR " + delayvar + " DEST: " + dest + " FORK: " + fork)
	forker(forkval, delay, dest)
}

func forker(f int, delay time.Duration, dest string) {
	for i := 0; i <= f; i++ {
		// fmt.Println("Forker")
		go worker(delay, dest)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func worker(delay time.Duration, dest string) {
	// fmt.Println("Worker")
	for {
		start := time.Now()
		GetRemoteStatus(dest)
		elapsed := time.Since(start)
		// log.Printf("%s", elapsed)
		fmt.Printf(" %s\n", elapsed)
		// fmt.Printf("Sleeping %s ", delay)
		time.Sleep(delay)
	}
}
func GetRemoteStatus(u string) ([]byte, int, error) {
	// url := fmt.Sprintf("%s/api/v2/cluster/issues", u)
	var res Res
	body, status, err := httpGetCall(u)
	if err != nil || status != http.StatusOK {
		return body, status, err
	}

	if err = json.Unmarshal(body, &res); err != nil {

	}

	// print(res.Respo)
	return body, status, nil
}

func httpGetCall(url string) ([]byte, int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, http.StatusInternalServerError, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, resp.StatusCode, err
		}
		return body, resp.StatusCode, nil
	}
	//Get response code only if response isn't nil.
	status := http.StatusInternalServerError
	if resp != nil {
		status = resp.StatusCode
	}
	return []byte{}, status, err
}

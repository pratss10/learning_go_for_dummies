package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type Result struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
}

func main() {
	url := os.Getenv("FOREX_API_KEY")
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result Result
			json.Unmarshal([]byte(body), &result)
			keys := make([]string, 0, len(result.Rates))
			for k := range result.Rates {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Println(k, result.Rates[k])
			}
			// for i, v := range result.Rates {
			// 	fmt.Println(i, v)
			// }
		} else {
			var err Error
			json.Unmarshal([]byte(body), &err)
			fmt.Println(err.Error.Info)
		}
	} else {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pisensor/pkg/models"
)

const (
	//addr   = "http://198.199.87.130:5555/"
	addr   = "http://localhost:5555/"
	serial = "123"
	model  = "DF-A1"
)

func main() {
	client := &http.Client{}

	for {
		time.Sleep(time.Second)
		num := models.SensorMessage{
			Num: float64(rand.Intn(100)) + rand.Float64(),
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(num)
		req, err := http.NewRequest("POST", addr+serial+"/"+model, b)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		if resp != nil {
			resp.Body.Close()
		}
		req.Close = true
	}
}

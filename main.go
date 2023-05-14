package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	port := ":8080"
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Allow-Methods", "*")
		if context.Request.Method == "OPTIONS" {
			context.Status(200)
			context.Abort()
		}
	})

	router.POST("/pfb", sendPfb)
	fmt.Println("Server is running")

	err := router.Run(port)

	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
}

type PFB struct {
	NamespaceID string `json:"namespace_id"`
	Data        string `json:"data"`
	GasLimit    int    `json:"gas_limit"`
	Fee         int    `json:"fee"`
	IPAddress   string `json:"ip_address"`
}

type ResponseData struct {
	Height    int64  `json:"height"`
	TxHash    string `json:"txhash"`
	Data      string `json:"data"`
	RawLog    string `json:"raw_log"`
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Events    []struct {
		Type       string `json:"type"`
		Attributes []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			Index bool   `json:"index"`
		} `json:"attributes"`
	} `json:"events"`
}

func sendPfb(c *gin.Context) {
	pfbBody := PFB{}
	rawInput, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(406, "Invalid input format")
		return
	}
	if err := json.Unmarshal(rawInput, &pfbBody); err != nil {
		c.AbortWithStatusJSON(400, "Struct mismatch")
		return
	}

	pfb := PFB{
		NamespaceID: pfbBody.NamespaceID,
		Data:        pfbBody.Data,
		GasLimit:    pfbBody.GasLimit,
		Fee:         pfbBody.Fee,
	}

	jsonData, err := json.Marshal(pfb)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	url := fmt.Sprintf("http://%s:26659/submit_pfb", pfbBody.IPAddress)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making the request:", err)
		c.AbortWithStatusJSON(400, "Make sure you have enough Tia use our tool to check it out")
		return
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	responseData := ResponseData{}
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("Error parsing the JSON response:", err)
		c.AbortWithStatusJSON(400, "Make sure you have enough Tia use our tool to check it out")
		return
	}

	c.JSON(200, responseData)
}

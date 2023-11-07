package main

import (
    "fmt"
	"os"
	"encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
)

func dictionry(apiKey, word string) string {

	fmt.Println("---")
	fmt.Println(word)
    // fmt.Println("---")

	url := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word, apiKey)

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error making the HTTP request:", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        fmt.Println("Request failed with status code:", response.StatusCode)
    }

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading the response data:", err)
    }

    var dictionaryEntries []map[string]interface{}
    err = json.Unmarshal(data, &dictionaryEntries)
    if err != nil {
        fmt.Println("Error unmarshalling JSON data:", err)
    }

    shortDefs := make([]string, 0)

    for _, entry := range dictionaryEntries {
        shortdefs, ok := entry["shortdef"].([]interface{})
        if ok {
            for _, sd := range shortdefs {
                if shortdef, isString := sd.(string); isString {
                    shortDefs = append(shortDefs, shortdef)
                }
            }
        }
    }

    term := shortDefs[0]
    fmt.Println("---")
    fmt.Println(term)
    // fmt.Println("---")

    return term
}

func main() {

	if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go <your_argument>")
        os.Exit(1)
    }

    apiKey := os.Args[1]

	r := gin.Default()
	r.GET("/api/parameter/:param", func(c *gin.Context) {
		param := c.Param("param")

		result := dictionry(apiKey, param)
		
		c.JSON(http.StatusOK, gin.H{
			param: result,
		})
	})

    err := r.Run(":5600")
	if err != nil {
		fmt.Println(err)
	}

}

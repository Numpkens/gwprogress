package main
	
	import (
		"github.com/gin-gonic/gin"
		"net/http"
		"io"
		"os"
		"log"
	)

type boss struct {
	ID string `json:"id"`
	Name string `json:"name"`
	IsRaid bool `json:"israid"`
}

var bosses []boss

func main() {
//get and read data establish endpoint
	file, err := os.Open("bosses.json")
	if err != nil {
		log.Fatalf("Error: Could not find bosses.json: %v", err)
	}
	
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error: Could not read file data: %v" err)
	}

	err = json.Unmarshal(byteValue, &bosses)
	if err != nil {
		log.Fatalf("Error: JSON structure mismatch: %v", err)
	}
//Start server
		router := gin.Default()

		router.GET("/bosses", getBosses)
		log.Println("Server is starting on localhost:8080...")

		router.Run("localhost:8080")
	}

	func getBosses(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, bosses)
	}

	//Search Logic
	func getBossById(c *gin.Context) {
		id := c.Query("id")

		for _, b := range bosses {
			if b.ID == id {
				c.IndentedJSON(http.StatusOK, b)
				return
			}
		}
		//Fallback if no boss is found
		c.JSON(http.StautusNotFound, gin.H{"message": "Boss encounter not found"})
	}


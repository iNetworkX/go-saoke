package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	//buoc so 1: loc ket qua
	// Loc_ket_qua()

	//buoc so 2: kiem tra ket qua
	// fmt.Println(Kiem_tra_ket_qua("5213.45946"))

	// buoc so 3: trien khai len website
	r := gin.Default()

	r.LoadHTMLGlob("web/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/saoke", func(c *gin.Context) {
		// Handle form submission
		data := c.PostForm("data")
		data = strings.ReplaceAll(data, " ", "")
		result := Kiem_tra_ket_qua(data)
		c.String(http.StatusOK, result)
	})

	r.Run(":8080")
}

func Loc_ket_qua() {
	content, err := os.ReadFile("saoke.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	saoKe := strings.Split(string(content), "|")
	for i := 0; i < len(saoKe); i++ {
		saoKe[i] = strings.Replace(saoKe[i], "\n", "@", 3)
		saoKe[i] = strings.Replace(saoKe[i], "\n", " ", -1)
	}

	// Join the modified parts back together
	modifiedContent := strings.Join(saoKe, "|")

	// Write the modified content to a new file
	err = os.WriteFile("saoke_modified.txt", []byte(modifiedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}

	fmt.Println("File has been modified and saved as saoke_modified.txt")
}

func Kiem_tra_ket_qua(a string) string {
	var ketQua string
	content, err := os.ReadFile("saoke_modified.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	saoKe := strings.Split(string(content), "|")
	for i := 0; i < len(saoKe); i++ {
		if strings.Contains(saoKe[i], a) {
			saoKe[i] = strings.Replace(saoKe[i], "@", "<br>", 3)
			ketQua = saoKe[i]
		}
	}
	return ketQua
}

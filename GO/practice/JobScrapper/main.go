package main

import (
	"os"
	"strings"

	scrapper "./Scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, Hello")
	return c.File("home.html")
}

func handelScrape(c echo.Context) error {
	// 함수 실행후 file 삭제
	defer os.Remove(fileName)

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handelScrape)
	e.Logger.Fatal(e.Start(":1323"))

}

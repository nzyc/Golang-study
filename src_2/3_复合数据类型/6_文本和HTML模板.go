package main

import (
	"fmt"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func main() {
	fmt.Println("hello")
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

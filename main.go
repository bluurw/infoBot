package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

//42:40

var url string

func main(){
	url = "https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal"
	response := request(url)

	//fmt.Println(response)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(len(collectTagBlocks(response, "div")))

}

// Faz a requisicao
func request(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Error? Unable to make request"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body) // retorna resposta em binario
		if err != nil {
			return "Error? Unable to read response body"
		}
		return string(body)
	}
	return fmt.Sprintf("Error? Recived code %d", resp.StatusCode)
}

// Encontra as tags
func collectTagBlocks(html, tag string) []string {
    regexPattern := fmt.Sprintf(`<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag)
    re := regexp.MustCompile(regexPattern)
    matches := re.FindAllString(html, -1)

    return matches
}

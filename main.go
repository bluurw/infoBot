package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func main(){
	var url string

	url = "https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal"
	response := request(url)

	/*
	response := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Página principal</title>
    </head>
    <body>
		<div class="main-page-responsive-columns main-page-first-row" id="example">
    		Content inside the tag.
		</div>
        <h1>Bem-vindo</h1>
    </body>
    </html>
    `
	*/

	tag := "div"
	//blocks := findAll(response, tag)
	blocks := findSelect(response, tag, "class", "main-page-block-contents")

	for _, block := range blocks {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(block)
	}
}

// requisicao
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

// Extrai todas as tags
func findAll(html, tag string) []string {
	//regexPattern := fmt.Sprintf(`(?s)<%[1]s\b[^/>]*/>`, tag)
	regexPattern := fmt.Sprintf(`(?s)<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag) //	self-auto tags
    re := regexp.MustCompile(regexPattern)
	fmt.Println(re)
    matches := re.FindAllString(html, -1)
    return matches
}

// Extrai a tag selecionada
func findSelect(html, tag string, attribute string, value string) []string {
	regexPattern := fmt.Sprintf(`(?s)<%[1]s\b[^>]*%[2]s=["']%s["'][^>]*>(.*?)</%[1]s>`, tag, attribute, value)
    re := regexp.MustCompile(regexPattern)
    matches := re.FindAllString(html, -1)
    return matches
}

// Extrai o texto de cada tag
func extractTextTag(html string) string {
    re := regexp.MustCompile(`<[^>]*>`)
    text := strings.TrimSpace(re.ReplaceAllString(html, ""))
	return text
}
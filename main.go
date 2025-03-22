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

	tag := "div"
	blocks := collectAllTagBlocks(response, tag)

	for _, block := range blocks {
		fmt.Println(extractTextTag(block))
	}
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

// Extrai todas tags
func collectAllTagBlocks(html, tag string) []string {
    regexPattern := fmt.Sprintf(`<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag)
    re := regexp.MustCompile(regexPattern)
    matches := re.FindAllString(html, -1)

    return matches
}

// Extrai o texto de cada tag
func extractTextTag(html string) string {
    re := regexp.MustCompile(`<[^>]*>`)
    return strings.TrimSpace(re.ReplaceAllString(html, ""))
}


/*
func extractTextTagS(html string) []string {
	tags := []string{"div", "span"}
	texts := []string{}
	for {
		for _, tag := range(tags) {
			if strings.Contains(html, tag) {
				fmt.Println(html)
				regexPattern := fmt.Sprintf(`<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag)
				re := regexp.MustCompile(regexPattern)

				// Encontra todas as correspondências
				matches := re.FindAllStringSubmatch(html, -1)

				// Extrair apenas o conteúdo interno da tag
				for _, match := range matches {
					if len(match) > 1 {
						texts = append(texts, match[1]) // match[1] contém o conteúdo interno
					}
				}
			} else {
				break
			}
		}
	}
    return texts
}
*/
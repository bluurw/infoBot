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
	//attribute := "alt"
	//value := "Wikivoyage"
	//blocks := findAll(response, tag, &attribute, &value)
	blocks := findAll(response, tag, nil, nil)

	for _, block := range blocks {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		result := extractTextTag(block)
		result = clearBlockBlanck(result)
		fmt.Println(result)
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

// Extrai a tag selecionada
func findAll(html, tag string, attribute *string, value *string) []string {
    var tagList = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "source", "track", "wbr"}
    var regexPattern string

    isSelfClosing := false
    for _, iTag := range tagList {
        if tag == iTag {
            isSelfClosing = true
            break
        }
    }

    if isSelfClosing {
		if attribute != nil && value != nil {
            regexPattern = fmt.Sprintf(`(?s)<%[1]s\b[^>]*\b%[2]s=["']%s["'][^>]*\/>`, tag, *attribute, *value)
		} else {
			regexPattern = fmt.Sprintf(`(?s)<%[1]s\b[^>]*\/>`, tag)
		}
    } else {
        if attribute != nil && value != nil {
            regexPattern = fmt.Sprintf(`(?s)<%[1]s\b[^>]*%[2]s=["']%s["'][^>]*>(.*?)</%[1]s>`, tag, *attribute, *value)
        } else {
            regexPattern = fmt.Sprintf(`(?s)<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag)
        }
    }
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

// Substitui espacos em branco
func clearBlockBlanck(term string) string {
	re := regexp.MustCompile(`\r?\n`)
	result := re.ReplaceAllString(term, "")
	return result
}
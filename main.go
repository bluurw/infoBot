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
	//attribute := "class"
	//value := "main-page-block-contents"
	//blocks := findAll(response, tag, &attribute, &value)
	blocks := findAll(response, tag, nil, nil)

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

/*
// Extrai todas as tags
func findAll(html, tag string) []string {
	regexPattern := fmt.Sprintf(`(?s)<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag)
    re := regexp.MustCompile(regexPattern)
	fmt.Println(re)
    matches := re.FindAllString(html, -1)
    return matches
}
*/

// Extrai a tag selecionada
func findAll(html, tag string, attribute *string, value *string) []string {
    var regexPattern string

    // Check if both attribute and value are provided
    if attribute != nil && value != nil {
        // Dereference pointers to access their values
        regexPattern = fmt.Sprintf(`(?s)<%[1]s\b[^>]*%[2]s=["']%s["'][^>]*>(.*?)</%[1]s>`, tag, *attribute, *value)
    } else {
        // Default pattern for tags without attributes
        regexPattern = fmt.Sprintf(`(?s)<%[1]s\b[^>]*>(.*?)</%[1]s>`, tag)
    }

    // Compile the regex
    re := regexp.MustCompile(regexPattern)

    // Find all matches
    matches := re.FindAllString(html, -1)

    return matches
}


//regexPattern := fmt.Sprintf(`(?s)<%[1]s\b[^/>]*/>`, tag) //self-auto tags

// Extrai o texto de cada tag
func extractTextTag(html string) string {
    re := regexp.MustCompile(`<[^>]*>`)
    text := strings.TrimSpace(re.ReplaceAllString(html, ""))
	return text
}
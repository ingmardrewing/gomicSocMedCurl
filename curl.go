package curl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	shared "github.com/ingmardrewing/gomicSocMedShared"
)

const (
	JSON_FORMAT = `{"Link":"%s","ImgUrl":"%s","Title":"%s","TagsCsvString":"%s","Description":"%s"}`
	CURL_FORMAT = `curl -X POST -H "Content-Type application/json; charset=utf-8" -d '%s' -u %s %s`
)

type Content struct {
	shared.Content
}

func Command(title, description, link, imgUrl, tags string) string {
	json := createJson(title, description, link, imgUrl, tags)
	credentials := createCredentials()
	target := createTargetUrl()
	return createCurl(json, credentials, target)
}

func createJson(title, description, link, imgUrl, tags string) string {
	return fmt.Sprintf(JSON_FORMAT, link, imgUrl, title, tags, description)
}

func createCredentials() string {
	user := shared.Env(shared.GOMIC_BASIC_AUTH_USER)
	pass := shared.Env(shared.GOMIC_BASIC_AUTH_PASS)
	return fmt.Sprintf("'%s:%s'", user, pass)
}

func createTargetUrl() string {
	basicUrl := shared.Env(shared.GOMIC_SOCMED_PROD_URL)
	port := shared.Env(shared.GOMIC_SOCMED_PROD_PORT)
	restVersion := shared.CURRENT_REST_VERSION
	restBasePath := shared.REST_BASE_PATH
	echoPath := shared.REST_PATH_ECHO

	return fmt.Sprintf("%s:%s/%s/%s%s", basicUrl, port, restVersion, restBasePath, echoPath)
}

func createCurl(json, credentials, target string) string {
	return fmt.Sprintf(CURL_FORMAT, json, credentials, target)
}

func askUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

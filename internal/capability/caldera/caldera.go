package caldera

import (
	"reflect"
	"soarca/logger"
	"soarca/models/cacao"
	"soarca/models/execution"
	"fmt"
	"net/http"
	"bytes"
	"io"
	"soarca/utils"
)

const calderaCapabilityName = "soarca-caldera"

type CalderaCapability struct {
	x int
}

type Empty struct{}

var component = reflect.TypeOf(Empty{}).PkgPath()
var log *logger.Log

func init() {
	log = logger.Logger(component, logger.Info, "", logger.Json)
}

func New() *CalderaCapability {
	return &CalderaCapability{x: 1}
}

func (CalderaCapability *CalderaCapability) GetType() string {
	return calderaCapabilityName
}

func (CalderaCapability *CalderaCapability) Execute(
	metadata execution.Metadata,
	command cacao.Command,
	authentication cacao.AuthenticationInformation,
	target cacao.AgentTarget,
	variables cacao.Variables) (cacao.Variables, error) {

		runAbility("80879fa1-765e-48c2-bcd3-6fa3b4b2e7a7", "djnsbg")

		return cacao.NewVariables(), nil

}

func runAbility(abilityId, agentId string) {
	calderaBaseUrl := utils.GetEnv("CALDERA_BASE_URL", "false")
	calderaApiToken := utils.GetEnv("CALDERA_API_TOKEN", "-")

    apiUrl := fmt.Sprintf("%s/plugin/access/exploit", calderaBaseUrl)
    userData := []byte(fmt.Sprintf(
		`{
		"paw": "%s",
		"ability_id": "%s",
		"facts": [
			{
			"trait": "file.sensitive.extension",
			"value": "xxx"
			}
		],
		"obfuscator": "base64"
		}`,
		agentId,
		abilityId,
	))

    request, error := http.NewRequest("POST", apiUrl, bytes.NewBuffer(userData))
    request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Cookie", fmt.Sprintf("API_SESSION=\"%s\";", calderaApiToken))

    client := &http.Client{}
    response, error := client.Do(request)

    if error != nil {
        fmt.Println(error)
    }

    responseBody, error := io.ReadAll(response.Body)

    if error != nil {
        fmt.Println(error)
    }

    fmt.Println("Status: ", response.Status)
    fmt.Println("Response body: ", responseBody)

    defer response.Body.Close()
	
}

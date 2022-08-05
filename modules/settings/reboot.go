package settings

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/digisynlink/debug-cli/utils"
	"github.com/urfave/cli/v2"
)

var endpoint = "/api/device/reboot"
var logger = utils.GetInstance()

type ReturnJson struct {
	Status string `json:"code"`
}

func RebootEntry(cCtx *cli.Context) error {
	host := cCtx.String("host")
	return reboot(host)
}

func reboot(host string) error {
	reqBody := make(map[string]interface{})
	b, _ := json.Marshal(reqBody)

	req, err := http.Post("http://"+host+endpoint, "application/json", bytes.NewBuffer(b))

	if err != nil {
		logger.Debug("Error while sending request")
		return err
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		logger.Debug("Error while reading response")
		return err
	}

	var rj ReturnJson
	err = json.Unmarshal(body, &rj)
	if err != nil {
		logger.Debug("Error while unmarshalling response")
		return err
	}

	if rj.Status != "succ" {
		logger.Debug("Error while rebooting")
		return err
	}

	return nil
}

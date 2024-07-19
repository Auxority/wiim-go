package wiim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func boolToInt(value bool) int {
	return map[bool]int{true: 1, false: 0}[value]
}

func (w *API) executeCommand(command string) error {
	resp, err := w.executeRequest(command)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (w *API) executeCommandJSON(command string, result interface{}) error {
	resp, err := w.executeRequest(command)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return failedToDecodeResponse(err)
	}
	return nil
}

func (w *API) executeRequest(command string) (*http.Response, error) {
	url := w.createCommandURL(command)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, failedToCreateRequest(err)
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, failedToExecuteRequest(err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, unexpectedStatusCode(resp.StatusCode)
	}

	return resp, nil
}

func (w *API) createCommandURL(command string) string {
	return w.baseURL + command
}

func (w *API) GetStatus() (*StatusResponse, error) {
	status := &StatusResponse{}

	err := w.executeCommandJSON("getPlayerStatus", status)
	if err != nil {
		return nil, failedToGetStatus(err)
	}

	return status, nil
}

func (w *API) Play() error {
	return w.executeCommand("setPlayerCmd:resume")
}

func (w *API) Pause() error {
	return w.executeCommand("setPlayerCmd:pause")
}

func (w *API) TogglePlay() error {
	return w.executeCommand("setPlayerCmd:onepause")
}

func (w *API) Stop() error {
	return w.executeCommand("setPlayerCmd:stop")
}

func (w *API) Next() error {
	return w.executeCommand("setPlayerCmd:next")
}

func (w *API) Previous() error {
	return w.executeCommand("setPlayerCmd:prev")
}

func (w *API) SetVolume(volume int) error {
	volume = limitValue(volume, minVolume, maxVolume)
	command := fmt.Sprintf("setPlayerCmd:vol:%d", volume)

	return w.executeCommand(command)
}

func (w *API) Mute() error {
	return w.executeCommand("setPlayerCmd:mute:1")
}

func (w *API) Unmute() error {
	return w.executeCommand("setPlayerCmd:mute:0")
}

func (w *API) ToggleMute() error {
	status, err := w.GetStatus()
	if err != nil {
		return failedToGetStatus(err)
	}

	oppositeState := boolToInt(!bool(status.IsMuted))
	command := fmt.Sprintf("setPlayerCmd:mute:%d", oppositeState)

	return w.executeCommand(command)
}

func (w *API) SetInput(input Input) error {
	command := fmt.Sprintf("setPlayerCmd:switchmode:%s", input)

	return w.executeCommand(command)
}

func (w *API) SelectPreset(preset int) error {
	preset = limitValue(preset, firstPreset, lastPreset)
	command := fmt.Sprintf("MCUKeyShortClick:%d", preset)

	return w.executeCommand(command)
}

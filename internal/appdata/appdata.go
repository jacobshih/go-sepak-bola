package appdata

import (
	"bytes"
	"go-sepak-bola/internal/dboxapi"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AppData struct {
	data []byte
	Api  dboxapi.DBoxAPI
}

type PostData struct {
	Category string      `json:"category"`
	Action   string      `json:"action,omitempty"`
	Params   interface{} `json:"params,omitempty"`
}

var appData AppData

func Instance() *AppData {
	return &appData
}

func (appdata *AppData) Data() []byte {
	return appdata.data
}

func (appdata *AppData) Get(path string) gjson.Result {
	return gjson.Get(string(appdata.data), path)
}

func (appdata *AppData) Set(path string, value interface{}) {
	appdata.data, _ = sjson.SetBytes(appdata.data, path, value)
}

func (appdata *AppData) Delete(path string) {
	appdata.data, _ = sjson.DeleteBytes(appdata.data, path)
}

func (appdata *AppData) String() string {
	return string(appdata.data)
}

func (appdata *AppData) Pull(path string) []byte {
	data := appdata.Api.Download(path)
	/*
	 * FIXME: check if data is validate json.
	 */
	if len(gjson.Get(string(data), "error.\\.tag").Str) == 0 {
		appdata.data = data
	}
	return appdata.data
}

func (appdata *AppData) Push(path string) {
	content := bytes.NewReader(appdata.data)
	appdata.Api.UploadContent(path, content)
	return
}

package dboxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	HTTP_POST string = "POST"
	HTTP_GET  string = "GET"
)

const (
	DOWNLOAD_FILE                    string = "https://content.dropboxapi.com/2/files/download"
	UPLOAD_FILE                      string = "https://content.dropboxapi.com/2/files/upload"
	LIST_FOLDER_GET_LATEST_CURSOR    string = "https://api.dropboxapi.com/2/files/list_folder/get_latest_cursor"
	LIST_FOLDER_CONTINUE             string = "https://api.dropboxapi.com/2/files/list_folder/continue"
	LIST_SHARED_LINKS                string = "https://api.dropboxapi.com/2/sharing/list_shared_links"
	CREATE_SHARED_LINK_WITH_SETTINGS string = "https://api.dropboxapi.com/2/sharing/create_shared_link_with_settings"
	REVOKE_SHARED_LINK               string = "https://api.dropboxapi.com/2/sharing/revoke_shared_link"
)

/*
 * DBoxAPI
 */
type DBoxAPI struct {
	token string
}

func (api *DBoxAPI) SetToken(token string) {
	api.token = token
}

func (api *DBoxAPI) GetToken() string {
	return api.token
}

/*
 * https://api.dropboxapi.com/2/files/list_folder/get_latest_cursor
 */
func (api *DBoxAPI) ListFolder_GetLatestCursor(path string) (cursor string) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	type Payload struct {
		Path                            string `json:"path"`
		Recursive                       bool   `json:"recursive"`
		IncludeMediaInfo                bool   `json:"include_media_info"`
		IncludeDeleted                  bool   `json:"include_deleted"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members"`
		IncludeMountedFolders           bool   `json:"include_mounted_folders"`
	}

	data := Payload{
		Path:                            path,
		Recursive:                       true,
		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return cursor
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(HTTP_POST, LIST_FOLDER_GET_LATEST_CURSOR, body)
	if err != nil {
		fmt.Println(err.Error())
		return cursor
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return cursor
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return cursor
		}
		var dict map[string]interface{}
		if err := json.Unmarshal(content, &dict); err != nil {
			panic(err)
		}
		cursor = dict["cursor"].(string)
	}
	return cursor
}

/*
 * https://api.dropboxapi.com/2/files/list_folder/continue
 */
func (api *DBoxAPI) ListFolder_Continue(cursor string) (dict map[string]interface{}) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	type Payload struct {
		Cursor string `json:"cursor"`
	}

	data := Payload{
		Cursor: cursor,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(HTTP_POST, LIST_FOLDER_CONTINUE, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		var dict map[string]interface{}
		if err := json.Unmarshal(content, &dict); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return dict
	} else {
		fmt.Println(resp.Status)
		return nil
	}
}

/*
 * https://api.dropboxapi.com/2/sharing/list_shared_links
 */
func (api *DBoxAPI) ListSharedLinks(cursor string) (links []interface{}) {
	type Payload struct {
		Cursor string `json:"cursor"`
	}

	var payloadBytes []byte
	var err error
	var has_more bool = true

	for has_more {
		if len(cursor) > 0 {
			data := Payload{
				Cursor: cursor,
			}
			payloadBytes, err = json.Marshal(data)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
		} else {
			payloadBytes = []byte("{}")
		}
		body := bytes.NewReader(payloadBytes)
		req, err := http.NewRequest(HTTP_POST, LIST_SHARED_LINKS, body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		req.Header.Set("Authorization", "Bearer "+api.token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			var dict map[string]interface{}
			if err := json.Unmarshal(content, &dict); err != nil {
				fmt.Println(err.Error())
				return nil
			}
			links = append(links, dict["links"].([]interface{})...)
			has_more = dict["has_more"].(bool)
			if has_more {
				/* FIXME:
				 * continue to list shared links if has_more is true.
				 */
				cursor = dict["cursor"].(string)
			}
		} else {
			has_more = false
		}
	}
	return links
}

/*
 * https://api.dropboxapi.com/2/sharing/create_shared_link_with_settings
 */
func (api *DBoxAPI) CreateSharedLinkWithSettings(path string) (newlink map[string]interface{}) {
	type Settings struct {
		RequestedVisibility string `json:"requested_visibility"`
	}

	type Payload struct {
		Path     string   `json:"path"`
		Settings Settings `json:"settings"`
	}

	settings := Settings{
		RequestedVisibility: "public",
	}
	data := Payload{
		Path:     path,
		Settings: settings,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(HTTP_POST, CREATE_SHARED_LINK_WITH_SETTINGS, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		if err := json.Unmarshal(content, &newlink); err != nil {
			fmt.Println(err.Error())
			return nil
		}
	} else {
		fmt.Println(resp.Status, path)
	}
	return newlink
}

/*
 * https://api.dropboxapi.com/2/sharing/revoke_shared_link
 */
func (api *DBoxAPI) RevokeSharedLink(url string) (ok bool) {
	type Payload struct {
		URL string `json:"url"`
	}

	data := Payload{
		URL: url,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(HTTP_POST, REVOKE_SHARED_LINK, body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer resp.Body.Close()
	return true
}

/*
 * https://content.dropboxapi.com/2/files/download
 */
func (api *DBoxAPI) Download(path string) []byte {

	req, err := http.NewRequest(HTTP_POST, DOWNLOAD_FILE, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Dropbox-Api-Arg", "{\"path\": \""+path+"\"}")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return content
}

/*
 * https://content.dropboxapi.com/2/files/upload
 */
func (api *DBoxAPI) UploadContent(path string, content io.Reader) {
	req, err := http.NewRequest(HTTP_POST, UPLOAD_FILE, content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	req.Header.Set("Dropbox-Api-Arg", "{\"path\": \""+path+"\",\"mode\": \"overwrite\",\"autorename\": true,\"mute\": false,\"strict_conflict\": false}")
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
}

func (api *DBoxAPI) Upload(from, to string) {
	f, err := os.Open(from)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	api.UploadContent(to, f)
}

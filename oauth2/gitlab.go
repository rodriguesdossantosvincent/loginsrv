package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/rodriguesdossantosvincent/loginsrv/model"
)

func init() {
	RegisterProvider(providerGitlab)
}

// GitlabUser is used for parsing the gitlab response
type GitlabUser struct {
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type GitlabGroup struct {
	FullPath string `json:"full_path,omitempty"`
}

var providerGitlab = Provider{
	Name:     "gitlab",
	GetUserInfo: func(token TokenInfo, InfoURL string) (model.UserInfo, string, error) {
		gu := GitlabUser{}
		url := fmt.Sprintf("%v/user?access_token=%v", InfoURL, token.AccessToken)

		var respUser *http.Response
		respUser, err := http.Get(url)
		if err != nil {
			return model.UserInfo{}, "", err
		}
		defer respUser.Body.Close()

		if !strings.Contains(respUser.Header.Get("Content-Type"), "application/json") {
			return model.UserInfo{}, "", fmt.Errorf("wrong content-type on gitlab get user info: %v", respUser.Header.Get("Content-Type"))
		}

		if respUser.StatusCode != 200 {
			return model.UserInfo{}, "", fmt.Errorf("got http status %v on gitlab get user info", respUser.StatusCode)
		}

		b, err := ioutil.ReadAll(respUser.Body)
		if err != nil {
			return model.UserInfo{}, "", fmt.Errorf("error reading gitlab get user info: %v", err)
		}

		err = json.Unmarshal(b, &gu)
		if err != nil {
			return model.UserInfo{}, "", fmt.Errorf("error parsing gitlab get user info: %v", err)
		}

		return model.UserInfo{
			Sub:     gu.Username,
			Picture: gu.AvatarURL,
			Name:    gu.Name,
			Email:   gu.Email,
			Origin:  "gitlab",
		}, `{"user":` + string(b) + `}`, nil
	},
}

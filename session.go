package devisesession

import (
	"net/http"
	"github.com/greasysock/railscookie"
	"encoding/json"
	"errors"
)

type UserSession struct {
	SessionID string `json:"session_id"`
	CsrfToken string `json:"_csrf_token"`
	SessionDetails []interface{} `json:"warden.user.user.key"`
	UserHash string
	UserID int
}

func CreateUserSession(cookie *http.Cookie, secret_key_base []byte, sql_connection SqlConnection) (user UserSession, err error) {
	var u UserSession
	c, uh_oh := railscookie.DecryptAndVerify(cookie, secret_key_base)
	if uh_oh != nil {return u, uh_oh}
	var msg UserSession
	uh_oh = json.Unmarshal(c, &msg)
	if uh_oh != nil{
		return u, uh_oh
	}
	if msg.SessionDetails == nil{return u, errors.New("Not signed in")}
	msg.UserID = int(msg.SessionDetails[0].([]interface{})[0].(float64))
	msg.UserHash = msg.SessionDetails[1].(string)
	return u, nil
}
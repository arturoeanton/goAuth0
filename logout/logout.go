// /routes/logout/logout.go
package logout

import (
	"net/http"
	"net/url"
	"time"

	"github.com/arturoeanton/goAuth0/app"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	domain := app.Config.Domain

	logoutUrl, err := url.Parse("https://" + domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	/*	var scheme string
		if r.TLS == nil {
			scheme = "http"
		} else {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + r.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}*/
	parameters.Add("returnTo", app.Config.Login)
	parameters.Add("client_id", app.Config.ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	c := &http.Cookie{
		Name:    "auth-session",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

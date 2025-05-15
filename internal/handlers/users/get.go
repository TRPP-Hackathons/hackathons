package users

import (
	"net/http"

	"hackathons/internal/services"
	httpErr "hackathons/pkg/http/error"
	"hackathons/pkg/http/writer"
)

func GetMe(uc services.Users) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		me, err := uc.GetUser(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, me)
	}
}

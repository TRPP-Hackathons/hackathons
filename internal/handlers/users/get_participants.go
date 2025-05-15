package users

import (
	"net/http"

	"hackathons/internal/services"
	httpErr "hackathons/pkg/http/error"
	"hackathons/pkg/http/writer"
)

func GetParticipants(uc services.Users) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		participants, err := uc.GetParticipants(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, participants)
	}
}

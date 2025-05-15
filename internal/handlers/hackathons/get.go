package hackathons

import (
	"net/http"

	"hackathons/internal/services"
	httpErr "hackathons/pkg/http/error"
	"hackathons/pkg/http/writer"
)

func GetHackathons(uc services.Hackathons) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hackathons, err := uc.GetHackathons(r.Context())
		if err != nil {
			httpErr.InternalError(w, err)
			return
		}

		writer.WriteJson(w, hackathons)
	}
}

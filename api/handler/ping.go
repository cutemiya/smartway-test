package handler

import (
	"encoding/json"
	"net/http"
	"smartway-test/client"
	responseSender "smartway-test/lib/response"
)

// Ping
//
//	@Description	for test
//	@Tags			Base
//	@Success		200	{object}	client.Pong
//	@Router			/ping [get]
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal(client.Pong{
			Pong: "pong",
		})

		responseSender.SendResponse(w, http.StatusOK, res)
	}
}

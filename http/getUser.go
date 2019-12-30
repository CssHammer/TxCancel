package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"txCancel/models/wrapper"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.userQ.GetByID(1)
	if err != nil {
		logrus.WithError(err).Error("failed to get user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wUser := wrapper.WrapUser(user)

	bytes, err := json.Marshal(wUser)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		logrus.WithError(err).Error("failed to write response")
		return
	}
}

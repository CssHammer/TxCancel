package http

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"txCancel/models/db"
	"txCancel/models/transport"
	"txCancel/models/wrapper"
)

func (h *Handler) PostTx(w http.ResponseWriter, r *http.Request) {
	var tx transport.Transaction
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		logrus.WithError(err).Error("failed to decode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tx.State != db.StateWin && tx.State != db.StateLost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uwTx, err := wrapper.UnwrapTx(&tx)
	if err != nil {
		logrus.WithError(err).Error("failed to unwrap tx")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.compositeQ.ApplyTransaction(uwTx)
	if err != nil {
		logrus.WithError(err).Infof("failed to apply tx. tx id: %v", tx.TransactionID)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)

}

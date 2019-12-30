package http

import (
	"bytes"
	jsonPkg "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"

	mock_dbqueries "txCancel/dbqueries/mocks"
	"txCancel/models/transport"
)

var (
	payloadInvalidState = transport.Transaction{
		State:         "wrong",
		Amount:        "3.12",
		TransactionID: "id",
	}

	payloadInvalidAmount = transport.Transaction{
		State:         "win",
		Amount:        "string",
		TransactionID: "id",
	}

	payloadValid = transport.Transaction{
		State:         "win",
		Amount:        "3.14",
		TransactionID: "id",
	}
)

func TestPostTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Convey("Test PostTx handler", t, func() {
		recorder := httptest.NewRecorder()

		userQ := mock_dbqueries.NewMockUserQI(ctrl)
		compositeQ := mock_dbqueries.NewMockCompositeQI(ctrl)
		h := NewHandler(userQ, compositeQ)

		r := chi.NewRouter()
		r.Post("/tx", h.PostTx)

		Convey("Invalid payload", func() {
			req := httptest.NewRequest(
				http.MethodPost,
				"/tx",
				bytes.NewReader([]byte("123")),
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("Invalid State", func() {
			b, err := jsonPkg.Marshal(payloadInvalidState)
			So(err, ShouldBeNil)

			req := httptest.NewRequest(
				http.MethodPost,
				"/tx",
				bytes.NewReader(b),
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("Invalid Amount", func() {
			b, err := jsonPkg.Marshal(payloadInvalidAmount)
			So(err, ShouldBeNil)

			req := httptest.NewRequest(
				http.MethodPost,
				"/tx",
				bytes.NewReader(b),
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusBadRequest)
		})

		Convey("ApplyTransaction Error", func() {
			compositeQ.EXPECT().ApplyTransaction(gomock.Any()).Return(errors.New("error"))

			b, err := jsonPkg.Marshal(payloadValid)
			So(err, ShouldBeNil)

			req := httptest.NewRequest(
				http.MethodPost,
				"/tx",
				bytes.NewReader(b),
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusNotAcceptable)
		})

		Convey("Success", func() {
			compositeQ.EXPECT().ApplyTransaction(gomock.Any()).Return(nil)

			b, err := jsonPkg.Marshal(payloadValid)
			So(err, ShouldBeNil)

			req := httptest.NewRequest(
				http.MethodPost,
				"/tx",
				bytes.NewReader(b),
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusOK)
		})

	})
}

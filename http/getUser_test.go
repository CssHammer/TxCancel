package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"

	mock_dbqueries "txCancel/dbqueries/mocks"
	"txCancel/models/db"
)

var user = db.User{
	ID:      1,
	Name:    "Dave",
	Balance: 10,
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Convey("Test GetUser handler", t, func() {
		recorder := httptest.NewRecorder()

		userQ := mock_dbqueries.NewMockUserQI(ctrl)
		compositeQ := mock_dbqueries.NewMockCompositeQI(ctrl)
		h := NewHandler(userQ, compositeQ)

		r := chi.NewRouter()
		r.Get("/user", h.GetUser)

		Convey("GetByID Error", func() {
			userQ.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("error"))

			req := httptest.NewRequest(
				http.MethodGet,
				"/user",
				nil,
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusInternalServerError)
		})

		Convey("Success", func() {
			userQ.EXPECT().GetByID(gomock.Any()).Return(&user, nil)

			req := httptest.NewRequest(
				http.MethodGet,
				"/user",
				nil,
			)

			r.ServeHTTP(recorder, req)

			So(recorder.Code, ShouldEqual, http.StatusOK)
		})

	})
}

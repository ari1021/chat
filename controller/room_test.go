package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ari1021/websocket/model/mock_model"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoomHandler_CreateRoom(t *testing.T) {
	e := echo.New()
	f := make(url.Values)
	f.Set("name", "test")
	req := httptest.NewRequest(http.MethodPost, "/rooms", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// mockの準備
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roomMock := mock_model.NewMockIRoom(ctrl)
	roomMock.EXPECT().Create(nil).Return(1)
	rh := RoomHandler{
		IRoom: roomMock,
	}
	if assert.NoError(t, rh.CreateRoom(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

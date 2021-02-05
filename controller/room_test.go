package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/ari1021/websocket/model"
	"github.com/ari1021/websocket/model/mock_model"
	"github.com/ari1021/websocket/server/validation"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRoomHandler_CreateRoom(t *testing.T) {
	expected := &model.Room{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: "test",
	}
	e := echo.New()
	e = validation.ValidateEcho(e)
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
	roomMock.EXPECT().Create("test").Return(expected, nil)
	rh := RoomHandler{
		IRoom: roomMock,
	}

	err := rh.CreateRoom(c)
	// error確認
	assert.NoError(t, err)
	// statusCode確認
	assert.Equal(t, http.StatusOK, rec.Code)
	// response確認
	got := &model.Room{}
	if err := json.Unmarshal(rec.Body.Bytes(), got); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, expected, got)
}

func TestRoomHandler_GetRooms(t *testing.T) {
	expected := &model.Rooms{
		model.Room{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Name: "test",
		},
		model.Room{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Name: "test",
		},
	}
	e := echo.New()
	e = validation.ValidateEcho(e)
	req := httptest.NewRequest(http.MethodGet, "/rooms", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// mockの準備
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roomMock := mock_model.NewMockIRoom(ctrl)
	roomMock.EXPECT().FindAll().Return(expected, nil)
	rh := RoomHandler{
		IRoom: roomMock,
	}

	err := rh.GetRooms(c)
	// error確認
	assert.NoError(t, err)
	// statusCode確認
	assert.Equal(t, http.StatusOK, rec.Code)
	// response確認
	got := &model.Rooms{}
	if err := json.Unmarshal(rec.Body.Bytes(), got); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, expected, got)
}

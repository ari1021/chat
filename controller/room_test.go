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
	"github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRoomHandler_CreateRoom(t *testing.T) {
	tests := []struct {
		title           string
		name            string
		wantMock        bool
		prepareRoomMock func(*mock_model.MockIRoom)
		want            *model.Room
		wantErr         bool
		wantCode        int
	}{
		{
			title:    "正しくルームを作成することができる",
			name:     "test",
			wantMock: true,
			prepareRoomMock: func(rm *mock_model.MockIRoom) {
				rm.EXPECT().Create("test").Return(&model.Room{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
						DeletedAt: gorm.DeletedAt{},
					},
					Name: "test",
				}, nil)
			},
			want: &model.Room{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{},
				},
				Name: "test",
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			title:           "validationに失敗したときはStatusBadRequest",
			name:            "",
			wantMock:        false,
			prepareRoomMock: nil,
			want:            nil,
			wantErr:         true,
			wantCode:        http.StatusBadRequest,
		},
		{
			title:    "DuplicateErrorでルーム作成に失敗したときはStatusBadRequest",
			name:     "test",
			wantMock: true,
			prepareRoomMock: func(rm *mock_model.MockIRoom) {
				rm.EXPECT().Create("test").Return(nil, &mysql.MySQLError{
					Number:  1452,
					Message: "Duplicate key error",
				})
			},
			want:     nil,
			wantErr:  true,
			wantCode: http.StatusBadRequest,
		},
		{
			title:    "DatabaseErrorでルーム作成に失敗したときはStatusInternalServerError",
			name:     "test",
			wantMock: true,
			prepareRoomMock: func(rm *mock_model.MockIRoom) {
				rm.EXPECT().Create("test").Return(nil, &mysql.MySQLError{
					Number:  1,
					Message: "Database error",
				})
			},
			want:     nil,
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
		{
			title:    "UnknownErrorでルーム作成に失敗したときはStatusInternalServerError",
			name:     "test",
			wantMock: true,
			prepareRoomMock: func(rm *mock_model.MockIRoom) {
				rm.EXPECT().Create("test").Return(nil, model.UnknownError)
			},
			want:     nil,
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			e := echo.New()
			// validatorの設定
			e = validation.ValidateEcho(e)
			// formの設定
			f := make(url.Values)
			f.Set("name", tt.name)
			req := httptest.NewRequest(http.MethodPost, "/rooms", strings.NewReader(f.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			var rh RoomHandler
			if tt.wantMock {
				// mockの準備
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				rm := mock_model.NewMockIRoom(ctrl)
				// return値の設定
				tt.prepareRoomMock(rm)
				rh = RoomHandler{
					IRoom: rm,
				}
			} else {
				rh = RoomHandler{nil}
			}

			// エラーチェック
			if err := rh.CreateRoom(c); err != nil {
				t.Errorf("CreateRoom() err = %v, want = %v", err, nil)
			}
			// ステータスコードのチェック
			if rec.Code != tt.wantCode {
				t.Errorf("CreateRoom() code = %d, want = %d", rec.Code, tt.wantCode)
			}
			// 返り値の中身チャック
			if !tt.wantErr {
				got := &model.Room{}
				if err := json.Unmarshal(rec.Body.Bytes(), got); err != nil {
					log.Fatal(err)
				}
				if !cmp.Equal(got, tt.want) {
					t.Errorf("CreateRoom() diff = %v", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}

func TestRoomHandler_GetRooms(t *testing.T) {
	tests := []struct {
		title           string
		name            string
		prepareRoomMock func(*mock_model.MockIRoom)
		want            *model.Rooms
		wantErr         bool
		wantCode        int
	}{
		{
			title: "正しくルームを取得することができる",
			name:  "test",
			prepareRoomMock: func(rm *mock_model.MockIRoom) {
				rm.EXPECT().FindAll().Return(&model.Rooms{
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
				}, nil)
			},
			want: &model.Rooms{
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
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			title: "UnknownErrorでルーム取得に失敗したときはStatusInternalServerError",
			name:  "test",
			prepareRoomMock: func(rm *mock_model.MockIRoom) {
				rm.EXPECT().FindAll().Return(nil, model.UnknownError)
			},
			want:     nil,
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			e := echo.New()
			e = validation.ValidateEcho(e)
			req := httptest.NewRequest(http.MethodGet, "/rooms", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			// mockの準備
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			rm := mock_model.NewMockIRoom(ctrl)
			tt.prepareRoomMock(rm)
			rh := RoomHandler{
				IRoom: rm,
			}

			// エラーチェック
			if err := rh.GetRooms(c); err != nil {
				t.Errorf("GetRooms() err = %v, want = %v", err, nil)
			}
			// ステータスコードのチェック
			if rec.Code != tt.wantCode {
				t.Errorf("GetRooms() code = %d, want = %d", rec.Code, tt.wantCode)
			}
			// 返り値の中身チャック
			if !tt.wantErr {
				got := &model.Rooms{}
				if err := json.Unmarshal(rec.Body.Bytes(), got); err != nil {
					log.Fatal(err)
				}
				if !cmp.Equal(got, tt.want) {
					t.Errorf("GetRooms() diff = %v", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
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
	req := httptest.NewRequest(http.MethodDelete, "/rooms/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/rooms/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// mockの準備
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	roomMock := mock_model.NewMockIRoom(ctrl)
	roomMock.EXPECT().Delete(uint(1)).Return(expected, nil)
	rh := RoomHandler{
		IRoom: roomMock,
	}

	err := rh.DeleteRoom(c)
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

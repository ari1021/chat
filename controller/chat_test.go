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
)

func TestChatHandler_CreateChat(t *testing.T) {
	tests := []struct {
		title           string
		id              string
		message         string
		userName        string
		wantMock        bool
		prepareChatMock func(*mock_model.MockIChat)
		want            *model.Chat
		wantErr         bool
		wantCode        int
	}{
		{
			title:    "正しくメッセージを作成できる",
			id:       "1",
			message:  "test",
			userName: "test",
			wantMock: true,
			prepareChatMock: func(cm *mock_model.MockIChat) {
				cm.EXPECT().Create("test", 1, "test").Return(&model.Chat{
					ID:        1,
					CreatedAt: time.Time{},
					RoomID:    1,
					Room:      model.Room{},
					Message:   "test",
					UserName:  "test",
				}, nil)
			},
			want: &model.Chat{
				ID:        1,
				CreatedAt: time.Time{},
				RoomID:    1,
				Room:      model.Room{},
				Message:   "test",
				UserName:  "test",
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			title:           "messageのvalidationに失敗したときはStatusBadRequest",
			id:              "",
			message:         "",
			userName:        "test",
			wantMock:        false,
			prepareChatMock: nil,
			want:            nil,
			wantErr:         true,
			wantCode:        http.StatusBadRequest,
		},
		{
			title:           "userNameのvalidationに失敗したときはStatusBadRequest",
			id:              "",
			message:         "test",
			userName:        "",
			wantMock:        false,
			prepareChatMock: nil,
			want:            nil,
			wantErr:         true,
			wantCode:        http.StatusBadRequest,
		},
		{
			title:    "DatabaseErrorでメッセージ作成に失敗したときはStatusInternalServerError",
			id:       "1",
			message:  "test",
			userName: "test",
			wantMock: true,
			prepareChatMock: func(cm *mock_model.MockIChat) {
				cm.EXPECT().Create("test", 1, "test").Return(nil, &mysql.MySQLError{
					Number:  1,
					Message: "Database error",
				})
			},
			want:     nil,
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
		{
			title:    "UnknownErrorでメッセージ作成に失敗したときはStatusInternalServerError",
			id:       "1",
			message:  "test",
			userName: "test",
			wantMock: true,
			prepareChatMock: func(cm *mock_model.MockIChat) {
				cm.EXPECT().Create("test", 1, "test").Return(nil, model.UnknownError)
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
			f.Set("message", tt.message)
			f.Set("user_name", tt.userName)
			req := httptest.NewRequest(http.MethodPost, "/rooms/:id/chats", strings.NewReader(f.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			// pathparamの設定
			c.SetPath("/rooms/:id/chats")
			c.SetParamNames("id")
			c.SetParamValues(tt.id)

			var ch ChatHandler
			if tt.wantMock {
				// mockの準備
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				cm := mock_model.NewMockIChat(ctrl)
				// return値の設定
				tt.prepareChatMock(cm)
				ch = ChatHandler{
					IChat: cm,
				}
			} else {
				ch = ChatHandler{nil}
			}

			// エラーチェック
			if err := ch.CreateChat(c); err != nil {
				t.Errorf("CreateChat() err = %v, want = %v", err, nil)
			}
			// ステータスコードのチェック
			if rec.Code != tt.wantCode {
				t.Errorf("CreateChat() code = %d, want = %d", rec.Code, tt.wantCode)
			}
			// 返り値の中身チャック
			if !tt.wantErr {
				got := &model.Chat{}
				if err := json.Unmarshal(rec.Body.Bytes(), got); err != nil {
					log.Fatal(err)
				}
				if !cmp.Equal(got, tt.want) {
					t.Errorf("CreateChat() diff = %v", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}

func TestChatHandler_GetChats(t *testing.T) {
	tests := []struct {
		title           string
		id              string
		limit           string
		offset          string
		wantMock        bool
		prepareChatMock func(*mock_model.MockIChat)
		want            *model.Chats
		wantErr         bool
		wantCode        int
	}{
		{
			title:    "正しくメッセージを取得できる",
			id:       "1",
			limit:    "2",
			offset:   "0",
			wantMock: true,
			prepareChatMock: func(cm *mock_model.MockIChat) {
				cm.EXPECT().Find(1, 2, 0).Return(&model.Chats{
					model.Chat{
						ID:        1,
						CreatedAt: time.Time{},
						RoomID:    1,
						Room:      model.Room{},
						Message:   "test",
						UserName:  "test",
					},
					model.Chat{
						ID:        2,
						CreatedAt: time.Time{},
						RoomID:    1,
						Room:      model.Room{},
						Message:   "test",
						UserName:  "test",
					},
				}, nil)
			},
			want: &model.Chats{
				model.Chat{
					ID:        1,
					CreatedAt: time.Time{},
					RoomID:    1,
					Room:      model.Room{},
					Message:   "test",
					UserName:  "test",
				},
				model.Chat{
					ID:        2,
					CreatedAt: time.Time{},
					RoomID:    1,
					Room:      model.Room{},
					Message:   "test",
					UserName:  "test",
				},
			},
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			title:           "limitのvalidationに失敗したときはStatusBadRequest",
			id:              "1",
			limit:           "",
			offset:          "0",
			wantMock:        false,
			prepareChatMock: nil,
			want:            nil,
			wantErr:         true,
			wantCode:        http.StatusBadRequest,
		},
		{
			title:    "DatabaseErrorでメッセージ取得に失敗したときはStatusInternalServerError",
			id:       "1",
			limit:    "2",
			offset:   "0",
			wantMock: true,
			prepareChatMock: func(cm *mock_model.MockIChat) {
				cm.EXPECT().Find(1, 2, 0).Return(nil, &mysql.MySQLError{
					Number:  1,
					Message: "database error",
				})
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
			// queryparamの設定
			q := make(url.Values)
			q.Set("limit", tt.limit)
			q.Set("offset", tt.offset)
			req := httptest.NewRequest(http.MethodPost, "/rooms/:id/chats?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			// pathparamの設定
			c.SetPath("/rooms/:id/chats")
			c.SetParamNames("id")
			c.SetParamValues(tt.id)

			var ch ChatHandler
			if tt.wantMock {
				// mockの準備
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				cm := mock_model.NewMockIChat(ctrl)
				// return値の設定
				tt.prepareChatMock(cm)
				ch = ChatHandler{
					IChat: cm,
				}
			} else {
				ch = ChatHandler{nil}
			}

			// エラーチェック
			if err := ch.GetChats(c); err != nil {
				t.Errorf("GetChats() err = %v, want = %v", err, nil)
			}
			// ステータスコードのチェック
			if rec.Code != tt.wantCode {
				t.Errorf("GetChats() code = %d, want = %d", rec.Code, tt.wantCode)
			}
			// 返り値の中身チャック
			if !tt.wantErr {
				got := &model.Chats{}
				if err := json.Unmarshal(rec.Body.Bytes(), got); err != nil {
					log.Fatal(err)
				}
				if !cmp.Equal(got, tt.want) {
					t.Errorf("GetChats() diff = %v", cmp.Diff(got, tt.want))
				}
			}
		})
	}
}

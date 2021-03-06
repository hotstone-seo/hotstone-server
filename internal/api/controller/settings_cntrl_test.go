package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hotstone-seo/hotstone-seo/internal/api/controller"
	"github.com/hotstone-seo/hotstone-seo/internal/api/repository"
	"github.com/hotstone-seo/hotstone-seo/internal/api/service_mock"
	"github.com/typical-go/typical-rest-server/pkg/echotest"
)

type (
	settingTestCase struct {
		testName string
		echotest.TestCase
		settingCntrlBuilder
	}
	settingCntrlBuilder struct {
		settingSvcFn func(*service_mock.MockSettingSvc)
	}
)

func (b *settingCntrlBuilder) build(mock *gomock.Controller) *controller.SettingCntrl {
	mockSvc := service_mock.NewMockSettingSvc(mock)
	if b.settingSvcFn != nil {
		b.settingSvcFn(mockSvc)
	}

	return &controller.SettingCntrl{
		SettingSvc: mockSvc,
	}
}

func TestSettingCntrl_Find(t *testing.T) {
	testcases := []settingTestCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
				},
				ExpectedCode: http.StatusOK,
				ExpectedBody: "[{\"key\":\"key-1\",\"value\":\"value-1\",\"updated_at\":\"0001-01-01T00:00:00Z\"},{\"key\":\"key-2\",\"value\":\"value-2\",\"updated_at\":\"0001-01-01T00:00:00Z\"}]\n",
			},
			settingCntrlBuilder: settingCntrlBuilder{
				settingSvcFn: func(svc *service_mock.MockSettingSvc) {
					svc.EXPECT().Find(gomock.Any()).Return([]*repository.Setting{
						{Key: "key-1", Value: "value-1"},
						{Key: "key-2", Value: "value-2"},
					}, nil)
				},
			},
		},

		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
				},
				ExpectedErr: "code=500, message=some-error",
			},
			settingCntrlBuilder: settingCntrlBuilder{
				settingSvcFn: func(svc *service_mock.MockSettingSvc) {
					svc.EXPECT().Find(gomock.Any()).Return(nil, errors.New("some-error"))
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).Find)
		})
	}
}

func TestSettingCntrl_FindOne(t *testing.T) {
	testcases := []settingTestCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
					URLParams: map[string]string{
						"key": "some-key",
					},
				},
				ExpectedErr: "code=500, message=some-error",
			},
			settingCntrlBuilder: settingCntrlBuilder{
				settingSvcFn: func(svc *service_mock.MockSettingSvc) {
					svc.EXPECT().FindOne(gomock.Any(), "some-key").Return(nil, errors.New("some-error"))
				},
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
					URLParams: map[string]string{
						"key": "some-key",
					},
				},

				ExpectedCode: http.StatusOK,
				ExpectedBody: "{\"key\":\"key-1\",\"value\":\"value-1\",\"updated_at\":\"0001-01-01T00:00:00Z\"}\n",
			},
			settingCntrlBuilder: settingCntrlBuilder{
				settingSvcFn: func(svc *service_mock.MockSettingSvc) {
					svc.EXPECT().FindOne(gomock.Any(), "some-key").Return(&repository.Setting{
						Key:   "key-1",
						Value: "value-1",
					}, nil)
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).FindOne)
		})
	}
}

func TestSettingCntrl_Update(t *testing.T) {
	testcases := []settingTestCase{
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
					URLParams: map[string]string{
						"key": "some-key",
					},
					Body:   "{\"key\":\"key-1\",\"value\":\"value-1\"}\n",
					Header: echotest.HeaderForJSON(),
				},

				ExpectedErr: "code=500, message=some-error",
			},
			settingCntrlBuilder: settingCntrlBuilder{
				settingSvcFn: func(svc *service_mock.MockSettingSvc) {
					svc.EXPECT().
						Update(
							gomock.Any(),
							"some-key",
							&repository.Setting{
								Key:   "key-1",
								Value: "value-1",
							},
						).
						Return(errors.New("some-error"))
				},
			},
		},
		{
			TestCase: echotest.TestCase{
				Request: echotest.Request{
					Method: http.MethodGet,
					Target: "/",
					URLParams: map[string]string{
						"key": "some-key",
					},
					Body:   "{\"key\":\"key-1\",\"value\":\"value-1\"}\n",
					Header: echotest.HeaderForJSON(),
				},
				ExpectedCode: http.StatusOK,
			},
			settingCntrlBuilder: settingCntrlBuilder{
				settingSvcFn: func(svc *service_mock.MockSettingSvc) {
					svc.EXPECT().
						Update(
							gomock.Any(),
							"some-key",
							&repository.Setting{
								Key:   "key-1",
								Value: "value-1",
							},
						).
						Return(nil)
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			mock := gomock.NewController(t)
			defer mock.Finish()
			tt.Execute(t, tt.build(mock).Update)
		})
	}
}

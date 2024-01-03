package tests

import (
	"context"
	"errors"
	"reflect"
	"short-url/pkg/domain"
	"short-url/pkg/repository/mocks"
	"short-url/pkg/usecase"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_SaveURL(t *testing.T) {
	type args struct {
		ctx context.Context
		url domain.URL
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fail: error saving url",
			args: args{
				ctx: context.TODO(),
				url: domain.URL{
					Key:         "9892121",
					ShortURL:    "https://shorturl/9892121",
					LongURL:     "https://askjnkansknaknkjsabnk/janxjkbxakjnz.com",
					IsKeySigned: false,
				},
			},
			wantErr: true,
		},
		{
			name: "success: save url successfull",
			args: args{
				ctx: context.TODO(),
				url: domain.URL{
					Key:         "9892121",
					ShortURL:    "https://shorturl/9892121",
					LongURL:     "https://askjnkansknaknkjsabnk/janxjkbxakjnz.com",
					IsKeySigned: true,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedRepo := mocks.MockURLRepo{}
			if tt.wantErr {
				mockedRepo.
					On("Save", mock.Anything, mock.Anything).
					Return(errors.New("unable to save url in db"))
			} else {
				mockedRepo.
					On("Save", mock.Anything, mock.Anything).
					Return(nil)
			}

			urlUsecase := usecase.URLInteractor{
				URL: &mockedRepo,
			}

			err := urlUsecase.URL.Save(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error: %v\n", err.Error())
				return
			}
		})
	}
}

func Test_FetchURL(t *testing.T) {
	type args struct {
		ctx context.Context
		url domain.URL
	}

	tests := []struct {
		name    string
		args    args
		want    domain.URL
		wantErr bool
	}{
		{
			name: "fail: error fetching url",
			args: args{
				ctx: context.TODO(),
				url: domain.URL{
					Key:         "9892121",
					IsKeySigned: false,
				},
			},
			wantErr: true,
		},
		{
			name: "success: fetched url successfull",
			args: args{
				ctx: context.TODO(),
				url: domain.URL{
					Key:         "9892121",
					IsKeySigned: true,
				},
			},
			want: domain.URL{
				Key:         "9892121",
				IsKeySigned: true,
				LongURL:     "https://jknkjnkjan/sjnkan.com",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedRepo := mocks.MockURLRepo{}
			if tt.wantErr {
				mockedRepo.
					On("Fetch", mock.Anything, mock.Anything).
					Return(domain.URL{}, errors.New("invalid_hash"))
			} else {
				mockedRepo.
					On("Fetch", mock.Anything, mock.Anything).
					Return(domain.URL{
						Key:         "9892121",
						IsKeySigned: true,
						LongURL:     "https://jknkjnkjan/sjnkan.com",
					}, nil)
			}

			urlUsecase := usecase.URLInteractor{
				URL: &mockedRepo,
			}

			got, err := urlUsecase.URL.Fetch(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error: %v\n", err.Error())
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unable to fetch desired result. Got: %v, Want: %v", got, tt.want)
			}
		})
	}
}

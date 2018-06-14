package jwt

import (
	"reflect"
	"testing"
)

func TestTokenRepository_GetAuthClaims(t *testing.T) {
	u := UserInfo{
		ID:        "abc123",
		Email:     "test@test.com",
		FirstName: "te",
		LastName:  "st",
		Name:      "test",
		Picture:   "imgur.com/test",

		Roles: map[string]string{
			"ab": "admin",
		},
		Permissions: map[string]string{
			"ac": "all",
		},
	}
	_ = u

	type fields struct {
		authKey    []byte
		refreshKey []byte
	}
	type args struct {
		user UserInfo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BaseClaims
	}{
		/*	{
				name: "I will pass",
				fields: nil,
				args: args{
					user: u,
					},
				},
				want: nil,
			},
		*/}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te := &TokenRepository{
				authKey:    tt.fields.authKey,
				refreshKey: tt.fields.refreshKey,
			}
			if got := te.GetAuthClaims(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenRepository.GetAuthClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}

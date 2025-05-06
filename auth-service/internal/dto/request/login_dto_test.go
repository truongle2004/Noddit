package request

import (
	"testing"
)

func TestLoginDto_Validate(t *testing.T) {
	type fields struct {
		Email    string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		errMsg  string
	}{
		{
			name:    "missing email",
			fields:  fields{Email: "", Password: "123"},
			wantErr: true,
			errMsg:  "email is required",
		},
		{
			name:    "missing password",
			fields:  fields{Email: "123", Password: ""},
			wantErr: true,
			errMsg:  "password is require",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoginDto{
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
			}
			if err := l.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("LoginDto.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

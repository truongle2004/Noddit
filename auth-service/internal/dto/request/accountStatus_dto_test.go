package request

import "testing"

func TestAccountStatusUpdateRequest_Validate(t *testing.T) {
	type fields struct {
		Status string
		ID     string
		Email  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		errMsg  string
	}{
		{
			name:    "missing status",
			fields:  fields{Status: ""},
			wantErr: true,
			errMsg:  "status is required",
		},
		{
			name:    "missing id",
			fields:  fields{Status: "ACT", ID: ""},
			wantErr: true,
			errMsg:  "id is required",
		},
		{
			name:   "missing email",
			fields: fields{Status: "ACT", ID: "123", Email: ""},

			wantErr: true,
			errMsg:  "email is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &AccountStatusUpdateRequest{
				Status: tt.fields.Status,
				ID:     tt.fields.ID,
				Email:  tt.fields.Email,
			}
			if err := req.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AccountStatusUpdateRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

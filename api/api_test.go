package api

import (
	"testing"
)

func Test_verifyEmail(t *testing.T) {
	type args struct {
		new_email     string
		existed_email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "email does not exists",
			args: args{
				new_email:     "xpto@gmail.com",
				existed_email: "topx@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "email does exists",
			args: args{
				new_email:     "wolverine@gmail.com",
				existed_email: "wolverine@gmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyEmail(tt.args.new_email, tt.args.existed_email); (err != nil) != tt.wantErr {
				t.Errorf("verifyEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

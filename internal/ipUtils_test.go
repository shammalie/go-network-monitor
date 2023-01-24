package internal

import "testing"

func TestIsPrivateIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "non-private ip",
			args: args{
				ip: "145.124.12.32",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "private ip 1",
			args: args{
				ip: "127.0.0.1",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "private ip 2",
			args: args{
				ip: "172.16.20.4",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "private ip 3",
			args: args{
				ip: "192.168.20.2",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "no ip passed",
			args: args{
				ip: "",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrivateIpCheck(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateIpCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrivateIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

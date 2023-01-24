package internal

import (
	"testing"
)

func TestIpProcessor_Add(t *testing.T) {
	type fields struct {
		counterMs int64
		status    string
		incoming  chan string
		outgoing  chan *IpDetail
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "rate limited",
			fields: fields{
				status: rateLimitedStatus,
			},
			args: args{
				ip: "8.8.8.8",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IpProcessor{
				status:   tt.fields.status,
				incoming: tt.fields.incoming,
				outgoing: tt.fields.outgoing,
			}
			if err := p.Add(tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("IpProcessor.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

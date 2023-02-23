package state

import (
	"reflect"
	"testing"
)

func TestNewFifoQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Fifo
	}{
		{
			name: "pass",
			want: &Fifo{
				queue: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFifoQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFifoQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFifo_findElement(t *testing.T) {
	type fields struct {
		queue []string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "arg not found",
			fields: fields{
				queue: []string{"test1", "test2", "test1", "test3"},
			},
			args: args{
				v: "test",
			},
			want: false,
		},
		{
			name: "arg found",
			fields: fields{
				queue: []string{"test1", "test2", "test", "test1", "test3"},
			},
			args: args{
				v: "test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fifo{
				queue: tt.fields.queue,
			}
			if got := f.findElement(tt.args.v); got != tt.want {
				t.Errorf("Fifo.findElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFifo_GetLen(t *testing.T) {
	type fields struct {
		queue []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "pass",
			fields: fields{
				queue: []string{"A", "B", "C", "D"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fifo{
				queue: tt.fields.queue,
			}
			if got := f.GetLen(); got != tt.want {
				t.Errorf("Fifo.GetLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFifo_Enqueue(t *testing.T) {
	type fields struct {
		queue []string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "add value",
			fields: fields{
				queue: []string{},
			},
			args: args{
				v: "some value",
			},
			wantErr: false,
		},
		{
			name: "value already exists",
			fields: fields{
				queue: []string{"some", "value", "some value"},
			},
			args: args{
				v: "some value",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fifo{
				queue: tt.fields.queue,
			}
			if err := f.Enqueue(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Fifo.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFifo_Dequeue(t *testing.T) {
	type fields struct {
		queue []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "pass",
			fields: fields{
				queue: []string{"1", "2", "3", "4"},
			},
			want: "1",
		},
		{
			name: "no values to dequeue",
			fields: fields{
				queue: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fifo{
				queue: tt.fields.queue,
			}
			if got := f.Dequeue(); got != tt.want {
				t.Errorf("Fifo.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

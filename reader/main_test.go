package main

import (
	"testing"
)

func saveMock(customer Customer, f Action) (int, error) {
	return 0, nil
}

func postMock(customer []byte) (string, error) {
	return "", nil
}

func Test_read(t *testing.T) {
	type args struct {
		fileName string
		save     ActionDB
		post     CrmPost
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"test", args{fileName: "/tmp/reader/test.csv", save: saveMock, post: postMock}, 30, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := read(tt.args.fileName, tt.args.save, tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_save(t *testing.T) {
	type args struct {
		customer Customer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"save",
			args{
				Customer{
					Id:        "1",
					Firstname: "Alberto",
					Lastname:  "Rodrigues",
					Email:     "alberto@rodrigues.fake",
					Phone:     "4543543543",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := action(tt.args.customer, insert)
			if err != nil {
				t.Errorf("save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

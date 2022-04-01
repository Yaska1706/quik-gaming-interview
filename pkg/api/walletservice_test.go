package api

import (
	"testing"
)

func Test_findbalance(t *testing.T) {
	type args struct {
		credit string
		debit  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				credit: "400",
				debit:  "500",
			},
			want: "-100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findbalance(tt.args.credit, tt.args.debit); got != tt.want {
				t.Errorf("findbalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSum(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				array: []string{"34", "45", "56.45", "45"},
			},
			want: "180.45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSum(tt.args.array); got != tt.want {
				t.Errorf("findSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walletService_validateAmount(t *testing.T) {
	type fields struct {
		storage WalletRepository
	}
	type args struct {
		amount string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "When amount is empty",
			args: args{
				amount: "",
			},
			wantErr: true,
		},
		{
			name: "When amount is negative",
			args: args{
				amount: "-3445",
			},
			wantErr: true,
		},
		{
			name: "When amount is not a number",
			args: args{
				amount: "hello",
			},
			wantErr: true,
		},
		{
			name: "When amount is correct",
			args: args{
				amount: "4000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &walletService{
				storage: tt.fields.storage,
			}
			if err := w.validateAmount(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("walletService.validateAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

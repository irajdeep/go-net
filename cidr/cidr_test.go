package main

import "testing"

func TestIsOverlap(t *testing.T) {
	tests := []struct {
		cidra string
		cidrb string
		want  bool
	}{
		{
			cidra: "172.22.2.0/24",
			cidrb: "172.22.2.0/24",
			want:  true,
		},
		{
			cidra: "172.22.1.0/24",
			cidrb: "172.22.2.0/24",
			want:  false,
		},
		{
			cidra: "192.168.1.0/24",
			cidrb: "192.168.2.0/24",
			want:  false,
		},
		{
			cidra: "192.168.1.0/24",
			cidrb: "192.168.2.0/25",
			want:  false,
		},
		{
			cidra: "192.168.2.0/24",
			cidrb: "192.168.2.0/25",
			want:  true,
		},
		{
			cidra: "192.168.2.0/24",
			cidrb: "192.168.2.0/10",
			want:  true,
		},
		{
			cidra: "192.168.2.0/24",
			cidrb: "192.168.3.0/24",
			want:  false,
		},
		{
			cidra: "192.168.2.0/23",
			cidrb: "192.168.3.0/24",
			want:  true,
		},
		{
			cidra: "10.0.0.0/22",
			cidrb: "10.0.1.0/24",
			want:  true,
		},
		{
			cidra: "172.20.0.0/16",
			cidrb: "172.31.0.0/16",
			want:  false,
		},
	}

	for _, tt := range tests {
		got := IsOverlap(tt.cidra, tt.cidrb)
		if got != tt.want {
			t.Errorf("got: %t, want: %t", got, tt.want)
		}
	}
}

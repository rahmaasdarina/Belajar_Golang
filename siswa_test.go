package services

import "testing"

func TestHitungNilai(t *testing.T) {
	type args struct {
		nilai int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case return C",
			args: args{
				nilai: 35,
			},
			want: "Nilai Siswa : C",
		},
		{
			name: "Case return B",
			args: args{
				nilai: 60,
			},
			want: "Nilai Siswa : B",
		},
		{
			name: "Case return A",
			args: args{
				nilai: 90,
			},
			want: "Nilai Siswa : A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HitungNilai(tt.args.nilai); got != tt.want {
				t.Errorf("HitungNilai() = %v, want %v", got, tt.want)
			}
		})
	}
}

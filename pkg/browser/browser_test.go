package browser

import "testing"

func TestIsWebUrlOrLocalFilePath(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test_https_true", args{"https://www.baidu.com"}, true},
		{"test_http_true", args{"http://www.baidu.com"}, true},
		{"test_file_true", args{"/home"}, true},
		{"test_file_false", args{"/home/zzzzzzzzzzzz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWebUrlOrLocalFilePath(tt.args.s); got != tt.want {
				t.Errorf("IsWebUrlOrLocalFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

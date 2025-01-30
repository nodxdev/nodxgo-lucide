package assert

import "testing"

func Equal(t *testing.T, want any, got any) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func NotEqual(t *testing.T, want any, got any) {
	t.Helper()
	if want == got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("want error, got nil")
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("want no error, got %v", err)
	}
}

func Nil(t *testing.T, got any) {
	t.Helper()
	if got != nil {
		t.Errorf("want nil, got %v", got)
	}
}

func True(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("want true, got false")
	}
}

func False(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("want false, got true")
	}
}

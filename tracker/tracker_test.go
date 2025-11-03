package tracker

import (
	"net/http"
	"testing"
)

func TestCreateFingerprintConsistency(t *testing.T) {
	ip := "192.168.1.1"
	ua := "TestAgent"

	fp1 := createFingerprint(ip, ua)
	fp2 := createFingerprint(ip, ua)

	if fp1 != fp2 {
		t.Fatalf("fingerprints differ for same input: %s vs %s", fp1, fp2)
	}
}

func TestCreateFingerprintUniqueness(t *testing.T) {
	ip := "192.168.1.1"
	ua := "TestAgent"
	differentUA := "OtherAgent"

	fp1 := createFingerprint(ip, ua)
	fp2 := createFingerprint(ip, differentUA)

	if fp1 == fp2 {
		t.Fatalf("fingerprints should differ for different input: %s", fp1)
	}
}

func TestGetRedisKey(t *testing.T) {
	fingerprint := "abc123"
	path := "homepage"
	expected := "tracker:abc123:homepage"

	result := getRedisKey(fingerprint, path)

	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestGetIPAddress(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string]string
		expected string
	}{
		{
			name:     "X-Forwarded-For with single IP",
			headers:  map[string]string{"X-Forwarded-For": "192.168.1.1"},
			expected: "192.168.1.1",
		},
		{
			name:     "X-Forwarded-For with multiple IPs",
			headers:  map[string]string{"X-Forwarded-For": "192.168.1.1, 10.0.0.1, 172.16.0.1"},
			expected: "192.168.1.1",
		},
		{
			name:     "X-Real-IP",
			headers:  map[string]string{"X-Real-IP": "192.168.1.2"},
			expected: "192.168.1.2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}

			result := getIPAddress(req)

			if result != tt.expected {
				t.Fatalf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// Benchmark tests to validate performance improvements
func BenchmarkCreateFingerprint(b *testing.B) {
	ip := "192.168.1.1"
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createFingerprint(ip, ua)
	}
}

func BenchmarkGetIPAddressSingleIP(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIPAddress(req)
	}
}

func BenchmarkGetIPAddressMultipleIPs(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.1, 10.0.0.1, 172.16.0.1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getIPAddress(req)
	}
}

func BenchmarkGetRedisKey(b *testing.B) {
	fingerprint := "a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6"
	path := "homepage"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getRedisKey(fingerprint, path)
	}
}

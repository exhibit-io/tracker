package tracker

import "testing"

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

// +build debug

package errors

import (
	"testing"
)

func TestSerializableInDebug(t *testing.T) {
	cause := New("causal error")
	const msg = "external error message for the client"
	err := New(msg, Permission, cause)
	got := testJSONRoundTrip(Serializable(err)).(*Error)
	if got.msg != msg {
		t.Errorf("Serializable(); got msg = %v, want %v", got.msg, msg)
	}
	if got.cause == nil || got.cause.Error() != cause.Error() {
		t.Errorf("Serializable(); got cause.Error() = %v, want %v", got.cause.Error(), cause.Error())
	}
	if got.kind != Permission {
		t.Errorf("Serializable(); got kind = %v, want %v", got.kind, Permission)
	}

}

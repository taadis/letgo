package errors

import "testing"

func TestErrors(t *testing.T) {
	tests := []*Error{
		&Error{
			Id:      "test-error-500",
			Code:    500,
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
		},
		// more...
	}

	for _, e := range tests {
		newErr := New(e.Id, e.Code, e.Status, e.Message)
		if newErr.Error() != e.Error() {
			t.Fatalf("want %s but got %s", e.Error(), newErr.Error())
		}

		parseErr := Parse(newErr.Error())
		if parseErr == nil {
			t.Fatalf("want error but got nil")
		}
		if parseErr.Id != e.Id {
			t.Fatalf("want id %s but got id %s", e.Id, parseErr.Id)
		}
		if parseErr.Code != e.Code {
			t.Fatalf("want code %d but got code %d", e.Code, parseErr.Code)
		}
		if parseErr.Status != e.Status {
			t.Fatalf("want status %s but got status %s", e.Status, parseErr.Status)
		}
		if parseErr.Message != e.Message {
			t.Fatalf("want message %s but got message %s", e.Message, parseErr.Message)
		}
	}
}

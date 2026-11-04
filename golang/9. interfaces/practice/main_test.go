package main

import (
	"fmt"
	"testing"
)

type MockWriter struct {
	Received []string
	Failed   bool
}

func (m *MockWriter) Write(data string) error {
	if m.Failed {
		return fmt.Errorf("mock failure")
	}

	m.Received = append(m.Received, data)
	return nil
}

func TestWriterMock(t *testing.T) {
	mock := &MockWriter{}

	err := mock.Write("Test Message")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(mock.Received) != 1 || mock.Received[0] != "Test Message" {
		t.Fatalf("message not received correctly, got: %v", mock.Received)
	}
}

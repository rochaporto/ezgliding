package collection

import "testing"

func TestEcho(t *testing.T) {
  res := Echo("Hello")
  if res != "Hello" {
    t.Errorf("Failed")
  }
}

func TestEchoEmpty(t *testing.T) {
  res := Echo("")
  if res != "" {
    t.Errorf("Failed to Echo empty string :: got '%v'", res)
  }
}

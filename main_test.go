package currency

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	setupTest()
	code := m.Run()
	shutdownTest()

	os.Exit(code)
}

func setupTest() {
}

func shutdownTest() {
}

// Create satment insert product.
func Test_CheckToFloat64(t *testing.T) {
	want := 123.45
	c, err := Parse("123.45", ".")
	if err != nil {
		t.Errorf("Could not parse string price")
	}

	cf := c.ToFloat64()
	if cf != want {
		t.Errorf("Received %f, want %f", cf, want)
	}
}

package letgo

import "testing"

func TestRun(t *testing.T) {

	app := NewApp(
		WithEnv(),
		WithPort(),
		WithConfigPath(),
		WithName(),
		WithVersion(),
	)
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}

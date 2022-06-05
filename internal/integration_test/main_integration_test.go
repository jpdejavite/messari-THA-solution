package integration_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestIntegration(t *testing.T) {
	executeIntegrationTestsEnvVar := os.Getenv("EXECUTE_INTEGRATION_TEST")
	if executeIntegrationTestsEnvVar == "" {
		t.Skip("set EXECUTE_INTEGRATION_TEST to run this test")
	}

	inputLines := []string{
		"BEGIN",
		`{"id":1,"market":8,"price":7.37,"volume":2633.63,"is_buy":true}`,
		`{"id":2,"market":7,"price":10,"volume":100,"is_buy":true}`,
		`{"id":3,"market":7,"price":20,"volume":200,"is_buy":false}`,
		`{"id":4,"market":7,"price":30,"volume":200,"is_buy":true}`,
		`{"id":5,"market":7,"price":40,"volume":500,"is_buy":false}`,
		"END",
	}
	result := `{"market":8,"total_volume":2633.63,"mean_price":7.37,"mean_volume":2633.63,"volume_weighted_average_price":7.37,"percentage_buy":1}
{"market":7,"total_volume":1000,"mean_price":25,"mean_volume":250,"volume_weighted_average_price":31,"percentage_buy":0.5}
`

	cmd := exec.Command("go", "run", "../../main.go")
	stdin, e := cmd.StdinPipe()
	if e != nil {
		t.Errorf("Error getting stdin pipe, %v", e)
	}

	for _, line := range inputLines {
		_, e = stdin.Write([]byte(line + "\n"))
		if e != nil {
			t.Errorf("Error writing on stdin pipe, %v", e)
		}
	}

	out, e := cmd.Output()
	if e != nil {
		t.Errorf("Error on command output, %v", e)
	}

	if result != string(out) {
		t.Errorf("IntegrationTest() got = %v, want %v", string(out), result)
	}
}

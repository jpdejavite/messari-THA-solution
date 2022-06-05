package command_line

import (
	"testing"
)

// An interface for the minimal api our code needs to work
type MockStringReader struct {
	stringMessages []string
}

func (sr *MockStringReader) Read(p []byte) (n int, err error) {
	if len(sr.stringMessages) == 0 {
		n = copy(p, []byte{'\n'})
		return
	}
	response := sr.stringMessages[0]
	n = copy(p, []byte(response))

	sr.stringMessages = sr.stringMessages[1:]
	return len(response), nil
}

func Test_Reader_ReadFromCommanLine(t *testing.T) {

	type args struct {
		stringMessages []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "begin, json and end",
			args: args{
				stringMessages: []string{"BEGIN\n", "TESTE\n", "END\n"},
			},
		},
		{
			name: "begin, invalid json and end", // no panic should be given
			args: args{
				stringMessages: []string{"BEGIN\n", "INVALID\n", "END\n"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := MockStringReader{
				stringMessages: tt.args.stringMessages,
			}
			ReadFromCommanLine(&reader)

		})
	}
}

package gemspec_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aquasecurity/go-dep-parser/pkg/ruby/gemspec"
	"github.com/aquasecurity/go-dep-parser/pkg/types"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		want      types.Library
		wantErr   string
	}{
		{
			name:      "happy",
			inputFile: "testdata/normal00.gemspec",
			want:      types.NewLibrary("rake", "13.0.3", "MIT"),
		},
		{
			name:      "another variable name",
			inputFile: "testdata/normal01.gemspec",
			want:      types.NewLibrary("async", "1.25.0", ""),
		},
		{
			name:      "license",
			inputFile: "testdata/license.gemspec",
			want:      types.NewLibrary("async", "1.25.0", "MIT"),
		},
		{
			name:      "multiple licenses",
			inputFile: "testdata/multiple_licenses.gemspec",
			want:      types.NewLibrary("test-unit", "3.3.7", "Ruby, BSDL, PSFL"),
		},
		{
			name:      "malformed variable name",
			inputFile: "testdata/malformed00.gemspec",
			wantErr:   "failed to parse gemspec",
		},
		{
			name:      "missing version",
			inputFile: "testdata/malformed01.gemspec",
			wantErr:   "failed to parse gemspec",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.inputFile)
			require.NoError(t, err)

			got, err := gemspec.Parse(f)
			if tt.wantErr != "" {
				require.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

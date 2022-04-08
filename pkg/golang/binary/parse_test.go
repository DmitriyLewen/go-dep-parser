package binary_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/aquasecurity/go-dep-parser/pkg/golang/binary"
	"github.com/aquasecurity/go-dep-parser/pkg/types"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		want      []types.Library
		wantErr   string
	}{
		{
			name:      "ELF",
			inputFile: "testdata/test.elf",
			want: []types.Library{
				types.NewLibrary("github.com/aquasecurity/go-pep440-version", "v0.0.0-20210121094942-22b2f8951d46", ""),
				types.NewLibrary("github.com/aquasecurity/go-version", "v0.0.0-20210121072130-637058cfe492", ""),
				types.NewLibrary("golang.org/x/xerrors", "v0.0.0-20200804184101-5ec99f83aff1", ""),
			},
		},
		{
			name:      "PE",
			inputFile: "testdata/test.exe",
			want: []types.Library{
				types.NewLibrary("github.com/aquasecurity/go-pep440-version", "v0.0.0-20210121094942-22b2f8951d46", ""),
				types.NewLibrary("github.com/aquasecurity/go-version", "v0.0.0-20210121072130-637058cfe492", ""),
				types.NewLibrary("golang.org/x/xerrors", "v0.0.0-20200804184101-5ec99f83aff1", ""),
			},
		},
		{
			name:      "Mach-O",
			inputFile: "testdata/test.macho",
			want: []types.Library{
				types.NewLibrary("github.com/aquasecurity/go-pep440-version", "v0.0.0-20210121094942-22b2f8951d46", ""),
				types.NewLibrary("github.com/aquasecurity/go-version", "v0.0.0-20210121072130-637058cfe492", ""),
				types.NewLibrary("golang.org/x/xerrors", "v0.0.0-20200804184101-5ec99f83aff1", ""),
			},
		},
		{
			name:      "with replace directive",
			inputFile: "testdata/replace.elf",
			want: []types.Library{
				types.NewLibrary("github.com/davecgh/go-spew", "v1.1.1", ""),
				types.NewLibrary("github.com/go-sql-driver/mysql", "v1.5.0", ""),
			},
		},
		{
			name:      "sad path",
			inputFile: "testdata/dummy",
			wantErr:   "unrecognized executable format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.inputFile)
			require.NoError(t, err)

			got, _, err := binary.Parse(f)
			if tt.wantErr != "" {
				require.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

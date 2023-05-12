package gitlab_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProject(t *testing.T) {
	cases := []struct {
		projectID string
		want      int
	}{
		{
			projectID: "45945944",
			want:      http.StatusOK,
		},
	}

	gitlab := NewGitlab()

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s - %d", tc.projectID, tc.want), func(t *testing.T) {
			res, err := gitlab.GetProject(tc.projectID)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, res.StatusCode)
		})
	}
}

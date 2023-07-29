package v1tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/lenguti/ezuzu/app/services/property/api/handlers/v1"
	"github.com/lenguti/ezuzu/business/core/manager"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateManager(t *testing.T) {
	// Setup.
	ctx := context.Background()
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	t.Run("success", func(t *testing.T) {
		input := v1.CreateManagerRequest{
			Entity: "Test Entity Enterprise",
		}
		want := v1.ClientManager{
			Entity: "Test Entity Enterprise",
		}
		ctrl := v1.Controller{
			Manager: manager.NewCore(&mockManagerStore{
				createFunc: func() error {
					return nil
				},
			}, log),
		}

		bs, err := json.Marshal(input)
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/managers", bytes.NewBuffer(bs))
		require.NoError(t, err)

		// Execute.
		err = ctrl.CreateManager(ctx, w, r)

		// Validate.
		require.NoError(t, err)

		var got v1.CreateManagerResponse
		err = json.Unmarshal(w.Body.Bytes(), &got)
		require.NoError(t, err)

		assert.Equal(t, want.Entity, got.Manager.Entity)
	})
}

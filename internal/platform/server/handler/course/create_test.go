package course

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jpastorm/hexagonalgoexample-cdtv/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course"))

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses" , CreateHandler(courseRepository))
	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createRequest := createRequest{
			ID: "129381923891823192",
			Name: "Demo course",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err )

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createRequest := createRequest{
			ID: "129381923891823192",
			Name: "Demo course",
			Duration: "10 months",
		}

		b, err := json.Marshal(createRequest)
		require.NoError(t, err )

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

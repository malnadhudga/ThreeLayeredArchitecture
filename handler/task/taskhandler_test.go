package task

import (
	models "ThreeLayeredArchitecture/models/task"
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestTaskHandler_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskservice(ctrl)
	handler := NewTaskHandler(mockService)

	testCases := []struct {
		id         int
		input      models.Task
		expected   models.Task
		mockErr    error
		expectCode int
		expectBody bool
	}{
		{
			id:         1,
			input:      models.Task{Description: "Test Task"},
			expected:   models.Task{ID: 1, Description: "Test Task"},
			mockErr:    nil,
			expectCode: http.StatusOK,
			expectBody: true,
		},
		{
			id:         2,
			input:      models.Task{Description: "Fail Task"},
			expected:   models.Task{},
			mockErr:    errors.New("failed to add"),
			expectCode: http.StatusInternalServerError,
			expectBody: false,
		},
	}

	for _, tc := range testCases {
		bodyBytes, _ := json.Marshal(tc.input)
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		mockService.EXPECT().Add(tc.input.Description).Return(tc.expected, tc.mockErr)

		handler.HandleTasks(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tc.expectCode {
			t.Errorf("[Test ID %d]: expected status %d, got %d", tc.id, tc.expectCode, resp.StatusCode)
		}

		if tc.expectBody {
			body, _ := io.ReadAll(resp.Body)

			var result models.Task
			json.Unmarshal(body, &result)

			if result.Description != tc.expected.Description {
				t.Errorf("[Test ID %d]: expected %+v, got %+v", tc.id, tc.expected, result)
			}
		}
	}
}

func TestTaskHandler_GetPending(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskservice(ctrl)
	handler := NewTaskHandler(mockService)

	testCases := []struct {
		id         int
		desc       string
		mockTasks  []models.Task
		mockErr    error
		expectCode int
	}{
		{
			id:         1,
			desc:       "Success fetching tasks",
			mockTasks:  []models.Task{{ID: 1, Description: "Pending Task"}},
			mockErr:    nil,
			expectCode: http.StatusOK,
		},
		{
			id:         2,
			desc:       "Error fetching tasks",
			mockTasks:  nil,
			mockErr:    errors.New("error"),
			expectCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", "/tasks", nil)
		w := httptest.NewRecorder()

		mockService.EXPECT().GetPending().Return(tc.mockTasks, tc.mockErr)

		handler.HandleTasks(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tc.expectCode {
			t.Errorf("[Test ID %d] %s: expected %d, got %d", tc.id, tc.desc, tc.expectCode, resp.StatusCode)
		}
	}
}

func TestTaskHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskservice(ctrl)
	handler := NewTaskHandler(mockService)

	testCases := []struct {
		id         int
		desc       string
		queryID    string
		mockErr    error
		expectCode int
	}{
		{
			id:         1,
			desc:       "Successful delete",
			queryID:    "1",
			mockErr:    nil,
			expectCode: http.StatusOK,
		},
		{
			id:         2,
			desc:       "Failed delete",
			queryID:    "2",
			mockErr:    errors.New("not found"),
			expectCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("DELETE", "/tasks?id="+tc.queryID, nil)
		w := httptest.NewRecorder()

		id, _ := strconv.Atoi(tc.queryID)
		mockService.EXPECT().Delete(id).Return(tc.mockErr)

		handler.HandleTasks(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tc.expectCode {
			t.Errorf("[Test ID %d] %s: expected %d, got %d", tc.id, tc.desc, tc.expectCode, resp.StatusCode)
		}
	}
}

func TestTaskHandler_MarkComplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskservice(ctrl)
	handler := NewTaskHandler(mockService)

	testCases := []struct {
		id         int
		desc       string
		queryID    string
		mockMsg    string
		mockErr    error
		expectCode int
	}{
		{
			id:         1,
			desc:       "Successful mark complete",
			queryID:    "1",
			mockMsg:    "Marked complete",
			mockErr:    nil,
			expectCode: http.StatusOK,
		},
		{
			id:         2,
			desc:       "Failed mark complete",
			queryID:    "2",
			mockMsg:    "",
			mockErr:    errors.New("update failed"),
			expectCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("PATCH", "/tasks?id="+tc.queryID, nil)
		w := httptest.NewRecorder()

		id, _ := strconv.Atoi(tc.queryID)
		mockService.EXPECT().MarkComplete(id).Return(tc.mockMsg, tc.mockErr)

		handler.HandleTasks(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tc.expectCode {
			t.Errorf("[Test ID %d] %s: expected %d, got %d", tc.id, tc.desc, tc.expectCode, resp.StatusCode)
		}
	}
}

func TestTaskHandler_HandleTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskservice(ctrl)
	handler := NewTaskHandler(mockService)

	testCases := []struct {
		id         int
		desc       string
		pathID     string
		expected   models.Task
		mockErr    error
		expectCode int
	}{
		{
			id:         1,
			desc:       "Successful fetch by ID",
			pathID:     "1",
			expected:   models.Task{ID: 1, Description: "Task by ID"},
			mockErr:    nil,
			expectCode: http.StatusOK,
		},
		{
			id:         2,
			desc:       "Task not found",
			pathID:     "2",
			expected:   models.Task{},
			mockErr:    errors.New("not found"),
			expectCode: http.StatusNotFound,
		},
		{
			id:         3,
			desc:       "Task not found",
			pathID:     "3e",
			expected:   models.Task{},
			mockErr:    errors.New("not found"),
			expectCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", "/tasks/"+tc.pathID, nil)
		req.SetPathValue("id", tc.pathID)
		w := httptest.NewRecorder()

		id, err := strconv.Atoi(tc.pathID)
		if err == nil {
			mockService.EXPECT().GetByID(id).Return(tc.expected, tc.mockErr)
		}
		handler.HandleTaskByID(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tc.expectCode {
			t.Errorf("[Test ID %d] %s: expected status %d, got %d", tc.id, tc.desc, tc.expectCode, resp.StatusCode)
		}
	}
}

func TestTaskHandler_MethodNotAllowed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskservice(ctrl)
	handler := NewTaskHandler(mockService)

	req := httptest.NewRequest(http.MethodPut, "/tasks/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()

	handler.HandleTaskByID(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d for PUT method, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	handler.HandleTasks(w, req)

	resp = w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d for PUT method, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

}

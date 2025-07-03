package task

import (
	model "ThreeLayeredArchitecture/models/task"
	"errors"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"testing"
)

var ctx *gofr.Context

func TestTaskService_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStore(ctrl)
	service := NewTaskService(mockStore)

	testCases := []struct {
		id       int
		desc     string
		input    model.Task
		expected model.Task
		mockErr  error
	}{
		{
			id:   1,
			desc: "Success adding task",
			input: model.Task{
				Description: "Write test cases",
			},
			expected: model.Task{ID: 1, Description: "Write test cases"},
			mockErr:  nil,
		},
		{
			id:   2,
			desc: "Failure adding task",
			input: model.Task{
				Description: "Write test cases",
			},
			expected: model.Task{},
			mockErr:  errors.New("failed to add task"),
		},
	}

	for _, tc := range testCases {
		mockStore.EXPECT().Add(ctx, tc.input).Return(tc.expected, tc.mockErr)
		result, err := service.Add(ctx, tc.input)

		if (err == nil && tc.mockErr != nil) || (err != nil && tc.mockErr == nil) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.mockErr, err)
		}

		if result.ID != tc.expected.ID || result.Description != tc.expected.Description {
			t.Errorf("[Test ID %d] %s: expected task %+v, got %+v", tc.id, tc.desc, tc.expected, result)
		}
	}
}

func TestTaskService_GetPending(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStore(ctrl)
	service := NewTaskService(mockStore)

	expected := []model.Task{
		{ID: 1, Description: "Task 1", Completed: false},
		{ID: 2, Description: "Task 2", Completed: false},
	}

	mockStore.EXPECT().GetPending(ctx).Return(expected, nil)
	result, err := service.GetPending(ctx)

	if err != nil {
		t.Errorf("[Success case] unexpected error: %v", err)
	}

	for i := range result {
		if result[i].ID != expected[i].ID || result[i].Description != expected[i].Description {
			t.Errorf("[Success case] mismatch task at index %d: expected %+v, got %+v", i, expected[i], result[i])
		}
	}

	mockErr := errors.New("fetch failed")
	mockStore.EXPECT().GetPending(ctx).Return(nil, mockErr)

	resultFail, errFail := service.GetPending(ctx)
	if errFail == nil || errFail.Error() != mockErr.Error() {
		t.Errorf("[Failure case] expected error %v, got %v", mockErr, errFail)
	}

	if resultFail != nil {
		t.Errorf("[Failure case] expected nil tasks, got %v", resultFail)
	}
}

func TestTaskService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStore(ctrl)
	service := NewTaskService(mockStore)

	testCases := []struct {
		id       int
		desc     string
		inputID  int
		expected model.Task
		mockErr  error
	}{
		{
			id:       1,
			desc:     "Task found",
			inputID:  1,
			expected: model.Task{ID: 1, Description: "Task 1"},
			mockErr:  nil,
		},
		{
			id:       2,
			desc:     "Task not found",
			inputID:  99,
			expected: model.Task{},
			mockErr:  errors.New("task not found"),
		},
	}

	for _, tc := range testCases {
		mockStore.EXPECT().GetByID(ctx, tc.inputID).Return(tc.expected, tc.mockErr)
		result, err := service.GetByID(ctx, tc.inputID)

		if (err == nil && tc.mockErr != nil) ||
			(err != nil && tc.mockErr == nil) ||
			(err != nil && tc.mockErr != nil &&
				err.Error() != tc.mockErr.Error()) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.mockErr, err)
		}

		if result.ID != tc.expected.ID || result.Description != tc.expected.Description {
			t.Errorf("[Test ID %d] %s: expected task %+v, got %+v", tc.id, tc.desc, tc.expected, result)
		}
	}
}

func TestTaskService_MarkComplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStore(ctrl)
	service := NewTaskService(mockStore)

	testCases := []struct {
		id          int
		desc        string
		taskID      int
		initialTask model.Task
		mockErr     error
		expectMsg   string
		expectErr   error
	}{
		{
			id:          1,
			desc:        "Success marking complete",
			taskID:      1,
			initialTask: model.Task{ID: 1, Completed: false},
			mockErr:     nil,
			expectMsg:   "Task marked as complete",
			expectErr:   nil,
		},
		{
			id:          2,
			desc:        "Task already completed",
			taskID:      2,
			initialTask: model.Task{ID: 2, Completed: true},
			mockErr:     nil,
			expectMsg:   "Task already completed",
			expectErr:   nil,
		},
		{
			id:          3,
			desc:        "GetByID error",
			taskID:      3,
			initialTask: model.Task{},
			mockErr:     errors.New("task not found"),
			expectMsg:   "",
			expectErr:   errors.New("task not found"),
		},
		{
			id:          4,
			desc:        "MarkComplete error",
			taskID:      4,
			initialTask: model.Task{ID: 4, Completed: false},
			mockErr:     errors.New("failed to mark complete"),
			expectMsg:   "",
			expectErr:   errors.New("failed to mark complete"),
		},
	}

	for _, tc := range testCases {
		if tc.desc == "GetByID error" {
			mockStore.EXPECT().GetByID(ctx, tc.taskID).Return(tc.initialTask, tc.mockErr)
		} else {
			mockStore.EXPECT().GetByID(ctx, tc.taskID).Return(tc.initialTask, nil)

			if !tc.initialTask.Completed {
				mockStore.EXPECT().MarkComplete(ctx, tc.taskID).Return(tc.mockErr)
			}
		}

		msg, err := service.MarkComplete(ctx, tc.taskID)

		if (err == nil && tc.expectErr != nil) || (err != nil && tc.expectErr == nil) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.expectErr, err)
		}

		if msg != tc.expectMsg {
			t.Errorf("[Test ID %d] %s: expected message '%s', got '%s'", tc.id, tc.desc, tc.expectMsg, msg)
		}
	}
}

func TestTaskService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStore(ctrl)
	service := NewTaskService(mockStore)

	// Set expectation for GetByID before Delete for id=1
	mockStore.EXPECT().GetByID(ctx, 1).Return(model.Task{}, nil)
	mockStore.EXPECT().Delete(ctx, 1).Return(nil)

	_, err := service.Delete(ctx, 1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	mockErr := errors.New("No task Found")

	mockStore.EXPECT().GetByID(ctx, 2).Return(model.Task{}, nil)
	mockStore.EXPECT().Delete(ctx, 2).Return(mockErr)

	_, err = service.Delete(ctx, 2)
	if err == nil || err.Error() != mockErr.Error() {
		t.Errorf("expected error %v, got %v", mockErr, err)
	}

	mockGetErr := errors.New("task not found")
	mockStore.EXPECT().GetByID(ctx, 3).Return(model.Task{}, mockGetErr)

	msg, err := service.Delete(ctx, 3)
	if err == nil || err.Error() != mockGetErr.Error() || msg != "No task Found" {
		t.Errorf("expected getByID error, got msg: %v, err: %v", msg, err)
	}
}

package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	ucase "github.com/meroedu/meroedu/internal/course/usecase"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	// mockUserRepo := new(mocks.UserRepository)
	// mockLessonRepo := new(mocks.LessonRepository)
	// mockAttachmentRepo:=new(mocks.AttachmentRepository)
	// mockCategoryRepo:=new(mocks.CategoryRepository)
	mockListCourse := []domain.Course{
		domain.Course{
			ID: 1, Title: "title-1",
			Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	t.Run("success", func(t *testing.T) {
		mockCourseRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(mockListCourse, nil).Once()

		start := int(0)
		limit := int(1)
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)
		list, err := u.GetAll(context.TODO(), start, limit)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListCourse))
		mockCourseRepo.AssertExpectations(t)

	})
	t.Run("error-failed", func(t *testing.T) {
		mockCourseRepo.On("GetAll", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)
		start := int(0)
		limit := int(1)
		list, err := u.GetAll(context.TODO(), start, limit)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockCourseRepo.AssertExpectations(t)
	})
}
func TestGetByID(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourse := domain.Course{
		Title:  "title-1",
		Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCourse, nil).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockCourse.ID)

		assert.NoError(t, err)
		assert.NotNil(t, *a)

		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockCourse.ID)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockCourseRepo.AssertExpectations(t)
	})
}
func TestGetByTitle(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourse := domain.Course{
		Title:  "title-1",
		Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		mockCourseRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(&mockCourse, nil).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		a, err := u.GetByTitle(context.TODO(), mockCourse.Title)

		assert.NoError(t, err)
		assert.NotNil(t, *a)

		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockCourseRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(nil, errors.New("Unexpected")).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		a, err := u.GetByTitle(context.TODO(), "random")

		assert.Error(t, err)
		assert.Nil(t, a)

		mockCourseRepo.AssertExpectations(t)
	})
}
func TestCreateCourse(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourse := domain.Course{
		Title:       "Hello",
		Description: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempMockCourse := mockCourse
		tempMockCourse.ID = 1
		mockCourseRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockCourseRepo.On("CreateCourse", mock.Anything, mock.AnythingOfType("*domain.Course")).Return(nil).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.CreateCourse(context.TODO(), &tempMockCourse)
		assert.NoError(t, err)
		assert.Equal(t, mockCourse.Title, tempMockCourse.Title)
		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingCourse := mockCourse
		mockCourseRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(&existingCourse, nil).Once()
		mockCourseRepo.On("CreateCourse", mock.Anything, mock.AnythingOfType("*domain.Course")).Return(domain.ErrConflict).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.CreateCourse(context.TODO(), &mockCourse)
		assert.Error(t, err)
	})
}
func TestUpdateCourse(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourse := domain.Course{
		ID:          1,
		Title:       "Hello",
		Description: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempMockCourse := mockCourse
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&domain.Course{}, nil).Once()
		mockCourseRepo.On("UpdateCourse", mock.Anything, mock.AnythingOfType("*domain.Course")).Return(nil).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.UpdateCourse(context.TODO(), &tempMockCourse, tempMockCourse.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockCourse.ID, tempMockCourse.ID)
		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("existing-title", func(t *testing.T) {
		existingCourse := mockCourse
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&existingCourse, nil).Once()
		mockCourseRepo.On("UpdateCourse", mock.Anything, mock.AnythingOfType("*domain.Course")).Return(domain.ErrNotFound).Once()
		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.UpdateCourse(context.TODO(), &mockCourse, existingCourse.ID)

		assert.Error(t, err)
		mockCourseRepo.AssertExpectations(t)
	})
}

func TestDeleteCourse(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)
	mockCourse := domain.Course{
		Title:       "Hello",
		Description: "Description here",
	}

	t.Run("success", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&mockCourse, nil).Once()

		mockCourseRepo.On("DeleteCourse", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.DeleteCourse(context.TODO(), mockCourse.ID)

		assert.NoError(t, err)
		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("course-is-not-exist", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.DeleteCourse(context.TODO(), mockCourse.ID)

		assert.Error(t, err)
		mockCourseRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockCourseRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(&domain.Course{}, errors.New("Unexpected Error")).Once()

		u := ucase.NewCourseUseCase(mockCourseRepo, time.Second*2)

		err := u.DeleteCourse(context.TODO(), mockCourse.ID)

		assert.Error(t, err)
		mockCourseRepo.AssertExpectations(t)
	})

}

package grpc

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/service/dto"
)

type service interface {
	GetCoursesList(ctx context.Context, req dto.GetCoursesListRequest) (res dto.GetCoursesListResponse, err error)
	GetCourseCategories(ctx context.Context, req dto.GetCourseCategoriesRequest) (res dto.GetCourseCategoriesResponse, err error)
	GetUserCourses(ctx context.Context, req dto.GetUserCoursesRequest) (res dto.GetUserCoursesResponse, err error)
	EnrollToCourse(ctx context.Context, req dto.EnrollToCourseRequest) (res dto.EnrollToCourseResponse, err error)

	GetCourseLessons(ctx context.Context, req dto.GetCourseLessonsRequest) (res dto.GetCourseLessonsResponse, err error)
	GetLessonDetails(ctx context.Context, req dto.GetLessonDetailsRequest) (res dto.GetLessonDetailsResponse, err error)
}

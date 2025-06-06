package grpc

import (
	"context"

	"github.com/glamostoffer/arete/learning/internal/service/dto"
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

func (h *handler) GetCoursesList(ctx context.Context, req *v1.GetCoursesListRequest) (res *v1.GetCoursesListResponse, err error) {
	in := dto.GetCoursesListRequest{
		UserID:     req.GetUserID(),
		Categories: nil,
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
	}

	if len(req.GetCategories()) != 0 {
		in.Categories = req.GetCategories()
	}

	out, err := h.service.GetCoursesList(ctx, in)
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

func (h *handler) GetCourseCategories(ctx context.Context, req *v1.GetCourseCategoriesRequest) (res *v1.GetCourseCategoriesResponse, err error) {
	out, err := h.service.GetCourseCategories(ctx, dto.GetCourseCategoriesRequest{})
	if err != nil {
		return res, err
	}

	return &v1.GetCourseCategoriesResponse{
		Categories: out.Categories,
	}, err
}

func (h *handler) EnrollToCourse(ctx context.Context, req *v1.EnrollToCourseRequest) (res *v1.EnrollToCourseResponse, err error) {
	_, err = h.service.EnrollToCourse(ctx, dto.EnrollToCourseRequest{
		UserID:   req.GetUserID(),
		CourseID: req.GetCourseID(),
	})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (h *handler) GetUserCourses(ctx context.Context, req *v1.GetUserCoursesRequest) (res *v1.GetUserCoursesResponse, err error) {
	out, err := h.service.GetUserCourses(ctx, dto.GetUserCoursesRequest{
		UserID: req.GetUserID(),
	})
	if err != nil {
		return res, err
	}

	return out.ToProto(), nil
}

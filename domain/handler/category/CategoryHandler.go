package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	validator "gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"microservice/domain/entitie"
	_interface "microservice/domain/interface"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CategoryHandler struct {
	CategoryService _interface.CategoryService
}

func NewCategoryHandler(echoInstance *echo.Echo, service _interface.CategoryService) {
	handler := &CategoryHandler{
		CategoryService: service,
	}
	echoInstance.PUT("/category", handler.Update)
	echoInstance.POST("/category", handler.Store)
	echoInstance.GET("/category/:id", handler.GetByID)
	echoInstance.DELETE("/category/:id", handler.Delete)
	echoInstance.GET("/category/getAll", handler.GetAll)
}

func (a *CategoryHandler) Delete(echoContext echo.Context) error {

	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, _interface.ErrNotFound.Error())
	}

	id := uint(idP)
	ctx := echoContext.Request().Context()

	category, err := a.CategoryService.GetByID(ctx, id)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	err = a.CategoryService.Delete(ctx, category.ID)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusOK, "")
}

func (a *CategoryHandler) GetByID(echoContext echo.Context) error {

	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, _interface.ErrNotFound.Error())
	}

	id := uint(idP)
	ctx := echoContext.Request().Context()

	art, err := a.CategoryService.GetByID(ctx, id)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusOK, art)
}
func (a *CategoryHandler) GetAll(echoContext echo.Context) error {

	ctx := echoContext.Request().Context()

	art, err := a.CategoryService.GetAll(ctx)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusOK, art)
}
func (a *CategoryHandler) Store(echoContext echo.Context) (err error) {
	var ent entitie.Category
	err = echoContext.Bind(&ent)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&ent); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := echoContext.Request().Context()
	err = a.CategoryService.Store(ctx, &ent)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusCreated, ent)
}

func (a *CategoryHandler) Update(echoContext echo.Context) (err error) {
	var ent entitie.Category
	err = echoContext.Bind(&ent)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&ent); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := echoContext.Request().Context()
	err = a.CategoryService.Update(ctx, &ent)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusCreated, ent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	zap.L().Error("error", zap.Error(err))

	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case _interface.ErrInternalServerError:
		return http.StatusInternalServerError
	case _interface.ErrNotFound:
		return http.StatusNotFound
	case _interface.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *entitie.Category) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

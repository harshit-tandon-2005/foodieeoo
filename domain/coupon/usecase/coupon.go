package usecase

import (
	"net/http"
	"strings"

	CouponModels "github.com/foodieeoo/domain/coupon/models"
	"github.com/foodieeoo/models"
	"github.com/foodieeoo/shared/constants"
	"github.com/foodieeoo/shared/util"
	"github.com/labstack/echo"
)

func (u *couponUsecase) AddCouponCodes(c echo.Context) models.ApiUsescaseResponse {

	request := []CouponModels.AddCouponRequest{}
	if err := c.Bind(&request); err != nil {
		return models.ApiUsescaseResponse{
			Error: err,
		}
	}

	for _, r := range request {
		couponCodes, err := util.ReadGzipFileFromResource(r.Filepath)
		if err != nil {
			return models.ApiUsescaseResponse{
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
				Error:      err,
				Message:    constants.INTERNAL_SERVER_ERROR,
				ErrorCode:  "FAILED_TO_READ_FILE",
			}
		}
		couponCode := string(couponCodes)
		coupons := strings.Split(couponCode, "\n")

		err = u.insertCouponCode(coupons, r.Filename)
		if err != nil {
			return models.ApiUsescaseResponse{
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
				Error:      err,
				Message:    constants.INTERNAL_SERVER_ERROR,
				ErrorCode:  "FAILED_TO_INSERT_COUPON_CODE",
			}
		}
	}

	return models.ApiUsescaseResponse{
		StatusCode: http.StatusOK,
		Data:       nil,
		Error:      nil,
		Message:    "Successfully inserted coupon codes",
		ErrorCode:  "SUCCESS",
	}
}

// Insert Coupon Code In Batches of 1000
func (u *couponUsecase) insertCouponCode(couponCode []string, filename string) error {
	c := []models.Coupon{}
	batchSize := 1000
	tx := u.db.Begin()
	defer tx.Rollback()
	for index, coupon := range couponCode {
		c = append(c, models.Coupon{
			Code:     coupon,
			Filename: filename,
			IsActive: true,
		})
		if len(c) == batchSize || index == len(couponCode)-1 {
			err := u.couponRepo.CreateCoupons(tx, c)
			if err != nil {
				tx.Rollback()
				return err
			}
			c = []models.Coupon{}
		}
	}
	err := tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

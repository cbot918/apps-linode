package internal

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	S *Service
}

func NewController() *Controller {

	return &Controller{
		S: NewService(),
	}
}

func (ctr *Controller) Singup(c *fiber.Ctx) error {
	// handle request
	req := &SignupRequest{}
	if err := c.BodyParser(req); err != nil {
		return err
	}

	err := ctr.S.SignupService(req)
	if err != nil {
		return err
	}

	// handle response
	res := &SignupResponse{
		Email: req.Email,
		Name:  req.Name,
	}
	return c.JSON(res)
}

func (ctr *Controller) Insert(c *fiber.Ctx) error {
	return ctr.S.InsertService()
}

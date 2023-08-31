package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/cbot918/liby/cmdy"
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

func download(oldUrl string) (string, error) {

	newUrl := strings.Replace(oldUrl, "tree/main", "trunk", 1)

	folderName := strings.Trim(regexp.MustCompile(`/trunk/(.*)$`).FindString(newUrl), "/trunk/")

	cmd := cmdy.New()
	cmd0 := fmt.Sprintf("apk add subversion ")
	cmd1 := fmt.Sprintf("svn checkout %s", newUrl)
	cmd2 := fmt.Sprintf("tar -cvf %s.tar %s", folderName, folderName)
	cmd3 := fmt.Sprintf("mv %s.tar files", folderName)
	cmd4 := fmt.Sprintf("rm -rf %s", folderName)

	err := cmd.Run([]string{cmd0})
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run([]string{cmd1})
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run([]string{cmd2})
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run([]string{cmd3})
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run([]string{cmd4})
	if err != nil {
		log.Fatal(err)
	}

	return folderName, nil
}

func (ctr *Controller) GetSub(c *fiber.Ctx) error {

	// oldUrl := c.Params("u")
	oldUrl := c.Queries()["url"]
	fmt.Println(oldUrl)
	name, err := download(oldUrl)
	if err != nil {
		fmt.Println("download error")
		return err
	}

	fmt.Println("name: ", name)

	// return nil
	path := fmt.Sprintf("./files/%s.tar", name)
	nname := fmt.Sprintf("%s.tar", name)
	return c.Download(path, nname)

	// handle request
	// req := &GetSubRequest{}
	// if err := c.BodyParser(req); err != nil {
	// 	return err
	// }

	// err := ctr.S.GetSubService(req)
	// if err != nil {
	// 	return err
	// }

	// // handle response
	// res := &GetSubResponse{
	// 	Email: req.Email,
	// 	Name:  req.Name,
	// }
	// return c.JSON(res)
}

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
	fmt.Println(newUrl)
	// folderName := strings.Trim(regexp.MustCompile(`/trunk/(.*)$`).FindString(newUrl), "/trunk/")

	parts := strings.Split(regexp.MustCompile(`trunk\/.*`).FindString(newUrl), "/")
	folderName := parts[1]

	cmd := cmdy.New()
	// 技術債
	cmd0 := fmt.Sprintf("apk add subversion")
	cmd1 := fmt.Sprintf("svn checkout %s", newUrl)
	cmd2 := fmt.Sprintf("tar -cvf %s.tar %s", folderName, folderName)
	cmd3 := fmt.Sprintf("mv %s.tar files", folderName)
	cmd4 := fmt.Sprintf("rm -rf %s", folderName)
	fmt.Println("cmd0: " + cmd0)
	err := cmd.Run([]string{cmd0})
	if err != nil {
		fmt.Println("cmd0 error")
		return "", err
	}

	fmt.Println("cmd1: " + cmd1)
	err = cmd.Run([]string{cmd1})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cmd2: " + cmd2)
	err = cmd.Run([]string{cmd2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cmd3: " + cmd3)
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
	fmt.Println("path: " + path)
	fmt.Println("nname: " + nname)
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

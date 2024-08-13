package users

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateHandler(c *gin.Context)
	ReadHandler(c *gin.Context)
	ReadAllHandler(c *gin.Context)
	UpdateHandler(c *gin.Context)
	DeleteHandler(c *gin.Context)
}

type UserHandler struct {
	userService IUserService
}

func NewUserHandler(userSercvice IUserService) IUserHandler {
	return &UserHandler{
		userService: userSercvice,
	}
}

func (h *UserHandler) CreateHandler(c *gin.Context) {
	var user UserModel
	if c.ShouldBind(&user) == nil {
		log.Println(user.Username)
		log.Println(user.ID)
	}

	user, err := h.userService.Create(user)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success create data",
		"data":    user,
	})
}

func (h *UserHandler) ReadHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  false,
			"message": "Error when parse",
		})
		return
	}

	user, err := h.userService.Read(int64(id))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    user,
	})
}

func (h *UserHandler) ReadAllHandler(c *gin.Context) {
	users, err := h.userService.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "ReadHandler",
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    users,
	})
}

func (h *UserHandler) UpdateHandler(c *gin.Context) {
	var user UserModel
	if c.ShouldBind(&user) == nil {
		log.Println(user.Username)
		log.Println(user.ID)
	}

	user, err := h.userService.Update(user)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success Update",
		"data":    user,
	})
}

func (h *UserHandler) DeleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  false,
			"message": "Error when parse",
		})
		return
	}

	err = h.userService.Delete(int64(id))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Success delete data",
	})
}

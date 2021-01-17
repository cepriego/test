package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Repo Postgresrepo
}

func NewHandlers(repo Postgresrepo) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}

func (h *Handlers) InitRoutes(router *gin.Engine) {
	router.GET("/getUsers", h.getUsers)
	router.GET("/getUser", h.getUser)
	router.GET("/getGroup", h.getGroup)
	router.GET("/getGroupAndUsers", h.GetGroupsUsers)
	router.POST("/newUser", h.getUsers)
	router.POST("/newGroup", h.CreateGroup)
	router.POST("/assignToGroup", h.AssignUserToGroup)
	router.POST("/deleteUser", h.DeleteUser)
	router.POST("/deleteUserFromGroup", h.DeleteUserFromGroup)
	router.POST("/deleteGroup", h.DeleteGroup)

	router.NoRoute(h.notFound)
}

func (h *Handlers) notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}

func (h *Handlers) getUsers(c *gin.Context) {
	users, err := h.Repo.GetUsers()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": users,
	})
}

//Gets the user specidied in the body of the request
func (h *Handlers) getUser(c *gin.Context) {
	var user Usuario
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	userRetrieved, err := h.Repo.GetUser(user.Id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": userRetrieved,
	})
}

func (h *Handlers) CreateUser(c *gin.Context) {
	var user Usuario
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	saved, err := h.Repo.CreateUser(user)
	if err != nil {
		return
	}

	if !saved {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error creating user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "user created",
		})
	}
}

func (h *Handlers) CreateGroup(c *gin.Context) {
	var grupo Grupo
	err := c.BindJSON(&grupo)
	if err != nil {
		return
	}
	saved, err := h.Repo.CreateGroup(grupo)
	if err != nil {
		return
	}

	if !saved {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error creating group",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "group created",
		})
	}
}

func (h *Handlers) getGroup(c *gin.Context) {
	var user Usuario
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	groupRetrieved, err := h.Repo.GetGroup(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": groupRetrieved,
	})

}

func (h *Handlers) AssignUserToGroup(c *gin.Context) {
	var usGr UsuarioGrupo
	err := c.BindJSON(&usGr)
	if err != nil {
		return
	}
	saved, err := h.Repo.AssignUserToGroup(usGr)
	if err != nil {
		return
	}

	if !saved {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error creating user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "user assigned to group",
		})
	}
}

func (h *Handlers) GetGroupsUsers(c *gin.Context) {
	groupRetrieved, err := h.Repo.GetGroupsAndUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": groupRetrieved,
	})
}

func (h *Handlers) DeleteUser(c *gin.Context) {
	var user Usuario
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	saved, err := h.Repo.DeleteUser(user.Id)
	if err != nil {
		return
	}

	if !saved {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error creating group",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "group created",
		})
	}
}

func (h *Handlers) DeleteUserFromGroup(c *gin.Context) {
	var usrgrp UsuarioGrupo
	err := c.BindJSON(&usrgrp)
	if err != nil {
		return
	}
	saved, err := h.Repo.DeleteUserFromGroup(usrgrp.IdUsuario, usrgrp.IdGrupo)
	if err != nil {
		return
	}

	if !saved {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error creating group",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "group created",
		})
	}
}

//Hides group
func (h *Handlers) DeleteGroup(c *gin.Context) {
	var grupo Grupo
	err := c.BindJSON(&grupo)
	if err != nil {
		return
	}
	saved, err := h.Repo.DeleteGroup(grupo.Id)
	if err != nil {
		return
	}

	if !saved {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error creating group",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "group deleted",
		})
	}
}

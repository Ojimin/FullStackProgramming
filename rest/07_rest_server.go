package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RestServer() {
	router := gin.Default()
	handler := NewMembershipHandler()

	router.POST("/membership_api/:member_id", handler.Create)
	router.GET("/membership_api/:member_id", handler.Read)
	router.PUT("/membership_api/:member_id", handler.Update)
	router.DELETE("/membership_api/:member_id", handler.Delete)
	router.Run(":5000")
}

type MembershipHandler struct {
	database map[string]string
}

func NewMembershipHandler() *MembershipHandler {
	return &MembershipHandler{
		database: make(map[string]string), //데이터베이스
	}
}

func (h *MembershipHandler) Create(c *gin.Context) {
	id := c.Param("member_id")
	val := c.PostForm(id)
	if _, exists := h.database[id]; exists {
		c.JSON(http.StatusOK, gin.H{id: "None"})
		return
	}
	h.database[id] = val
	c.JSON(http.StatusOK, gin.H{id: val})
}

func (h *MembershipHandler) Read(c *gin.Context) {
	id := c.Param("member_id")
	if val, exists := h.database[id]; exists {
		c.JSON(http.StatusOK, gin.H{id: val})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "None"})
	}
}

func (h *MembershipHandler) Update(c *gin.Context) {
	id := c.Param("member_id")
	val := c.PostForm(id)
	_, exists := h.database[id]
	if !exists { //기존에 없는 member id이면 none 출력
		c.JSON(http.StatusOK, gin.H{id: "None"})
		return
	}
    if val != "" {
        h.database[id] = val
        c.JSON(http.StatusOK, gin.H{id: val})
    } else {
        c.JSON(http.StatusOK, gin.H{id: "None"})
    }
}

func (h *MembershipHandler) Delete(c *gin.Context) {
	id := c.Param("member_id")
	if _, exists := h.database[id]; exists {
		delete(h.database, id)
		c.JSON(http.StatusOK, gin.H{id: "Removed"})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "None"})
	}
}
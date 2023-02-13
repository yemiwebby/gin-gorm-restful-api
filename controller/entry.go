package controller

import (
	"diary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntry(context *gin.Context) {
	var input model.Entry
	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// user, err := helper.CurrentUser(context)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// user has been set in auth middleware
	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured internally"})
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
	// user, err := helper.CurrentUser(context)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// user has been set in auth middleware
	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured internally"})
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

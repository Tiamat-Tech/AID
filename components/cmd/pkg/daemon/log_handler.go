// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getLogs : Get /logs -> returns all logs
func getLogs(c *gin.Context) {
	logs := entities.FetchLogs()
	c.JSON(http.StatusOK, logs)
}

// getlog : Get /logs/:logid -> returns the specified log
func getLog(c *gin.Context) {
	requestedID, err := strconv.Atoi(c.Param("logid"))
	if err != nil {
		utilities.CheckError(err, "cannot convert the string")
	}
	requestedFilename := entities.GetLog(requestedID).Filepath
	fileContent := utilities.ReadFileContent(requestedFilename)
	c.JSON(http.StatusOK, logContent{
		Content: fileContent,
	})
}

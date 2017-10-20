package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"workspace/SpiderService/models"
	"workspace/SpiderService/utils"
)

// 查找或添加 pname 并返回
func preparePublishName(name, url, refer string) (pName *models.PublishName, err error) {
	pName = &models.PublishName{Name: name, InsertTime: time.Now(), Icon: sql.NullString{String: "", Valid: false}}
	err = pName.NameSelect()
	_icon := ""
	if err != nil {
		if _, ok := err.(models.NoSuchObjectError); !ok { // 不是查询不到pname的错误
			return
		} else {
			err = nil
		}
		if len(url) > 10 { // 如果url链接有效，则下载该图片并上传至oss，更新pName的Icon
			_icon, err = utils.UploadImage(url, refer)
			if err != nil {
				log.Warn(err)
				err = nil
			} else {
				pName.Icon.String = _icon
				pName.Icon.Valid = true
			}
		}
		err = pName.Store()                // 即使icon下载不成功pname也要存储
		if err != nil && len(_icon) > 10 { // 若图像上传成功而pname存储失败则删除图像
			arr := strings.Split(_icon, "/")
			err := utils.DeleteImage(arr[len(arr)-1]) // 保留外部 err 的状态
			if err != nil {
				log.Warn(err)
			}
		}
	}
	return pName, err
}

// 存储新闻并处理其中pname的icon
func StoreNewsHandler(c *gin.Context) {
	var an models.ApiNews
	if err := c.BindJSON(&an); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	news := an.ToNews()
	pName, err := preparePublishName(news.PublishName, an.SiteIcon, news.PublishUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	news.Icon = pName.Icon
	err = news.Store()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	nnews := an.ToNNews(news.Nid)
	nnews.Icon = pName.Icon
	nnews.ChannelID = news.ChannelID
	nnews.SecondChannelID = news.SecondChannelID
	nnews.Style = news.Style
	err = nnews.Store()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		err = news.Delete()
		if err != nil {
			log.Error(err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": news.Nid})
}

// 存储视频并处理其中pname的icon
func StoreVideoHandler(c *gin.Context) {
	var av models.ApiVideo
	if err := c.BindJSON(&av); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	news := av.ToNews()
	pName, err := preparePublishName(news.PublishName, av.SiteIcon, news.PublishUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	news.Icon = pName.Icon
	err = news.Store()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	video := av.ToVideo(news.Nid)
	video.Icon = pName.Icon
	err = video.Store()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		err = news.Delete()
		if err != nil {
			log.Error(err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": news.Nid})
}

// 存储段子并处理其中 pname 的 icon
func StoreJokeHandler(c *gin.Context) {
	var aj models.ApiJoke
	if err := c.BindJSON(&aj); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	news := aj.ToNews()
	pName, err := preparePublishName(news.PublishName, aj.SiteIcon, news.PublishUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	news.Icon = pName.Icon
	err = news.Store()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	joke := aj.ToJoke(news.Nid)
	joke.Icon = pName.Icon
	err = joke.Store()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		err = news.Delete()
		if err != nil {
			log.Error(err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": news.Nid})
}

// 存储评论
func StoreCommentHandler(c *gin.Context) {
	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := comment.Store(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id": comment.ID})
	}
}

// 根据 docid 或 nid 更新评论数
func UpdateCommentHandler(c *gin.Context) {
	docid, ok := c.GetPostForm("docid")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	n, ok := c.GetPostForm("n")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	number, err := strconv.Atoi(n)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = models.CommentNumberUpdate(docid, number)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, map[string]string{"message": "success"})
	}
}

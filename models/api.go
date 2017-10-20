package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx/types"
)

const (
	ZeroStyle   = 0
	OneStyle    = 1
	ThreeStyle  = 3
	VideoStyle  = 6
	VideoType   = 6
	JokeType    = 8
	JokeChannel = 45
)

type ApiNews struct {
	Title       string         `json:"title" binding:"required"`
	UniqueID    string         `json:"unique_id" binding:"required"`
	PublishUrl  string         `json:"publish_url" binding:"required"`
	PublishSite string         `json:"publish_site" binding:"required"`
	PublishTime time.Time      `json:"publish_time" binding:"required"`
	InsertTime  time.Time      `json:"insert_time" binding:"required"`
	Author      string         `json:"author" binding:"omitempty"`
	AuthorIcon  string         `json:"author_icon" binding:"omitempty"`
	SiteIcon    string         `json:"site_icon" binding:"exists"`
	Images      []string       `json:"images" binding:"exists"`
	Province    string         `json:"province" binding:"omitempty"`
	City        string         `json:"city" binding:"omitempty"`
	District    string         `json:"district" binding:"omitempty"`
	SourceID    int            `json:"source_id" binding:"omitempty"`
	Online      bool           `json:"online" binding:"exists"`
	Content     types.JSONText `json:"content" binding:"required"`
	ImageNumber int            `json:"image_number" binding:"exists"`
	Tags        []string       `json:"tags" json:"tags" binding:"exists"`
	Like        int            `json:"like" binding:"exists"`
	Dislike     int            `json:"dislike" binding:"exists"`

	ChannelID       int `json:"channel_id" binding:"omitempty"`
	SecondChannelID int `json:"second_channel_id" binding:"omitempty"`
	Read            int `json:"read" binding:"omitempty"`

	// add for trace spider source info
	SpiderSourceID string `json:"spider_source_id" binding:"omitempty"`
}

func (an ApiNews) ToNews() (news *News) {
	var style int
	switch len(an.Images) {
	case 1, 2:
		style = OneStyle
	case 3:
		style = ThreeStyle
	default:
		style = ZeroStyle
	}
	offline := 1
	if an.Online {
		offline = 0
	}
	source := Source{ID: int64(an.SourceID), ChannelID: an.ChannelID, SecondChannelID: sql.NullInt64{Int64: int64(an.SecondChannelID), Valid: true}}
	if an.ChannelID == 0 { // 如果没有上传 channel 信息，则通过 source id 查询相关的 channel 信息
		err := source.Select()
		if err != nil {
			offline = 1
		}
	} else {
		if an.SecondChannelID == 0 {
			source.SecondChannelID.Valid = false
		}
	}
	icon := sql.NullString{String: an.SiteIcon, Valid: true}
	news = &News{
		Url:             an.PublishUrl,
		UniqueID:        an.UniqueID,
		Title:           an.Title,
		Content:         an.Content,
		Author:          an.Author,
		PublishTime:     an.PublishTime,
		PublishName:     an.PublishSite,
		PublishUrl:      an.PublishUrl,
		Tags:            an.Tags,
		Province:        an.Province,
		City:            an.City,
		District:        an.District,
		ImageNumber:     an.ImageNumber,
		Style:           style,
		Images:          an.Images,
		InsertTime:      an.InsertTime,
		ChannelID:       source.ChannelID,
		SecondChannelID: int(source.SecondChannelID.Int64),
		SourceID:        int(source.ID),
		SourceState:     1,
		Icon:            icon,
		Offline:         offline,
		Concern:         an.Like,
		Unconcern:       an.Dislike,
		ClickTimes:      an.Read,
		SpiderSourceID:  an.SpiderSourceID,
	}
	return news
}

func (an ApiNews) ToNNews(nid int64) *NNews {
	return &NNews{
		Nid:         nid,
		Url:         an.PublishUrl,
		UniqueID:    an.UniqueID,
		Title:       an.Title,
		Content:     an.Content,
		Author:      an.Author,
		PublishTime: an.PublishTime,
		PublishName: an.PublishSite,
		PublishUrl:  an.PublishUrl,
		Tags:        an.Tags,
		Province:    an.Province,
		City:        an.City,
		District:    an.District,
		ImageNumber: an.ImageNumber,
		Images:      an.Images,
		InsertTime:  an.InsertTime,
		SourceID:    an.SourceID,
	}
}

type ApiVideo struct {
	Title           string    `json:"title" binding:"required"`
	UniqueID        string    `json:"unique_id" binding:"required"`
	PublishUrl      string    `json:"publish_url" binding:"required"`
	PublishSite     string    `json:"publish_site" binding:"required"`
	PublishTime     time.Time `json:"publish_time" binding:"required"`
	InsertTime      time.Time `json:"insert_time" binding:"required"`
	Author          string    `json:"author" binding:"omitempty"`
	AuthorIcon      string    `json:"author_icon" binding:"omitempty"`
	SiteIcon        string    `json:"site_icon" binding:"exists"`
	ChannelID       int       `json:"channel_id" binding:"required"`
	SecondChannelID int       `json:"second_channel_id" binding:"exists"`
	SourceID        int       `json:"source_id" binding:"omitempty"`
	Online          bool      `json:"online" binding:"exists"`
	VideoUrl        string    `json:"video_url" binding:"required"`
	VideoThumbnail  string    `json:"video_thumbnail" binding:"required"`
	VideoDuration   int       `json:"video_duration" binding:"exists"`
	PlayTimes       int       `json:"play_times" binding:"exists"`
	Tags            []string  `json:"tags" json:"tags" binding:"omitempty"`
	Like            int       `json:"like" binding:"omitempty"`
	Dislike         int       `json:"dislike" binding:"omitempty"`
	Comment         int       `json:"comment" binding:"omitempty"`

	// add for trace spider source info
	SpiderSourceID string `json:"spider_source_id" binding:"omitempty"`
}

func (av ApiVideo) ToNews() (news *News) {
	offline := 1
	if av.Online {
		offline = 0
	}
	icon := sql.NullString{String: av.SiteIcon, Valid: true}
	news = &News{
		Url:             av.PublishUrl,
		UniqueID:        av.UniqueID,
		Title:           av.Title,
		Author:          av.Author,
		PublishTime:     av.PublishTime,
		PublishName:     av.PublishSite,
		PublishUrl:      av.PublishUrl,
		Style:           VideoStyle,
		Offline:         offline,
		InsertTime:      av.InsertTime,
		ChannelID:       av.ChannelID,
		SecondChannelID: av.SecondChannelID,
		SourceID:        av.SourceID,
		SourceState:     0,
		Icon:            icon,
		VideoUrl:        av.VideoUrl,
		Thumbnail:       av.VideoThumbnail,
		Duration:        av.VideoDuration,
		ReturnType:      VideoType,
		Content:         types.JSONText("[]"),
		ClickTimes:      av.PlayTimes,
		Tags:            av.Tags,
		Concern:         av.Like,
		Unconcern:       av.Dislike,
		Comment:         av.Comment,
		SpiderSourceID:  av.SpiderSourceID,
	}
	return news
}

func (av ApiVideo) ToVideo(nid int64) (video *Video) {
	images := []string{av.VideoThumbnail}
	icon := sql.NullString{String: av.SiteIcon, Valid: true}
	video = &Video{
		Nid:             nid,
		Url:             av.PublishUrl,
		UniqueID:        av.UniqueID,
		Title:           av.Title,
		Author:          av.Author,
		PublishTime:     av.PublishTime,
		PublishName:     av.PublishSite,
		Style:           VideoStyle,
		Images:          images,
		InsertTime:      av.InsertTime,
		ChannelID:       av.ChannelID,
		SourceID:        av.SourceID,
		SecondChannelID: av.SecondChannelID,
		Icon:            icon,
		VideoUrl:        av.VideoUrl,
		Duration:        av.VideoDuration,
		Tags:            av.Tags,
	}
	return video
}

type ApiJoke struct {
	Title       string         `json:"title" binding:"required"`
	UniqueID    string         `json:"unique_id" binding:"required"`
	PublishSite string         `json:"publish_site" binding:"required"`
	PublishTime time.Time      `json:"publish_time" binding:"required"`
	InsertTime  time.Time      `json:"insert_time" binding:"required"`
	Author      string         `json:"author" binding:"omitempty"`
	AuthorIcon  string         `json:"author_icon" binding:"omitempty"`
	SiteIcon    string         `json:"site_icon" binding:"exists"`
	SourceID    int            `json:"source_id" binding:"omitempty"`
	Online      bool           `json:"online" binding:"exists"`
	Content     types.JSONText `json:"content" binding:"required"`
	Like        int            `json:"like" binding:"exists"`
	Dislike     int            `json:"dislike" binding:"exists"`
	Comment     int            `json:"comment" binding:"exists"`

	Style       int      `json:"style" binding:"omitempty"` // 20170809 for new channel joke
	ImageNumber int      `json:"image_number" binding:"omitempty"`
	Images      []string `json:"images" binding:"omitempty"`

	// add for trace spider source info
	SpiderSourceID string `json:"spider_source_id" binding:"omitempty"`
}

func (aj ApiJoke) ToNews() (news *News) {
	offline := 1
	if aj.Online {
		offline = 0
	}
	icon := sql.NullString{String: aj.SiteIcon, Valid: true}
	news = &News{
		Url:            aj.UniqueID,
		UniqueID:       aj.UniqueID,
		Title:          aj.Title,
		Author:         aj.Author,
		PublishTime:    aj.PublishTime,
		PublishName:    aj.PublishSite,
		Style:          aj.Style,
		Offline:        offline,
		InsertTime:     aj.InsertTime,
		ChannelID:      JokeChannel,
		SourceID:       aj.SourceID,
		SourceState:    0,
		Icon:           icon,
		ReturnType:     JokeType,
		Content:        types.JSONText("[]"),
		Concern:        aj.Like,
		Unconcern:      aj.Dislike,
		ImageNumber:    aj.ImageNumber,
		Images:         aj.Images,
		SpiderSourceID: aj.SpiderSourceID,
	}
	return news
}

func (aj ApiJoke) ToJoke(nid int64) *Joke {
	icon := sql.NullString{String: aj.SiteIcon, Valid: true}
	return &Joke{
		Nid:         nid,
		UniqueID:    aj.UniqueID,
		Content:     aj.Content,
		Author:      aj.Author,
		Avatar:      aj.AuthorIcon,
		PublishTime: aj.PublishTime,
		PublishName: aj.PublishSite,
		Style:       aj.Style,
		InsertTime:  aj.InsertTime,
		ChannelID:   JokeChannel,
		SourceID:    aj.SourceID,
		Icon:        icon,
		Images:      aj.Images,
	}
}

package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"

	S "workspace/SpiderService/storage"
)

type News struct {
	Nid             int64          `db:"nid" json:"nid"`
	Url             string         `db:"url" json:"url"`
	UniqueID        string         `db:"docid" json:"unique_id"`
	Title           string         `db:"title" json:"title"`
	Content         types.JSONText `db:"content" json:"content"`
	Author          string         `db:"author" json:"author"`
	PublishTime     time.Time      `db:"ptime" json:"publish_time"`
	PublishName     string         `db:"pname" json:"publish_name"`
	PublishUrl      string         `db:"purl" json:"publish_url"`
	Tags            pq.StringArray `db:"tags" json:"tags"`
	Province        string         `db:"province" json:"province"`
	City            string         `db:"city" json:"city"`
	District        string         `db:"district" json:"district"`
	ImageNumber     int            `db:"inum" json:"image_number"`
	Style           int            `db:"style" json:"style"`
	Images          pq.StringArray `db:"imgs" json:"images"`
	InsertTime      time.Time      `db:"ctime" json:"insert_time"`
	ChannelID       int            `db:"chid" json:"channel_id"`
	SecondChannelID int            `db:"sechid" json:"second_channel_id"`
	SourceID        int            `db:"srid" json:"source_id"`
	Icon            sql.NullString `db:"icon" json:"icon"`
	Html            string         `db:"html" json:"html"`
	Collect         int            `db:"collect" json:"collect"`
	Concern         int            `db:"concern" json:"concern"`
	Comment         int            `db:"comment" json:"comment"`
	Offline         int            `db:"state" json:"offline"`
	SourceState     int            `db:"srstate" json:"source_state"`
	VideoUrl        string         `db:"videourl" json:"video_url"`
	// VideoType int
	Thumbnail  string `db:"thumbnail" json:"thumbnail"`
	Duration   int    `db:"duration" json:"duration"`
	ReturnType int    `db:"rtype" json:"return_type"`
	Unconcern  int    `db:"un_concern" json:"unconcern"`
	ClickTimes int    `db:"clicktimes" json:"click_times"`

	// add for trace spider source info
	SpiderSourceID string `db:"sec_purl_id" json:"spider_source_id"`
}

func (news *News) Store() (err error) {
	table := `newslist_v2`
	fields := []string{"url", "docid", "title", "author", "ptime", "pname", "purl",
		"tags", "province", "city", "district", "comment", "inum", "style", "imgs", "state", "ctime", "chid",
		"sechid", "srid", "icon", "videourl", "thumbnail", "duration", "rtype", "concern", "un_concern",
		"clicktimes", "srstate", "sec_purl_id",
	}
	returns := []string{"nid"}
	statement := FormatInsertStatement(table, fields, returns)
	var id int64
	var sechid interface{}
	if news.SecondChannelID == 0 {
		sechid = nil
	} else {
		sechid = news.SecondChannelID
	}
	err = S.PG.QueryRowx(statement, news.Url, news.UniqueID, news.Title, news.Author,
		news.PublishTime, news.PublishName, news.PublishUrl, news.Tags, news.Province,
		news.City, news.District, news.Comment, news.ImageNumber, news.Style, news.Images, news.Offline,
		news.InsertTime, news.ChannelID, sechid, news.SourceID, news.Icon, news.VideoUrl, news.Thumbnail,
		news.Duration, news.ReturnType, news.Concern, news.Unconcern, news.ClickTimes,
		news.SourceState, news.SpiderSourceID).Scan(&id)
	if err == nil {
		news.Nid = id
	}
	return err
}

func (news *News) Delete() (err error) {
	statement := `DELETE FROM newslist_v2 where nid=$1;`
	_, err = S.PG.Exec(statement, news.Nid)
	return err
}

type NNews struct {
	Nid             int64          `db:"nid" json:"nid"`
	Url             string         `db:"url" json:"url"`
	UniqueID        string         `db:"docid" json:"unique_id"`
	Title           string         `db:"title" json:"title"`
	Content         types.JSONText `db:"content" json:"content"`
	Author          string         `db:"author" json:"author"`
	PublishTime     time.Time      `db:"ptime" json:"publish_time"`
	PublishName     string         `db:"pname" json:"publish_name"`
	PublishUrl      string         `db:"purl" json:"publish_url"`
	Tags            pq.StringArray `db:"tags" json:"tags"`
	Province        string         `db:"province" json:"province"`
	City            string         `db:"city" json:"city"`
	District        string         `db:"district" json:"district"`
	ImageNumber     int            `db:"inum" json:"image_number"`
	Style           int            `db:"style" json:"style"`
	Images          pq.StringArray `db:"imgs" json:"images"`
	InsertTime      time.Time      `db:"ctime" json:"insert_time"`
	ChannelID       int            `db:"chid" json:"channel_id"`
	SecondChannelID int            `db:"sechid" json:"second_channel_id"`
	SourceID        int            `db:"srid" json:"source_id"`
	Icon            sql.NullString `db:"icon" json:"icon"`
}

func (n *NNews) Store() (err error) {
	table := `info_news`
	fields := []string{"nid", "url", "docid", "title", "content", "author", "ptime", "pname", "purl", "tags",
		"province", "city", "district", "inum", "style", "imgs", "ctime", "chid", "sechid", "srid", "icon"}
	returns := []string{"nid"}
	statement := FormatInsertStatement(table, fields, returns)
	var sechid interface{}
	if n.SecondChannelID == 0 {
		sechid = nil
	} else {
		sechid = n.SecondChannelID
	}
	err = S.PG.QueryRowx(statement, n.Nid, n.Url, n.UniqueID, n.Title, n.Content, n.Author, n.PublishTime,
		n.PublishName, n.PublishUrl, n.Tags, n.Province, n.City, n.District, n.ImageNumber, n.Style, n.Images,
		n.InsertTime, n.ChannelID, sechid, n.SourceID, n.Icon).Scan(&n.Nid)
	return err
}

func (n *NNews) Delete() (err error) {
	statement := `DELETE FROM info_news where nid=$1;`
	_, err = S.PG.Exec(statement, n.Nid)
	return err
}

type Video struct {
	Nid             int64          `db:"nid" json:"nid"`
	Url             string         `db:"url" json:"url"`
	UniqueID        string         `db:"docid" json:"unique_id"`
	Title           string         `db:"title" json:"title"`
	Author          string         `db:"author" json:"author"`
	PublishTime     time.Time      `db:"ptime" json:"publish_time"`
	PublishName     string         `db:"pname" json:"publish_name"`
	Style           int            `db:"style" json:"style"`
	Images          pq.StringArray `db:"imgs" json:"images"`
	InsertTime      time.Time      `db:"ctime" json:"insert_time"`
	ChannelID       int            `db:"chid" json:"channel_id"`
	SourceID        int            `db:"srid" json:"source_id"`
	SecondChannelID int            `db:"sechid" json:"second_channel_id"`
	Icon            sql.NullString `db:"icon" json:"icon"`
	VideoUrl        string         `db:"videourl" json:"video_url"`
	Duration        int            `db:"duration" json:"duration"`
	Tags            pq.StringArray `db:"tags" json:"tags"`
}

func (v *Video) Store() (err error) {
	statement := `
	INSERT INTO info_video (nid, url, docid, title, author, ptime, pname, style, imgs, ctime, chid, srid, sechid, icon, videourl, duration, tags)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	RETURNING nid
	`
	var id int64
	err = S.PG.QueryRowx(statement, v.Nid, v.Url, v.UniqueID, v.Title, v.Author, v.PublishTime, v.PublishName, v.Style, v.Images, v.InsertTime, v.ChannelID,
		v.SourceID, v.SecondChannelID, v.Icon, v.VideoUrl, v.Duration, v.Tags).Scan(&id)
	if err == nil {
		v.Nid = id
	}
	return err
}

type Joke struct {
	Nid         int64          `db:"nid" json:"nid"`
	UniqueID    string         `db:"docid" json:"unique_id"`
	Content     types.JSONText `db:"content" json:"content"`
	Author      string         `db:"author" json:"author"`
	Avatar      string         `db:"avatar" json:"avatar"`
	PublishTime time.Time      `db:"ptime" json:"publish_time"`
	PublishName string         `db:"pname" json:"publish_name"`
	Style       int            `db:"style" json:"style"`
	Images      pq.StringArray `db:"imgs" json:"images"`
	InsertTime  time.Time      `db:"ctime" json:"insert_time"`
	ChannelID   int            `db:"chid" json:"channel_id"`
	SourceID    int            `db:"srid" json:"source_id"`
	Icon        sql.NullString `db:"icon" json:"icon"`
}

func (j *Joke) Store() (err error) {
	statement := `
	INSERT INTO info_joke (nid, docid, content, author, avatar, ptime, pname, style, imgs, ctime, chid, srid, icon)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING nid
	`
	var id int64
	err = S.PG.QueryRowx(statement, j.Nid, j.UniqueID, j.Content, j.Author, j.Avatar, j.PublishTime,
		j.PublishName, j.Style, j.Images, j.InsertTime, j.ChannelID, j.SourceID, j.Icon).Scan(&id)
	if err == nil {
		j.Nid = id
	}
	return err
}

type PublishName struct {
	Id         int64          `db:"id" json:"id"`
	InsertTime time.Time      `db:"ctime" json:"insert_time"`
	Name       string         `db:"name" json:"name"`
	Icon       sql.NullString `db:"icon" json:"icon"`
}

func (p *PublishName) Store() (err error) {
	statement := `
	INSERT INTO newspublisherlist_v2 (ctime, name, icon)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var id int64
	err = S.PG.QueryRowx(statement, p.InsertTime, p.Name, p.Icon).Scan(&id)
	if err == nil {
		p.Id = id
	}
	return err
}

func (p *PublishName) NameSelect() (err error) {
	ps := []PublishName{}
	statement := `SELECT id, ctime, name, icon FROM newspublisherlist_v2 WHERE name=$1`
	err = S.PG.Select(&ps, statement, p.Name)
	if err != nil {
		return
	}
	if len(ps) == 0 {
		return NewNoSuchObjectError(fmt.Sprintf("No such pname: %s", p.Name))
	}
	p.Id = ps[0].Id
	p.InsertTime = ps[0].InsertTime
	p.Name = ps[0].Name
	p.Icon = ps[0].Icon
	return
}

type Comment struct {
	ID         int64     `db:"id" json:"id"`
	Content    string    `db:"content" json:"content" binding:"required"`
	Commend    int       `db:"commend" json:"commend" binding:"exists"`
	InsertTime time.Time `db:"ctime" json:"insert_time" binding:"required"`
	UserName   string    `db:"uname" json:"user_name" binding:"required"`
	Avatar     string    `db:"avatar" json:"avatar" binding:"exists"`
	ForeignID  string    `db:"docid" json:"foreign_id" binding:"required"` // 评论关联的外键
	UniqueID   string    `db:"cid" json:"unique_id" binding:"required"`
}

func (c *Comment) Store() (err error) {
	table := "commentlist_v2"
	fields := []string{"content", "commend", "ctime", "uname", "avatar", "docid", "cid"}
	returns := []string{"id"}
	statement := FormatInsertStatement(table, fields, returns)
	var id int64
	err = S.PG.QueryRowx(statement, c.Content, c.Commend, c.InsertTime, c.UserName, c.Avatar, c.ForeignID, c.UniqueID).Scan(&id)
	if err == nil {
		c.ID = id
	}
	return err
}

type Source struct {
	ID              int64         `db:"id" json:"id"`
	ChannelID       int           `db:"cid" json:"channel_id"`
	SecondChannelID sql.NullInt64 `db:"scid" json:"second_channel_id"`
	Name            string        `db:"sname" json:"name"`
	State           int           `db:"state" json:"state"`
}

func (s *Source) Select() (err error) {
	ss := []Source{}
	statement := `SELECT id, cid, scid, sname, state FROM sourcelist_v2 WHERE id=$1`
	err = S.PG.Select(&ss, statement, s.ID)
	if err == nil {
		if len(ss) == 0 {
			err = NoSuchObjectError{Info: "No such object in sourcelist_v2"}
		} else {
			*s = ss[0]
		}
	}
	return err
}

func FormatInsertStatement(table string, fields, returns []string) string {
	keys := strings.Join(fields, ", ")
	ids := make([]string, len(fields))
	for i := range fields {
		ids[i] = fmt.Sprintf("$%d", i+1)
	}
	values := strings.Join(ids, ", ")
	suffix := ""
	if len(returns) != 0 {
		suffix = fmt.Sprintf("RETURNING %s", strings.Join(returns, ", "))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) %s", table, keys, values, suffix)
}

// 根据 field 字段检索新闻并更新 comment 字段
func CommentNumberUpdate(docid string, n int) (err error) {
	statement := `UPDATE newslist_v2 SET comment=comment+$1 WHERE docid=$2`
	stms, err := S.PG.Prepare(statement)
	if err != nil {
		return err
	}
	_, err = stms.Exec(n, docid)
	return err
}

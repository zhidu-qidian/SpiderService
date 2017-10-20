## API 文档

api 格式说明: METHOD PATH PARAMS

METHOD: 表示http请求的方法, GET, POST, PUT 等

PATH: api的路径

PARAMS: 传入参数的方式 QUERY(正常的query参数), FORM(form表单), JSON(json数据)

所有返回的数据都是 json 格式

#### POST /api/store/news	JSON

存储news数据到postgres

== required 表示不为空字符串,0,false; exists 表示存在该字段即可; omitempty 表示可以不传 ==
- title string required
- unique_id string required
- publish_url string required
- publish_site string required
- publish_time datetime.datetime required
- insert_time datetime.datetime required
- author string omitempty
- author_icon string omitempty
- images \[str, str, ...\] exists
- province string omitempty
- city string omitempty
- district string omitempty
- source_id int omitempty
- online bool exists
- content list of dict required
- image_number int exists
- tags list of str exists
- like int exists
- dislike int exists
- channel_id int omitempty
- second_channel_id int omitempty
- read int omitempty
- spider_source_id string omitempty (该条数据的二级抓取源id)

```json
{
  "id": 18322322
}
```

#### POST /api/store/video	JSON

存储video数据到postgres

== required 表示不为空字符串,0,false; exists 表示存在该字段即可 ==
- title string required
- unique_id string required (对应docid字段)
- publish_url string required
- publish_site string required
- publish_time datetime.datetime required
- insert_time datetime.datetime required (插入时间)
- author string omitempty
- author_icon string omitempty (用户头像)
- site_icon string exists (发布源的图标)
- channel_id int required
- second_channel_id int exists
- source_id int omitempty
- online bool exists (这条数据是否上线)
- video_url string required
- video_thumbnail string required
- video_duration int exists
- play_times int exists
- tags \[str, str, ...\] omitempty
- like int omitempty
- dislike int omitempty
- comment int omitempty
- spider_source_id int omitempty (该条数据的二级抓取源id)

```json
{
  "id": 15888232
}
```

#### POST /api/store/joke	JSON

存储joke数据到postgres

- title string required
- unique_id string required
- publish_site string required
- publish_time datetime.datetime required
- insert_time datetime.datetime required
- author string omitempty
- author_icon string omitempty
- site_icon string exists
- source_id source_id omitempty
- online bool exists
- content list of dict required
- like int exists
- dislike int exists
- comment int exists
- style int omitempty (为了段子的各种不同格式,包含普通图片的,包含gif图的等)
- image_number int omitempty
- images \[str, str, ...\] omitempty
- spider_source_id string omitempty (该条数据的二级抓取源id)

```json
{
	"id": 323251235
}
```

#### POST /api/store/comment	JSON

存储评论到postgres

- content string required
- commend int exists
- insert_time datetime.datetime required
- user_name string required
- avatar string exists
- foreign_id string required (评论关联的外键)
- unique_id string required (评论的唯一性id)

```json
{
	"id": 23421324
}
```

#### PUT /api/update/comment    FORM

更新指定文章的评论数

```json
{"docid": "video_599ff0954215ef23bf0e4bc9"}
```

```json
{"message": "success"}
```

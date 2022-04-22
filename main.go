package main

import (
	"github.com/gin-gonic/gin"
)

type Comment struct {
	PostId int    `json:"postid"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

var comments = []Comment{
	{
		PostId: 1,
		Id:     1,
		Name:   "Id labore ex et quam laborum",
		Email:  "Eliseo@gardner.biz",
		Body:   "laudantium enim quasi est quIdem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium",
	},
	{
		PostId: 1,
		Id:     2,
		Name:   "quo vero reiciendis velit similique earum",
		Email:  "Jayne_Kuhic@sydney.com",
		Body:   "est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et",
	},
	{
		PostId: 1,
		Id:     3,
		Name:   "odio adipisci rerum aut animi",
		Email:  "Nikita@garfield.biz",
		Body:   "quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione",
	},
	{
		PostId: 1,
		Id:     4,
		Name:   "alias odio sit",
		Email:  "Lew@alysha.tv",
		Body:   "non et atque\noccaecati deserunt quas accusantium unde odit nobis qui voluptatem\nquia voluptas consequuntur itaque dolor\net qui rerum deleniti ut occaecati",
	},
	{
		PostId: 1,
		Id:     5,
		Name:   "vero eaque aliquId doloribus et culpa",
		Email:  "Hayden@althea.biz",
		Body:   "harum non quasi et ratione\ntempore iure ex voluptates in ratione\nharum architecto fugit inventore cupIditate\nvoluptates magni quo et",
	},
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H{
				"msg": "Hello World",
			},
		)
	})
	router.GET("/comment", func(ctx *gin.Context) {
		ctx.IndentedJSON(
			200,
			comments,
		)
	})
	router.POST("/comment", func(ctx *gin.Context) {
		var body Comment
		err := ctx.BindJSON(&body)
		if err != nil {
			return
		}
		comments = append(comments, body)
		ctx.IndentedJSON(
			201,
			comments,
		)
	})
	router.Run()
}

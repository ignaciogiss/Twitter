package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/ignaciogiss/twitter/src/domain"
	"github.com/ignaciogiss/twitter/src/service"
	"net/http"
	"strconv"
)

type RestServer interface {
	listarTweets(c * gin.Context)
	StartRestServer()
}
type RestTweetServer struct {
	tweetManager *service.TweetManager
}
type TweetJson struct {
	User string `json:"user"`
	Text string `json:"text"`
	ImgUrl string `json:"img_url"`
	QuotedId int `json:"quoted_id"`
}

func (restTweetServer RestTweetServer) listTweets(c * gin.Context) {
	for _, t := range restTweetServer.tweetManager.GetTweets() {
		c.String(http.StatusOK, t.PrintableTweet() )
	}
}

func (restTweetServer RestTweetServer) listTweetsByUser(c * gin.Context) {
	for _, t := range restTweetServer.tweetManager.GetTweetsByUser( c.Param("user") ) {
		c.String(http.StatusOK, t.PrintableTweet() )
	}
}

func (restTweetServer RestTweetServer) listTweetsById(c * gin.Context) {
	id, _ := strconv.Atoi( c.Param("id") )
	tweet := restTweetServer.tweetManager.GetTweetById( id )
	if tweet != nil {
		c.String(http.StatusOK, tweet.PrintableTweet() )
	}
}

func (restTweetServer RestTweetServer) publishTweet(c * gin.Context) {

	var tweetData TweetJson
	c.BindJSON(&tweetData)

	tweet := domain.NewTextTweet( tweetData.User, tweetData.Text )
	_, err := restTweetServer.tweetManager.PublishTweet(tweet)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not ok")
	}
	c.JSON(http.StatusOK, "ok")
}

func (restTweetServer RestTweetServer) publishQuotedTweet(c * gin.Context) {

	var tweetData TweetJson
	c.BindJSON(&tweetData)

	tweet_ori := restTweetServer.tweetManager.GetTweetById( tweetData.QuotedId )
	quotedTweet := domain.NewTextTweet(tweet_ori.GetUser(), tweet_ori.GetText())

	tweet := domain.NewQuoteTweet( tweetData.User, tweetData.Text, quotedTweet )
	_, err := restTweetServer.tweetManager.PublishTweet( tweet )
	if err != nil {
		c.JSON(http.StatusBadRequest, "not ok")
	}
	c.JSON(http.StatusOK, "ok")
}

func (restTweetServer RestTweetServer) publishImgTweet(c * gin.Context) {

	var tweetData TweetJson
	c.BindJSON(&tweetData)

	tweet := domain.NewImageTweet( tweetData.User, tweetData.Text, tweetData.ImgUrl )
	_, err := restTweetServer.tweetManager.PublishTweet(tweet)
	if err != nil {
		c.JSON(http.StatusBadRequest, "not ok")
	}
	c.JSON(http.StatusOK, "ok")
}


func NewRestTweetServer(t *service.TweetManager) *RestTweetServer {

	return &RestTweetServer{tweetManager: t}
}

func (restTweetServer RestTweetServer) StartRestServer() {
	router := gin.Default()

	router.GET("/listTweets/", restTweetServer.listTweets )
	router.GET("/listTweets/:user", restTweetServer.listTweetsByUser )
	router.GET("/listTweetsById/:id", restTweetServer.listTweetsById )

	router.POST("publishTweet/", restTweetServer.publishTweet )
	router.POST("publishQuotedTweet/", restTweetServer.publishQuotedTweet )
	router.POST("publishImgTweet/", restTweetServer.publishImgTweet )

	go router.Run()
}

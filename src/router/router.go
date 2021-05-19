package router

import (
	"forum/src/controller"
	"forum/src/middleware"

	"github.com/gin-gonic/gin"
)

func RunAPP() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/register", controller.Register) //注册
		api.POST("/login", controller.Login)       //登录
		api.GET("/section/list")                   //所有板块
		api.GET("/section/:section_id/topic/list") //板块页帖子列表
		api.GET("/search")                         //搜索
		api.GET("/topic/:topic_id/detail")         //帖子详情
		api.GET("/topic/:topic_id/comment/list")   //帖子回复列表

		user := api.Group("/user", middleware.UserAuth)
		{
			user.GET("/detail")     //用户详情
			user.PATCH("/nickname") //改用户昵称
			user.PATCH("/password") //改用户密码

			topic := user.Group("/topic")
			{
				topic.POST("/")                           //发帖
				topic.POST("/:topic_id/comment")          //回帖
				topic.GET("/:topic_id/record")            //用户某一帖子记录
				topic.PUT("/:topic_id/thumb")             //点赞贴子
				topic.DELETE("/:topic_id/thumb")          //取消点赞
				topic.PUT("/:topic_id/favor")             //收藏贴子
				topic.DELETE("/:topic_id/favor")          //取消收藏
				topic.POST("/:topic_id/tipoff")           //举报帖子
				topic.POST("/:topic_id")                  //举报回帖
				topic.POST("/comment/:comment_id/tipoff") //用户发布的贴子列表

				topic.GET("/list")         //删自己的帖
				topic.DELETE("/:topic_id") //浏览历史
			}

			user.GET("/history/list") //浏览历史
			user.GET("/thumb/list")   //点赞记录
			user.GET("/favor/list")   //收藏夹
		}

		admin := api.Group("/admin")
		{
			tipoff := admin.Group("/tipoff")
			{
				tipoff.GET("/topic/list")      //帖子举报列表
				tipoff.GET("/comment/list")    //回帖举报列表
				tipoff.PUT("/process/:tip_id") //处理举报工单
			}

			admin.DELETE("/topic/:topic_id")     //封贴
			admin.DELETE("/comment/:comment_id") //封回帖
		}
	}

	_ = r.Run("")
}

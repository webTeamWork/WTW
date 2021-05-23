package router

import (
	"forum/src/controller"
	"forum/src/middleware"
	"forum/src/model"

	"github.com/gin-gonic/gin"
)

func RunAPP() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/register", controller.Register)                                 //注册
		api.POST("/login", controller.Login)                                       //登录
		api.GET("/section/list", controller.GetAllSection)                         //所有板块
		api.GET("/section/:section_id/topic/list", controller.GetSectionTopicList) //板块页帖子列表
		api.GET("/search", controller.Search)                                      //搜索
		api.GET("/topic/:topic_id/detail", controller.GetTopicDetail)              //帖子详情
		api.GET("/topic/:topic_id/comment/list", controller.GetTopicCommentList)   //帖子回帖列表

		user := api.Group("/user", middleware.UserAuth)
		{
			user.GET("/detail", controller.GetUserDetail)          //用户详情
			user.PATCH("/nickname", controller.ChangeUserNickname) //改用户昵称
			user.PATCH("/password", controller.ChangeUserPassword) //改用户密码

			topic := user.Group("/topic")
			{
				topic.POST("/", controller.PostTopic)                               //发帖
				topic.POST("/:topic_id/comment", controller.CommentTopic)           //回帖
				topic.GET("/:topic_id/record", controller.GetUserTopicRecord)       //用户某一帖子记录
				topic.PUT("/:topic_id/thumb", controller.ThumbTopic)                //点赞贴子
				topic.DELETE("/:topic_id/thumb", controller.CancelThumbTopic)       //取消点赞
				topic.PUT("/:topic_id/favor", controller.FavorTopic)                //收藏贴子
				topic.DELETE("/:topic_id/favor", controller.CancelFavorTopic)       //取消收藏
				topic.POST("/:topic_id/tipoff", controller.TipoffTopic)             //举报帖子
				topic.POST("/comment/:comment_id/tipoff", controller.TipoffComment) //举报回帖

				topic.GET("/list", controller.GetUserTopicList) //用户发布的贴子列表
				topic.DELETE("/:topic_id")                      //删自己的帖
			}

			user.GET("/history/list", controller.GetUserRecordList(model.RecordTypeView)) //浏览历史
			user.GET("/thumb/list", controller.GetUserRecordList(model.RecordTypeThumb))  //点赞记录
			user.GET("/favor/list", controller.GetUserRecordList(model.RecordTypeFavor))  //收藏夹
		}

		admin := api.Group("/admin", middleware.AdminAuth)
		{
			tipoff := admin.Group("/tipoff")
			{
				tipoff.GET("/topic/list", controller.GetTipoffList(model.TipoffTargetTypeTopic))     //帖子举报列表
				tipoff.GET("/comment/list", controller.GetTipoffList(model.TipoffTargetTypeComment)) //回帖举报列表
				tipoff.PUT("/process/:tip_id", controller.ProcessTipoff)                             //处理举报工单
			}

			admin.DELETE("/topic/:topic_id", controller.BanTopic)       //封贴
			admin.DELETE("/comment/:comment_id", controller.BanComment) //封回帖
		}
	}

	_ = r.Run("")
}

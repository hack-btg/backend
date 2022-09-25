package models

type Social struct {
	User    string `json:"user"`
	Youtube struct {
		YoutubeChannelName          string `json:"youtube_channel_name"`
		Subscribers                 int    `json:"subscribers"`
		Views                       int    `json:"views"`
		Uploads                     int    `json:"uploads"`
		VarianceSubscribersLastYear int    `json:"variance_subscribers_last_year"`
		ViewsLastYear               int    `json:"views_last_year"`
		RankingSocialBladeYoutube   int    `json:"ranking_social_blade_youtube"`
		Grade                       string `json:"grade"`
	} `json:"youtube"`
	Facebook struct {
		FacebookUsername      string `json:"facebook_username"`
		Likes                 int    `json:"likes"`
		VarianceLikesLastYear int    `json:"variance_likes_last_year"`
		RankingSocialBlade    int    `json:"ranking_social_blade"`
		Grade                 string `json:"grade"`
	} `json:"facebook"`
	Twitter struct {
		TwitterUsername           string `json:"twitter_username"`
		Followers                 int    `json:"followers"`
		Following                 int    `json:"following"`
		Tweets                    int    `json:"tweets"`
		VarianceFollowersLastYear int    `json:"variance_followers_last_year"`
		VarianceTweetsLastYear    int    `json:"variance_tweets_last_year"`
		RankingSocialBlade        int    `json:"ranking_social_blade"`
		Grade                     string `json:"grade"`
	} `json:"twitter"`
	Instagram struct {
		InstagramUsername         string  `json:"instagram_username"`
		Followers                 int     `json:"followers"`
		Following                 int     `json:"following"`
		Media                     int     `json:"media"`
		EngagementRate            float64 `json:"engagement_rate"`
		VarianceFollowersLastYear int     `json:"variance_followers_last_year"`
		VarianceTweetsLastYear    int     `json:"variance_tweets_last_year"`
		RankingSocialBlade        int     `json:"ranking_social_blade"`
		Grade                     string  `json:"grade"`
		RankingEngagement         int     `json:"ranking_engagement"`
		AverageLikes              float64 `json:"average_likes"`
		AverageComments           float64 `json:"average_comments"`
	} `json:"instagram"`
}

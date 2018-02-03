package suggestion_service

func init() {
	go job_TopTrendingUsersCalculate_Infinite()
	go job_TopTrendingPosts_Infinite()
}

package suggestion_service

func init() {
	go job_TopTrendingUsersCalculate_Infinixte()
	go job_TopTrendingPosts_Infinite()
}

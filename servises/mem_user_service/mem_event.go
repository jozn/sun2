package mem_user_service

import "ms/sun2/servises/event_service"

func listen() {
	sub := event_service.SubParam{
		Added_Post_Event:     true,
		Deleted_Post_Event:   true,
		Followed_User_Event:  true,
		UnFollwed_User_Event: true,
	}
	listener := event_service.NewSub(sub)

	//posts add
	go func() {
		for e := range listener.Added_Post_Event {
			mu := allMemUserMap.GetForUser(e.Event.ByUserId)
			if mu != nil && len(mu.lastPosts) > 0 { //} && e.Post != nil {
				mu.isPostsLoaded = false
				//mu.lastPosts = append(mu.lastPosts, e.Post)
			}
		}
	}()

	//posts delete
	go func() {
		for e := range listener.Deleted_Post_Event {
			mu := allMemUserMap.GetForUser(e.Event.ByUserId)
			if mu != nil && len(mu.lastPosts) > 0 {
				mu.isPostsLoaded = false
				/*for idx, p := range mu.lastPosts {
					if p != nil && p.PostId == e.Event.PostId {
						mu.lastPosts[idx] = nil
					}
				}*/
			}
		}
	}()

	//followed
	cleanFollowed := func(e event_service.GeneralEvent) {
		mu1 := allMemUserMap.GetForUser(e.Event.ByUserId)
		mu1.isFollowedLoaded = false //force reload

		mu2 := allMemUserMap.GetForUser(e.Event.PeerUserId)
		mu2.isFollowersLoaded = false //force reload
	}

	go func() {
		for e := range listener.Followed_User_Event {
			cleanFollowed(e)
		}
	}()

	go func() {
		for e := range listener.UnFollwed_User_Event {
			cleanFollowed(e)
		}
	}()

}

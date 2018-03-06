package mem_user_service

import "ms/sun2/servises/event_service"

func listen() {
	sub := event_service.SubParam{
		ADDED_POST_EVENT:      true,
		DELETED_POST_EVENT:    true,
		FOLLOWED_USER_EVENT:   true,
		UNFOLLOWED_USER_EVENT: true,
	}
	listener := event_service.NewSub(sub)

	//posts add
	go func() {
		for e := range listener.ADDED_POST_EVENT {
			mu := AllMemUserMap.GetForUser(e.Event.ByUserId)
			if mu != nil && len(mu.lastPosts) > 0 { //} && e.Post != nil {
				mu.isPostsLoaded = false
				//mu.lastPosts = append(mu.lastPosts, e.Post)
			}
		}
	}()

	//posts delete
	go func() {
		for e := range listener.DELETED_POST_EVENT {
			mu := AllMemUserMap.GetForUser(e.Event.ByUserId)
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
		mu1 := AllMemUserMap.GetForUser(e.Event.ByUserId)
		mu1.isFollowedLoaded = false //force reload

		mu2 := AllMemUserMap.GetForUser(e.Event.PeerUserId)
		mu2.isFollowersLoaded = false //force reload
	}

	go func() {
		for e := range listener.FOLLOWED_USER_EVENT {
			cleanFollowed(e)
		}
	}()

	go func() {
		for e := range listener.UNFOLLOWED_USER_EVENT {
			cleanFollowed(e)
		}
	}()

}

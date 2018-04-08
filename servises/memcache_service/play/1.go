package main

import "ms/sun/servises/memcache_service"

func main()  {
    for i:=0;i<1000 ; i++ {
        r:=memcache_service.DidUserLikedPost(6,i)
        print(r)
    }
}




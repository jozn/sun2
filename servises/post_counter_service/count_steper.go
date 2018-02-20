package post_counter_service

import (
    "ms/sun2/shared/x"
    "ms/sun/base"
    "fmt"
    "ms/sun2/shared/config"
    "github.com/labstack/gommon/log"
)

type postCount struct {
	postId    int
	seenCount int
}

type stepper struct {
	incrSize int
	postIds  []int
}

func (s *stepper)proccedPost(postCnt *postCount)  {
    if postCnt.seenCount >= s.incrSize {
        postCnt.seenCount -= s.incrSize
        s.postIds = append(s.postIds,postCnt.postId)
    }
}

//todo in futuer limit len of postIds to 10000 each update to not face limitaions of databases length
func (s *stepper)updatePostsViewsCounts()  {
    if len(s.postIds) > 0 {
        x.NewPost_Updater().PostId_In(s.postIds).ViewsCount_Increment(s.incrSize).Update(base.DB)
    }
}
///////////////////////////////////
func stepperProceedMap(mp map[int]int) {
	if len(mp) == 0 {
		return
	}
    fmt.Println("stepperProceedMap: ", mp)

	arrCnts := make([]*postCount, 0, len(mp))
	for pid, cnt := range mp {
		s := &postCount{
			postId:    pid,
			seenCount: cnt,
		}
		arrCnts = append(arrCnts, s)
	}
    steps := getAllSteppers()
    for _, postCnt := range arrCnts {
        for _, step := range steps {
            step.proccedPost(postCnt)
        }
    }

    for _, step := range steps {
        step.updatePostsViewsCounts()
    }

    //check data integrity
    if config.IS_DEBUG{
        for _, postCnt := range arrCnts {
            if postCnt.seenCount != 0 {
                log.Panic("in post_counter_service post seen count is not 0 after procceding via stepper.", postCnt.postId)
            }
        }
    }
}

func getAllSteppers() (steps []*stepper) {
	incers := []int{100000, 50000, 30000, 20000, 10000, 5000, 3000, 2000, 1000, 500, 300, 200, 100, 50, 30, 20, 10, 5, 3, 2, 1}

	for _, incer := range incers {
		s := &stepper{
			incrSize: incer,
			postIds:  make([]int,0, 1000),
		}
		steps = append(steps, s)
	}
	return
}

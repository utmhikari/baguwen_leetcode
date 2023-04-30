package solution_351_400

import "container/heap"

//'355. 设计推特'

type Tweet355 struct {
	tweetId int
	timeId int
}

type heap355 struct {
	IsDesc bool
	arr []Tweet355
}

func (h heap355) Get(i int) Tweet355 {
	return h.arr[i]
}

func (h heap355) Len() int {
	return len(h.arr)
}

func (h heap355) Less(i, j int) bool {
	if h.IsDesc {
		return h.arr[i].timeId > h.arr[j].timeId
	} else {
		return h.arr[i].timeId < h.arr[j].timeId
	}
}

func (h heap355) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *heap355) Pop() interface{} {
	n := h.Len()
	x := h.arr[n - 1]
	h.arr = h.arr[0 : n - 1]
	return x
}

func (h *heap355) Push(x interface{}) {
	h.arr = append(h.arr, x.(Tweet355))
}


func (h heap355) Peek() Tweet355 {
	if h.Len() == 0 {
		return Tweet355{
			tweetId: 0,
			timeId:  0,
		}
	}
	return h.arr[0]
}


type Twitter struct {
	// user follow map
	mpFollow map[int]map[int]bool

	// tweets
	userTweets map[int][]Tweet355

	// time counter
	timeId int
}


/** Initialize your data structure here. */
func Constructor355() Twitter {
	return Twitter{
		mpFollow: make(map[int]map[int]bool),
		userTweets: make(map[int][]Tweet355),
		timeId: 0,
	}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	tweets, ok := this.userTweets[userId]
	this.timeId++
	if !ok {
		this.userTweets[userId] = []Tweet355{{tweetId, this.timeId}}
	} else {
		this.userTweets[userId] = append(tweets, Tweet355{tweetId, this.timeId})
	}
}


/** Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &heap355{}
	userIds := []int{userId}
	followSet, ok := this.mpFollow[userId]
	if ok {
		for followee, _ := range followSet {
			userIds = append(userIds, followee)
		}
	}
	for _, uid := range userIds {
		tweetIds, ok := this.userTweets[uid]
		if ok {
			for _, tweetId := range tweetIds {
				if h.Len() < 10 {
					heap.Push(h, tweetId)
				} else if tweetId.timeId > h.Peek().timeId {
					heap.Pop(h)
					heap.Push(h, tweetId)
				}
			}
		}
	}
	news := make([]int, h.Len())
	for i := h.Len(); i > 0; i-- {
		news[i - 1] = heap.Pop(h).(Tweet355).tweetId
	}
	return news
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	followSet, ok := this.mpFollow[followerId]
	if !ok {
		this.mpFollow[followerId] = map[int]bool{
			followeeId: true,
		}
	} else {
		followSet[followeeId] = true
	}
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	followSet, ok := this.mpFollow[followerId]
	if !ok {
		return
	}
	_, followed := followSet[followeeId]
	if followed {
		delete(followSet, followeeId)
	}
}

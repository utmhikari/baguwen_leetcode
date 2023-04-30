package leetcode_301_400;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;

/**
 * Created by 36191 on 2019/1/14
 */
public class No_355_Design_Twitter {
    class Twitter {
        
        private HashMap<Integer, HashSet<Integer>> followerMap;
        private HashMap<Integer, Integer> tweetMap;
        private List<Integer> tweetIds;
        
        /** Initialize your data structure here. */
        public Twitter() {
            followerMap = new HashMap<>();
            tweetMap = new HashMap<>();
            tweetIds = new ArrayList<>();
        }

        /** Compose a new tweet. */
        public void postTweet(int userId, int tweetId) {
            tweetIds.add(tweetId);
            tweetMap.put(tweetId, userId);
        }

        /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
        public List<Integer> getNewsFeed(int userId) {
            List<Integer> news = new ArrayList<>();
            for (int i = tweetIds.size() - 1; i >= 0; i--) {
                int tweetId = tweetIds.get(i);
                int id = tweetMap.get(tweetId);
                if (id == userId || (followerMap.containsKey(id) && followerMap.get(id).contains(userId))) {
                    news.add(tweetId);
                    if (news.size() == 10) { break; }
                }
            }
            return news;
        }

        /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
        public void follow(int followerId, int followeeId) {
            if (!followerMap.containsKey(followeeId)) {
                followerMap.put(followeeId, new HashSet<>());
            }
            followerMap.get(followeeId).add(followerId);
        }

        /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
        public void unfollow(int followerId, int followeeId) {
            if (followerMap.containsKey(followeeId)) {
                followerMap.get(followeeId).remove(followerId);
            }
        }
    }
    
    
}

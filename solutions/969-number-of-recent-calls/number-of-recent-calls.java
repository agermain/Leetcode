// Write a class RecentCounter to count recent requests.
//
// It has only one method: ping(int t), where t represents some time in milliseconds.
//
// Return the number of pings that have been made from 3000 milliseconds ago until now.
//
// Any ping with time in [t - 3000, t] will count, including the current ping.
//
// It is guaranteed that every call to ping uses a strictly larger value of t than before.
//
//  
//
// Example 1:
//
//
// Input: inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]
// Output: [null,1,2,3,3]
//
//  
//
// Note:
//
//
// 	Each test case will have at most 10000 calls to ping.
// 	Each test case will call ping with strictly increasing values of t.
// 	Each call to ping will have 1 <= t <= 10^9.
//
//
//
//  
//


class RecentCounter {
    List<Integer> l = new ArrayList<>();
    int lb = 0;
    public RecentCounter() {}
    public int ping(int t) {
        l.add(t);
        int r = l.size()-1;
        int target = t - 3000;
        while (lb < r) {
            int middle = lb + (r-lb) / 2;
            if (l.get(middle) >= target) {
                r = middle;
            } else {
                lb = middle+1;
            }
        }
        return l.size() - lb;
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */

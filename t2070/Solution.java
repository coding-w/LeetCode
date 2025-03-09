package t2070;

import java.util.Arrays;
import java.util.Comparator;


class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        Arrays.sort(items, Comparator.comparingInt(a -> a[0]));

        int[] res = new int[queries.length];

        int[] dp = new int[items.length];
        dp[0] = items[0][1];
        for (int i = 1; i < items.length; i++) {
            dp[i] = Math.max(dp[i - 1], items[i][1]);
        }

        for (int i = 0; i < queries.length; i++) {
            int l = 0, r = items.length - 1;
            int best = -1;
            while (l <= r) {
                int mid = r - (r - l) / 2;
                if (queries[i] >= items[mid][0]) {
                    best = mid;
                    l = mid;
                } else {
                    r = mid -1;
                }
            }
            res[i] = best == -1 ? 0 : dp[best];
        }
        return res;
    }
}
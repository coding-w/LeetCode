package t3091;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/7 09:27
 */
class Solution {
    public int countKeyChanges(String s) {
        s = s.toLowerCase();
        int res = 0;
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) != s.charAt(i - 1)) {
                res++;
            }
        }
        return res;
    }
}
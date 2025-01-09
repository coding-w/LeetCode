package t3297;

import java.util.Arrays;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/9 10:08
 */
class Solution {
    public long validSubstringCount(String word1, String word2) {
        long count = 0;
        if (word1.length() < word2.length()) {
            return count;
        }
        int[] count2 = new int[26];
        for (int i = 0; i < word2.length(); i++) {
            count2[word2.charAt(i) - 'a']++;
        }
        int size2 = word2.length();
        int[] count1 = new int[26];
        int cur = 0;
        for (int i = 0; i < word1.length(); i++) {
            count1[word1.charAt(i) - 'a']++;
            while (compare(count1, count2)) {
                count += word1.length() - i;
                count1[word1.charAt(cur) - 'a']--;
                cur++;
            }
        }
        return count;
    }

    private boolean compare(int[] count1, int[] count2) {
        for (int i = 0; i < 26; i++) {
            if (count1[i] < count2[i]) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.validSubstringCount("dcbdcdccb", "cdd"));
    }
}
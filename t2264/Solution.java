package t2264;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/8 14:12
 */
public class Solution {
    public static void main(String[] args) {

    }

    // https://leetcode.cn/problems/largest-3-same-digit-number-in-string/?envType=daily-question&envId=2025-01-08
    // 每日一题
    public String largestGoodInteger(String num) {
        char result = '0'-1;
        for (int i = 2; i < num.length(); i++) {
            if (num.charAt(i - 1) == num.charAt(i) && num.charAt(i - 2) == num.charAt(i)) {
                if (result < num.charAt(i)) {
                    result = num.charAt(i);
                }
            }
        }
        if (result == '0'-1) {
            return "";
        }
        return String.format("%c%c%c", result, result, result);
    }
}

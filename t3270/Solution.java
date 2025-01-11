package t3270;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/11 19:36
 */
class Solution {
    public static void main(String[] args) {

    }
    // https://leetcode.cn/problems/find-the-key-of-the-numbers/?envType=daily-question&envId=2025-01-11
    public int generateKey(int num1, int num2, int num3) {
        int res = 0;
        res += Math.min(Math.min(num1/1000, num2/1000), num3/1000)*1000;
        num1 = num1%1000;
        num2 = num2%1000;
        num3 = num3%1000;
        res += Math.min(Math.min(num1/100, num2/100), num3/100)*100;
        num1 = num1%100;
        num2 = num2%100;
        num3 = num3%100;
        res += Math.min(Math.min(num1/10, num2/10), num3/10)*10;
        num1 = num1%10;
        num2 = num2%10;
        num3 = num3%10;
        res += Math.min(Math.min(num1, num2), num3);
        return res;
    }
}
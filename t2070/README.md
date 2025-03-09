# 每日一题：2070. 每一个查询的最大美丽值

## 题目描述

给定一个二维整数数组 `items`，其中 `items[i] = [price_i, beauty_i]` 分别表示第 *i* 个物品的价格和美丽值。同时给定一个下标从 0 开始的整数数组 `queries`，对于每个查询 `queries[j]`，请你求出所有价格小于等于 `queries[j]` 的物品中最大的美丽值。如果不存在符合条件的物品，则该查询的答案为 `0`。

请返回一个与 `queries` 长度相同的数组 `answer`，其中 `answer[j]` 表示第 *j* 个查询的答案。

### 示例

- **示例 1：**

  - **输入：**
    - `items = [[1,2],[3,2],[2,4],[5,6],[3,5]]`
    - `queries = [1,2,3,4,5,6]`
  - **输出：** `[2,4,5,5,6,6]`
  - **解释：**
    - 对于 `queries[0]=1`，只有 `[1,2]` 的价格小于等于 1，所以答案为 `2`。
    - 对于 `queries[1]=2`，符合条件的物品有 `[1,2]` 和 `[2,4]`，最大美丽值为 `4`。
    - 对于 `queries[2]=3` 和 `queries[3]=4`，符合条件的物品有 `[1,2]`、`[3,2]`、`[2,4]` 和 `[3,5]`，最大美丽值为 `5`。
    - 对于 `queries[4]=5` 和 `queries[5]=6`，所有物品均符合条件，所以答案为 `6`。

- **示例 2：**

  - **输入：**
    - `items = [[1,2],[1,2],[1,3],[1,4]]`
    - `queries = [1]`
  - **输出：** `[4]`
  - **解释：**
    - 所有物品的价格均为 `1`，因此选择最大美丽值 `4`。

- **示例 3：**

  - **输入：**
    - `items = [[10,1000]]`
    - `queries = [5]`
  - **输出：** `[0]`
  - **解释：**
    - 没有物品的价格小于等于 `5`，所以查询结果为 `0`。

## 解题思路

1. **排序处理**  
   将 `items` 数组按照价格进行升序排序；若价格相同，则按美丽值降序排序。这样可以保证在二分查找时，能够快速找到满足条件的物品。

2. **动态规划预处理**  
   构建一个 `dp` 数组，其中 `dp[i]` 表示从第 0 个物品到第 *i* 个物品中出现的最大美丽值。这样，在查找到满足条件的物品下标后，可以直接利用 `dp` 数组获取对应的最大美丽值。

3. **二分查找**  
   对于每个查询 `queries[j]`，使用二分查找在排序后的 `items` 数组中寻找价格小于等于 `queries[j]` 的最右侧物品的下标，然后利用 `dp` 数组得到该位置的最大美丽值作为答案。

```java
import java.util.Arrays;
import java.util.Comparator;


class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        Arrays.sort(items, (int[]a, int[]b)->{
            if(a[0] != b[0]) return a[0] - b[0];
            return b[1] - a[1];
        });

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
                    l = mid + 1;
                } else {
                    r = mid -1;
                }
            }
            res[i] = best == -1 ? 0 : dp[best];
        }
        return res;
    }
}
```


主要思想是通过排序和二分查找相结合，利用动态规划预处理记录前缀最大美丽值
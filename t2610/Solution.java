package t2610;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class Solution {
    public List<List<Integer>> findMatrix(int[] nums) {
        HashMap<Integer, Integer> map = new HashMap<>();
        List<List<Integer>> res = new ArrayList<>();
        for (int num : nums) {
            Integer index = map.getOrDefault(num, 0);
            if (index == res.size()) {
                res.add(new ArrayList<>());
            }
            res.get(index).add(num);
            map.put(num, index + 1);
        }

        return res;
    }
}
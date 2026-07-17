class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        // sort the vector so we can easily track down the pair 
        // make sure we keep the original vector untouched:
        // we want indices from it, not the values, so it's immutable
        vector<int> sorted = nums;
        std::sort(sorted.begin(), sorted.end());

        // pointers for left and right, to close in on solution
        // also values to store the correct numbers, so we can search for
        // them in nums afte
        int i = 0;
        int j = sorted.size() - 1;
        int val1 = -1;
        int val2 = -1;
        while (i < j) {
            int a = sorted[i];
            int b = sorted[j];

            if (a+b == target) {
                // have a hit! we just need to find what their index was in original array
                val1 = a;
                val2 = b; 
                std::cout << "match\n" << val1 << " " << val2;
                break;
            } else if (a+b < target) {
                // move right on i, to get bigger
                i++;
            } else { 
                // move left on j, to get smaller
                j--;
            }
        }
        
        // these give pointers, get distance after for raw index
        auto first = std::find(nums.begin(), nums.end(), val1);
        // we can't pick up the same I twice, and second will always be bigger, because b > a
        // so if we got a duplicate, find the second instance
        auto second = std::find(nums.begin(), nums.end(), val2); 
        if (second == first) { 
            second = std::find(first + 1, nums.end(), val2);
        }

        int firstInt = std::distance(nums.begin(), first);
        int secondInt = std::distance(nums.begin(), second);
        std::cout << "found\n" << firstInt << " " << secondInt;


        vector<int> toReturn = {firstInt, secondInt};
        std::sort(toReturn.begin(), toReturn.end()); // smaller index first!
        return toReturn;
    }
};

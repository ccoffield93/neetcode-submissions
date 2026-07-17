class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        std::unordered_map<int, int> inventory;

        for (int i = 0; i < nums.size(); i++) {
            if (inventory.find(nums[i]) != inventory.end()) {
                return true;
            }
            inventory[nums[i]] = 1;
        }
        
        return false;
    }
};
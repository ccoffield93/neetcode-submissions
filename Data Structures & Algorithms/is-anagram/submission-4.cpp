class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.length() != t.length()) return false; 

        std::unordered_map<int, int> lettersS;
        std::unordered_map<int, int> lettersT;
        for (int i=0; i < s.length(); i++) {
            if (lettersS.find(s[i]) != lettersT.end()) {
                lettersS[s[i]] = lettersS[s[i]] + 1;
            } else {
                lettersS[s[i]] = 1;
            } 
        }

        for (int i=0; i < t.length(); i++) {
            if (lettersT.find(t[i]) != lettersT.end()) {
                lettersT[t[i]] = lettersT[t[i]] + 1;
            } else {
                lettersT[t[i]] = 1;
            } 
        }

        for (const auto& [key, value] : lettersS) {
            if (lettersT[key] != value) {
                return false;
            }
        }
        return true;
    }
};

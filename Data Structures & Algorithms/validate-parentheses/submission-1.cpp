class Solution {
public:
    bool isValid(string s) {
         std::stack<char> parenthesisStack;

         int curlyOpen = 0;
         int squareOpen = 0;
         int roundOpen = 0;

         for (size_t i = 0; i < s.size(); ++i) { 
           std::cout << "Index " << i << ": " << s[i] << "\n";
           if (i == 0) {
            // first item, we just need to push it onto the stack and track what it was
            parenthesisStack.push(s[i]);
            switch(s[i]) {
                case '(':
                    roundOpen++;
                    break;
                case '[':
                    squareOpen++;
                    break;
                case '{':
                    curlyOpen++;
                    break;
                default: 
                    return false; 
            }
           } else {
             // what we expect, and what we do now, depends entirely on what we're trying to add 
             char last;
             switch(s[i]) {
                // opening parenthesis is free, you can do that without a problem
                case '(':
                    roundOpen++;
                    parenthesisStack.push(s[i]);
                    break;
                case '[':
                    squareOpen++;
                    parenthesisStack.push(s[i]);
                    break;
                case '{':
                    curlyOpen++;
                    parenthesisStack.push(s[i]);
                    break;

                // closing is NOT free 
                // if we're closing parenthesis, the last parenthesis we 
                // opened must be of the same type 
                case ')':
                    if ( parenthesisStack.size() == 0) return false;
                    last = parenthesisStack.top();
                    if (last == '(') {
                        roundOpen--;
                        parenthesisStack.pop();
                    } else {
                        return false;
                    }
                    break;
                case ']':
                    if ( parenthesisStack.size() == 0) return false;
                    last = parenthesisStack.top();
                    if (last == '[') {
                        squareOpen--;
                        parenthesisStack.pop();
                    } else {
                        return false;
                    }
                    break;
                case '}':
                    if ( parenthesisStack.size() == 0) return false;
                    last = parenthesisStack.top();
                    if (last == '{') {
                        curlyOpen--;
                        parenthesisStack.pop();
                    } else {
                        return false;
                    }
                    break;
                default: 
                    return false; 
                }
           }
        } // end for loop

        if (curlyOpen != 0 || roundOpen != 0 || squareOpen != 0) {
            return false;
        }

        return true;
    } // end function
};

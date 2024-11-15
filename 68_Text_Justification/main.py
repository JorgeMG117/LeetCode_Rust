"""
Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.
 

Example 1:

Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Example 2:

Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
Note that the second line is also left-justified because it contains only one word.
Example 3:

Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]
"""
class Solution:
    def fullJustify(self, words, maxWidth):
        current_space = maxWidth
        added_words = []
        r = []
        res = ""
        len_words = 0
        i = 0
        while i < len(words):
            l = len(words[i])
            added_words.append(words[i]) 
            len_words += l

            current_space = maxWidth - (len_words + (len(added_words) - 1))

            if current_space >= 0:
                print("Length of new word: ", words[i], l)
                print("Adding word: ", words[i], " current space: ", current_space)

                if i == len(words)-1:
                    print(words[i])
                    for j, word in enumerate(added_words):
                        res += word
                        if len(added_words)-1 != j: 
                            res += " "

                    print(maxWidth)
                    print(len(res))
                    res += " " * (maxWidth - len(res))
                    r.append(res)


                    # Escribir linea
                i += 1
            else: 
                added_words.pop()
                len_words -= l
                current_space = maxWidth - (len_words + (len(added_words) - 1))


                num_words = 1
                if len(added_words) != 1:
                    num_words = len(added_words) - 1
                num_spaces_all = current_space // num_words
                num_spaces_extra = current_space % num_words
                j = 0
                while j < num_words:
                    added_word = added_words[j]
                    res += added_word
                    if num_spaces_extra > 0:
                        space_string = " " * (num_spaces_all + 2)
                        num_spaces_extra -= 1
                    else:
                        if len(added_words) != 1:
                            space_string = " " * (num_spaces_all + 1) 
                        else:
                            space_string = " " * (num_spaces_all) 
                    res += space_string 
                    j += 1

                if len(added_words) != 1:
                    res += added_words[j]

                r.append(res)
                res = ""
                
                added_words = []
                len_words = 0
                current_space = maxWidth
            
        return r
        

def main():
    sol = Solution()
    #words = ["This", "is", "an", "example", "of", "text", "justification."]
    words = ["What","must","be","acknowledgment","shall","be"]
    #words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]

    maxWidth = 16
    s = sol.fullJustify(words, maxWidth)
    print(s)







if __name__ == "__main__":
    main()

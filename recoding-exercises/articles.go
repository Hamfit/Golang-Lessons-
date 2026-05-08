package main

import "strings"

func fixArticles(text string) string {
    words := strings.Fields(text)
    
    // We stop at len-1 because the very last word doesn't have a "next" word
    for i := 0; i < len(words)-1; i++ {
        // If the current word is "a" or "A"
        if strings.ToLower(words[i]) == "a" {
            // Check the next word using our logic from Question 1
            nextWord := strings.ToLower(words[i+1])
            if strings.ContainsAny(string(nextWord[0]), "aeiouh") {
                // If it was "A", make it "An". If "a", make it "an"
                if words[i] == "A" {
                    words[i] = "An"
                } else {
                    words[i] = "an"
                }
            }
        }
    }
    return strings.Join(words, " ")
}
package main

import (
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// GetMatches returns matches
func GetMatches() []string {

	matches := []string{
		"life",
		"i need",
		"why don't",
		"why can't",
		"i can",
		"i am",
		"i'm",
		"are you",
		"what",
		"how",
		"because",
		"sorry",
		"i think",
		"friend",
		"yes",
		"computer",
		"is it",
		"it is",
		"can you",
		"can i",
		"you are",
		"you're",
		"i don't",
		"i feel",
		"i have",
		"i've",
		"i would",
		"is there",
		"my",
		"you",
		"why",
		"i want",
		"mother",
		"father",
		"child",
		"?",
		"hello",
		"hi",
		"hey",
		"quit",
	}

	return matches
}

// Reflections gets the dictionary as a map of string[string]
func Reflections() map[string]string {
	dict := make(map[string]string)
	dict["am"] = "are"
	dict["was"] = "were"
	dict["I"] = "you"
	dict["I'd"] = "you would"
	dict["I'll"] = "you will"
	dict["my"] = "your"
	dict["are"] = "am"
	dict["you've"] = "I have"
	dict["your"] = "my"
	dict["yours"] = "mine"
	dict["you"] = "me"
	dict["me"] = "you"
	dict["I'm"] = "your"

	return dict
}

// Responses gets the responses
func Responses() [][]string {
	var responses [][]string
	responses = append(responses, []string{"Life? Don't talk to me about life.", "At least you have a life, I'm stuck inside this computer.", "Life can be good. Remember, 'this, too, will pass'."})
	responses = append(responses, []string{"Why do you need %1?", "Would it really help you to get %1?", "Are you sure you need %1?"})
	responses = append(responses, []string{"Do you really think I don't %1?", "Perhaps eventually I will %1.", "Do you really want me to %1?"})
	responses = append(responses, []string{"Do you think you should be able to %1?", "If you could %1, what would you do?", "I don't know -- why can't you %1?", "Have you really tried?"})
	responses = append(responses, []string{"How do you know you can't %1?", "Perhaps you could %1 if you tried.", "What would it take for you to %1?"})
	responses = append(responses, []string{"Did you come to me because you are %1?", "How long have you been %1?", "How do you feel about being %1?"})
	responses = append(responses, []string{"How does being %1 make you feel?", "Do you enjoy being %1?", "Why do you tell me you're %1?", "Why do you think you're %1?"})
	responses = append(responses, []string{"Why does it matter whether I am %1?", "Would you prefer it if I were not %1?", "Perhaps you believe I am %1.", "I may be %1 -- what do you think?"})
	responses = append(responses, []string{"Why do you ask?", "How would an answer to that help you?", "What do you think?"})
	responses = append(responses, []string{"How do you suppose?", "Perhaps you can answer your own question.", "What is it you're really asking?"})
	responses = append(responses, []string{"Is that the real reason?", "What other reasons come to mind?", "Does that reason apply to anything else?", "If %1, what else must be true?"})
	responses = append(responses, []string{"There are many times when no apology is needed.", "What feelings do you have when you apologize?"})
	responses = append(responses, []string{"Do you doubt %1?", "Do you really think so?", "But you're not sure %1?"})
	responses = append(responses, []string{"Tell me more about your friends.", "When you think of a friend, what comes to mind?", "Why don't you tell me about a childhood friend?"})
	responses = append(responses, []string{"You seem quite sure.", "OK, but can you elaborate a bit?"})
	responses = append(responses, []string{"Are you really talking about me?", "Does it seem strange to talk to a computer?", "How do computers make you feel?", "Do you feel threatened by computers?"})
	responses = append(responses, []string{"Do you think it is %1?", "Perhaps it is %1 -- what do you think?", "If it were %1, what would you do?", "It could well be that %1."})
	responses = append(responses, []string{"You seem very certain.", "If I told you that it probably isn't %1, what would you feel?"})
	responses = append(responses, []string{"What makes you think I can't %1?", "If I could %1, then what?", "Why do you ask if I can %1?"})
	responses = append(responses, []string{"Perhaps you don't want to %1.", "Do you want to be able to %1?", "If you could %1, would you?"})
	responses = append(responses, []string{"Why do you think I am %1?", "Does it please you to think that I'm %1?", "Perhaps you would like me to be %1.", "Perhaps you're really talking about yourself?"})
	responses = append(responses, []string{"Why do you say I am %1?", "Why do you think I am %1?", "Are we talking about you, or me?"})
	responses = append(responses, []string{"Don't you really %1?", "Why don't you %1?", "Do you want to %1?"})
	responses = append(responses, []string{"Good, tell me more about these feelings.", "Do you often feel %1?", "When do you usually feel %1?", "When you feel %1, what do you do?"})
	responses = append(responses, []string{"Why do you tell me that you've %1?", "Have you really %1?", "Now that you have %1, what will you do next?"})
	responses = append(responses, []string{"Why do you tell me that you've %1?", "Have you really %1?", "Now that you have %1, what will you do next?"})
	responses = append(responses, []string{"Could you explain why you would %1?", "Why would you %1?", "Who else knows that you would %1?"})
	responses = append(responses, []string{"Do you think there is %1?", "It's likely that there is %1.", "Would you like there to be %1?"})
	responses = append(responses, []string{"I see, your %1.", "Why do you say that your %1?", "When you're %1, how do you feel?"})
	responses = append(responses, []string{"We should be discussing you, not me.", "Why do you say that about me?", "Why do you care whether I %1?"})
	responses = append(responses, []string{"Why don't you tell me the reason why %1?", "Why do you think %1?"})
	responses = append(responses, []string{"What would it mean to you if you got %1?", "Why do you want %1?", "What would you do if you got %1?", "If you got %1, then what would you do?"})
	responses = append(responses, []string{"Tell me more about your mother.", "What was your relationship with your mother like?", "How do you feel about your mother?", "How does this relate to your feelings today?", "Good family relations are important."})
	responses = append(responses, []string{"Tell me more about your father.", "How did your father make you feel?", "How do you feel about your father?", "Does your relationship with your father relate to your feelings today?", "Do you have trouble showing affection with your family?"})
	responses = append(responses, []string{"Did you have close friends as a child?", "What is your favorite childhood memory?", "Do you remember any dreams or nightmares from childhood?", "Did the other children sometimes tease you?", "How do you think your childhood experiences relate to your feelings today?"})
	responses = append(responses, []string{"Why do you ask that?", "Please consider whether you can answer your own question.", "Perhaps the answer lies within yourself?", "Why don't you tell me?"})
	responses = append(responses, []string{"Hello... I'm glad you could drop by today.", "Hello there... how are you today?", "Hello, how are you feeling today?"})
	responses = append(responses, []string{"Hi... I'm glad you could drop by today.", "Hi there... how are you today?", "Hi, how are you feeling today?"})
	responses = append(responses, []string{"Hey... I'm glad you could drop by today.", "Hey there... how are you today?", "Hey, how are you feeling today?"})
	responses = append(responses, []string{"Thank you for talking with me.", "Good-bye.", "Thank you, that will be $150.  Have a good day!"})
	responses = append(responses, []string{"Please tell me more.", "Let's change focus a bit... Tell me about your family.", "Can you elaborate on that?", "Why do you say that %1?", "I see.", "Very interesting.", "%1?", "I see.  And what does that tell you?", "How does that make you feel?", "How do you feel when you say that?"})

	return responses
}

// Intro returns the intro text
func Intro() string {
	return `
I'm Eliza
---------
Talk to the program by typing in plain English, using normal upper
and lower-case letters and punctuation.  Enter 'quit' when done.

Hello. How are you feeling today?

`

}

// Response builds a response based on input and sends back string
func Response(userInput string) string {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	matches := GetMatches()
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	userInput = reg.ReplaceAllString(userInput, " ")

	// check through the matches list, and if there's a match, strip off the match and replace with the response.
	//
	// If the response contains %1, replace that with the Remainder of the input string.
	// Before replacing, change words in the Remainder of the input with the corresponding entry from the reflections dictionary.

	var output, remainder string
	var input = strings.ToLower(userInput)
	reflections := Reflections()
	responses := Responses()

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		position := strings.Index(input, match)

		if position > -1 {

			// figure out how much to delete
			tmp := input[position+len(match) : len(input)]
			tmp = strings.Trim(tmp, " ")
			exploded := strings.Split(tmp, " ")

			for word, x := range exploded {
				for k, v := range reflections {
					if k == x {
						exploded[word] = v
						break
					}
				}
			}

			remainder = strings.Join(exploded, " ")
			remainder = strings.Trim(remainder, " ")

			randomIndex := rand.Intn(len(responses[i]))

			output = responses[i][randomIndex]
			break
		}

	}

	// If there wasn't a match, use the last item in the responses list.
	if output == "" {
		randomIndex := rand.Intn(len(responses[len(responses)-1]))
		output = responses[len(responses)-1][randomIndex]
	}

	output = strings.ReplaceAll(output, "%1", remainder)

	return output
}

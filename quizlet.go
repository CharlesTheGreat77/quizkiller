package main

import (
    "fmt"
	"bufio"
	"flag"
	"math/rand"
    "net/http"
    "net/url"
    "strings"
    "sort"
	"sync"
    "os"
	"regexp"
	"time"

    "github.com/gocolly/colly/v2"
    "github.com/imroc/req/v3"
)


func main() {
	var questions []string

	question := flag.String("question", "", "specify a question/file to query")
	flag.StringVar(question, "q", "", "alias for --question")
	help := flag.Bool("h", false, "show usage")
	flag.Parse()

	// handle quotes
	args := flag.Args()
	if len(args) > 0 {
		*question = args[0]
	}

	if *help {
		flag.Usage()
		return
	}

	if *question == "" {
		flag.Usage()
		return
	}

	*question = strings.TrimSpace(*question)

	// check if arg is file
	if _, err := os.Stat(*question); err == nil {
		// get questions from file
		questions, err = openQuizFile(*question); if err != nil {
			fmt.Printf("[-] Error occurred opening quiz questions\n -> Error: %v", err)
			return
		}
	} else {
		questions = append(questions, *question)
	}

	// replace special chars
	questions = replaceSpecialChars(questions)

	// search google for the questions
	quizletLinks := searchGoogle(questions)

	// scrape question and answer bank for each quizlet link 
	questionAndAnswerBank := scrapeQuizletBank(quizletLinks)

	// search for answers in question bank
	answeredQuestions, unansweredQuestions := searchQuestionBank(questions, questionAndAnswerBank)

	// output question and answers found (if any)
	if len(answeredQuestions) != 0 {
		for question, answer := range answeredQuestions {
			fmt.Printf("\n[*] Question: %v\n -> Answer: %v\n\n", question, answer)
		}
	}

	// output unanswered questions (if any)
	if len(unansweredQuestions) != 0 {
		fmt.Printf("\n[*] Questions not found..\n")
		for _, unansweredQuestion := range unansweredQuestions {
			fmt.Printf("[*] Question: %v\n\n", unansweredQuestion)
		}
	}
}


// function to search for the answer for each question in the dict
func searchQuestionBank(questions []string, questionAndAnswerBank map[string]string) (map[string]string, []string) {
    answerResults := make(map[string]string)
	var unansweredQuestions []string

	// easier than regex personally..
	for _, question := range questions {
		question = strings.TrimSpace(question)
		found := false
		for questionBank, answerBank := range questionAndAnswerBank {
			if strings.Contains(questionBank, strings.ToLower(question)) {
				answerResults[questionBank] = answerBank
				found = true
			}
		}

		if !found {
			unansweredQuestions = append(unansweredQuestions, question) // if question is not found, append to slice
		}
	}

    return answerResults, unansweredQuestions
}


// function to scrape quizlet for each link
func scrapeQuizletBank(quizletLinks []string) (map[string]string) {
	var wg sync.WaitGroup

	var questions []string
    questionsAndAnswers := make(map[string]string)

    c := createCollector()
	setCollyBehavior(c)

	// question and answers are stored here
    c.OnHTML("span.TermText.notranslate.lang-en", func(e *colly.HTMLElement) {
            questions = append(questions, strings.TrimSpace(strings.ToLower(e.Text))) // make everything lowercase
    })

    for _, link := range quizletLinks {
		wg.Add(1)

		// go routine to scrape quizlet links.. but there's a sleep
		go func(quizletLink string) {
			defer wg.Done()

			err := c.Visit(quizletLink); if err != nil {
				fmt.Printf("[-] Error occurred visiting: %v\n -> Error: %v\n\n", link, err)
			}
		}(link)

		// sleep for rand time
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second) // one can get rid of this if you wish for turbo!
    }

	wg.Wait()

	// parse question and answers to dict
    for i := 0; i < len(questions); i += 2 {
            if i+1 < len(questions) {
                    questionsAndAnswers[questions[i]] = questions[i+1]
            }
    }

    return questionsAndAnswers
}


// function to search google and scrape the links of the search results
func searchGoogle(questions []string) []string {
	var wg sync.WaitGroup

    quizletLinks := []string{}

    c := createCollector()
	setCollyBehavior(c)

	// extract links that contain quizlet.com from google search results
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
            links := e.Attr("href")
            if strings.Contains(e.Request.AbsoluteURL(links), "quizlet.com") {
                    quizletLinks = append(quizletLinks, e.Request.AbsoluteURL(links))
            }
    })

	for _, question := range questions {
		wg.Add(1)
    	link := fmt.Sprintf("https://google.com/search?q=%s", url.QueryEscape(question))

		// we want the turbo!
		go func(searchQuery string) {
			defer wg.Done()
			err := c.Visit(searchQuery); if err != nil {
				fmt.Printf("\n[-] Error visiting %v\n -> Error: %v\n\n", searchQuery, err)
			}
		}(link)
	}

	wg.Wait()

	return removeDuplicates(quizletLinks)
}


// function to create a colly collector impersonating chrome browser
func createCollector() *colly.Collector {
    fakeChrome := req.DefaultClient().ImpersonateChrome()
    c := colly.NewCollector(
            colly.UserAgent(fakeChrome.Headers.Get("user-agent")),
    )
    c.SetClient(&http.Client{Transport: fakeChrome.Transport,})
    return c
}


// set behavior of colly for each request and error response
func setCollyBehavior(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\r[*] Requesting: %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
			fmt.Printf("\n[-] Error occurred\n -> Response Code: %v\n --> Error: %v\n\n", r.StatusCode, err)
	})
}

// function to replace stupid chars that prevent matching question with answers..
func replaceSpecialChars(slices []string) []string {
	re := regexp.MustCompile(`[“”‘’]`)
	for i, question := range slices {
		slices[i] = re.ReplaceAllString(question, `"`)
		slices[i] = strings.TrimSpace(slices[i])
	}
	return slices
}

// function to remove duplicates in a slice
func removeDuplicates(slices []string) []string {
    if len(slices) < 1 {
            return slices
    }
    sort.Strings(slices)

    prev := 1
	// I honestly don't know an easier way.. :( +1 for python
    for curr := 1; curr < len(slices); curr++ {
            if slices[curr-1] != slices[curr] {
                    slices[prev] = slices[curr]
                    prev++
            }
    }

    return slices[:prev]
}


// function to open file and grab the questions
func openQuizFile(fileName string) ([]string, error) {
	var questions []string

	file, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		if scanner.Text() != "" {
			questions = append(questions, scanner.Text())
		}
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("\n[-] Error scanning file: %v\n", err)
		return nil, err
    }

	return questions, nil
}
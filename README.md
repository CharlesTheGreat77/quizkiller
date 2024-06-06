# quizkiller
Quizcrawler scrapes quizlet for question(s), and obtains the answer(s). This was made for one to quickly go through quizzes that are of no significants to ones exam(s).

```
             .--.           .---.        .-.
         .---|--|   .-.     | A |  .---. |~|    .--.
      .--|===|Ch|---|_|--.__| S |--|:::| |~|-==-|==|---.
      |%%|NT2|oc|===| |~~|%%| C |--|   |_|~|CATS|  |___|-.
      |  |   |ah|===| |==|  | I |  |:::|=| |    |GB|---|=|
      |  |   |ol|   |_|__|  | I |__|   | | |    |  |___| |
      |~~|===|--|===|~|~~|%%|~~~|--|:::|=|~|----|==|---|=|
      ^--^---'--^---^-^--^--^---'--^---^-^-^-==-^--^---^-'
```

# Description
I wrote this to save me the time from a billion CTRL-F's to google--then quizlet, for me to get the answers, I was not planning on writing so but I tend to do this for quizzes more often than I would like. With that.. we may not be able to read a book.. but we can sure write some code.

# Prerequisite
| Prerequisite | Version |
|--------------|---------|
| Go           |  <=1.21 |


# Install Go 🚀
```
apt install golang-go || brew install go
```

# usage 😈
```
Usage of ./quizlet:
  -h    show usage
  -q string
        alias for --question
  -question string
        specify a question/file to query
  ```

# examples 👀
Single question look up
```
./quizlet -q "What is top-down editing?"
```

# Multiple questions 
Multiple question look up
```
./quizlet -q quiz.txt
```
- questions must be seperated by line

# Potential Addons
  (1) Just copy and paste whole canvas and auto grab questions to search

  (2) AI for questions unanswered maybe? 👀

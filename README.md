# quizcrawler
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

# usage
```
usage: python3 quizlet.py -q <question>

Quizlet search for quick answers for quizzes

options:
  -h, --help            show this help message and exit
  -q QUESTION, --question QUESTION
                        specify question [or file] to search quizlet for the answers
  -o OUTPUT, --output OUTPUT
                        specify output file to save question and answers
  -v, --verbose         enable verbosity to output the requests sent
  ```

# example
Single question look up
```
python3 quizlet.py -q "what is a covalent bond?"
```

# Multiple questions
Multiple question look up
```
python3 quizlet.py -q quiz.txt
```
- questions must be seperated by line

# Potential Addons
  (1) Just copy and paste whole canvas and auto grab questions to search

  (2) AI for questions unanswered maybe? 👀
# quizkiller
# quizkiller
# quizkiller

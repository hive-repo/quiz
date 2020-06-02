# QuizApp

Date: 09 May, 2020

The concept of this QuizApp is very simple. This will be the same as our vocab app we’re using, but in a more general way. There will be a question followed by a number of options including one correct answer. The main purpose of this app will be remembering the quiz with ease.

## Some of the features that needs to be included:

- Quiz contains list questions.
- A set of questions will be staged in a session. And only staged questions will appear in a particular session.
- Users can Mask or Master a question.
- Masked question will not appear in the session, but continue appearing in next session
- Mastered questions will not appear until it is reset. Mastering a question will remove the mastered question from stage, and add another one from the list.
- Masked Questions, Staged Questions and Mastered Questions' states will be stored, so that can be retrieved in a continuing session.
- Resets will flush all the states and start from the beginning.
- There will be a view option, which will be the correct answer.
- Stats of Masked, Mastered, and Staged Questions will be displayed.

Other options can be added and suggestions from other members are appreciated.

## TODO

- [x] Basic functionality as described in feature list
- [x] External Config
- [x] External Data
- [x] Save State in External file/db
- [ ] Testing
- [ ] Modular
- [ ] Use Cobra cli
- [ ] Use environment variables and arguments to get quizes. Support URLs

package api

import (
	"errors"
	"fasttrack_api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var rate = make(map[int]int)

var users = []model.Registred_user{
	{ID: "1", Name: "John Doe", Email: "doe.jonh@hotmail.com", Quiz: questions, Number_corrected_answers: 5, User_rated: 0.00},
	{ID: "2", Name: "Jane Doe", Email: "janedoe1989@gmail.com", Quiz: questions, Number_corrected_answers: 4, User_rated: 0.00},
	{ID: "3", Name: "John Does", Email: "does.jonh@hotmail.com", Quiz: questions, Number_corrected_answers: 3, User_rated: 0.00},
	{ID: "4", Name: "Jane Does", Email: "janedoes1999@gmail.com", Quiz: questions, Number_corrected_answers: 2, User_rated: 0.00},
	{ID: "5", Name: "John Foe", Email: "foe.jonh@hotmail.com", Quiz: questions, Number_corrected_answers: 1, User_rated: 0.00},
	{ID: "6", Name: "Jane Foe", Email: "janefoe1979@gmail.com", Quiz: questions, Number_corrected_answers: 0, User_rated: 0.00},
	{ID: "7", Name: "John Toe", Email: "toe.jonh@hotmail.com", Quiz: questions, Number_corrected_answers: 5, User_rated: 0.00},
	{ID: "8", Name: "Jane Toe", Email: "jantoe1984@gmail.com", Quiz: questions, Number_corrected_answers: 4, User_rated: 0.00},
	{ID: "9", Name: "John Hoe", Email: "hoe.jonh@hotmail.com", Quiz: questions, Number_corrected_answers: 3, User_rated: 0.00},
	{ID: "10", Name: "Jane Hoe", Email: "janehoe1988@gmail.com", Quiz: questions, Number_corrected_answers: 2, User_rated: 0.00},
}

var questions = []model.Question{
	{ID: "1", Question: "Question 1", Answers: []string{"A1", "A2", "A3", "A4", "A5"}, Correct_answer: "A4", Answered: false, Is_corrected: false},
	{ID: "2", Question: "Question 2", Answers: []string{"B1", "B2", "B3", "B4", "B5"}, Correct_answer: "B1", Answered: false, Is_corrected: false},
	{ID: "3", Question: "Question 3", Answers: []string{"C1", "C2", "C3", "C4", "C5"}, Correct_answer: "C5", Answered: false, Is_corrected: false},
	{ID: "4", Question: "Question 4", Answers: []string{"D1", "D2", "D3", "D4", "D5"}, Correct_answer: "D2", Answered: false, Is_corrected: false},
	{ID: "5", Question: "Question 5", Answers: []string{"E1", "E2", "E3", "E4", "E5"}, Correct_answer: "E3", Answered: false, Is_corrected: false},
}

//returns an error with emails are equals
func verifyEmail(new_email string, existed_email string) error {
	if new_email == existed_email {
		return errors.New("this e-mail already exists")
	}
	return nil
}

//getQuestions responds with the list of all questions as JSON.
func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}

//getUsers list of all users as JSON.
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

//getUsers list of all users as JSON.
func getUserById(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("userid"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	message := users[userid-1]
	c.JSON(http.StatusOK, message)
}

//getUsers list of all users as JSON.
func getQuestionById(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("userid"), 10, 0)
	questionid, err := strconv.ParseInt(c.Param("questionid"), 10, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	user := users[userid-1]
	message := user.Quiz[questionid-1]

	c.JSON(http.StatusOK, message)
}

//getUsers list of all users as JSON.
func getUserByEmail(c *gin.Context) {
	email := c.Param("userid")

	//find user by email.
	for _, existed_user := range users {
		if existed_user.Email == email {
			c.JSON(http.StatusOK, existed_user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user does not exist."})
}

//Post method to register a new user
func registerUser(c *gin.Context) {
	var newUser model.Registred_user

	//Call a BindJson to bind  the received JSON to new user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	//add a new user to the users slice.
	for _, existed_user := range users {
		if err := verifyEmail(existed_user.Email, newUser.Email); err != nil {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
	}
	newUser.ID = strconv.Itoa(len(users) + 1)
	newUser.Quiz = questions
	newUser.Number_corrected_answers = 0
	newUser.User_rated = 0.00
	users = append(users, newUser)
	// c.JSON(http.StatusCreated, "user created sucessfully")
	c.JSON(http.StatusCreated, "user created sucessfully")
}

//Post method to store user response
func updateUserQuestions(c *gin.Context) {
	email := c.Param("userid")
	var existed_user model.Registred_user

	//Call a BindJson to bind  the received JSON to questions
	if err := c.BindJSON(&existed_user); err != nil {
		return
	}

	//find user by email and set questions and numbers of corrected answers
	for i := 0; i < len(users); i++ {
		if users[i].Email == email {
			users[i].Quiz = existed_user.Quiz
			users[i].Number_corrected_answers = existed_user.Number_corrected_answers
			calculateRateUsers()
			c.JSON(http.StatusOK, users[i])
			return
		}

	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user does not exist."})
}

var port string = "8080"

func SetPortFlag(serverPort string) {
	port = serverPort
}

//Calculate users rate
func calculateRateUsers() {
	rate = make(map[int]int)
	total_answered_question := 0
	for i := 0; i < len(users); i++ {
		rate[users[i].Number_corrected_answers]++
		total_answered_question++
	}
	for i := 0; i < len(users); i++ {
		users[i].User_rated = 0.0
		if users[i].Number_corrected_answers != 0 {
			for j := 0; j < users[i].Number_corrected_answers; j++ {
				users[i].User_rated += float64(rate[j])
			}
			users[i].User_rated = users[i].User_rated / float64(total_answered_question) * 100
		}
	}
}

func StartServer() {
	calculateRateUsers()
	router := gin.Default()
	router.GET("/questions", getQuestions)
	router.GET("/users", getUsers)
	router.GET("/user/:userid", getUserById)
	router.GET("/user/:userid/email", getUserByEmail)
	router.POST("/user/:userid/questions", updateUserQuestions)
	router.GET("/user/:userid/:questionid", getQuestionById)
	router.POST("/user", registerUser)
	router.Run("localhost:" + port)
}

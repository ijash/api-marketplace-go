package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"ijash-jwt-auth/src/configs"
	"ijash-jwt-auth/src/helpers"
	"ijash-jwt-auth/src/middleware"
	"ijash-jwt-auth/src/models"
	"ijash-jwt-auth/src/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	log.Println("GetUsersController accessed")
	var users []models.User

	return c.JSON(http.StatusOK, helpers.BaseResponseOk(users))
}
func LoginController(c echo.Context) error {
	log.Println("LoginController accessed")

	// Define the UserLogin struct
	type UserLogin struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}

	// Parse the request body to get the user login data
	var loginRequest UserLogin
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.BadRequest("Invalid request body"))
	}

	// Check if required fields (userName, password) are provided
	if loginRequest.UserName == "" || loginRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, helpers.BadRequest("Please provide userName and password"))
	}

	// Query the database to find the user by their username
	var user models.User
	result := configs.DB.Where("user_name = ?", loginRequest.UserName).First(&user)
	if result.Error != nil {
		// User with the given username not found
		return c.JSON(http.StatusUnauthorized, helpers.Unauthorized("Invalid credentials"))
	}

	// Hash the provided plain text password for comparison
	hashedPassword, _ := generateSHA256Hash(loginRequest.Password)

	// Compare the hashed password with the stored hashed password
	if hashedPassword != user.Password {
		// Passwords do not match
		return c.JSON(http.StatusUnauthorized, helpers.Unauthorized("Invalid credentials"))
	}

	// Passwords match, generate a JWT token for authentication
	token, err := middleware.GenerateJwt(user.UserName, user.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError("Failed to generate token"))
	}

	// Return the JWT token as a response
	loginResponse := models.UserLoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, helpers.BaseResponseOk(loginResponse, "Success login"))
}

func RegisterController(c echo.Context) error {
	log.Println("RegisterController accessed")
	// Parse the request body to get the user registration data
	var userRequest models.User
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.BadRequest("Invalid request body"))
	}

	// Check if required fields (userName, fullName, password) are provided
	if userRequest.UserName == "" || userRequest.FullName == "" || userRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, helpers.BadRequest("Please provide userName, fullName, and password"))
	}

	// Hash the user's password
	userRequest.Password, _ = generateSHA256Hash(userRequest.Password)
	userRequest.Id = utils.GenStringUUID()

	// Check if a user with the same username already exists
	var existingUser models.User
	userNameResult := configs.DB.Where("user_name = ?", userRequest.UserName).First(&existingUser)
	if userNameResult.Error == nil {
		// A user with the same username already exists
		return c.JSON(http.StatusConflict, helpers.BadRequest("User with the same username already exists"))
	}

	result := configs.DB.Create(&userRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError("Failed to insert user"))
	}

	var userResponse = models.UserResponse{
		Id:       userRequest.Id,
		UserName: userRequest.UserName,
		FullName: userRequest.FullName,
	}

	// Generate a JWT token upon successful registration
	userResponse.Token, _ = middleware.GenerateJwt(
		userResponse.FullName,
		userResponse.Id,
	)

	return c.JSON(http.StatusCreated, helpers.BaseResponseOk(userResponse, "Success add user"))
}

func generateSHA256Hash(input string) (string, error) {
	log.Println("generateSHA256Hash accessed")
	// Membuat sebuah objek hasher SHA-256
	hasher := sha256.New()

	// Menambahkan data string ke hasher
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}

	// Mendapatkan hasil hash sebagai byte
	hashBytes := hasher.Sum(nil)

	// Mengonversi hasil hash menjadi string hexadecimal
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}

func GetUserProfileController(c echo.Context) error {

	log.Println("GetUserProfileController accessed")
	// Extract the JWT token from the authorization header
	rawToken := c.Request().Header.Get("Authorization")

	tokenString, jwtErr := utils.ExtractJWTFromBearerToken(rawToken)
	if jwtErr == true {
		return c.JSON(http.StatusUnauthorized, helpers.Unauthorized("Invalid token"))
	}
	// Extract and validate the user ID from the JWT token
	userID, err := middleware.ExtractUserIDFromJWT(tokenString)
	fmt.Println(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.Unauthorized("Invalid token"))
	}

	// Query the database to retrieve user data based on the userID
	var user models.User
	result := configs.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		// Handle database query error (user not found or other issues)
		return c.JSON(http.StatusInternalServerError, helpers.InternalServerError("User not found"))
	}
	user.Password = "[hidden]"
	// Return the user data as a response
	return c.JSON(http.StatusOK, helpers.BaseResponseOk(user))
}

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/teamhanko/hanko-sdk-golang/webauthn"
	"gitlab.com/hanko/simple-webauthn-example/config"
	"gitlab.com/hanko/simple-webauthn-example/models"
	"net/http"
	"strings"
)

var apiClient *webauthn.Client

func init() {
	// Create a hanko api client
	apiClient = webauthn.NewClient(config.C.ApiUrl, config.C.ApiSecret).
		WithHmac(config.C.ApiKeyId).
		WithLogLevel(log.DebugLevel)
}

func main() {
	r := gin.Default()

	// init session cookie store
	sessionStore := cookie.NewStore([]byte("superSecureSecret"))
	r.Use(sessions.Sessions("webauthn-session", sessionStore))

	r.Static("/assets", "./assets")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.StaticFile("/", "./index.html")

	// Load html files to use as an template
	r.LoadHTMLFiles("./content.html")

	// Get the registration request from the Hanko-API
	r.POST("/registration_initialize", func(c *gin.Context) {
		userName := strings.TrimSpace(c.Query("user_name"))

		// If no username was given return an error
		if userName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user name not provided"})
			return
		}

		// If the username already exists return an error
		userModel, err := models.FindUserByName(userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if userModel != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
			return
		}

		// If the username doesn't already exists, create a new user
		userModel = models.NewUser(uuid.NewV4().String(), userName)
		err = userModel.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create the request options for the Hanko-API
		user := webauthn.NewRegistrationInitializationUser(userModel.ID, userModel.Name)

		authenticatorSelection := webauthn.NewAuthenticatorSelection().
			WithUserVerification(webauthn.VerificationRequired).
			WithAuthenticatorAttachment(webauthn.Platform).
			WithRequireResidentKey(true)

		request := webauthn.NewRegistrationInitializationRequest(user).
			WithAuthenticatorSelection(authenticatorSelection).
			WithConveyancePreference(webauthn.PreferNoAttestation)

		// Get the registration request from the Hanko-API with the given request options
		response, apiErr := apiClient.InitializeRegistration(request)
		if apiErr != nil {
			c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Error()})
			return
		}

		// return the registration request
		c.JSON(http.StatusOK, response)
	})

	// Send the authenticator response to the Hanko-API
	r.POST("/registration_finalize", func(c *gin.Context) {
		// parse the authenticator response
		request, err := webauthn.ParseRegistrationFinalizationRequest(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// send the authenticator response to the Hanko-API
		response, apiErr := apiClient.FinalizeRegistration(request)
		if apiErr != nil {
			c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Error()})
			return
		}

		// If no error occurred during the authenticator response validation, create a session for the given user
		session := sessions.Default(c)
		session.Set("userId", response.Credential.User.ID)
		session.Save()

		c.JSON(http.StatusOK, response)
	})

	// Get an authentication request from the Hanko-API
	r.POST("/authentication_initialize", func(c *gin.Context) {
		// Create the request options
		request := webauthn.NewAuthenticationInitializationRequest().
			WithUserVerification("required").
			WithAuthenticatorAttachment("platform")

		// Get the authentication request from the Hanko-API with the given request options
		response, apiErr := apiClient.InitializeAuthentication(request)
		if apiErr != nil {
			c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	})

	// TODO: Send the authenticator response to the Hanko-API
	r.POST("/authentication_finalize", func(c *gin.Context) {
		// Parse the authenticator response
		request, err := webauthn.ParseAuthenticationFinalizationRequest(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Send the authenticator reponse to the Hanko-API
		response, apiErr := apiClient.FinalizeAuthentication(request)
		if apiErr != nil {
			c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Error()})
			return
		}

		// If no error occurred during the authenticator response validation, create a session for the given user
		session := sessions.Default(c)
		session.Set("userId", response.Credential.User.ID)
		session.Save()

		c.JSON(http.StatusOK, response)
	})

	// Return the content page, if the user as a valid session else redirect to sign in/register page
	r.GET("/content", func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")
		if userId == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		user, err := models.FindUserById(userId.(string))

		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		// Render the content page with the username
		c.HTML(http.StatusOK, "content.html", gin.H{
			"username": user.Name,
		})
	})

	// Delete the session fo a user
	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("userId", nil)
		session.Save()

		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	log.Println("Starting server on localhost:3000")
	log.Fatal(r.Run(":3000"))
}
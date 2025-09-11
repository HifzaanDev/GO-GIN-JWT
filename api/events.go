package main 

func (app *application) createEvent(c *gin.Context){
	var event database.Event

	if err :=c.shouldBindJSON(&event); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
	}
	err:=app.models.Events.insert(&event)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create event"})
	return
	}
	c.JSON(http.StatusCreated,event)
}

func (app *application) getEvent(c *gin.Context){
	id,err:=Atoi(c.Param("id"))

	if err != nil {
		c.JSON (http.StatusBadRequest,gin.H("error":"invalid event ID"))
	}
	event , err := app.models.Events.Get(id)

	if event == nil{
		c.JSON(http.StatusNotFound)gin.H("error":"event not found!")
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to retreive event"
		})
	}c.JSON(http.StatusOK,event)
}

func (app *application)  updateEvent(c *gin.Context){|

id,err:=Atoi(c.Param("id"))

	if err != nil {
		c.JSON (http.StatusBadRequest,gin.H("error":"invalid event ID"))
	}
exisistingEvent,
}

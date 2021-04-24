## Endpoints
**Note:** This API it's still an _ongoing_ project

### Player
`GET /players/{id}`: Get player by ID  

### Match
`GET /matches`: Get all world cup matches  
`GET /matches/{id}`: Get match by ID  

### Tournaments
`GET /tournaments/{id}`: Get tournament by ID  
`GET /tournaments/{id}/scores`:  Get tournament current scores (players standings)  
`GET /players/{playerID}/tournaments`: Get tournaments a player has joined  

### Match predictions
`GET /predictions/{id}`: Get specific prediction by ID    
`GET /tournaments/{tournamentID}/players/{playerID}/predictions`: asd  
`GET /tournaments/{tournamentID}/matches/{matchID}/predictions`: asd  
`GET /tournaments/{tournamentID}/players/{playerID}/matches/{matchID}/predictions`: asd  

`POST /predictions`: Create a new prediction  
`POST /tournaments/{tournamentID}/predictions`: Create a new prediction  
`POST /tournaments/{tournamentID}/players/{playerID}/predictions`: Create a new prediction  
`POST /tournaments/{tournamentID}/matches/{matchID}/predictions`: Create a new prediction  
`POST /tournaments/{tournamentID}/players/{playerID}/matches/{matchID}/predictions`: Create a new prediction

`PUT /predictions/{id}`: Update prediction  

### Misc
`GET /ping`: server healthcheck

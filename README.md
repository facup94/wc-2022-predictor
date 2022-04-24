## WorldCup Predictor
WorldCup Predictor is an online game where players can predict World Cup 2022 matches scores and compete with other players.

Players can create private tournaments where the invite friends to compete.
Tournaments can change their scoring system and prediction strategies.

This API solves the backend part of the game, and it is consumed only by an Android app (game app) and a web admin.

### Game definitions
 - Scores are updated live, but points are only added after a match ends.
 - Multiple scoring strategies
 - Multiple prediction strategies
 - Private tournaments require a short password to join
 - Public tournaments can be joined by anyone who knows the tournament code (id)
 - Players can have different predictions for different tournaments
 - Player's predictions can be hidden (tournament setting)
 - Tournaments have settings:
   - scoring strategy
   - prediction strategy
   - prediction visibility
 - Tournaments have a max size of 10 players
 - Tournament owners can remove (and ban) players from tournaments

#### Scoring Strategies:
- Winner: players predict which team wins or if they tie
- Score: players predict the exact match result
- Combined: players predict the exact match result, but they also get points if they miss the final result but guess the winner correctly

#### Prediction Strategies:
- All: players have to predict all cup matches
- Current phase: players have to predict only current phase matches 

#### Prediction Visibility:
- Hidden: other players predictions can only be seen after a match ends
- Visible: other players predictions can be seen at anytime

### Technicalities
- All data will be stored on a relational database
- The API should send push notifications to players with match score updates
- We should consume a public API with matches scores
  - If available, we should implement a webhook to receive match scores updates from this service
- All endpoints should have authentication implemented to make sure players can only see, make, and update their own predictions / tournaments
- There is a Docker compose file that spins up a local MySQL instance

### Endpoints
[link](endpoints.md)

### Progress
[link](todos.md)

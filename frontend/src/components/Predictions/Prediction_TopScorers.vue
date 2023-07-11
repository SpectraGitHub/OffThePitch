<!--Component which make it possible for the user to predict three topscorers-->
<template>
  <!--Displays if playersloaded-->
  <div class="TopScorerComp" v-if="playersloaded">
    <!--Displays if haspredictions-->
    <div class="teamNameAndFlagBox" v-if="haspredictions">
      <!--Using the SelectedTopScorers compontent with the players props to display the selected top scorers-->
      <SelectedTopScorers :players="prediction.topscorers" />
    </div>
    <!--Checking if the first round has started-->
    <div class="TopScorerTitle" v-if="!this.roundOneStarted">
      <h1>SELECT 3 DIFFERENT TOPSCORERS</h1>
    </div>
    <div class="topScorers" v-if="!this.roundOneStarted">
      <!--Using the TopScorerChoice component with props to display all the possible top scorers-->
      <TopScorerChoice
        class="TopScorerChoice"
        :AllPossiblePlayers="teams"
        :PresetPredictions="prediction.topscorers[0]"
        :Index="0"
        v-if="playersloaded"
        @saveSelectedTopscorers="savePlayer"
      />
      <!--Using the TopScorerChoice component with props to display all the possible top scorers-->
      <TopScorerChoice
        class="TopScorerChoice"
        :AllPossiblePlayers="teams"
        :PresetPredictions="prediction.topscorers[1]"
        :Index="1"
        v-if="playersloaded"
        @saveSelectedTopscorers="savePlayer"
      />
      <!--Using the TopScorerChoice component with props to display all the possible top scorers-->
      <TopScorerChoice
        class="TopScorerChoice"
        :AllPossiblePlayers="teams"
        :PresetPredictions="prediction.topscorers[2]"
        :Index="2"
        v-if="playersloaded"
        @saveSelectedTopscorers="savePlayer"
      />
    </div>
    <!--Not possible to set top scorers after the first round has started-->
    <div v-if="this.roundOneStarted">
      <h2>
        The first round has already started. You can no longer choose your
        topscorers.
      </h2>
    </div>

    <div class="button" v-if="!this.roundOneStarted">
      <button class="submitButton" @click="submitPredictions">
        Sumbit Top Scorers
      </button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import TopScorerChoice from "./TopScorerChoice.vue";
import SelectedTopScorers from "./Prediction_SelectedTopScorers.vue";

export default {
  name: "TopScorers",
  props: {
    roundOneStarted: {
      required: true,
    },
  },
  data() {
    return {
      //returns a teams scruct
      teams: [
        {
          team_name: undefined,
          team_id: undefined,
          players: [
            {
              id: undefined,
            },
          ],
        },
      ],
      //returns a prediction struct with three topscorers
      prediction: {
        topscorers: [
          {
            goals: undefined,
            player_id: undefined,
            team_id: undefined,
            player_name: undefined,
          },
          {
            goals: undefined,
            player_id: undefined,
            team_id: undefined,
            player_name: undefined,
          },
          {
            goals: undefined,
            player_id: undefined,
            team_id: undefined,
            player_name: undefined,
          },
        ],
      },
      playersloaded: false,
      haspredictions: false,
    };
  },
  async created() {
    try {
      //Getting the savedPlayers
      let savedPlayers = await this.getPlayers();
      this.teams = savedPlayers.teams;
      //Getting the topscorers
      let pred = await this.getTopScorers();

      //Checking if there are any topscorers
      if (pred.topscorers != null) {
        //Looping through topscorers and setting the different variables
        for (let i = 0; i < pred.topscorers.length; i++) {
          this.prediction.topscorers[i] = pred.topscorers[i];
          this.prediction.topscorers[i].player_name = this.getPlayerName(
            this.prediction.topscorers[i].player_id,
            this.prediction.topscorers[i].team_id
          );
          this.prediction.topscorers[i].team_name = this.getTeamName(
            this.prediction.topscorers[i].team_id
          );
          //Setting has predictions to true
          this.haspredictions = true;
        }
      }
      //setting playersloaded to true
      this.playersloaded = true;
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    //Function that gets players from the backend
    async getPlayers() {
      try {
        //get request to backend
        const response = await axios.get(
          this.backendUrl + "/otp/getallplayers"
        );
        return response.data;
      } catch (e) {
        console.error(e);
      }
    },
    async getTopScorers() {
      try {
        //Setting the headers for get
        axios.defaults.headers.get["token"] = this.getToken();
        //Get request to backend to get the topscorers
        const response = await axios.get(
          this.backendUrl + "/otp/gettopscorers"
        );
        return response.data;
      } catch (e) {
        console.error(e);
        //Session expired, redirect to login, clearing session storage and printing a message to the user
        location.href = "/login";
        sessionStorage.clear();
        alert("Session expired, please login again");
      }
    },
    //gets the current user's session token
    getToken() {
      return sessionStorage.getItem("token");
    },
    //Function that saves the players id and team id
    savePlayer(args) {
      let pred = args.p;
      let index = args.i;

      this.prediction.topscorers[index].player_id = pred.player_id;
      this.prediction.topscorers[index].team_id = pred.team_id;
    },
    //If the submit predictions button is pressed, this function submits the topscorers
    submitPredictions() {
      const data = {
        topscorers: [],
      };
      //Looping through the prediction.topscorers
      for (let i = 0; i < this.prediction.topscorers.length; i++) {
        data.topscorers.push(this.prediction.topscorers[i]);
      }
      //Checking for duplicate choices
      if (this.duplicateChoice(data)) {
        return;
      }

      const predJSON = JSON.stringify(data);
      //Setting the headers for the post request
      axios.defaults.headers.post["token"] = this.getToken();
      //Post requst to backend witht the topscorers data
      axios
        .post(this.backendUrl + "/otp/savetopscorers", predJSON)
        .then(function () {}.bind(this), alert("Players saved"));
    },
    //Function to get the players name when the playerID and teamID is giben
    getPlayerName(playerID, teamID) {
      //Looping through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checking for match
        if (this.teams[i].team_id == teamID) {
          //Looping through the players on the team
          for (let y = 0; y < this.teams[i].players.length; y++) {
            //Checking for match
            if (this.teams[i].players[y].id == playerID) {
              //returning the name
              return (
                this.teams[i].players[y]["first-name"] +
                " " +
                this.teams[i].players[y]["last-name"]
              );
            }
          }
        }
      }
    },
    //Function that gets the teamname from the teamID
    getTeamName(teamID) {
      //Looping through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checking for a match
        if (this.teams[i].team_id == teamID) {
          //Returning the name of the team
          return this.teams[i].team_name;
        }
      }
    },
    //Checking for duplicate choices
    duplicateChoice(data) {
      // If any two of the topscorers have the same ID, return true, else return false
      if (data.topscorers.length < 3) {
        alert("Need three topscorers");
        return true;
      }
      //Checking for duplicates
      if (
        data.topscorers[0].player_id == data.topscorers[1].player_id ||
        data.topscorers[0].player_id == data.topscorers[2].player_id ||
        data.topscorers[1].player_id == data.topscorers[2].player_id
      ) {
        alert("You can not choose the same topscorer twice.");
        return true;
      }
      return false;
    },
  },
  components: { TopScorerChoice, SelectedTopScorers },
};
</script>

<style scoped>
.TopScorerTitle {
  display: flex;
  justify-content: center;
}

.button {
  display: flex;
  justify-content: center;
  align-items: center;
}
.submitButton {
  margin-top: 40px;
  border-radius: 20px;
  font-size: 24px;
  background-color: rgb(127, 87, 200);
}

.submitButton:hover {
  cursor: pointer;
  font-size: 30px;
}

.TopScorerComp {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
  flex-direction: column;
}

.teamNameAndFlagBox {
  width: 25%;
}

@media (max-width: 1250px) {
  .teamNameAndFlagBox {
    width: 40%;
  }
}

@media (max-width: 750px) {
  .teamNameAndFlagBox {
    width: 60%;
  }
}

@media (max-width: 500px) {
  .teamNameAndFlagBox {
    width: 80%;
  }
}

@media (max-width: 1250px) {
  .topScorers {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 40%;
  }
}

@media (max-width: 750px) {
  .topScorers {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 60%;
  }
}

@media (max-width: 500px) {
  .topScorers {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 80%;
  }
}

@media (min-width: 1250px) {
  .topScorers {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 25%;
  }
}
</style>

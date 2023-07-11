<!--Component which display the topscorers-->
<template>
  <div class="TopScorerComp" v-if="playersloaded">
    <div class="teamNameAndFlagBox" v-if="haspredictions">
      <!--Using the SelectedTopScorers component to get the topscorers-->
      <SelectedTopScorers :players="prediction.topscorers" />
    </div>
  </div>
</template>

<script>
import axios from "axios";
import SelectedTopScorers from "../Predictions/Prediction_SelectedTopScorers.vue";

export default {
  name: "TopScorers",
  data() {
    return {
      //Struct with the topscorer predictions
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
  //Setting the props
  props: {
    userid: {
      required: true,
    },
  },
  async created() {
    try {
      //setting savedPlayers to getPlayers()
      let savedPlayers = await this.getPlayers();
      //Updating this.teams
      this.teams = savedPlayers.teams;
      //Setting pred to getTopScorers()
      let pred = await this.getTopScorers();

      if (pred.topscorers != null) {
        //Looping through the topscorers and setting the variables
        for (let i = 0; i < pred.topscorers.length; i++) {
          this.prediction.topscorers[i] = pred.topscorers[i];
          this.prediction.topscorers[i].player_name = this.getPlayerName(
            this.prediction.topscorers[i].player_id,
            this.prediction.topscorers[i].team_id
          );
          this.prediction.topscorers[i].team_name = this.getTeamName(
            this.prediction.topscorers[i].team_id
          );
          this.haspredictions = true;
        }
      }

      this.playersloaded = true;
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    //Function that get all players from backend with a get request
    async getPlayers() {
      try {
        const response = await axios.get(
          this.backendUrl + "/otp/getallplayers"
        );
        return response.data;
      } catch (e) {
        console.error(e);
      }
    },
    //Function that get the topscorers from backend with a post request
    async getTopScorers() {
      const data = {
        userid: parseInt(this.userid),
      };
      const predJSON = JSON.stringify(data);

      const response = await axios.post(
        this.backendUrl + "/otp/gettopscorers",
        predJSON
      );
      return response.data;
    },
    //A function to get the player name from their playerid and teamid
    getPlayerName(playerID, teamID) {
      //Looping through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checking for a match
        if (this.teams[i].team_id == teamID) {
          //Looping through the players
          for (let y = 0; y < this.teams[i].players.length; y++) {
            //Checking for a match
            if (this.teams[i].players[y].id == playerID) {
              //returns the name
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
    //Getting the team name from teamid
    getTeamName(teamID) {
      //Loops through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checking for a match
        if (this.teams[i].team_id == teamID) {
          //Returns the team name
          return this.teams[i].team_name;
        }
      }
    },
  },
  components: { SelectedTopScorers },
};
</script>

<style scoped>
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
</style>

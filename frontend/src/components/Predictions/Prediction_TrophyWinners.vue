<!--Compontent for displaying the prediciton trophywinners-->
<template>
  <!--Checking if the first round has started-->
  <div class="trophy" v-if="!this.roundOneStarted">
    <div class="trophyboxes">
      <h2>Gold</h2>
      <select v-model="this.predictions.gold">
        <!--Looping through the goldMedalOptions and displaying the team names in a option selection-->
        <option v-for="team in this.goldMedalOptions" :key="team.id">
          {{ team.name }}
        </option>
      </select>
      <div>
        <!--Using the FlagLogo component to display the flag if predsloaded and if there are any predicitons for the gold medal-->
        <FlagLogo
          id="goldflag"
          :TeamName="this.predictions.gold"
          v-if="predsloaded && this.predictions.gold"
        />
      </div>
    </div>
    <div class="trophyboxes">
      <h2>Silver</h2>
      <select v-model="this.predictions.silver">
        <!--Looping through the silverMedalOptions and displaying the team names in a option selection-->
        <option v-for="team in this.silverMedalOptions" :key="team.id">
          {{ team.name }}
        </option>
      </select>
      <div>
        <!--Using the FlagLogo component to display the flag if predsloaded and if there are any predicitons for the silver medal-->
        <FlagLogo
          :TeamName="this.predictions.silver"
          v-if="predsloaded && this.predictions.silver"
        />
      </div>
    </div>
    <div class="trophyboxes">
      <h2>Bronze</h2>
      <select v-model="this.predictions.bronze">
        <!--Looping through the bronzeMedalOptions and displaying the team names in a option selection-->
        <option v-for="team in this.bronzeMedalOptions" :key="team.id">
          {{ team.name }}
        </option>
      </select>
      <div>
        <!--Using the FlagLogo component to display the flag if predsloaded and if there are any predicitons for the bronze medal-->
        <FlagLogo
          :TeamName="this.predictions.bronze"
          v-if="renderComponent && this.predictions.bronze"
        />
      </div>
    </div>
  </div>
  <div class="button" v-if="!this.roundOneStarted">
    <button class="submitButton" @click="submitPredictions">
      Submit Trophy Winners
    </button>
  </div>
  <!--Checking if roundOne has started and then using the show_Prediction_TrophyWinners component to display the selected trophywinners-->
  <div v-if="this.roundOneStarted">
    <show_Prediction_TrophyWinners :userid="this.predictions.userid" />
    <h2 class="noChangeMsg">You can no longer change your medal predictions</h2>
  </div>
</template>

<script>
import axios from "axios";
import FlagLogo from "../Flag/FlagLogo.vue";
import show_Prediction_TrophyWinners from "../showPredictions/show_Prediction_TrophyWinners.vue";

export default {
  name: "Prediction_TrophyWinners",
  //Props to use the component in another component
  props: {
    roundOneStarted: {
      required: true,
    },
  },
  computed: {
    //Filter out the teams that could not be selected as options for each of the gold, silver and bronze. For example is a team have been already selected for gold, it should not be possible to select the same team for silver.
    goldMedalOptions: function () {
      return this.teams.filter(
        (i) =>
          i.name != this.predictions.silver && i.name != this.predictions.bronze
      );
    },
    silverMedalOptions: function () {
      return this.teams.filter(
        (i) =>
          i.name != this.predictions.gold && i.name != this.predictions.bronze
      );
    },
    bronzeMedalOptions: function () {
      return this.teams.filter(
        (i) =>
          i.name != this.predictions.silver && i.name != this.predictions.gold
      );
    },
  },
  data() {
    return {
      renderComponent: true,
      //teams struct that will be filled with information from backend
      teams: [
        {
          name: undefined,
          id: undefined,
        },
      ],
      //predictions struct which will be filled with information from backend
      predictions: {
        userid: undefined,
        gold: undefined,
        silver: undefined,
        bronze: undefined,
      },
      predsloaded: false,
    };
  },
  async created() {
    try {
      //Get request from backend to get all the teams
      const response = await axios.get(this.backendUrl + "/otp/getteams");
      this.teams = response.data;
      //Getting the medal predictions
      let savedMedals = await this.getMedalPredictions();
      //Setting the gold, silver and bronze variables
      this.predictions = {
        gold: this.getTeamName(savedMedals.gold),
        silver: this.getTeamName(savedMedals.silver),
        bronze: this.getTeamName(savedMedals.bronze),
      };
      //Setting predsloaded to true
      this.predsloaded = true;
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    //Function that gets the medal predictions from backend
    async getMedalPredictions() {
      try {
        axios.defaults.headers.get["token"] = this.getToken();
        //get request to backend to get the medals
        const response = await axios.get(this.backendUrl + "/otp/getmedals");
        return response.data;
      } catch (e) {
        console.error(e);
        //If session expired, redirect to login, clearing the session storage and printing a message to the user
        location.href = "/login";
        sessionStorage.clear();
        alert("Session expired, please login again");
      }
    },
    getToken() {
      return sessionStorage.getItem("token");
    },
    //Function to get the team id from the team name
    getTeamID(teamName) {
      //Looping through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checking for match
        if (this.teams[i].name == teamName) {
          //Returns the teams[i].id
          return this.teams[i].id;
        }
      }
    },
    //Function to get the team name from teh team id
    getTeamName(teamID) {
      //Loops through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checking for a match
        if (this.teams[i].id == teamID) {
          //returning the name
          return this.teams[i].name;
        }
      }
    },
    //Function that submits the predictions
    submitPredictions() {
      const data = {
        gold: this.getTeamID(this.predictions.gold),
        silver: this.getTeamID(this.predictions.silver),
        bronze: this.getTeamID(this.predictions.bronze),
      };
      const predJSON = JSON.stringify(data);
      axios.defaults.headers.post["token"] = this.getToken();
      //Post request to save the medals to backend
      axios
        .post(this.backendUrl + "/otp/savemedals", predJSON)
        .then(function () {}.bind(this), alert("Predictions saved"));
    },
  },
  components: { FlagLogo, show_Prediction_TrophyWinners },
};
</script>

<style scoped>
.trophy {
  margin-top: 10px;
  display: flex;
  justify-content: space-around;
  flex-direction: row;
  width: 100%;
  height: auto;
}

.trophyboxes {
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 100px;
  border-color: black;
  outline-style: solid;
  height: 130px;
  width: 30%;
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
  margin-top: 20px;
}

.trophyboxes > select {
  height: 30%;
  width: 70%;
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

.noChangeMsg {
  display: flex;
  justify-content: center;
}

@media (max-width: 750px) {
  .trophy {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .trophyboxes {
    width: 60%;
  }
}
@media (max-width: 450px) {
  .trophyboxes {
    width: 80%;
  }
}
</style>

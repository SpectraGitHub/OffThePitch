<!--Component to display the trophy winners for the different users in the leagues-->
<template>
  <div class="trophy">
    <div class="trophyboxes" @change="updateFlag">
      <h2>Gold</h2>
      <h4>{{ this.predictions.gold }}</h4>
      <div>
        <!--Setting the team logo with the FlagLogo component-->
        <FlagLogo
          id="goldflag"
          :TeamName="this.predictions.gold"
          v-if="predsloaded && this.predictions.gold"
        />
      </div>
    </div>
    <div class="trophyboxes">
      <h2>Silver</h2>
      <h4>{{ this.predictions.silver }}</h4>
      <div>
        <!--Setting the team logo with the FlagLogo component-->
        <FlagLogo
          :TeamName="this.predictions.silver"
          v-if="predsloaded && this.predictions.silver"
        />
      </div>
    </div>
    <div class="trophyboxes">
      <h2>Bronze</h2>
      <h4>{{ this.predictions.bronze }}</h4>
      <div>
        <!--Setting the team logo with the FlagLogo component-->
        <FlagLogo
          :TeamName="this.predictions.bronze"
          v-if="predsloaded && this.predictions.bronze"
        />
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import FlagLogo from "../Flag/FlagLogo.vue";

export default {
  name: "show_Prediction_TrophyWinners",
  data() {
    return {
      renderComponent: true,
      teams: [
        {
          name: undefined,
          id: undefined,
        },
      ],
      predictions: {
        gold: undefined,
        silver: undefined,
        bronze: undefined,
      },
      predsloaded: false,
    };
  },
  //Setting the props so this component could recive variables if it's used in another component
  props: {
    userid: {
      required: true,
    },
    isownpreds: {
      required: false,
    },
  },
  async created() {
    try {
      //Getting teams from backend on create
      const response = await axios.get(this.backendUrl + "/otp/getteams");
      this.teams = response.data;
      //Getting the savedMedals
      let savedMedals = await this.getMedalPredictions();
      this.predictions = {
        gold: this.getTeamName(savedMedals.gold),
        silver: this.getTeamName(savedMedals.silver),
        bronze: this.getTeamName(savedMedals.bronze),
      };
      this.predsloaded = true;
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    //Function to get the medal prediction with a get request from backend
    async getMedalPredictions() {
      if (this.isownpreds) {
        axios.defaults.headers.get["token"] = this.getToken();
        const response = await axios.get(this.backendUrl + "/otp/getmedals");
        return response.data;
      }
      //The userid of the user which is pressed
      const data = {
        userid: parseInt(this.userid),
      };
      const predJSON = JSON.stringify(data);

      //post request
      const response = await axios.post(
        this.backendUrl + "/otp/getmedals",
        predJSON
      );
      return response.data;
    },
    //Function to get the team name with the team id
    getTeamName(teamID) {
      //Loops through the teams
      for (let i = 0; i < this.teams.length; i++) {
        //Checks for match
        if (this.teams[i].id == teamID) {
          //Returns the name
          return this.teams[i].name;
        }
      }
    },
    //To get the token
    getToken() {
      return sessionStorage.getItem("token");
    },
  },
  components: { FlagLogo },
};
</script>

<style scoped>
.trophy {
  margin-top: 10px;
  margin-bottom: 10px;
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
  height: 120px;
  width: 30%;
  display: flex;
  justify-content: center;
  margin-top: 20px;
  margin-bottom: 20px;
}

.trophyboxes > select {
  height: 30%;
  width: 70%;
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

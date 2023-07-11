<!--This compontent creates a new league where the user can type in a league name-->
<template>
  <Top_Bar />
  <div class="createNewLeagueDiv">
    <div class="createNewLeague">
      <h2>Create a new league</h2>
      <p>League name</p>
      <input
        type="text"
        placeholder="Type the league name"
        v-model="leaguename"
      />
      <button class="createNewButton" @click="getInfo()">Create League</button>
      <p>
        When the league is created, the league is given an unique league code,
        <br />
        which you could use to invite players to the league.
      </p>
    </div>
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";
import axios from "axios";

export default {
  components: { Top_Bar },
  data: function () {
    return {
      leaguename: "",
    };
  },
  methods: {
    //A function that checks if the input league name
    async getInfo() {
      //If name shorter than 5, an alert will be shown on the screen for the user
      if (this.leaguename.length < 5) {
        alert("Leaguename has to be at least 5 characters");
        return;
      }
      const createLeague = {
        league_name: this.leaguename,
      };
      const createLeagueJSON = JSON.stringify(createLeague);
      //Setting the headers for a post
      axios.defaults.headers.post["token"] = this.getToken();
      //Post request to create a league
      axios.post(this.backendUrl + "/otp/createleague", createLeagueJSON).then(
        function (response) {
          //Checks for valid response
          if (response.status < 300) {
            alert("League created!");
          } else {
            alert("Something went wrong. Please try again later");
          }
        }.bind(this),
        this.relocate()
      );
    },
    //Function used for getting the current users token
    getToken() {
      return sessionStorage.getItem("token");
    },
    //function to relocate the user window
    relocate() {
      location.href = "/leagues";
    },
  },
};
</script>

<style scoped>
.createNewLeagueDiv {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 50px;
}

.createNewButton {
  font-size: large;
  border-radius: 20px;
  border-color: grey;
}

.createNewButton:hover {
  cursor: pointer;
}
</style>

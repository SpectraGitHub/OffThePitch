<!--The component displays the join leauge page, where the user can join a leauge with a league code-->
<template>
  <Top_Bar />
  <div class="joinLeagueDiv">
    <div class="joinLeague">
      <h2>Join a leauge</h2>
      <p>League code</p>
      <input
        type="text"
        v-model="joinleague"
        placeholder="Type in the league code"
      />
      <button class="joinLeagueButton" @click="getInfo()">Join League</button>
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
      joinleague: "",
    };
  },
  methods: {
    //Function that lets the user join a league
    async getInfo() {
      const joinLeague = {
        league_code: this.joinleague,
      };
      const joinLeagueJSON = JSON.stringify(joinLeague);
      //Setting the headers for post
      axios.defaults.headers.post["token"] = this.getToken();
      //Post request to backend to join a leauge
      var response = await axios
        .post(this.backendUrl + "/otp/joinleague", joinLeagueJSON)
        .then(function (response) {
          //Checks the status code
          if (response.status < 300) {
            return response.status;
          }
        })
        //Checking for an error status code
        .catch(function (error) {
          if (error.response.status >= 400) {
            return error.response.status;
          }
        });
      //If the response status code is below 300, then the league is joined and the user will be relocated.
      if (response < 300) {
        alert("League joined!");
        this.relocate();
      }
      //Wrong league code returns an error message to the user
      if (response >= 400) {
        alert("No leagues with the given league code!");
      }
    },
    //Function to get the current users token for this session
    getToken() {
      return sessionStorage.getItem("token");
    },
    //Relocate the user to the leagues page
    relocate() {
      location.href = "/leagues";
    },
  },
};
</script>

<style scoped>
.joinLeagueDiv {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 50px;
}

.joinLeagueButton {
  font-size: large;
  border-radius: 20px;
  border-color: grey;
}

.joinLeagueButton:hover {
  cursor: pointer;
}
</style>

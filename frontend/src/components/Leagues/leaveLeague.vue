<!--Leave league compontent which makes the user leave a league if the leave league button is pressed-->
<template>
  <div class="leaveLeague">
    <button class="leaveButton" @click="getInfo()">Leave League</button>
  </div>
</template>

<script>
import axios from "axios";

export default {
  //Props so that the component could be used in another component and pass a variable back to this component
  props: {
    League_ID: {
      required: true,
    },
  },
  methods: {
    async getInfo() {
      const leaveLeague = {
        league_id: this.League_ID,
      };

      const leaveLeagueJSON = JSON.stringify(leaveLeague);
      //Setting the headers for post
      axios.defaults.headers.post["token"] = this.getToken();
      //POst request to backend
      axios
        .post(this.backendUrl + "/otp/leaveleague", leaveLeagueJSON)
        .then(
          function () {}.bind(this),
          alert("You left the league!"),
          this.refresh()
        );
    },
    //Getting the session token
    getToken() {
      return sessionStorage.getItem("token");
    },
    //relocate the user to the leagues page
    refresh() {
      location.href = "/leagues";
    },
  },
};
</script>

<style scoped>
.leaveButton:hover {
  cursor: pointer;
}

.leaveButton {
  font-size: large;
  border-radius: 20px;
  border-color: grey;
}
</style>

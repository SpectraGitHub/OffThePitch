<!--A component which display all the leagues on the leauges page-->
<template>
  <!-- This div loops through all leagues in leagueinfo reverse and splits up the leagueinfo -->
  <div
    class="leagueTables"
    v-for="leagues in leagueinfo.slice().reverse()"
    :key="leagues.league_id"
  >
    <!--Displays the league name at the top of each table-->
    <h2 class="leagueName">{{ leagues.league_name }}</h2>
    <table class="eachTable">
      <!--A table row with 3 table headers, rank, name and points-->
      <tr>
        <th class="rankTableHeader">Rank</th>
        <th class="nameTableHeader">Name</th>
        <th class="pointsTableHeader">Points</th>
      </tr>
      <!--Looping through and setting all the players rank, name and points-->
      <tr v-for="(numRank, index) in leagues.players.length" :key="index">
        <td class="numbers">
          {{ numRank }}
        </td>
        <!--Router-link on the players name so that is possible to click it-->
        <td class="names">
          <router-link
            class="link"
            :to="'/predictions/' + leagues.players[index].userid"
            >{{ leagues.players[index].user_name }}</router-link
          >
        </td>
        <td class="points">
          {{ leagues.players[index].score }}
        </td>
      </tr>
    </table>
    <!--Checking if the league id is not the id of the overall league which everyone is joining automatically-->
    <h4 v-if="leagues.league_id != 1">
      <!--Displays the league code-->
      <br />Invite others to the league:
      {{ leagues.league_code }}
    </h4>
    <!--Using the leaveLeague component and checks if the league id is not the id of the overall league-->
    <leaveLeague v-if="leagues.league_id != 1" :League_ID="leagues.league_id" />
  </div>
</template>

<script>
import axios from "axios";
import leaveLeague from "./leaveLeague.vue";
export default {
  components: { leaveLeague },
  data() {
    return {
      //A leagueinfo struct which will be updated from backend
      leagueinfo: [
        {
          league_id: undefined,
          league_code: undefined,
          league_name: undefined,
          players: [
            {
              userid: undefined,
              user_name: undefined,
              score: undefined,
            },
          ],
        },
      ],
      leaguesloaded: false,
      showMatchPred: false,
    };
  },
  async created() {
    try {
      //Getting the user leagues
      let savedUserLeagues = await this.getUserLeagues();
      //Checks if the user is not part of any leagues
      if (savedUserLeagues != null) {
        //Looping through savedUserLeagues and setting the leagueinfo[i] to the savedUserLeagues[i]
        for (let i = 0; i < savedUserLeagues.length; i++) {
          this.leagueinfo[i] = savedUserLeagues[i];
        }
        //Setting the leaguesloaded variable to true. This was because we had some problems that the function was done without
        //all the leagues was loaded.
        this.leaguesloaded = true;
      }
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    //Function that get the user leagues from backend
    async getUserLeagues() {
      try {
        //Sets the headers for get
        axios.defaults.headers.get["token"] = this.getToken();
        //Get request from backend
        const response = await axios.get(this.backendUrl + "/otp/getleagues");
        return response.data;
      } catch (e) {
        console.error(e);
        //Relocates to login is session has expired, also clear the sessionStorage and output a message to the user
        location.href = "/login";
        sessionStorage.clear();
        alert("Session expired, please login again");
      }
    },
    //Gets the current users token
    getToken() {
      return sessionStorage.getItem("token");
    },
  },
};
</script>

<style scoped>
.leagueTables td,
.leagueTables th {
  border: 1px solid #ddd;
  padding: 8px;
}

.leagueTables tr:nth-child(even) {
  background-color: #f2f2f2;
}

.leagueTables tr:hover {
  background-color: #ddd;
}

.leagueTables th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: rgb(47, 62, 163);
  color: white;
}

.leagueName {
  display: flex;
  justify-content: center;
  overflow: auto;
}
.numbers {
  display: flex;
  justify-content: center;
}

.points {
  display: flex;
  justify-content: center;
}

.leagueTables {
  padding-bottom: 30px;
  width: 40%;
  margin-left: auto;
  margin-right: auto;
}

.rankTableHeader {
  width: 20%;
}
.nameTableHeader {
  width: 50%;
}
.pointsTableHeader {
  width: 30%;
}

.eachTable {
  width: 100%;
}

.link {
  text-decoration: none;
  color: black;
}

@media (max-width: 650px) {
  .leagueTables {
    width: 60%;
  }
}

@media (max-width: 450px) {
  .leagueTables {
    width: 95%;
  }
}
</style>

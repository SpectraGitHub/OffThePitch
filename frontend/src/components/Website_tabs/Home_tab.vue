<!--Component which displays the home tab-->
<template>
  <!--Header component-->
  <Top_Bar />
  <div>
    <div class="main">
      <!--This displays the leagues section of the home page, and it uses the LeagueHomePage component-->
      <div class="leaguehomebox" v-if="leaguesloaded">
        <h2 class="title">Your leagues:</h2>
        <div v-for="league in this.leagues" :key="league">
          <LeagueHomepage :League="league" />
        </div>
      </div>
      <!--This displays the selected topscorers section of the home page-->
      <div class="PredictedTopScorersComp" v-if="topscorersloaded">
        <SelectedTopScorers :players="this.prediction.topscorers" />
      </div>

      <!--MatchDisplay is used to display the 4 upcoming matches-->
      <div class="upcomingMatches" v-if="matchesloaded">
        <div class="list">
          <h2>Upcoming Matches:</h2>
          <div v-for="match in this.matches" :key="match" class="item">
            <MatchDisplay class="matchDisplay" :Fixture="match" />
          </div>
        </div>
      </div>

      <!--show_Predictions_TrophyWinners component is used so display the trophywinners on the home screen-->
      <div class="box" v-if="medalloaded">
        <div class="trophybox">
          <h2>Selected Trophy Winners:</h2>
          <show_Prediction_TrophyWinners :userid="this.id" :isownpreds="true" />
        </div>
      </div>
    </div>
    <div class="welcomemessage" v-if="welcome">
      <h2 class="borderwelcometext">
        Welcome to Off The Pitch! Login or Register to start predicting!
      </h2>
    </div>
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";
import SelectedTopScorers from "../Predictions/Prediction_SelectedTopScorers.vue";
import MatchDisplay from "../matchFixture/MatchDisplay.vue";
import LeagueHomepage from "../Leagues/LeagueHomepage.vue";
import axios from "axios";
import show_Prediction_TrophyWinners from "../showPredictions/show_Prediction_TrophyWinners.vue";

export default {
  name: "Home_Screen",
  data: function () {
    return {
      id: undefined,
      matches: undefined,
      //prediction struct with the topscorers
      prediction: {
        topscorers: [
          {
            goals: undefined,
            player_id: undefined,
            team_id: undefined,
            player_name: undefined,
            team_name: undefined,
          },
          {
            goals: undefined,
            player_id: undefined,
            team_id: undefined,
            player_name: undefined,
            team_name: undefined,
          },
          {
            goals: undefined,
            player_id: undefined,
            team_id: undefined,
            player_name: undefined,
            team_name: undefined,
          },
        ],
      },
      leagues: undefined,

      topscorersloaded: false,
      matchesloaded: false,
      leaguesloaded: false,
      medalloaded: false,
      welcome: false,
    };
  },
  components: {
    Top_Bar,
    SelectedTopScorers,
    MatchDisplay,
    LeagueHomepage,
    show_Prediction_TrophyWinners,
  },
  async created() {
    let data = await this.homeData();

    if (data == undefined) {
      this.welcome = true;
    }

    // selected topscorers
    if (data.topscorers.topscorers) {
      this.prediction.topscorers = data.topscorers.topscorers;
      this.topscorersloaded = true;
    }

    // upcoming matches
    if (data.matches.matches) {
      this.matches = data.matches.matches;
      this.matchesloaded = true;
    }

    // leagues
    if (data.leagues) {
      this.leagues = data.leagues;
      this.leaguesloaded = true;
    }

    // medal predictions
    if (
      data.medals.bronze != 0 &&
      data.medals.silver != 0 &&
      data.medals.gold != 0
    ) {
      this.id = data.userid;
      this.medalloaded = true;
    }
  },
  methods: {
    //A function that gets the homepageinfo from backend with a get request
    async homeData() {
      try {
        axios.defaults.headers.get["token"] = this.getToken();
        const response = await axios.get(
          this.backendUrl + "/otp/gethomepageinfo"
        );
        return response.data;
      } catch (e) {
        console.error(e); // SHOULD BE REMOVED ON RELEASE
        sessionStorage.clear();
      }
    },
    getToken() {
      return sessionStorage.getItem("token");
    },
  },
};
</script>

<style scoped>
.main {
  display: grid;
  gap: 30px;
  margin-top: 30px;
  margin-right: 20px;
  grid-template-areas: "z a f b b" "z c c b b";
  grid-template-columns: 20px 300px 430px 250px auto;
  width: 100%;
  height: 100%;
  background: white;
  padding-bottom: 0;
  border: solid 2px #fff;
}

.PredictedTopScorersComp {
  grid-area: a;
  width: 50%;
  margin-left: auto;
  margin-right: auto;
}

.upcomingMatches {
  margin-top: auto;
  margin-bottom: auto;
  width: 50%;
  grid-area: b;
  border-color: black;
  margin-left: auto;
  margin-right: auto;
}

.matchDisplay {
  border-color: black;
}

.item {
  margin-right: 20px;
}

.trophybox {
  border: 5px;
  border-color: black;
  border-style: solid;
  width: 50%;
  height: 250px;
  margin-left: auto;
  margin-right: auto;
}

.box {
  grid-area: f;
  width: 100%;
}

.trophybox > h2 {
  display: flex;
  justify-content: center;
}

.leaguehomebox {
  border: 5px;
  border-color: black;
  border-style: solid;
  width: 50%;
  grid-area: c;
  margin-left: auto;
  margin-right: auto;
}

.title {
  display: flex;
  justify-content: center;
}

.welcomemessage {
  display: flex;
  text-align: center;
  color: black;
  margin-top: 200px;
}
.borderwelcometext {
  border: 10px solid rgb(47, 62, 163);
  width: min(100% - 2rem);
  margin-inline: auto;
}

@media (min-width: 1200px) {
  .upcomingMatches {
    width: 100%;
  }
  .leaguehomebox {
    width: 100%;
  }

  .PredictedTopScorersComp {
    width: 100%;
  }
  .box {
    width: 100%;
  }

  .trophybox {
    width: 100%;
  }

  .welcomemessage {
    display: flex;
    justify-content: center;
    font-size: 25px;
  }
}

@media (max-width: 1200px) {
  .main {
    grid-template-areas: "a" "b" "c" "d";
    grid-template-columns: 1fr;
  }
  .upcomingMatches {
    grid-area: a;
  }
  .leaguehomebox {
    grid-area: b;
  }

  .PredictedTopScorersComp {
    grid-area: c;
  }
  .box {
    grid-area: d;
  }

  .welcomemessage {
    display: flex;
    justify-content: center;
    font-size: 18px;
  }
}

@media (max-width: 800px) {
  .main {
    grid-template-areas: "a" "b" "c" "d";
    grid-template-columns: 1fr;
  }
  .upcomingMatches {
    grid-area: a;
    width: 100%;
  }
  .leaguehomebox {
    grid-area: b;
    width: 100%;
  }

  .PredictedTopScorersComp {
    grid-area: c;
    width: 100%;
  }
  .box {
    grid-area: d;
    width: 100%;
  }
  .trophybox {
    width: 100%;
    height: 600px;
  }

  .welcomemessage {
    display: flex;
    justify-content: center;
    font-size: 12px;
  }
}
</style>

<!--This component will display all the matches in the fixture with their flag, name, and the rest of the information about the match-->
<template>
  <div class="matchDisplay">
    <!--The fixture date-->
    <h4 class="fixtureDateDisplay">{{ Fixture.date }}</h4>
    <h2 class="teamNamesDisplay">
      <div class="homeTeamName">
        <!--Using the FlagLogo component and passing in the home team name to display this flag on the screen-->
        <FlagLogo :TeamName="Fixture['home-team'].name" />
        <!--Outputting the home team name-->
        {{ Fixture["home-team"].name }}
      </div>
      <div class="dash">
        <!--Checking if it should be displayed the score of the match or the time-->
        {{ getScore(Fixture) }}
      </div>
      <div class="awayTeamName">
        <!--Outputting the away team name-->
        {{ Fixture["away-team"].name }}
        <!--Using the FlagLogo component and passing in the away team name to display this flag on the screen-->
        <FlagLogo :TeamName="Fixture['away-team'].name" />
      </div>
    </h2>
  </div>
</template>

<script scoped>
import FlagLogo from "../Flag/FlagLogo.vue";
export default {
  data() {
    return {
      image: FlagLogo,
    };
  },
  //Props so that the component could be used in another component and pass a variable back to this component
  props: {
    Fixture: {
      required: true,
    },
  },
  methods: {
    //Getting the scores or time from the fixture
    getScore(Fixture) {
      //Checking if it should display time
      if (Fixture.status.short != "FT") {
        return Fixture.time;
      } else {
        //Else displaying the score
        return Fixture["home-team"].score + "-" + Fixture["away-team"].score;
      }
    },
  },

  components: { FlagLogo },
};
</script>

<style scoped>
.matchDisplay {
  border-style: solid;
  border-width: 5px;
  border-color: blue;
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  margin: 7px;
  overflow: hidden;
}
.teamNamesDisplay {
  display: flex;
  width: 100%;
  padding-left: 10px;
  padding-right: 10px;
  flex-direction: row;
  justify-content: space-between;
  font-family: "Franklin Gothic Medium", "Arial Narrow", Arial, sans-serif;
  font-size: 10px;
  color: black;
}

.homeTeamName {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  font-size: large;
}

.awayTeamName {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
  font-size: large;
}

@media (max-width: 1350px) {
  .homeTeamName {
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
  }
  .awayTeamName {
    flex-direction: column;
    align-items: flex-end;
  }

  .awayTeamName > * {
    order: -1;
  }
}

.dash {
  display: flex;
  position: absolute;
  font-size: xx-large;
  width: 100%;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.fixtureDateDisplay {
  color: black;
  align-self: center;
}

.goalsScored {
  color: black;
  width: 100%;
}

.homeAndAwayGoals {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
}

.homeGoals {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  padding-left: 10px;
  align-items: flex-end;
  width: 200px;
  height: 60px;
  overflow: auto;
  direction: rtl;
}

.awayGoals {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-end;
  padding-right: 10px;
  width: 200px;
  height: 60px;
  overflow: auto;
  direction: ltr;
}

.goalsHeader {
  display: flex;
  justify-content: center;
}
</style>

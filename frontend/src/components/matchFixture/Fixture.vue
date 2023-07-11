<!--Component which display the different matchdays if the buttons are clicked, it's using the MatchDisplay compontent-->
<template>
  <div class="matchdayButtons">
    <button type="button" class="md" @click="chosenMatchday = 1">
      Matchday 1
    </button>
    <button type="button" class="md" @click="chosenMatchday = 2">
      Matchday 2
    </button>
    <button type="button" class="md" @click="chosenMatchday = 3">
      Matchday 3
    </button>
    <button type="button" class="md" @click="chosenMatchday = 4">
      Playoffs
    </button>
  </div>

  <ul v-if="chosenMatchday != 4" class="list">
    <li
      v-for="fixture in fixtures.slice(
        (chosenMatchday - 1) * 16,
        chosenMatchday * 16
      )"
      :key="fixture.id"
      class="item"
    >
      <!--Using the MatchDisplay compontent and sending in the fixture variable to display all the correct matches-->
      <MatchDisplay class="matchDisplay" :Fixture="fixture" />
    </li>
  </ul>
  <!--Playoffs-->
  <div v-if="chosenMatchday === 4">
    <h1 class="playoffTitle">Round of 16</h1>
    <ul class="list">
      <li
        class="item"
        v-for="fixture in fixtures.slice(48, 56)"
        :key="fixture.id"
      >
        <!--Using the MatchDisplay compontent and sending in the fixture variable to display all the correct matches-->
        <MatchDisplay class="matchDisplay" :Fixture="fixture" />
      </li>
    </ul>

    <h1 class="playoffTitle">Quarter Finals</h1>
    <ul class="list">
      <li
        class="item"
        v-for="fixture in fixtures.slice(56, 60)"
        :key="fixture.id"
      >
        <!--Using the MatchDisplay compontent and sending in the fixture variable to display all the correct matches-->
        <MatchDisplay class="matchDisplay" :Fixture="fixture" />
      </li>
    </ul>

    <h1 class="playoffTitle">Semi Finals</h1>
    <ul class="list">
      <li
        class="item"
        v-for="fixture in fixtures.slice(60, 62)"
        :key="fixture.id"
      >
        <!--Using the MatchDisplay compontent and sending in the fixture variable to display all the correct matches-->
        <MatchDisplay class="matchDisplay" :Fixture="fixture" />
      </li>
    </ul>

    <h1 class="playoffTitle">Third Place Finals</h1>
    <ul class="list">
      <li
        class="item"
        v-for="fixture in fixtures.slice(62, 63)"
        :key="fixture.id"
      >
        <!--Using the MatchDisplay compontent and sending in the fixture variable to display all the correct matches-->
        <MatchDisplay class="matchDisplay" :Fixture="fixture" />
      </li>
    </ul>

    <h1 class="playoffTitle">Finals</h1>
    <ul class="list">
      <li
        class="item"
        v-for="fixture in fixtures.slice(63, 64)"
        :key="fixture.id"
      >
        <!--Using the MatchDisplay compontent and sending in the fixture variable to display all the correct matches-->
        <MatchDisplay class="matchDisplay" :Fixture="fixture" />
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
import MatchDisplay from "../matchFixture/MatchDisplay.vue";

// get request fixtures, for å prøve post request kommenter vekk koden under
export default {
  name: "FixtureData",
  data() {
    return {
      fixtures: [],
      //Set to 1 at the start so the user will always see the first matchday
      chosenMatchday: 1,
    };
  },
  async created() {
    try {
      //get request to backend
      const response = await axios.get(this.backendUrl + "/otp/fixtures");
      //Getting all the fixtures
      this.fixtures = response.data;
    } catch (e) {
      console.error(e);
    }
  },

  components: { MatchDisplay },
};
</script>

<style scoped>
.item {
  color: white;
  width: 100%;
  display: flex;
  justify-content: center;
}

.playoffTitle {
  width: 100%;
  text-align: center;
}

.list {
  display: grid;
  grid-template-columns: 1fr 1fr;
  place-items: center;
  width: 100%;
  margin-top: 20px;
  padding: 0px 0px 0px 0px;
}

@media (max-width: 850px) {
  .list {
    grid-template-columns: 1fr;
  }

  .matchdayButtons {
    display: flex;
    flex-direction: column;
    width: 100%;
    align-items: center;
  }
  .matchdayButtons > button {
    display: flex;
    justify-content: center;
    margin-bottom: 5px;
  }
}

@media (min-width: 850px) {
  .matchdayButtons {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    width: 100%;
  }
}
.md {
  font-size: 26px;
  border-radius: 12px;
  color: black;
  border: 2px black;
}

.md:hover {
  cursor: pointer;
  background-color: rgb(47, 62, 163);
  box-shadow: 0 12px 16px 0 rgba(0, 0, 0, 0.24),
    0 17px 50px 0 rgba(0, 0, 0, 0.19);
  color: white;
}
</style>

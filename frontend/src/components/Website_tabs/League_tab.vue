<!--A component which use the LeagueDisplay component to display the leagues-->
<template>
  <Top_Bar />
  <h1 class="leaguesTitle">My leagues:</h1>
  <div class="leagueOptionsButtons">
    <button @click="joinLeague()">Join a new league</button>
    <button @click="createLeague()">Create a new league</button>
  </div>
  <div class="Leagues">
    <LeagueDisplay />
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";
import LeagueDisplay from "../Leagues/LeagueDisplay.vue";

export default {
  name: "Leagues_Screen",
  components: {
    Top_Bar,
    LeagueDisplay,
  },
  methods: {
    joinLeague() {
      location.href = "/joinleague";
    },
    createLeague() {
      location.href = "/createleague";
    },
  },
};
</script>

<style scoped>
.Leagues {
  margin-top: 30px;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
}

.leagueOptionsButtons {
  display: flex;
  justify-content: center;
}

.leagueOptionsButtons > button {
  font-size: large;
  border-radius: 20px;
  border-color: grey;
  margin: 5px;
}

.leagueOptionsButtons:hover > button {
  cursor: pointer;
}

.leaguesTitle {
  display: flex;
  justify-content: center;
}
</style>

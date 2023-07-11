<!--This component displays the leagues for the given player on the home screen-->
<template>
  <div class="leaguehomemain">
    <!--Looping through league.players and displays the leagues for the player on the home page-->
    <div v-for="(player, index) in League.players" :key="player">
      <!--Checking if the the username-->
      <div class="leaguesBox" v-if="player.user_name == getUser()">
        <div class="leagueName">
          <!--Display the league name-->
          <h3>{{ "League name: " + League.league_name }} <br /></h3>
        </div>
        <div class="place">
          <!--Setting the place in each league-->
          <h3>{{ "  Place: " + (index + 1) }} <br /></h3>
        </div>
        <div class="points">
          <!--Setting the players points-->
          <h3>{{ " Points: " + player.score }} <br /></h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  //Props so that the component could be used in another component and pass a variable back to this component
  props: {
    League: {
      required: true,
    },
  },
  methods: {
    //Gets the current user
    getUser() {
      return sessionStorage.getItem("currentUser");
    },
  },
};
</script>

<style scoped>
.leaguesBox {
  display: grid;
  grid-template-areas: "a b c";
  grid-template-columns: auto 70px 70px;
  gap: 15px;
}

.leagueName {
  grid-area: a;
  overflow: auto;
  align-items: start;
}
.place {
  grid-area: b;
}

.points {
  grid-area: c;
  align-items: end;
}
</style>

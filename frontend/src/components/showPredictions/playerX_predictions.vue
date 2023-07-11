<!--Component to display the predictions for all players in the league-->
<template>
  <div>
    <!--Using the show_Predictions component with the id of the player that has been pressed-->
    <show_Predicitons :id="this.$route.params.id" />
  </div>
</template>

<script>
import show_Predicitons from "./show_Predictions.vue";

export default {
  components: { show_Predicitons },
};
</script>

<style scoped></style>

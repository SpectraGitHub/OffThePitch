<!--Compontent to display the flag logo for each team-->
<template>
  <img class="logo" :src="'/Flags/' + img + '.png'" />
</template>

<script>
export default {
  name: "FlagLogo",
  //Props so this component could be used in another component and take a value TeamName
  props: {
    TeamName: {
      required: true,
    },
  },

  data: function () {
    return {
      img: this.TeamName.replace(" ", " "), // the module request
    };
  },
};
</script>

<style scoped>
.logo {
  height: 50px;
  width: 50px;
  border-radius: 50%;
  margin: 4px;
  border: 1px solid;
}
</style>

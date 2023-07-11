<!--Compontent to display the header at the top of each screen, all the bars in the header are router links and will redirect the user to the correct pages-->
<template>
  <div class="header">
    <nav class="navbar">
      <ul class="bar">
        <li>
          <router-link to="/">
            <img class="image" :src="image" />
          </router-link>
        </li>
        <li class="list_item">
          <router-link class="link" to="/">Home</router-link>
        </li>
        <li class="list_item">
          <router-link class="link" to="/leagues">Leagues</router-link>
        </li>
        <li class="list_item">
          <router-link class="link" to="/fixtures">Fixtures</router-link>
        </li>
        <li class="list_item">
          <router-link class="link" to="/my_predictions"
            >My Predictions</router-link
          >
        </li>
        <li class="list_item">
          <router-link class="link" to="/info_rules">Info/rules</router-link>
        </li>
        <li class="list_item">
          <h3 v-if="getState()" class="link">{{ getUser() }}</h3>
          <router-link v-if="getState()" class="link" to="/" @click="logOut()"
            >Log out
          </router-link>
          <router-link v-if="!getState()" class="link" to="/login"
            >Log in
          </router-link>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
import image from "../../assets/OTP_Logo.png";

export default {
  name: "Top_Bar",
  data: function () {
    return {
      image: image,
      state: Boolean,
      user: "",
    };
  },

  methods: {
    //Function that redirect to login
    onPress() {
      location.href = "/login";
    },
    //Clears session storage and reloads the window if the user logs out
    logOut() {
      sessionStorage.clear();
      window.location.reload();
    },
    openStorage(name) {
      return sessionStorage.getItem(name);
    },
    //Getting the current user
    getUser() {
      this.user = this.openStorage("currentUser");
      return this.user;
    },
    getState() {
      this.state = this.openStorage("currentState");
      if (this.state) {
        return this.state;
      } else {
        return false;
      }
    },
    alert() {
      alert("You have to log in to view this page!");
    },
  },

  components: {},
};
</script>

<style scoped>
.header {
  color: black;
  background-color: rgb(47, 62, 163);
  height: 100px;
  width: 100%;
}

.image {
  height: 60px;
  padding-left: 10px;
}

.image:hover {
  cursor: pointer;
}

.link {
  text-decoration: none;
  display: flex;
  justify-content: center;
  align-items: center;
  padding-left: 5px;
  color: white;
}
.list_item {
  list-style: none;
  display: flex;
  flex-direction: row;
}

.bar {
  padding: 0px 30px 0px 0px;
}

.navbar {
  height: 100%;
  width: 100%;
  display: flex;
  align-items: center;
}
.navbar a {
  color: #fff;
}

.navbar li:hover {
  font-size: 30px;
  border-radius: 20px;
  transition: 0.1s ease;
}

.navbar ul {
  display: flex;
  align-items: center;
  flex-direction: row;
  justify-content: space-between;
  width: 100%;
  font-size: 18px;
}

@media (max-width: 900px) {
  .list_item {
    font-size: medium;
  }

  .navbar li:hover {
    font-size: large;
    border-radius: 20px;
    transition: 0.1s ease;
  }

  .image {
    height: 40px;
  }
}

@media (max-width: 750px) {
  .list_item {
    font-size: small;
  }

  .navbar li:hover {
    font-size: medium;
    border-radius: 20px;
    transition: 0.1s ease;
  }

  .image {
    height: 20px;
  }
}

@media (max-width: 550px) {
  .navbar ul {
    display: flex;
    flex-direction: column;
    list-style-type: none;
  }
  .header {
    height: 200px;
  }
  .list_item {
    font-size: medium;
  }
  .navbar li:hover {
    font-size: large;
    border-radius: 20px;
    transition: 0.1s ease;
  }
}
</style>

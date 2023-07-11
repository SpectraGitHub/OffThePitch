<!--This component displays the login page where the user can log in-->
<template>
  <Top_Bar />
  <form>
    <div class="logInComp">
      <div class="logInTextComp">
        <h2 class="logInHeader">
          <h2>Log in or create a new user</h2>
          <img class="image" @click="onClick()" :src="image" />
        </h2>
      </div>
      <div class="logInForm">
        <div class="logInUserFormField">
          <label>Username </label>
          <input type="text" id="username" name="username" v-model="username" />
        </div>
        <div class="logInPasswordFormField">
          <label>Password </label>
          <input
            type="password"
            id="password"
            name="password"
            v-model="password"
          />
        </div>
        <div class="buttonDivForLogInPage">
          <button type="button" class="logInButton" @click="logIn()">
            Log in
          </button>
          <button
            type="button"
            class="createNewUserButtonForLogIn"
            @click="onClickButton()"
          >
            Create new user
          </button>
        </div>
      </div>
    </div>
  </form>
</template>

<script>
import image from "/icons/favicon.ico";
import Top_Bar from "../Header/Header.vue";
import axios from "axios";

export default {
  name: "LogIn_Screen",
  data: function () {
    return {
      image: image,
      //form struct
      form: {
        username: "",
        password: "",
      },
      //state struct to check the users state
      state: {
        currentUser: "",
        currentState: false,
      },
    };
  },
  methods: {
    //Function to redirect to home page
    onClick() {
      location.href = "/";
    },
    //Function to redirect to newuser page
    onClickButton() {
      location.href = "/newuser";
    },
    //Function that checks for the correct log in credentials
    async logIn() {
      const logindata = {
        USERNAME: this.username,
        PASSWORD: this.password,
      };

      const loginJSON = JSON.stringify(logindata);
      let token = "";
      //post request from backend
      var resp = await axios
        .post(this.backendUrl + "/otp/login", loginJSON)
        .then(
          async function (response) {
            token = await axios
              .post(this.backendUrl + "/otp/login", loginJSON)
              .then(function (response) {
                if (response.status == "202") {
                  return response;
                }
              })
              .catch(function (error) {
                return error.response;
              });
            return response;
          }.bind(this)
        )
        .catch(function (error) {
          alert("Username and/or Password Incorrect!");
          return error.response;
        });

      if (resp.status < 300) {
        // sets token, current user and state
        sessionStorage.setItem("token", token.data);
        sessionStorage.setItem("currentUser", this.username.toLowerCase());
        sessionStorage.setItem("currentState", true);
        this.onClick();
      } else {
        return;
      }
    },
  },

  components: { Top_Bar },
};
</script>

<style scoped>
.logInComp {
  top: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

.logInHeader {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 40px;
}

.image {
  left: 10px;
  height: 50px;
  border-radius: 10px;
}

.buttonDivForLogInPage {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.logInButton {
  top: 40px;
  height: 40px;
  width: 100%;
  border-radius: 20px;
  border-color: rgb(100, 18, 200);
  background-color: rgb(136, 136, 136);
}

.logInButton:hover {
  cursor: pointer;
}

.createNewUserButtonForLogIn {
  top: 60px;
  height: 40px;
  width: 100%;
  border-radius: 20px;
  border-color: rgb(100, 18, 200);
  background-color: rgb(136, 136, 136);
}

.createNewUserButtonForLogIn:hover {
  cursor: pointer;
}

.logInForm {
  font-size: 30px;
}

.logInUserFormField {
  top: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.logInPasswordFormField {
  display: flex;
  justify-content: center;
  align-items: center;
  top: 20px;
}

.logInUserFormField > input {
  border-radius: 50px;
  height: 50px;
  font-size: 20px;
  border-radius: 20px;
  left: 3px;
  border-color: rgb(15, 10, 117);
}

.logInPasswordFormField > input {
  border-radius: 50px;
  height: 50px;
  font-size: 20px;
  border-radius: 20px;
  left: 3px;
  border-color: rgb(15, 10, 117);
}

@media (max-width: 750px) {
  .logInHeader {
    font-size: x-large;
    margin-left: auto;
    margin-right: auto;
  }
}

@media (max-width: 500px) {
  .logInHeader {
    font-size: medium;
    margin-left: auto;
    margin-right: auto;
  }
  .logInForm {
    font-size: 20px;
  }
  .logInUserFormField > input {
    font-size: 15px;
  }
  .logInPasswordFormField > input {
    font-size: 15px;
  }
}
</style>

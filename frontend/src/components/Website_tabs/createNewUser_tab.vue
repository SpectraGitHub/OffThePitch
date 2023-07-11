<!--Component which display the form where the user can create a new user-->
<template>
  <Top_Bar />
  <form>
    <div class="createNewUserComp">
      <div class="createNewUserTextComp">
        <h2 class="newUserHeader">
          Create new user <img class="img" @click="onClick()" :src="image" />
        </h2>
      </div>
      <div class="newUserForm">
        <div class="newUserEmailFormField">
          <label>E-mail</label>
          <input type="text" id="email" name="email" v-model="email" />
        </div>
        <div class="newUserUsernameFormField">
          <label>Username </label>
          <input type="text" id="username" name="username" v-model="username" />
        </div>
        <div class="newUserPasswordFormField">
          <label>Password </label>
          <input
            type="password"
            id="password"
            name="password"
            v-model="password"
          />
        </div>
        <div class="newUserRepeatedPasswordFormField">
          <label>Repeat password </label>
          <input
            type="password"
            id="repeatedpassword"
            name="repeatedpassword"
            v-model="repeatedpassword"
          />
        </div>
        <button
          type="button"
          class="createNewUserButton"
          v-on:click="getInfo()"
        >
          Create new user
        </button>
      </div>
    </div>
  </form>
</template>

<script>
import image from "/icons/favicon.ico";
import Top_Bar from "../Header/Header.vue";
import axios from "axios";

export default {
  name: "CreateNew_Screen",
  data: function () {
    return {
      image: image,
      //form struct
      form: {
        email: "",
        username: "",
        password: "",
        repeatedpassword: "",
      },
    };
  },
  methods: {
    //relocate function to home screen
    onClick() {
      location.href = "/";
    },
    //Function to get the create new user details
    async getInfo() {
      const registerdata = {
        USERNAME: this.username,
        EMAIL: this.email,
        PASSWORD: this.password,
      };
      const logindata = {
        USERNAME: this.username,
        PASSWORD: this.password,
      };

      const registerUserJSON = JSON.stringify(registerdata);
      const loginUserJSON = JSON.stringify(logindata);

      //Checks if the details are accepted
      if (this.acceptedDetails()) {
        let token = "";
        //Post request to backend with a new user, when the user is created, the user will automatically be logged in
        var resp = await axios
          .post(this.backendUrl + "/otp/newuser", registerUserJSON)
          .then(
            async function (response) {
              token = await axios
                .post(this.backendUrl + "/otp/login", loginUserJSON)
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
            alert(error.response.data);
            return error.response;
          });

        if (resp.status < 300) {
          // sets token, current user and state
          sessionStorage.setItem("token", token.data);
          sessionStorage.setItem("currentUser", this.username.toLowerCase());
          sessionStorage.setItem("currentState", true);
          this.onClick();
        }
      } else {
        return;
      }
    },
    //Function to check the details
    acceptedDetails() {
      let smallChar = false;
      let upperChar = false;
      let num = false;
      //Looping throught the password
      for (let i = 0; i < this.password.length; i++) {
        //Checking for small character
        if (
          96 < this.password.charCodeAt(i) &&
          this.password.charCodeAt(i) < 123
        ) {
          smallChar = true;
        }
        //Checking for upper character
        if (
          64 < this.password.charCodeAt(i) &&
          this.password.charCodeAt(i) < 91
        ) {
          upperChar = true;
        }
        //Checking for number
        if (
          47 < this.password.charCodeAt(i) &&
          this.password.charCodeAt(i) < 58
        ) {
          num = true;
        }
      }
      //Output message to the user
      if (!smallChar || !upperChar || !num) {
        alert(
          "Password needs an uppercase character a lowercase character and a number"
        );
        return false;
      }
      //Checks if the email is on the correct format
      if (!this.checkEmail()) {
        alert("Please enter a real email.");
        return false;
      }
      //Checks if the username and password are entered
      if (!this.username && !this.password) {
        alert("Please enter a username and a password");
        return false;
      }

      //Checks the username length
      if (this.username.length < 5) {
        alert("Username needs to be at least 5 characters");
        return false;
      }

      //Checks the password length
      if (this.password.length < 7) {
        alert("Password needs to be at least 7 characters");
        return false;
      }
      //Checks if the repeated password match
      if (this.password != this.repeatedpassword) {
        alert("The passwords do not match");
        return false;
      }

      return true;
    },

    //Fucntion to check the email
    checkEmail() {
      //Checks for empty email
      if (!this.email) {
        return false;
      }
      //Email length longer than 6
      if (this.email.length < 6) {
        return false;
      }
      //Checks if it includes a @ and a dot after @
      if (!this.email.includes("@")) {
        return false;
      } else {
        let index = this.email.indexOf("@");
        if (!this.email.substring(index).includes(".")) {
          return false;
        }
      }
      return true;
    },
  },
  components: { Top_Bar },
};
</script>

<style scoped>
.createNewUserComp {
  top: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

.newUserHeader {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 40px;
}

.img {
  left: 10px;
  height: 50px;
  border-radius: 10px;
}

.img:hover {
  cursor: pointer;
}

.createNewUserButton {
  top: 60px;
  height: 40px;
  width: 100%;
  border-radius: 20px;
  border-color: rgb(100, 18, 200);
  background-color: rgb(136, 136, 136);
}

.createNewUserButton:hover {
  cursor: pointer;
}

.newUserForm {
  font-size: 30px;
}

.newUserEmailFormField {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  top: 20px;
}

.newUserUsernameFormField {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  top: 30px;
}

.newUserPasswordFormField {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  top: 40px;
}

.newUserRepeatedPasswordFormField {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  top: 50px;
}

.newUserEmailFormField > input {
  border-radius: 50px;
  height: 50px;
  font-size: 20px;
  border-radius: 20px;
  left: 3px;
  border-color: rgb(15, 10, 117);
  float: right;
}

.newUserUsernameFormField > input {
  border-radius: 50px;
  height: 50px;
  font-size: 20px;
  border-radius: 20px;
  left: 3px;
  border-color: rgb(15, 10, 117);
}

.newUserPasswordFormField > input {
  border-radius: 50px;
  height: 50px;
  font-size: 20px;
  border-radius: 20px;
  left: 3px;
  border-color: rgb(15, 10, 117);
}

.newUserRepeatedPasswordFormField > input {
  border-radius: 50px;
  height: 50px;
  font-size: 20px;
  border-radius: 20px;
  left: 3px;
  border-color: rgb(15, 10, 117);
}

@media (max-width: 500px) {
  .newUserForm {
    font-size: large;
  }
  .newUserEmailFormField > input {
    font-size: medium;
  }
  .newUserUsernameFormField > input {
    font-size: medium;
  }

  .newUserPasswordFormField > input {
    font-size: medium;
  }
  .newUserRepeatedPasswordFormField > input {
    font-size: medium;
  }
}
</style>

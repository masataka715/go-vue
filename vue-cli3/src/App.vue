<template>
  <div id="app" class="container">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <nav>
      <div class="nav nav-tabs" id="nav-tab" role="tablist">
        <router-link to="/" class="nav-item nav-link font-weight-bold">ホーム</router-link>
        <router-link to="/game" class="nav-item nav-link font-weight-bold">ゲーム</router-link>
        <router-link to="/tool" class="nav-item nav-link font-weight-bold">ツール</router-link>
      </div>

      <div v-if="is_auth" class="text-right">
        <!-- <span class=" font-weight-bold bg-light">ログイン中</span> -->
        <span class="badge badge-secondary mr-3">ログイン中</span>
        <button @click="logout()" class="btn btn-primary">ログアウト</button>
      </div>
      <div v-if="!is_auth" class="text-right">
        <small class="mr-2">ログインしていません</small>
        <button @click="login()" class="btn btn-primary">ログイン</button>
        <button @click="testLogin()" class="btn btn-danger" type="button">テストでログイン</button>
      </div>
    </nav>
    <div v-if="show_login_form">
      <LoginForm />
    </div>
    <!-- ここにパスと一致したコンポーネントが埋め込まれる -->
    <router-view />
  </div>
</template>

<script>
import LoginForm from "./components/LoginForm.vue";
// import axios from "axios";
export default {
  name: "app",
  components: {
    LoginForm
  },
  data() {
    return {
      //   show_login_form: false
    };
  },
  computed: {
    is_auth: {
      get: function() {
        return this.$store.getters.checkAuth;
      }
    },
    show_login_form: {
      get: function() {
        return this.$store.state.show_login_form;
      }
    }
  },
  methods: {
    testLogin() {
      window.sessionStorage.setItem(["user_id"], [9185491]);
      this.$store.commit("login");
    },
    login() {
      this.$store.commit("showLoginForm");
    },
    logout() {
      window.sessionStorage.clear();
      this.$store.commit("logout");
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
}
</style>

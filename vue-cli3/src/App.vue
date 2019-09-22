<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <div class="wrap-video position-relative">
      <h1 class="title position-absolute">動画共有サイト</h1>
      <video autoplay loop muted src="./assets/vr.mp4" width="850" height="500">
      </video>
    </div>
    <!-- ナビリンク -->
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
        <button @click="login()" class="btn btn-sm btn-primary">ログイン</button>
        <button @click="testLogin()" class="btn btn-sm btn-danger" type="button">テストでログイン</button>
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
export default {
  name: "app",
  components: {
    LoginForm
  },
  data() {
    return {};
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
}

.wrap-video video{
  width: 100%;
  height:100%;
}

.wrap-video .title{
  top: 50%;
  left: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  text-align:center;
  font-weight:bold;
  color:white;
}
</style>

<template>
  <div id="app" class="container">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <nav>
      <div class="nav nav-tabs" id="nav-tab" role="tablist">
        <router-link to="/" class="nav-item nav-link font-weight-bold">ホーム</router-link>
        <router-link to="/game" class="nav-item nav-link font-weight-bold">ゲーム</router-link>
        <router-link to="/tool" class="nav-item nav-link font-weight-bold">ツール</router-link>
      </div>

      <div v-if="is_auth" class="text-right font-weight-bold">
        <small class="mr-2">ログイン中</small>
        <button @click="logout()" class="btn btn-primary">ログアウト</button>
      </div>
      <div v-if="!is_auth" class="text-right">
        <small class="mr-2">ログインしていません</small>
        <button class="btn btn-primary">ログイン</button>
        <button @click="testLogin()" class="btn btn-danger" type="button">テストでログイン</button>
      </div>
    </nav>
    <!-- ここにパスと一致したコンポーネントが埋め込まれる -->
    <router-view />
  </div>
</template>

<script>
export default {
  name: "app",
  components: {},
  data() {
    return {
    };
  },
  computed: {
    is_auth: {
      get: function() {
        return this.$store.getters.checkAuth;
      }
    }
  },
  methods: {
    testLogin() {
      window.sessionStorage.setItem(["user_id"], [9185491]);
      this.$store.commit("login");
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
  color: #2c3e50;
  margin-top: 60px;
}
</style>

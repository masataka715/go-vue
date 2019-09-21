<template>
  <div id="app" class="container">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <nav>
      <div class="nav nav-tabs" id="nav-tab" role="tablist">
        <router-link to="/" class="nav-item nav-link font-weight-bold">ホーム</router-link>
        <router-link to="/game" class="nav-item nav-link font-weight-bold">ゲーム</router-link>
        <router-link to="/tool" class="nav-item nav-link font-weight-bold">ツール</router-link>
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
      goDomain: "http://localhost:5000"
    };
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

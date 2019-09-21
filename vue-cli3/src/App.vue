<template>
  <div id="app" class="container">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <Battle :goDomain="goDomain"/>
    <Auth :goDomain="goDomain"/>
    <Twitter :goDomain="goDomain"/>
    <Todo :goDomain="goDomain"/>
    <Shiritori :goDomain="goDomain"/>
  </div>
</template>

<script>
import Battle from './components/Battle.vue'
import Auth from './components/Auth.vue'
import Twitter from './components/Twitter.vue'
import Todo from './components/Todo.vue'
import Shiritori from './components/Shiritori.vue'

export default {
  name: 'app',
  components: {
    Battle,
    Auth,
    Twitter,
    Todo,
    Shiritori,
  },
  data() {
      return {
          goDomain: "http://localhost:5000"
      }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

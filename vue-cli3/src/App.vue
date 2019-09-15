<template>
  <div id="app" class="container">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <Todo/>
    <Shiritori/>
  </div>
</template>

<script>
import Todo from './components/Todo.vue'
import Shiritori from './components/Shiritori.vue'

export default {
  name: 'app',
  components: {
    Todo,
    Shiritori,
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

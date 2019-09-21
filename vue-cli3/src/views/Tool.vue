<template>
  <div>
    <h1 class="my-4">ツール</h1>
    <div class="text-center">
      <Todo :goDomain="goDomain" />
    </div>
  </div>
</template>

<script>
import Todo from "./../components/Todo.vue";
export default {
  components: {
    Todo
  },
  data() {
    return {
      goDomain: "http://localhost:5000"
    };
  }
};
</script>
<template>
  <div>
    <h1 class="my-4">ツール</h1>
    <div v-if="is_auth">
      <Todo />
    </div>
    <div v-if="!is_auth">
      <p>ログインが必要です</p>
      <button @click="testLogin()" class="btn btn-danger" type="button">テストでログイン</button>
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
    return {};
  },
  mounted: function() {},
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
    }
  }
};
</script>
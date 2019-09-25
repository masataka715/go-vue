<template>
  <div>
    <h1 class="mt-0 mb-2">ノート</h1>
    <div v-if="is_auth" class="row">
      <div class="col-3 p-0 bg-dark" style="height:700px">
        <NoteLeft />
      </div>
      <div class="col-8 p-0 ml-3">
        <router-view />
      </div>
    </div>
    <div v-if="!is_auth">
      <p>ログインが必要です</p>
      <button @click="testLogin()" class="btn btn-danger" type="button">テストでログイン</button>
    </div>
  </div>
</template>

<script>
import NoteLeft from "./../../components/NoteLeft.vue";
export default {
  components: {
    NoteLeft
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
      window.sessionStorage.setItem(["user_id"], [this.$store.state.test_auth_id]);
      this.$store.commit("login");
    }
  }
};
</script>
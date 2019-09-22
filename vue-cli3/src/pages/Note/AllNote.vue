<template>
  <div>
    <h2 class="mt-0 mb-2">ノート</h2>
    <div v-if="is_auth" class="row">
      <div class="col-3 p-0 bg-dark" style="height:700px">
        <NoteLeft />
      </div>
      <div class="col-9">
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
      window.sessionStorage.setItem(["user_id"], [9185491]);
      this.$store.commit("login");
    }
  }
};
</script>
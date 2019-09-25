<template>
  <div>
    <div class="row">
      <h4 class="col-5 offset-3">新規ノート</h4>
      <div class="col-2 offset-1">
        <router-link to="/note" @click.native="newStore()" class="badge badge-pill badge-light">保存</router-link>
      </div>
    </div>
    <div>
      <div>
        <input v-model="note.title" class="col-10 m-0 h4" type="text" placeholder="ノートのタイトル" />
      </div>
      <textarea v-model="note.content" class="col-10" rows="25" placeholder="ここから入力"></textarea>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      go_url: this.$store.state.go_domain + "/note/new",
      note: {
        id: 0,
        auth_id: 0,
        title: "",
        content: ""
      },
      message: ""
    };
  },
  computed: {},
  mounted: function() {
    const session_id = window.sessionStorage.getItem(["user_id"]);
    this.note.auth_id = Number(session_id);
  },
  methods: {
    newStore() {
      const json = JSON.stringify(this.note);
      axios
        .post(this.go_url, json)
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
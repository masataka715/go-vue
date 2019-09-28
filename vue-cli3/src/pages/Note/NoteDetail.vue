<template>
  <div>
    <div class="col p-0">
      <div class="row">
        <span class="col-5 offset-3 text-primary mr-2 mb-0 p-0">{{ message }}</span>
        <button
          @click="store()"
          class="col-2 mb-1 badge badge-pill badge-light"
          style="outline: none;"
        >保存</button>
      </div>
      <div class="text-left">
        <input v-model="note.title" class="col-11 m-0 h4" type="text" placeholder="ノートのタイトル" />
      </div>
      <textarea v-model="note.content" class="col-11 m-0 d-block" rows="27" placeholder="ここから入力"></textarea>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      go_url: this.$store.state.go_domain + "/note/details/",
      note: {},
      message: ""
    };
  },
  watch: {
    // ルートの変更の検知
    $route(to) {
      // go_url1と2は分けないと不都合
      const go_url1 = this.go_url + to.params.id;
      axios.get(go_url1).then(res => {
        this.note = res.data;
      });
      this.message = "";
    }
  },
  mounted: function() {
    const go_url2 = this.go_url + this.$route.params.id;
    axios
      .get(go_url2)
      .then(res => {
        this.note = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    store() {
      this.note.id = Number(this.$route.params.id);
      const json = JSON.stringify(this.note);
      axios.post(this.go_url, json).catch(err => {
        console.log(err);
      });
      location.reload();
    }
  }
};
</script>
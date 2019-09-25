<template>
  <div>
    <div class="col p-0">
      <div class="row">
        <span class="col-5 offset-3 text-primary mr-2 mb-0 p-0">{{ message }}</span>
        <button
          @click="noteDel()"
          class="col-2 mb-1 badge badge-pill badge-light"
          style="outline: none;"
        >削除</button>
      </div>
      <div class="text-left">
        <h4 class="col-11 m-0">{{ note.title }}</h4>
      </div>
      <p class="col-11 m-0 d-block">{{ note.content}}</p>
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
      // go_urlは分けないと不都合
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
    noteDel() {
      console.log("工事中")
      
    }
  }
};
</script>
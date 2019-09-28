<template>
  <div>
    <div class="col p-0">
      <div class="row">
        <span class="col-7 text-primary mr-2 mb-0 p-0">{{ message }}</span>
        <button
          @click="garbageDel()"
          class="col-4 mb-1 badge badge-pill badge-light"
          style="outline: none;"
        >完全に削除</button>
      </div>
      <div class="text-left">
        <h4 class="col-11 m-0">{{ garbage.title }}</h4>
      </div>
      <p class="col-11 m-0 d-block">{{ garbage.content}}</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      go_url_details: this.$store.state.go_domain + "/note/garbage_details/",
      go_url_del: this.$store.state.go_domain + "/note/garbage_del/",
      garbage: {},
      message: ""
    };
  },
  watch: {
    // ルートの変更の検知
    $route(to) {
      // go_url1と2は分けないと不都合
      const garbage_details_url = this.go_url_details + to.params.id;
      axios.get(garbage_details_url).then(res => {
        this.garbage = res.data;
      });
      this.message = "";
    }
  },
  mounted: function() {
    const garbage_details_url2 = this.go_url_details + this.$route.params.id;
    axios
      .get(garbage_details_url2)
      .then(res => {
        this.garbage = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    garbageDel() {
      const garbage_del_url = this.go_url_del + this.$route.params.id;
      console.log(garbage_del_url)
      axios.get(garbage_del_url).catch(err => {
        console.log(err);
      });
      location.reload();
    }
  }
};
</script>
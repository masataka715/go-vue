<template>
  <div>
    <div class="col p-0">
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
  }
};
</script>
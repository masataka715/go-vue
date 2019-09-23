<template>
  <div>
    <p>{{ note.title }}</p>
    <p>{{ note.content }}</p>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      goUrl: this.$store.state.go_domain + "/note/details/",
      note: {}
    };
  },
  watch: {
    // ルートの変更の検知
    $route(to) {
      // goUrlは分けないと不都合
      const goUrl1 = this.goUrl + to.params.id;
      axios.get(goUrl1).then(res => {
        this.note = res.data;
      });
    }
  },
  mounted: function() {
    const goUrl2 = this.goUrl + this.$route.params.id;
    axios
      .get(goUrl2)
      .then(res => {
        this.note = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    newStore() {}
  }
};
</script>
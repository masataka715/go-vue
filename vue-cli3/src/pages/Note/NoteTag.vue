<template>
  <div class="text-left">
    <p>新規タグの作成</p>
    <div>
      <input v-model="tag.tag_name" class="mx-2" type="text" placeholder="タグ名" />
      <button @click="tagCreate()" class="badge badge-pill badge-light">完了</button>
    </div>
    <hr />
    <p>タグ一覧</p>
    <div v-for="t in tags" :key="t.id" class="list-group">
      <a href="#" class="list-group-item list-group-item-action">{{ t.tag_name }}</a>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      go_url_all: this.$store.state.go_domain + "/note/tag_all/",
      go_url_create: this.$store.state.go_domain + "/note/tag_create",
      tag: {
        auth_id: Number(window.sessionStorage.getItem(["user_id"])),
        tag_name: ""
      },
      tags: []
    };
  },
  watch: {},
  mounted: function() {
    const url_all = this.go_url_all + this.tag.auth_id;
    axios
      .get(url_all)
      .then(res => {
        this.tags = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    tagCreate() {
      const json = JSON.stringify(this.tag);
      axios
        .post(this.go_url_create, json)
        .then(res => {
          this.tags = res.data;
        })
        .catch(err => {
          console.log(err);
        });
      //   location.reload();
    }
  }
};
</script>
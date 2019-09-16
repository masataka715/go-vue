<template>
  <div class="py-4" style="background-color:lavenderblush">
    <h2>ツイッター機能</h2>

    <div class="input-group mb-3 w-75 mx-auto">
      <input
        v-model="twitter.content"
        type="text"
        class="form-control"
        placeholder="つぶやいてみよう"
        aria-describedby="button-addon2"
      />
      <div class="input-group-append">
        <button @click="tweet()" class="btn btn-danger" type="button" id="button-addon2">ツイート</button>
      </div>
    </div>
    <!-- ツイート一覧 -->
    <div v-for="t in twitters" :key="t.ID" class="text-left w-75 mx-auto bg-white rounded">
      <span class="small py-0 text-nowrap font-weight-bold">ゲストユーザー</span>
      <!-- <span>{{ t.created_at}}</span> -->
      <button @click="del(t.id)" type="button" class="close" aria-label="Close">
        <span aria-hidden="true">&times;</span>
      </button>
      <p>{{ t.content }}</p>
      <div class="text-center">
        <button @click="good()" class="badge badge-secondary">いいね</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      twitter: {
        id: 0,
        content: ""
      },
      twitters: {},
      goUrl: "http://localhost:5000/twitter"
    };
  },

  mounted: function() {
    axios
      .get(this.goUrl)
      .then(res => {
        this.twitters = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },

  methods: {
    tweet() {
      const json = JSON.stringify(this.twitter);
      axios
        .post(this.goUrl, json)
        .then(res => {
          this.twitters = res.data;
          this.twitter.content = "";
        })
        .catch(err => {
          console.log(err);
        });
    },
    del(id) {
      axios
        .post(this.goUrl + "/delete/" + id)
        .then(res => {
          this.twitters = res.data;
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
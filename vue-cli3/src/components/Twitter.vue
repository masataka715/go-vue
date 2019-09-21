<template>
  <div class="py-4" style="background-color:lavenderblush">
    <h2>ツイッター機能</h2>

    <div class="input-group mb-3 w-75 mx-auto">
      <textarea
        v-model="twitter.content"
        class="form-control"
        placeholder="つぶやいてみよう"
        aria-describedby="button-addon2"
        autocomplete="on"
      ></textarea>
      <div class="input-group-append">
        <button @click="tweet()" class="btn btn-danger" type="button" id="button-addon2">ツイート</button>
      </div>
    </div>
    <!-- ツイート一覧 -->
    <div v-for="t in twitters" :key="t.ID" class="text-left w-75 mx-auto bg-white rounded border">
      <span class="small py-0 text-nowrap font-weight-bold">ゲストユーザー</span>
      <!-- <span>{{ t.created_at}}</span> -->
      <button @click="del(t.id)" type="button" class="close" aria-label="Close">
        <span aria-hidden="true">&times;</span>
      </button>
      <p>{{ t.content }}</p>
      <div class="text-center">
        <button @click="nice(t.id)" class="badge badge-secondary">
          いいね
          <span>{{ t.nice_count }}</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  props: ["goDomain"],
  data() {
    return {
      twitter: {
        id: 0,
        content: "",
        nice_count: 0
      },
      twitters: {},
      goUrl: this.goDomain + "/twitter"
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
    },
    nice(id) {
      axios
        .post(this.goUrl + "/nice/" + id)
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
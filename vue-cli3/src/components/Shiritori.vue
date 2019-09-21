<template>
  <div class="my-4 py-4 text-center" style="background-color: mintcream">
    <h2>しりとり</h2>
    <p>最初の文字：しりとりの「り」からです</p>
    <p>平仮名で入力して下さい</p>
    <p v-if="shiritori.lastLetter">次は「{{ shiritori.lastLetter }}」</p>
    <p>
      <input type="text" v-model="shiritori.word" />
      <button @click="submit()">送信</button>
    </p>
    <p v-if="shiritori.message" class="text-danger">{{ shiritori.message }}</p>
  </div>
</template>

<script>
import axios from "axios";
export default {
  props: ["goDomain"],
  data() {
    return {
      shiritori: {
        word: "り",
        lastLetter: "",
        message: null
      },
      goUrl: this.goDomain + "/shiri"
    };
  },
  methods: {
    submit() {
      const json = JSON.stringify(this.shiritori);
      axios
        .post(this.goUrl, json)
        .then(res => {
          this.shiritori = res.data;
          this.shiritori.word = "";
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
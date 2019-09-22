<template>
  <div class="py-4 my-4" style="background-color:lavenderblush">
    <div class="w-75 mx-auto">
      <h2>戦闘シーン</h2>
      <div class="bg-dark">
        <div>
          <table class="table table-dark my-0 w-25 border border-white">
            <tr>{{ member.name }}</tr>
            <tr>HP：{{ member.hp }}</tr>
            <tr>MP：{{ member.mp }}</tr>
          </table>
        </div>
        <div class="p-4 m-4 text-white">敵</div>
        <!-- コマンド -->
        <div class="border border-white row">
          <div class="col-4">
            <table class="table table-dark my-0 border border-white">
              <tr>
                <button @click="battle()" class="btn btn-dark btn-block">たたかう</button>
              </tr>
              <tr>
                <button @click="spell()" class="btn btn-dark btn-block">じゅもん</button>
              </tr>
              <tr>
                <button @click="escape()" class="btn btn-dark btn-block">にげる</button>
              </tr>
            </table>
          </div>
          <!-- バトルメッセージ -->
          <div class="col-8 mt-2 text-left text-white">
            <transition name="bounce">
              <p v-if="show">{{ message }}</p>
            </transition>
            <p v-if="!show">{{ message }}</p>
          </div>
          <!-- <p v-if="show">{{ message }}</p> -->
          <!-- <button @click="show = !show">Toggle show</button> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      member: {
        name: "勇者",
        hp: 100,
        mp: 50
      },
      enemy: "敵",
      message: "敵があらわれた！",
      goUrl: this.$store.state.go_domain + "/battle",
      show: false
    };
  },
  methods: {
    battle() {
      this.show = !this.show;
      // this.sleep(2000);
      this.message = this.member.name + "は、敵をこうげきした";
      const json = JSON.stringify(this.member);
      axios
        .post(this.goUrl, json)
        .then(res => {
          this.member = res.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    sleep(wait_msec) {
      var start_msec = new Date();
      // 指定ミリ秒間だけループさせる（CPUは常にビジー状態）
      while (new Date() - start_msec < wait_msec);
    }
  }
};
</script>

<style>
.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>
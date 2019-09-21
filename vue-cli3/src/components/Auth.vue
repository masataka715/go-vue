<template>
  <div class="my-4 py-4 text-center" style="background-color: mintcream">
    <h2>認証機能</h2>
    <div>
      <div class="form-group row">
        <label for="user_name" class="col-2 offset-1 col-form-label">ユーザー名</label>
        <div class="col-6">
          <input
            v-model="register.user_name"
            type="text"
            class="form-control"
            id="user_name"
            placeholder="ユーザー名"
          />
        </div>
      </div>
      <div class="form-group row">
        <label for="password" class="col-2 offset-1 col-form-label">パスワード</label>
        <div class="col-6">
          <input
            v-model="register.password"
            type="password"
            class="form-control"
            id="password"
            placeholder="パスワード"
          />
        </div>
      </div>
        <button @click="newRegister()">新規登録</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  props: ["goDomain"],
  data() {
    return {
      goUrl: this.goDomain + "/register",
      register: {
        user_name: "",
        password: ""
      }
    };
  },
  methods: {
    newRegister() {
      const json = JSON.stringify(this.register);
      axios
        .post(this.goUrl, json)
        .then(res => {
          this.register = res.data;
          console.log(this.register)
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
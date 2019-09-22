<template>
  <div class="my-4 py-4 text-center" style="background-color: mintcream">
    <h2>認証機能</h2>
    <div v-if="!is_register + !is_auth == 2">
      <div class="form-group row">
        <label for="email_address" class="col-3 offset-1 col-form-label">メールアドレス</label>
        <div class="col-6">
          <input
            v-model="register.email_address"
            type="email"
            class="form-control"
            id="email_address"
            placeholder="メールアドレス"
          />
        </div>
      </div>
      <div class="form-group row">
        <label for="password" class="col-3 offset-1 col-form-label">パスワード</label>
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
    <!-- メッセージ -->
    <div v-if="is_register" class="my-3">新規登録しました</div>
    <div v-if="is_auth" class="my-3">ログイン中</div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      goUrl: this.$store.state.go_domain + "/register",
      register: {
        id: 0,
        email_address: "",
        password: ""
      },
      is_register: false
    };
  },
  computed: {
    is_auth: {
      get: function() {
        return this.$store.getters.checkAuth;
      }
    }
  },
  methods: {
    newRegister() {
      const json = JSON.stringify(this.register);
      axios
        .post(this.goUrl, json)
        .then(res => {
          this.is_register = true;
          this.register = res.data;
          window.sessionStorage.setItem(["user_id"], [this.register.id]);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
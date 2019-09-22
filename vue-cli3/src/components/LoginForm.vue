<template>
    <div
      class="mt-2 mb-4 pb-4 text-center"
      style="background-color: aliceblue"
    >
      <div class="text-right">
        <button @click="close()">✖︎</button>
      </div>
      <h2>ログイン画面</h2>
      <div class="form-group row">
        <label for="email_address" class="col-3 offset-1 col-form-label">メールアドレス</label>
        <div class="col-6">
          <input
            v-model="auth.email_address"
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
            v-model="auth.password"
            type="password"
            class="form-control"
            id="password"
            placeholder="パスワード"
          />
        </div>
      </div>
      <button @click="login()">ログイン</button>
      <div class="mt-4">{{ message }}</div>
    </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      goUrl: this.$store.state.go_domain + "/login",
      auth: {
        email_address: "",
        password: ""
      },
      message: ""
    };
  },
  computed: {
    show_login_form: {
      get: function() {
        return this.$store.state.show_login_form;
      }
    }
  },
  methods: {
    login() {
      const json = JSON.stringify(this.auth);
      axios
        .post(this.goUrl, json)
        .then(res => {
          this.auth = res.data;
          if (this.auth.id == 0) {
            this.message = "ログインできませんでした";
          } else {
            window.sessionStorage.setItem(["user_id"], [this.auth.id]);
            this.$store.commit("login");
            this.$store.commit("hideLoginForm");
          }
        })
        .catch(err => {
          console.log(err);
        });
    },
    close() {
      this.$store.commit("hideLoginForm");
    }
  }
};
</script>
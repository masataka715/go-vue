<template>
  <div>
    <div class="col p-0">
      <div class="row">
        <span class="col-5 offset-3 text-primary mr-2 mb-0 p-0">{{ store_message }}</span>
        <button
          @click="store()"
          class="col-2 mb-1 badge badge-pill badge-light"
          style="outline: none;"
        >保存</button>
      </div>
      <div class="text-left">
        <input v-model="note.title" class="col-11 m-0 h4" type="text" placeholder="ノートのタイトル" />
      </div>
      <textarea v-model="note.content" class="col-11 m-0 d-block" rows="20" placeholder="ここから入力"></textarea>
    </div>
    <div class="text-left mt-1">
      <button
        type="button"
        class="badge badge-pill badge-light"
        data-toggle="modal"
        data-target="#addTag"
      >タグを追加</button>
    </div>
    <!-- Modal -->
    <div
      class="modal fade"
      id="addTag"
      tabindex="-1"
      role="dialog"
      aria-labelledby="ModalCenterTitle"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="ModalLongTitle">タグの追加</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <div v-for="t in tags" :key="t.id" class="form-check text-left">
              <input class="form-check-input" type="checkbox" value :id="'tagCheck' + t.id" />
              <label class="form-check-label" :for="'tagCheck' + t.id">{{ t.tag_name }}</label>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-dismiss="modal">キャンセル</button>
            <button type="button" class="btn btn-primary">追加</button>
          </div>
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
      go_url: this.$store.state.go_domain + "/note/details/",
      tag_all_url: this.$store.state.go_domain + "/note/tag_all/",
      auth_id: Number(window.sessionStorage.getItem(["user_id"])),
      note: {},
      store_message: "",
      tags: []
    };
  },
  watch: {
    // ルートの変更の検知
    $route(to) {
      // go_url1と2は分けないと不都合
      const go_url1 = this.go_url + to.params.id;
      axios.get(go_url1).then(res => {
        this.note = res.data;
      });
      this.store_message = "";
    }
  },
  mounted: function() {
    const go_url2 = this.go_url + this.$route.params.id;
    axios
      .get(go_url2)
      .then(res => {
        this.note = res.data;
      })
      .catch(err => {
        console.log(err);
      });
    // タグ一覧取得
    const tag_all_url = this.tag_all_url + this.auth_id;
    axios
      .get(tag_all_url)
      .then(res => {
        this.tags = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    store() {
      this.note.id = Number(this.$route.params.id);
      const json = JSON.stringify(this.note);
      axios.post(this.go_url, json).catch(err => {
        console.log(err);
      });
      location.reload();
    }
  }
};
</script>
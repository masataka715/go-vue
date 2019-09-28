<template>
  <div class="row">
    <div class="col p-0" style="height:700px; overflow: scroll;">
      <div v-for="n in notes" :key="n.id" class="bg-light border" style="height:80px;">
        <router-link
          :to="{ name: 'garbage_details', params: {id: n.id}}"
          class="d-block text-decoration-none text-dark text-left px-3"
        >
          <div class="float-right">
            <button @click="noteRestore(n.id)" class="badge badge-pill badge-light mr-1">復元</button>
            <button @click="garbageDel(n.id)" class="badge badge-pill badge-light">削除</button>
          </div>
          <p class="py-1 m-0 font-weight-bold">{{ n.title }}</p>
          <p class="pb-1 m-0" style="height: 50px; overflow: hidden;">{{ n.content }}</p>
        </router-link>
      </div>
    </div>
    <div class="col p-0">
      <router-view />
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      go_url: this.$store.state.go_domain + "/note/garbage_all/",
      go_url_restore: this.$store.state.go_domain + "/note/garbage_restore/",
      go_url_del: this.$store.state.go_domain + "/note/garbage_del/",
      session_id: Number(window.sessionStorage.getItem(["user_id"])),
      notes: []
    };
  },
  mounted: function() {
    // const session_id = window.sessionStorage.getItem(["user_id"]);
    this.go_url += this.session_id;

    axios
      .get(this.go_url)
      .then(res => {
        this.notes = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    noteRestore(garbage_id) {
        const garbage_restore =
        this.go_url_restore + garbage_id + "/" + this.session_id;
      axios
        .get(garbage_restore)
        .then(res => {
          this.notes = res.data;
        })
        .catch(err => {
          console.log(err);
        });
        if (this.$route.fullPath != "/note/garbage") {
            this.$router.push("/note/garbage");
        }
      
    },
    garbageDel(garbage_id) {
      const garbage_del_url =
        this.go_url_del + garbage_id + "/" + this.session_id;
      axios
        .get(garbage_del_url)
        .then(res => {
          this.notes = res.data;
        })
        .catch(err => {
          console.log(err);
        });
      this.$router.push("/note/garbage");
    }
  }
};
</script>
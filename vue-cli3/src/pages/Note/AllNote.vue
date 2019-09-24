<template>
  <div class="row">
    <div class="col p-0" style="height:700px; overflow: scroll;">
      <div v-for="n in notes" :key="n.id" class="bg-light border" style="height:80px;">
        <router-link
          :to="{ name: 'note_details', params: {id: n.id}}"
          class="d-block text-decoration-none text-dark text-left px-3"
        >
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
      go_url: this.$store.state.go_domain + "/note/all/",
      notes: []
    };
  },
  mounted: function() {
    const session_id = window.sessionStorage.getItem(["user_id"]);
    this.go_url += Number(session_id);

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
  }
};
</script>
<template>
  <div>
    <h1 class="my-4">ホーム</h1>
    <div class="text-center">
      <Twitter :goDomain="goDomain" />
      <Auth :goDomain="goDomain" />
    </div>
  </div>
</template>

<script>
import Twitter from "./../components/Twitter.vue";
import Auth from "./../components/Auth.vue";
export default {
  components: {
    Twitter,
    Auth
  },
  data() {
    return {
      goDomain: "http://localhost:5000"
    };
  }
};
</script>
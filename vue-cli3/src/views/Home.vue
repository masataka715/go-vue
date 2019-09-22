<template>
  <div>
    <h1 class="my-4">ホーム</h1>
    <div class="text-center">
      <Twitter/>
      <Auth/>
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
    };
  }
};
</script>
<template>
  <div>
    <h1 class="mt-0 mb-2">ゲーム</h1>
    <div class="text-center">
      <Battle/>
      <Shiritori/>
    </div>
  </div>
</template>

<script>
import Battle from "./../components/Battle.vue";
import Shiritori from "./../components/Shiritori.vue";
export default {
  components: {
    Battle,
    Shiritori
  },
  data() {
    return {
    };
  }
};
</script>
<template>
  <div class="nav flex-column nav-pills">
    <router-link to="/note/new" exact class="nav-link text-white note-left-link" data-toggle="pill">＋新規ノート</router-link>
    <router-link to="/note/all" exact class="nav-link active text-white" data-toggle="pill">すべてのノート</router-link>
    <router-link to="/note/book" exact class="nav-link text-white" data-toggle="pill">ノートブック</router-link>
    <router-link to="/note/tag" class="nav-link text-white" data-toggle="pill">タグ</router-link>
    <router-link to="/note/garbage" class="nav-link text-white" data-toggle="pill">ゴミ箱</router-link>
  </div>
</template>

<script>
export default {
    mounted: function() {
        if (this.$route.name != "note_all") {
            this.$router.push("/note/all");
        }
    }
}
</script>
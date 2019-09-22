<template>
  <div class="nav flex-column nav-pills">
    <router-link to="/note/new" class="nav-link text-white" data-toggle="pill">＋新規ノート</router-link>
    <router-link to="/note/all" class="nav-link active text-white" data-toggle="pill">すべてのノート</router-link>
    <router-link to="/note/book" class="nav-link text-white" data-toggle="pill">ノートブック</router-link>
    <router-link to="/note/tag" class="nav-link text-white" data-toggle="pill">タグ</router-link>
    <router-link to="/note/garbage" class="nav-link text-white" data-toggle="pill">ゴミ箱</router-link>
  </div>
</template>
<template>
  <div class="my-4 py-4 text-center" style="background-color:aliceblue">
    <h2>ToDo機能</h2>
    <table class="table table-sm w-75 my-4 mx-auto bg-white">
      <thead>
        <tr>
          <th scope="col-2">内容</th>
          <th scope="col-2">状態</th>
          <th scope="col-2"></th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>
            <input v-model="todo.text" type="text" name="text" placeholder="入力してください" />
          </td>
          <td>
            <select v-model="todo.status" name="status">
              <option value="未実行">未実行</option>
              <option value="実行中">実行中</option>
              <option value="終了">終了</option>
            </select>
          </td>
          <td>
            <button @click="post()">新規保存</button>
          </td>
        </tr>
      </tbody>
    </table>
    <table class="table table-sm w-75 mx-auto bg-white">
      <thead>
        <tr>
          <th scope="col-2">内容</th>
          <th scope="col-2">状態</th>
          <th scope="col-2"></th>
          <th scope="col-2"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="t in todos" :key="t.ID">
          <!-- 編集フォーム -->
          <template v-if="editID == t.ID">
            <td>
              <input v-model="t.Text" placeholder="入力してください" />
            </td>
            <td>
              <select v-model="t.Status" name="status">
                <option value="未実行">未実行</option>
                <option value="実行中">実行中</option>
                <option value="終了">終了</option>
              </select>
            </td>
            <td>
              <button @click="editSubmit(t)">編集終了</button>
            </td>
          </template>
          <template v-else>
            <td>{{ t.Text }}</td>
            <td>{{ t.Status }}</td>
            <td>
              <button @click="showEditForm(t.ID)">編集</button>
            </td>
          </template>

          <td>
            <button @click="del(t.ID)">削除</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      todo: {
        id: 0,
        text: "",
        status: "未実行"
      },
      todos: [],
      url: "http://localhost:5000/todo",
      editID: 0
    };
  },

  mounted: function() {
    axios
      .get(this.url)
      .then(res => {
        this.todos = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  },

  methods: {
    post() {
      const url = "http://localhost:5000/todo";
      const json = JSON.stringify(this.todo);
      axios
        .post(url, json)
        .then(res => {
          this.todos = res.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    showEditForm(id) {
      this.editID = id;
    },
    editSubmit(todo) {
      const json = JSON.stringify(todo);
      axios
        .post(this.url + "/edit/" + todo.ID, json)
        .then(res => {
          this.todos = res.data;
          this.editID = 0;
        })
        .catch(err => {
          console.log(err);
        });
    },
    del(id) {
      axios
        .post(this.url + "/delete/" + id)
        .then(res => {
          this.todos = res.data;
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>
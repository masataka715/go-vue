<template>
  <div>
    Message:<input type='text' v-model="form.message"><br/>
    Number:<input type='number' v-model="form.number"><br/>
    <button type="primary" @click="submit()">Submit</button>
    <p>{{ result }}</p>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      form: {
        message: '百',
        number: 100,
      },
      result: "",
    }
  },
  methods: {
    submit() {
      const url = 'http://localhost:5000/test';
      const json = JSON.stringify(this.form);
      console.log(axios.post(url, json))
      axios.post(url, json).then(res => {
        this.result = JSON.stringify(res.data);
      }).catch(err => {
        alert(err);
      });
    }
  }
}
</script>
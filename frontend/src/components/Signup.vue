<template>
  <div class="">
    <h2>Sign Up</h2>
    <p>Sign up for a free account</p>
    <div class="alert alert-danger" v-if="error">
      <p>{{ error }}</p>
    </div>
    <div class="form-group">
      <input 
        type="text" 
        class="form-control"
        placeholder="Enter your email"
        v-model="credentials.email"
      >
    </div>
    <div class="form-group">
      <input 
        type="text" 
        class="form-control"
        placeholder="Enter your username"
        v-model="credentials.username"
      >
    </div>
    <div class="form-group">
      <input
        type="password"
        class="form-control"
        placeholder="Enter your password"
        v-model="credentials.password"
      >
    </div>
    <button class="btn btn-primary" @click="submit()">Create</button>
  </div>
</template>

<script>
import auth from '../auth'
export default {
  data() {
    return {
      credentials: {
        email: '',
        username: '',
        password: ''
      },
      error: ''
    }
  },
  methods: {
    submit() {
      auth.signup(this.credentials).then(() => {
        this.$router.push({ path:'/' });
      }).catch(() => {
        alert("was an error creating user")
      })
    }
  }
}
</script>
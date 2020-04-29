<template>
  <div id="app">
    <header>
      <span>Vue.js Auth</span>
    </header>
    <main>
      <router-link to="/login">Login</router-link>
      <router-link to="/signup">Signup</router-link>
      <router-link @click.native="logout" to="">Logout</router-link>
      <router-link @click.native="check" to="">Validate</router-link>
      <router-link @click.native="refresh" to="">Refresh</router-link>
      <router-view></router-view>
    </main>
    <footer>
        <h2>Token</h2>
        <vue-json-pretty
          :path="'res'"
          :data="token"
          :showLine="true">
        </vue-json-pretty>
    </footer>
  </div>
</template>

<script>
import VueJsonPretty from 'vue-json-pretty'
import auth from './auth'

export default {
  name: 'app',
  components: {
    VueJsonPretty,
  },
  data () {
    return {
      token: auth.getToken()
    }
  },
  methods: {
    logout: function() {
      auth.logout()
      this.$nextTick(() => {
          this.$router.go();
        })
      this.$router.push({ path:'/login' });
    },
    refresh: function() {
      auth.refresh().then(() => {
        this.$nextTick(function () {
          this.$router.go()
        })
      })
    },
    check: function() {
      auth.checkAuth().then(rsp => {
        console.log(rsp)
        alert("token is valid")
      }).catch(() => {
        alert("token is invalid")
      })
    }
  },
  watch: {
    token: function () {
      this.token = auth.getToken()
    }
  }
}
</script>

<style>
body {
  margin: 0;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

main {
  text-align: center;
  margin-top: 40px;
}

header {
  margin: 0;
  height: 56px;
  padding: 0 16px 0 24px;
  background-color: #35495E;
  color: #ffffff;
}

header span {
  display: block;
  position: relative;
  font-size: 20px;
  line-height: 1;
  letter-spacing: .02em;
  font-weight: 400;
  box-sizing: border-box;
  padding-top: 16px;
}
</style>
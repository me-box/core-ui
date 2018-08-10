<template>
  <div id="app">
    <div id="logo"><img  src="./assets/logo.png"></div>
    <div v-bind="authenticated" v-if="authenticated" id="menu">
      <ul>
        <li v-on:click="goto('/')">Home</li>
        <li v-on:click="goto('AppStore')">App Store</li>
        <li v-on:click="goto('DataSources')">My Data Sources</li>
        <li v-on:click="goto('MyApps')">My Apps</li>
        <li v-on:click="goto('Settings')">Settings</li>
      </ul>
    </div>
    <div id="contentWrapper">
      <router-view ></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'app',
  data: function () {
    return {
      authToken: null,
      authenticated: false,
     }
  },
  created: function () {
    let authenticated = localStorage.getItem('databoxAuthenticated')
    if(authenticated === "true") {
      this.authenticated = true
    } else {
      this.$router.push('/login')
    }
  },
  methods: {
    goto: function (page) {
      this.$router.push(page)
    }
  }
}
</script>

<style>
html {
  height: 100%;
}
body {
  height: 100%;
}

#contentWrapper {
  display: flex;
  justify-content: center;
  width:95%;
  box-shadow: 2px 2px #888888;
  border: 1px solid  #888888;
  margin: auto;
  height:100%;
}
#logo {
  text-align: center;
  width: 100%;
}
#logo img {
  max-width: 1024px;
  width: 100%;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 60px;
  width: 100%;
  height:100%;
}
#menu {
    text-align: center;
}
#menu li {
  font-size: larger;
  font-weight: bold;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
  color: #42b983;
}
li:hover {
  color: #4293b9;
  cursor: pointer;
}
</style>

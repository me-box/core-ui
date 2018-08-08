<template>
 <div>
   <form class="login" @submit.prevent="login">
     <h1>Log in to databox</h1>
     <label>Databox url:&nbsp;</label>
     <input required v-model="url" type="text" /><br/>
     <label>Password:&nbsp;</label>
     <input required v-model="password" autocomplete="current-password" type="password" placeholder="Password"/>
     <hr/>
     <button type="submit" v-on:click="login">Login</button>
   </form>
 </div>
</template>
<script>
export default {
  name: 'logIn',
  props: {},
  data: function () {
      return {
          password: "",
          url: "127.0.0.1"
      }
  },
  methods: {
      login: function () {
            fetch('https://'+this.url+'/api/connect', {
                    method: "GET",
                    credentials: "include",
                    mode: "cors",
                    headers: {
                        'Authorization':"Token " + this.password,
                    },
                })
                .then( (response) => {
                    console.log(response)
                    if (response.status != 200) {
                        localStorage.setItem('databoxAuthenticated', false)
                        alert("Error trying to log in. Sorry.")
                    } else {
                        localStorage.setItem('databoxURL', this.url)
                        localStorage.setItem('databoxAuthenticated', true)
                        this.$router.push("/")
                    }
                })
                .catch(error => console.error(error));
      }
  }
}
</script>
<style scoped>
div {
  padding: 5px;
}
section {
  display: block;
}
</style>
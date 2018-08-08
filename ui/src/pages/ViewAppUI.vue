<template>
<div id="wrapper">
    <div id="toolbar">
        <ul>
            <li v-on:click="viewLogs"><font-awesome-icon icon="eye" />&nbsp;View Logs</li>
            <li v-on:click="restart"><font-awesome-icon icon="sync-alt" />&nbsp;Restart</li>
            <li v-on:click="remove" ><font-awesome-icon icon="times-circle" />&nbsp;Uninstall</li>
        </ul>
    </div>
    <iframe :src="'/' + this.ui + '/ui'" />
</div>
</template>
<script>

export default {
  name: 'appUI',
  props: {},
  data: function () {
     return { ui: ""}
  },
  created() {
    //get app name from Query string
    this.ui = this.$route.query['ui']
  },
  methods: {
      viewLogs: function () {
          alert("Comming soon!!")
      },
      restart: function () {
          let data = {
                name: this.ui,
            }
            fetch(`/core-ui/ui/api/restart`, {
                    method: "POST",
                    credentials: "same-origin",
                    headers: {
                        'Accept': 'application/json, text/plain, */*',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data),
                })
                .then( (response) => {
                    if (response.status == 401) {
                        localStorage.setItem('databoxAuthenticated', false)
                        this.$router.push('/')
                        return
                    }
                    if (response.status != 200) {
                        alert("Error restarting app. Sorry.")
                    } else {
                        this.$router.push("MyApps")
                    }
                })
                .catch(error => console.error(error));
      },
      remove: function () {
          let data = {
                name: this.ui,
            }
            fetch(`/core-ui/ui/api/uninstall`, {
                    method: "POST",
                    credentials: "same-origin",
                    headers: {
                        'Accept': 'application/json, text/plain, */*',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data),
                })
                .then( (response) => {
                    if (response.status == 401) {
                        localStorage.setItem('databoxAuthenticated', false)
                        this.$router.push('/')
                        return
                    }
                    if (response.status != 200) {
                        alert("Error restarting app. Sorry.")
                    } else {
                        this.$router.push("MyApps")
                    }
                })
                .catch(error => console.error(error));
      }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#wrapper {
    width: 100%;
}
#toolbar {
    width: 100%;
    text-align: right;
    height:3em;
}
#toolbar li {
    color: #42b983;
    cursor: pointer;
}
#toolbar li:hover {
    color: #4293b9;
    cursor: pointer;
}
.fa-times-circle {
    color: rgba(255, 91, 91, 0.849);
}
iframe {
    width: 100%;
    height: 100%;
    border:none;
}
</style>
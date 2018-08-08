<template>
  <div id="databoxStatus">
    <h1>Installed Apps</h1>
    <p>
      Hear is a list of the installed and running Apps.
    </p>
    <ul v-if="status">
      <li v-for="item in status"  v-if="item.type == 'app' || item.type == 'driver'" :key="item.name" v-on:click="GoToUI(item.name)">
        <icon v-bind:name="item.name"
              v-bind:displayName="true"
              :key="item"
        />
      </li>
    </ul>
  </div>
</template>
<script>
import testdata from '../testData/status.json'
import Icon from '../components/renderAppDriverIcon.vue'

export default {
  name: 'databoxStatus',
  props: {},
  components: {
		Icon,
	},
  data() {
    //get data from api later
    return { status: {}, timerID: 0}
  },
  mounted() {
    let _this = this

    this.loadData(_this);

    this.timerID = setInterval(() => {
      this.loadData(_this);
    }, 1000);
  },
  destroyed: function () {
    clearInterval(this.timerID)
  },
  methods: {
    loadData: (_this) => {
      //console.log("loadData")
      fetch('/core-ui/ui/api/containerStatus', {credentials: "same-origin"})
      .then((response) => {
            if (response.status == 401) {
                localStorage.setItem('databoxAuthenticated', false)
                this.$router.push('/')
                return
            }
            return response.json()
        })
        .then(json => {
            _this.status = json;
        })
        .catch(()=>{
            //console.log("loadData: error using test data")
            _this.status = testdata;
        });
    },
     GoToUI: function (appName) {
        this.$router.push("view?ui="+appName)
     },
     Restart: function (appName) {
       alert("Install " + appName)
     },
     Uninstall: function (appName) {
       alert("Uninstall" + appName)
     }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
li {
    padding: .5em;
}
</style>

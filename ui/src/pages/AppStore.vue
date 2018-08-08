<template>
  <div class="appStore">
    <div id="wrapper">
    <div>
    <h1>{{ msg }}</h1>
    <p>
      Here is a list of the available Apps for your databox.
    </p>
    <h3>Available Apps</h3>
    <ul v-if="apps">
      <li v-for="item in apps" :key="item" v-on:click="installApp(item)">
        <icon
          v-bind:name="item"
          v-bind:displayName="true"
          :key="item"
        />
      </li>
      <li v-for="item in drivers" :key="item" v-on:click="installApp(item)">
        <icon
          v-bind:name="item"
          v-bind:displayName="true"
          :key="item"
        />
      </li>
    </ul>

  </div>
  </div>
  </div>
</template>
<script>
import testdata from '../testData/apps.json'
import Icon from '../components/renderAppDriverIcon.vue'

export default {
  name: 'AppStore',
  props: {
    msg: String,
  },
  components: {
		Icon,
	},
  data() {
    //get data from api later
    return { apps: [], drivers:[], timerID: 0}
  },
  mounted() {
    let _this = this

    this.loadData(_this);

    this.timerID = setInterval(() => {
      this.loadData(_this);
    }, 5000);
  },
  destroyed: function () {
    clearInterval(this.timerID)
  },
  methods: {
    loadData: (_this) => {
      fetch('/core-ui/ui/api/appStore', {credentials: "same-origin"})
      .then((response) => {
            if (response.status == 401) {
                localStorage.setItem('databoxAuthenticated', false)
                this.$router.push('/')
                return
            }
            return response.json()
        })
        .then(json => {
            _this.apps = json.apps;
            _this.drivers = json.drivers;
        })
        .catch(()=>{
            _this.apps = testdata.apps;
            _this.drivers = testdata.drivers;
        });
    },
     installApp: function (appName) {
       this.$router.push("Install?manifest="+appName)
     }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
div {
    display: block;
    text-align: left;
}
</style>

<template>
  <div class="appStore">
    <h1>{{ msg }}</h1>
    <p>
      Hear is a list of the available Apps for your databox.
    </p>
    <h3>Available Apps</h3>
    <ul v-if="apps">
      <li v-for="item in apps" :key="item" v-on:click="installApp(item)">
        <icon
          v-bind:name="item"

        />
      </li>
    </ul>

    <h3>Available Drivers</h3>
    <ul v-if="apps">
      <li v-for="item in drivers" :key="item" v-on:click="installApp(item)">
        <icon
          v-bind:name="item"
        />
      </li>
    </ul>

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
    return { apps: testdata.apps, drivers:testdata.drivers}
    },
  methods: {
      installApp: function (appName) {
        alert("Install " + appName)
      }
    }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

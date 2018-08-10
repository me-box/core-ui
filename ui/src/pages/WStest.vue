<template>
  <div id="ws">
    <h1>Websockets Test</h1>
    <ul>
     <li class="revived" v-bind:key="r" v-for="r in received">{{ r}}</li>
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
    return { received: [], timerID: 0}
  },
  mounted() {

    this.startWS();

  },
  destroyed: function () {
    clearInterval(this.timerID)
  },
  methods: {
    startWS: function () {
        let _this = this
      //setup a test websocket connection
      let ws = new WebSocket("wss://" + location.host + "/core-ui/ui/api/ws");
      ws.onopen = function(evt) {
          console.log("OPEN");
      }
      ws.onclose = function(evt) {
          console.log("CLOSE");
          ws = null;
      }
      ws.onmessage = function(evt) {
          console.log("RESPONSE: " + evt.data);
          _this.received.push(evt.data)
      }
      ws.onerror = function(evt) {
          console.log("ERROR: " + evt.data);
      }

      this.timerID = setInterval(() => {
        if (ws != null) {
            ws.send("Ping!");
        } else {
            _this.received.push("WS NOT CONNECTED")
        }
      }, 500);
    },
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
